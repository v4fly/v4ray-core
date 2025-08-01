package tun

import (
	"gvisor.dev/gvisor/pkg/tcpip"
	"gvisor.dev/gvisor/pkg/tcpip/stack"

	"github.com/v4fly/v4ray-core/v0/app/tun/device"
	"github.com/v4fly/v4ray-core/v0/app/tun/tunsorter"
)

func NewDeviceWithSorter(overlay device.Device, sorter *tunsorter.TunSorter) device.Device {
	return &packetAddrDevice{
		Device: overlay,
		sorter: sorter,
	}
}

type packetAddrDevice struct {
	device.Device
	sorter *tunsorter.TunSorter

	secondaryDispatcher stack.NetworkDispatcher
}

func (p *packetAddrDevice) DeliverNetworkPacket(protocol tcpip.NetworkProtocolNumber, pkt stack.PacketBufferPtr) {
	buf := pkt.ToBuffer()
	_, err := p.sorter.OnPacketReceived(buf.Flatten())
	if err != nil {
		p.secondaryDispatcher.DeliverNetworkPacket(protocol, pkt)
	}
}

func (p *packetAddrDevice) DeliverLinkPacket(protocol tcpip.NetworkProtocolNumber, pkt stack.PacketBufferPtr) {
	// TODO implement me
	panic("implement me")
}

func (p *packetAddrDevice) Attach(dispatcher stack.NetworkDispatcher) {
	p.secondaryDispatcher = dispatcher
	p.Device.Attach(p)
}
