package scenarios

import (
	"testing"
	"time"

	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/types/known/anypb"

	core "github.com/v4fly/v4ray-core/v0"
	"github.com/v4fly/v4ray-core/v0/app/log"
	"github.com/v4fly/v4ray-core/v0/app/proxyman"
	"github.com/v4fly/v4ray-core/v0/common"
	clog "github.com/v4fly/v4ray-core/v0/common/log"
	"github.com/v4fly/v4ray-core/v0/common/net"
	"github.com/v4fly/v4ray-core/v0/common/protocol"
	"github.com/v4fly/v4ray-core/v0/common/serial"
	"github.com/v4fly/v4ray-core/v0/common/uuid"
	"github.com/v4fly/v4ray-core/v0/proxy/dokodemo"
	"github.com/v4fly/v4ray-core/v0/proxy/freedom"
	"github.com/v4fly/v4ray-core/v0/proxy/vmess"
	"github.com/v4fly/v4ray-core/v0/proxy/vmess/inbound"
	"github.com/v4fly/v4ray-core/v0/proxy/vmess/outbound"
	"github.com/v4fly/v4ray-core/v0/testing/servers/tcp"
	"github.com/v4fly/v4ray-core/v0/testing/servers/udp"
)

func TestDokodemoTCP(t *testing.T) {
	tcpServer := tcp.Server{
		MsgProcessor: xor,
	}
	dest, err := tcpServer.Start()
	common.Must(err)
	defer tcpServer.Close()

	userID := protocol.NewID(uuid.New())
	serverPort := tcp.PickPort()
	serverConfig := &core.Config{
		App: []*anypb.Any{
			serial.ToTypedMessage(&log.Config{
				Error: &log.LogSpecification{Level: clog.Severity_Debug, Type: log.LogType_Console},
			}),
		},
		Inbound: []*core.InboundHandlerConfig{
			{
				ReceiverSettings: serial.ToTypedMessage(&proxyman.ReceiverConfig{
					PortRange: net.SinglePortRange(serverPort),
					Listen:    net.NewIPOrDomain(net.LocalHostIP),
				}),
				ProxySettings: serial.ToTypedMessage(&inbound.Config{
					User: []*protocol.User{
						{
							Account: serial.ToTypedMessage(&vmess.Account{
								Id: userID.String(),
							}),
						},
					},
				}),
			},
		},
		Outbound: []*core.OutboundHandlerConfig{
			{
				ProxySettings: serial.ToTypedMessage(&freedom.Config{}),
			},
		},
	}

	clientPort := uint32(tcp.PickPort())
	clientPortRange := uint32(5)
	clientConfig := &core.Config{
		App: []*anypb.Any{
			serial.ToTypedMessage(&log.Config{
				Error: &log.LogSpecification{Level: clog.Severity_Debug, Type: log.LogType_Console},
			}),
		},
		Inbound: []*core.InboundHandlerConfig{
			{
				ReceiverSettings: serial.ToTypedMessage(&proxyman.ReceiverConfig{
					PortRange: &net.PortRange{From: clientPort, To: clientPort + clientPortRange},
					Listen:    net.NewIPOrDomain(net.LocalHostIP),
				}),
				ProxySettings: serial.ToTypedMessage(&dokodemo.Config{
					Address: net.NewIPOrDomain(dest.Address),
					Port:    uint32(dest.Port),
					NetworkList: &net.NetworkList{
						Network: []net.Network{net.Network_TCP},
					},
				}),
			},
		},
		Outbound: []*core.OutboundHandlerConfig{
			{
				ProxySettings: serial.ToTypedMessage(&outbound.Config{
					Receiver: []*protocol.ServerEndpoint{
						{
							Address: net.NewIPOrDomain(net.LocalHostIP),
							Port:    uint32(serverPort),
							User: []*protocol.User{
								{
									Account: serial.ToTypedMessage(&vmess.Account{
										Id: userID.String(),
									}),
								},
							},
						},
					},
				}),
			},
		},
	}

	servers, err := InitializeServerConfigs(serverConfig, clientConfig)
	common.Must(err)
	defer CloseAllServers(servers)

	for port := clientPort; port <= clientPort+clientPortRange; port++ {
		if err := testTCPConn(net.Port(port), 1024, time.Second*2)(); err != nil {
			t.Error(err)
		}
	}
}

func TestDokodemoUDP(t *testing.T) {
	udpServer := udp.Server{
		MsgProcessor: xor,
	}
	dest, err := udpServer.Start()
	common.Must(err)
	defer udpServer.Close()

	userID := protocol.NewID(uuid.New())
	serverPort := tcp.PickPort()
	serverConfig := &core.Config{
		Inbound: []*core.InboundHandlerConfig{
			{
				ReceiverSettings: serial.ToTypedMessage(&proxyman.ReceiverConfig{
					PortRange: net.SinglePortRange(serverPort),
					Listen:    net.NewIPOrDomain(net.LocalHostIP),
				}),
				ProxySettings: serial.ToTypedMessage(&inbound.Config{
					User: []*protocol.User{
						{
							Account: serial.ToTypedMessage(&vmess.Account{
								Id: userID.String(),
							}),
						},
					},
				}),
			},
		},
		Outbound: []*core.OutboundHandlerConfig{
			{
				ProxySettings: serial.ToTypedMessage(&freedom.Config{}),
			},
		},
	}

	clientPort := uint32(udp.PickPort())
	clientPortRange := uint32(5)
	clientConfig := &core.Config{
		Inbound: []*core.InboundHandlerConfig{
			{
				ReceiverSettings: serial.ToTypedMessage(&proxyman.ReceiverConfig{
					PortRange: &net.PortRange{From: clientPort, To: clientPort + clientPortRange},
					Listen:    net.NewIPOrDomain(net.LocalHostIP),
				}),
				ProxySettings: serial.ToTypedMessage(&dokodemo.Config{
					Address: net.NewIPOrDomain(dest.Address),
					Port:    uint32(dest.Port),
					NetworkList: &net.NetworkList{
						Network: []net.Network{net.Network_UDP},
					},
				}),
			},
		},
		Outbound: []*core.OutboundHandlerConfig{
			{
				ProxySettings: serial.ToTypedMessage(&outbound.Config{
					Receiver: []*protocol.ServerEndpoint{
						{
							Address: net.NewIPOrDomain(net.LocalHostIP),
							Port:    uint32(serverPort),
							User: []*protocol.User{
								{
									Account: serial.ToTypedMessage(&vmess.Account{
										Id: userID.String(),
									}),
								},
							},
						},
					},
				}),
			},
		},
	}

	servers, err := InitializeServerConfigs(serverConfig, clientConfig)
	common.Must(err)
	defer CloseAllServers(servers)

	var errg errgroup.Group
	for port := clientPort; port <= clientPort+clientPortRange; port++ {
		errg.Go(testUDPConn(net.Port(port), 1024, time.Second*5))
	}
	if err := errg.Wait(); err != nil {
		t.Error(err)
	}
}
