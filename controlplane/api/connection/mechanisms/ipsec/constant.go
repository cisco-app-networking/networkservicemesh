package ipsec

import (
	"github.com/networkservicemesh/networkservicemesh/controlplane/api/connection/mechanisms/common"
)

const (
	// Mechanism string
	MECHANISM = "IPSEC"

	// Mechanism parameters
	// SrcIP - source IP
	SrcIP = common.SrcIP
	// DstIP - destination IP
	DstIP = common.DstIP
	// SrcOriginalIP - original src IP
	SrcOriginalIP = common.SrcOriginalIP
	// DstExternalIP - external destination ip
	DstExternalIP = common.DstExternalIP
	// Protocol
	Protocol = "protocol"
	// available SA pool for remote to choose from
	SAIndexPoolStart = "SAIndexPoolStart"
	SAIndexPoolEnd = "SAIndexPoolEnd"
	// LocalSAOutIndex
	LocalSAOutIndex = "localSAOutIdx"
	// LocalSAInIndex
	LocalSAInIndex = "localSAInIdx"
	// RemoteSAOutIndex
	RemoteSAOutIndex = "localSAOutIdx"
	// RemoteSAInIndex
	RemoteSAInIndex = "localSAInIdx"
	// LocalEncrKey
	LocalEncrKey = "localEncrKey"
	// RemoteEncrKey
	RemoteEncrKey = "remoteEncrKey"
	// LocalEncrKey
	LocalIntegKey = "localIntegKey"
	// LocalEncrKey
	RemoteIntegKey = "remoteIntegKey"
	// LocalEspSPI
	LocalEspSPI = "localEspSpi"
	// RemoteEspSPI
	RemoteEspSPI = "remoteEspSpi"
	// EncrAlgo
	EncrAlgo = "encrAlgo"
	// IntegAlgo
	IntegAlgo = "integAlgo"
	// UdpEncap
	UDPEncap = "udpEncap"
	// UseESN
	UseESN = "useEsn"

	// MTUOverhead - maximum transmission unit overhead for VXLAN encapsulation
	MTUOverhead = 70
)
