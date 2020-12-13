package tests

import (
	"cisco-app-networking.github.io/networkservicemesh/controlplane/api/connection"
	"cisco-app-networking.github.io/networkservicemesh/controlplane/api/connection/mechanisms/kernel"
	"cisco-app-networking.github.io/networkservicemesh/controlplane/api/connection/mechanisms/vxlan"
	"cisco-app-networking.github.io/networkservicemesh/controlplane/pkg/model"
)

var testForwarder1 = &model.Forwarder{
	RegisteredName: "test_data_plane",
	SocketLocation: "tcp:some_addr",
	LocalMechanisms: []*connection.Mechanism{
		&connection.Mechanism{
			Type: kernel.MECHANISM,
		},
	},
	RemoteMechanisms: []*connection.Mechanism{
		&connection.Mechanism{
			Type: vxlan.MECHANISM,
			Parameters: map[string]string{
				vxlan.SrcIP: "127.0.0.1",
			},
		},
	},
	MechanismsConfigured: true,
}
var testForwarder1_1 = &model.Forwarder{
	RegisteredName: "test_data_plane_11",
	SocketLocation: "tcp:some_addr",
	LocalMechanisms: []*connection.Mechanism{
		{
			Type: kernel.MECHANISM,
		},
	},
	RemoteMechanisms: []*connection.Mechanism{
		&connection.Mechanism{
			Type: vxlan.MECHANISM,
			Parameters: map[string]string{
				vxlan.SrcIP: "127.0.0.7",
			},
		},
	},
	MechanismsConfigured: true,
}

var testForwarder2 = &model.Forwarder{
	RegisteredName: "test_data_plane2",
	SocketLocation: "tcp:some_addr",
	LocalMechanisms: []*connection.Mechanism{
		&connection.Mechanism{
			Type: kernel.MECHANISM,
		},
	},
	RemoteMechanisms: []*connection.Mechanism{
		&connection.Mechanism{
			Type: vxlan.MECHANISM,
			Parameters: map[string]string{
				vxlan.SrcIP: "127.0.0.2",
			},
		},
	},
	MechanismsConfigured: true,
}
