package common

import (
	"net"
	"strconv"

	"github.com/pkg/errors"

	"cisco-app-networking.github.io/networkservicemesh/controlplane/api/connection"
)

// SrcIP returns the source IP parameter of the Mechanism
func GetSrcIP(m *connection.Mechanism) (string, error) {
	return getIPParameter(m, SrcIP)
}

// DstIP returns the destination IP parameter of the Mechanism
func GetDstIP(m *connection.Mechanism) (string, error) {
	return getIPParameter(m, DstIP)
}

func getIPParameter(m *connection.Mechanism, name string) (string, error) {
	if m == nil {
		return "", errors.New("mechanism cannot be nil")
	}

	if m.GetParameters() == nil {
		return "", errors.Errorf("mechanism.Parameters cannot be nil: %v", m)
	}

	ip, ok := m.Parameters[name]
	if !ok {
		return "", errors.Errorf("mechanism.Type %s requires mechanism.Parameters[%s]", m.GetType(), name)
	}

	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return "", errors.Errorf("mechanism.Parameters[%s] must be a valid IPv4 or IPv6 address, instead was: %s: %v", name, ip, m)
	}

	return ip, nil
}

// SetMTUOverhead sets the MTU overhead parameter in the mechanism
func SetMTUOverhead(m *connection.Mechanism, mtu uint32) error {
	if m == nil {
		return errors.New("mechanism cannot be nil")
	}
	return SetMTUOverheadParameter(m.Parameters, mtu)
}

// SetMTUOverhead sets the MTU overhead parameter in the parameter map
func SetMTUOverheadParameter(parameters map[string]string, mtu uint32) error {
	if parameters == nil {
		return errors.Errorf("mechanism parameters cannot be nil")
	}
	parameters[MTUOverhead] = strconv.FormatUint(uint64(mtu), 10)
	return nil
}

// GetMTUOverhead returns the MTU overhead value from the mechanism
func GetMTUOverhead(m *connection.Mechanism) (uint32, error) {
	if m == nil {
		return 0, errors.New("mechanism cannot be nil")
	}
	if m.Parameters == nil {
		return 0, errors.Errorf("mechanism parameters cannot be nil")
	}
	overheadParam, ok := m.Parameters[MTUOverhead]
	if !ok {
		return 0, nil // parameter not found - return 0
	}
	overhead, err := strconv.Atoi(overheadParam)
	if err != nil {
		return 0, errors.Errorf("cannot convert mechanism.Parameters[%s] to number: %v", MTUOverhead, err)
	}
	return uint32(overhead), nil
}
