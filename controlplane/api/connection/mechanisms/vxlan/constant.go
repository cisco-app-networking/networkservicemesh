package vxlan

import (
	"cisco-app-networking.github.io/networkservicemesh/controlplane/api/connection/mechanisms/common"
)

const (
	// Mechanism string
	MECHANISM = "VXLAN"

	// Mechanism parameters
	// SrcIP - source IP
	SrcIP = common.SrcIP
	// DstIP - destitiona IP
	DstIP = common.DstIP
	// SrcOriginalIP - original src IP
	SrcOriginalIP = common.SrcOriginalIP
	// DstExternalIP - external destination ip
	DstExternalIP = common.DstExternalIP
	// VNI - vni
	VNI = "vni"

	// MTUOverhead - maximum transmission unit overhead for VXLAN encapsulation
	MTUOverhead = 50
)
