package ipsec

import (
	"strconv"

	"github.com/pkg/errors"

	"github.com/networkservicemesh/networkservicemesh/controlplane/api/connection"
	"github.com/networkservicemesh/networkservicemesh/controlplane/api/connection/mechanisms/common"
	"github.com/networkservicemesh/networkservicemesh/controlplane/api/connection/mechanisms/vxlan"
)

// Mechanism - a vxlan mechanism utility wrapper
type Mechanism interface {
	// SrcIP -  src ip
	SrcIP() (string, error)
	// DstIP - dst ip
	DstIP() (string, error)
	// SAOutIndex - SAOut index
	LocalSAOutIndex() (uint32, error)
	LocalSAInIndex() (uint32, error)
	RemoteSAOutIndex() (uint32, error)
	RemoteSAInIndex() (uint32, error)
	LocalSPI() (string, error)
	RemoteSPI() (string, error)
	LocalIntegKey() (string, error)
	RemoteIntegKey() (string, error)
	LocalEncrKey() (string, error)
	RemoteEncrKey() (string, error)
	EnableUdpEncap() (bool, error)
	UseEsn() (bool, error)
	VNI() (uint32, error)
}

type mechanism struct {
	*connection.Mechanism
}

func (m *mechanism) VNI() (uint32, error) {
	if m == nil {
		return 0, errors.New("mechanism cannot be nil")
	}

	if m.GetParameters() == nil {
		return 0, errors.Errorf("mechanism.Parameters cannot be nil: %v", m)
	}

	vxlanvni, ok := m.Parameters[vxlan.VNI]
	if !ok {
		return 0, errors.Errorf("mechanism.Type %s requires mechanism.Parameters[%s]", m.GetType(), vxlan.VNI)
	}

	vni, err := strconv.ParseUint(vxlanvni, 10, 24)

	if err != nil {
		return 0, errors.Wrapf(err, "mechanism.Parameters[%s] must be a valid 24-bit unsigned integer, instead was: %s: %v", vxlan.VNI, vxlanvni, m)
	}

	return uint32(vni), nil
}

// ToMechanism - convert unified mechanism to useful wrapper
func ToMechanism(m *connection.Mechanism) Mechanism {
	if m.Type == MECHANISM {
		return &mechanism{
			m,
		}
	}
	return nil
}

func (m *mechanism) SrcIP() (string, error) {
	return common.GetSrcIP(m.Mechanism)
}

func (m *mechanism) DstIP() (string, error) {
	return common.GetDstIP(m.Mechanism)
}

// LocalSAInIndex returns the local SAIn index of the Mechanism
func (m *mechanism) LocalSAInIndex() (uint32, error) {
	saInIdx, ok := m.Parameters[LocalSAInIndex]
	if !ok {
		return 0, errors.Errorf("mechanism.Type %s requires mechanism.Parameters[%s]", m.GetType(), LocalSAInIndex)
	}

	saIn, err := strconv.Atoi(saInIdx)
	if err != nil {
		return 0, errors.Errorf("mechanism.Type %s requires mechanism.Parameters[%s]", m.GetType(), LocalSAInIndex)

	}
	return uint32(saIn), nil
}

// LocalSAOutIndex returns the local SAOut index of the Mechanism
func (m *mechanism) LocalSAOutIndex() (uint32, error) {
	saOutIdx, ok := m.Parameters[LocalSAOutIndex]
	if !ok {
		return 0, errors.Errorf("mechanism.Type %s requires mechanism.Parameters[%s]", m.GetType(), LocalSAOutIndex)
	}

	saOut, err := strconv.Atoi(saOutIdx)
	if err != nil {
		return 0, errors.Errorf("mechanism.Type %s requires mechanism.Parameters[%s]", m.GetType(), LocalSAOutIndex)

	}
	return uint32(saOut), nil
}

// RemoteSAOutIndex returns the remote SAOut index of the Mechanism
func (m *mechanism) RemoteSAOutIndex() (uint32, error) {
	saOutIdx, ok := m.Parameters[LocalSAOutIndex]
	if !ok {
		return 0, errors.Errorf("mechanism.Type %s requires mechanism.Parameters[%s]", m.GetType(), RemoteSAOutIndex)
	}

	saOut, err := strconv.Atoi(saOutIdx)
	if err != nil {
		return 0, errors.Errorf("mechanism.Type %s requires mechanism.Parameters[%s]", m.GetType(), RemoteSAOutIndex)

	}
	return uint32(saOut), nil
}

// RemoteSAInIndex returns the remote SAIn index of the Mechanism
func (m *mechanism) RemoteSAInIndex() (uint32, error) {
	saInIdx, ok := m.Parameters[RemoteSAInIndex]
	if !ok {
		return 0, errors.Errorf("mechanism.Type %s requires mechanism.Parameters[%s]", m.GetType(), RemoteSAInIndex)
	}

	saIn, err := strconv.Atoi(saInIdx)
	if err != nil {
		return 0, errors.Errorf("mechanism.Type %s requires mechanism.Parameters[%s]", m.GetType(), RemoteSAInIndex)

	}
	return uint32(saIn), nil
}

func (m *mechanism) LocalSPI() (string, error) {
	localSpi, ok := m.Parameters[LocalEspSPI]
	if !ok {
		return "", errors.Errorf("mechanism.Type %s requires mechanism.Parameters[%s]", m.GetType(), LocalEspSPI)
	}
	return localSpi, nil
}

func (m *mechanism) RemoteSPI() (string, error) {
	remoteSpi, ok := m.Parameters[RemoteEspSPI]
	if !ok {
		return "", errors.Errorf("mechanism.Type %s requires mechanism.Parameters[%s]", m.GetType(), RemoteEspSPI)
	}
	return remoteSpi, nil
}

func (m *mechanism) LocalEncrKey() (string, error) {
	localEncrKey, ok := m.Parameters[LocalEncrKey]
	if !ok {
		return "", errors.Errorf("mechanism.Type %s requires mechanism.Parameters[%s]", m.GetType(), LocalEncrKey)
	}
	return localEncrKey, nil
}

func (m *mechanism) RemoteEncrKey() (string, error) {
	remoteEncrKey, ok := m.Parameters[RemoteEncrKey]
	if !ok {
		return "", errors.Errorf("mechanism.Type %s requires mechanism.Parameters[%s]", m.GetType(), RemoteEncrKey)
	}
	return remoteEncrKey, nil
}

func (m *mechanism) LocalIntegKey() (string, error) {
	localIntegKey, ok := m.Parameters[LocalIntegKey]
	if !ok {
		return "", errors.Errorf("mechanism.Type %s requires mechanism.Parameters[%s]", m.GetType(), LocalIntegKey)
	}
	return localIntegKey, nil
}

func (m *mechanism) RemoteIntegKey() (string, error) {
	remoteIntegKey, ok := m.Parameters[RemoteIntegKey]
	if !ok {
		return "", errors.Errorf("mechanism.Type %s requires mechanism.Parameters[%s]", m.GetType(), RemoteIntegKey)
	}
	return remoteIntegKey, nil
}

func (m *mechanism) EnableUdpEncap() (bool, error) {
	return true, nil
}

func (m *mechanism) UseEsn() (bool, error) {
	return true, nil
}
