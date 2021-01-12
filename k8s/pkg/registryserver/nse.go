package registryserver

import (
	"os"
	"strings"
	"time"

	"cisco-app-networking.github.io/networkservicemesh/pkg/tools/spanhelper"

	"cisco-app-networking.github.io/networkservicemesh/controlplane/pkg/serviceregistry"

	"github.com/pkg/errors"

	"cisco-app-networking.github.io/networkservicemesh/controlplane/pkg/nsmd"

	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"

	"cisco-app-networking.github.io/networkservicemesh/controlplane/api/registry"
	v1 "cisco-app-networking.github.io/networkservicemesh/k8s/pkg/apis/networkservice/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// ForwardingTimeout - Timeout waiting for Proxy NseRegistryClient
	ForwardingTimeout = 15 * time.Second
	// ProxyRegistryReconnectInterval - reconnect interval to Proxy NSMD-K8S if connection refused
	ProxyRegistryReconnectInterval = 15 * time.Second
)

type nseRegistryService struct {
	nsmName string
	cache   RegistryCache
}

func newNseRegistryService(nsmName string, cache RegistryCache) *nseRegistryService {
	return &nseRegistryService{
		nsmName: nsmName,
		cache:   cache,
	}
}

func (rs *nseRegistryService) RegisterNSE(ctx context.Context, request *registry.NSERegistration) (*registry.NSERegistration, error) {
	st := time.Now()

	span := spanhelper.FromContext(ctx, "nsmgr.RegisterNSE")
	defer span.Finish()
	logger := span.Logger()

	logger.Infof("Received RegisterNSE(%v)", request)

	labels := request.GetNetworkServiceEndpoint().GetLabels()
	if labels == nil {
		labels = make(map[string]string)
	}
	labels["networkservicename"] = request.GetNetworkService().GetName()
	if request.GetNetworkServiceEndpoint() != nil && request.GetNetworkService() != nil {
		_, err := rs.cache.AddNetworkService(&v1.NetworkService{
			ObjectMeta: metav1.ObjectMeta{
				Name: request.NetworkService.GetName(),
			},
			Spec: v1.NetworkServiceSpec{
				Payload: request.NetworkService.GetPayload(),
			},
			Status: v1.NetworkServiceStatus{},
		})
		if err != nil {
			logger.Errorf("Failed to register nsm: %s", err)
			return nil, err
		}

		var objectMeta metav1.ObjectMeta
		if request.GetNetworkServiceEndpoint().GetName() == "" {
			objectMeta = metav1.ObjectMeta{
				GenerateName: request.GetNetworkService().GetName(),
				Labels:       labels,
			}
		} else {
			objectMeta = metav1.ObjectMeta{
				Name:   request.GetNetworkServiceEndpoint().GetName(),
				Labels: labels,
			}
		}

		nseResponse, err := rs.cache.AddNetworkServiceEndpoint(&v1.NetworkServiceEndpoint{
			ObjectMeta: objectMeta,
			Spec: v1.NetworkServiceEndpointSpec{
				NetworkServiceName: request.GetNetworkService().GetName(),
				Payload:            request.GetNetworkService().GetPayload(),
				NsmName:            rs.nsmName,
			},
			Status: v1.NetworkServiceEndpointStatus{
				State: v1.RUNNING,
			},
		})
		if err != nil {
			return nil, err
		}

		request.NetworkServiceEndpoint = mapNseFromCustomResource(nseResponse)
		nsm, err := rs.cache.GetNetworkServiceManager(rs.nsmName)
		if err != nil {
			return nil, err
		}
		request.NetworkServiceManager = mapNsmFromCustomResource(nsm)

		go func() {
			if forwardErr := rs.forwardRegisterNSE(context.Background(), request); forwardErr != nil {
				logger.Errorf("Cannot forward NSE Registration: %v", forwardErr)
			}
		}()
	}
	logger.Infof("Returned from RegisterNSE: time: %v request: %v", time.Since(st), request)
	return request, nil
}

func (rs *nseRegistryService) BulkRegisterNSE(srv registry.NetworkServiceRegistry_BulkRegisterNSEServer) error {
	span := spanhelper.FromContext(srv.Context(), "ProxyNsmgr.BulkRegisterNSE")
	defer span.Finish()
	logger := span.Logger()

	logger.Infof("Forwarding Bulk Register NSE stream...")

	nsrURL := os.Getenv(ProxyNsmdK8sAddressEnv)
	if strings.TrimSpace(nsrURL) == "" {
		nsrURL = ProxyNsmdK8sAddressDefaults
	}

	ctx, cancel := context.WithCancel(span.Context())
	defer cancel()

	remoteRegistry := nsmd.NewServiceRegistryAt(nsrURL)
	defer remoteRegistry.Stop()

	for {
		stream, err := requestBulkRegisterNSEStream(ctx, remoteRegistry, nsrURL)
		if err != nil {
			logger.Warnf("Cannot connect to Proxy NSMGR %s : %v", nsrURL, err)
			<-time.After(ProxyRegistryReconnectInterval)
			continue
		}

		for {
			request, err := srv.Recv()
			if err != nil {
				err = errors.Wrapf(err, "error receiving BulkRegisterNSE request : %v", err)
				return err
			}

			logger.Infof("Forward BulkRegisterNSE request: %v", request)
			err = stream.Send(request)
			if err != nil {
				logger.Warnf("Error forwarding BulkRegisterNSE request to %s : %v", nsrURL, err)
				break
			}
		}

		<-time.After(ProxyRegistryReconnectInterval)
	}
}

func requestBulkRegisterNSEStream(ctx context.Context, remoteRegistry serviceregistry.ServiceRegistry, nsrURL string) (registry.NetworkServiceRegistry_BulkRegisterNSEClient, error) {
	nseRegistryClient, err := remoteRegistry.NseRegistryClient(ctx)
	if err != nil {
		err = errors.Wrapf(err, "error forwarding BulkRegisterNSE request to %s : %v", nsrURL, err)
		return nil, err
	}

	stream, err := nseRegistryClient.BulkRegisterNSE(ctx)
	if err != nil {
		err = errors.Wrapf(err, "error forwarding BulkRegisterNSE request to %s : %v", nsrURL, err)
		return nil, err
	}

	return stream, nil
}

func (rs *nseRegistryService) RemoveNSE(ctx context.Context, request *registry.RemoveNSERequest) (*empty.Empty, error) {
	st := time.Now()

	span := spanhelper.FromContext(ctx, "ProxyNsmgr.BulkRegisterNSE")
	defer span.Finish()
	logger := span.Logger()

	logger.Infof("Received RemoveNSE(%v)", request)

	if err := rs.cache.DeleteNetworkServiceEndpoint(request.GetNetworkServiceEndpointName()); err != nil {
		return nil, err
	}

	go func() {
		if forwardErr := rs.forwardRemoveNSE(context.Background(), request); forwardErr != nil {
			logger.Errorf("Cannot forward Remove NSE: %v", forwardErr)
		}
	}()

	logger.Infof("RemoveNSE done: time %v", time.Since(st))
	return &empty.Empty{}, nil
}

func (rs *nseRegistryService) forwardRegisterNSE(ctx context.Context, request *registry.NSERegistration) error {
	span := spanhelper.FromContext(ctx, "ProxyNsmgr.forwardRegisterNSE")
	defer span.Finish()
	logger := span.Logger()

	logger.Infof("Forwarding Register NSE request (%v)", request)

	nsrURL := os.Getenv(ProxyNsmdK8sAddressEnv)
	if strings.TrimSpace(nsrURL) == "" {
		nsrURL = ProxyNsmdK8sAddressDefaults
	}

	spanCtx, cancel := context.WithTimeout(span.Context(), ForwardingTimeout)
	defer cancel()

	done := make(chan registry.NetworkServiceRegistryClient)
	quit := make(chan error)

	remoteRegistry := nsmd.NewServiceRegistryAt(nsrURL)
	defer remoteRegistry.Stop()

	go func() {
		nseRegistryClient, err := remoteRegistry.NseRegistryClient(spanCtx)
		if err != nil {
			quit <- err
			return
		}

		done <- nseRegistryClient
	}()

	var nseRegistryClient registry.NetworkServiceRegistryClient
	select {
	case nseRegistryClient = <-done:
		break
	case err := <-quit:
		return err
	case <-time.After(ForwardingTimeout):
		return errors.Errorf("timeout requesting NseRegistryClient")
	}

	service, err := rs.cache.GetNetworkService(request.NetworkService.Name)
	if err != nil {
		return err
	}

	request.NetworkService.Payload = service.Spec.Payload

	for _, m := range service.Spec.Matches {
		var routes []*registry.Destination

		for _, r := range m.Routes {
			destination := &registry.Destination{
				DestinationSelector: r.DestinationSelector,
				Weight:              r.Weight,
			}
			routes = append(routes, destination)
		}

		match := &registry.Match{
			SourceSelector: m.SourceSelector,
			Routes:         routes,
		}
		request.NetworkService.Matches = append(request.NetworkService.Matches, match)
	}

	_, err = nseRegistryClient.RegisterNSE(spanCtx, request)
	if err != nil {
		return err
	}

	return nil
}

func (rs *nseRegistryService) forwardRemoveNSE(ctx context.Context, request *registry.RemoveNSERequest) error {
	span := spanhelper.FromContext(ctx, "ProxyNsmgr.forwardRemoveNSE")
	defer span.Finish()
	logger := span.Logger()

	logger.Infof("Forwarding Remove NSE request (%v)", request)

	nsrURL := os.Getenv(ProxyNsmdK8sAddressEnv)
	if strings.TrimSpace(nsrURL) == "" {
		nsrURL = ProxyNsmdK8sAddressDefaults
	}

	spanCtx, cancel := context.WithTimeout(span.Context(), ForwardingTimeout)
	defer cancel()

	done := make(chan registry.NetworkServiceRegistryClient)
	quit := make(chan error)

	remoteRegistry := nsmd.NewServiceRegistryAt(nsrURL)
	defer remoteRegistry.Stop()

	go func() {
		nseRegistryClient, err := remoteRegistry.NseRegistryClient(spanCtx)
		if err != nil {
			quit <- err
			return
		}

		done <- nseRegistryClient
	}()

	var nseRegistryClient registry.NetworkServiceRegistryClient
	select {
	case nseRegistryClient = <-done:
		break
	case err := <-quit:
		return err
	case <-time.After(ForwardingTimeout):
		return errors.Errorf("timeout requesting NseRegistryClient")
	}

	_, err := nseRegistryClient.RemoveNSE(spanCtx, request)
	if err != nil {
		return err
	}

	return nil
}
