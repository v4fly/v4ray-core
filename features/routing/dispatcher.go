package routing

import (
	"context"

	"github.com/v4fly/v4ray-core/v0/common/net"
	"github.com/v4fly/v4ray-core/v0/features"
	"github.com/v4fly/v4ray-core/v0/transport"
)

// Dispatcher is a feature that dispatches inbound requests to outbound handlers based on rules.
// Dispatcher is required to be registered in a V2Ray instance to make V2Ray function properly.
//
// v2ray:api:stable
type Dispatcher interface {
	features.Feature

	// Dispatch returns a Ray for transporting data for the given request.
	Dispatch(ctx context.Context, dest net.Destination) (*transport.Link, error)
}

// DispatcherType returns the type of Dispatcher interface. Can be used to implement common.HasType.
//
// v2ray:api:stable
func DispatcherType() interface{} {
	return (*Dispatcher)(nil)
}
