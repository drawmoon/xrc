package xray

import (
	"errors"

	"github.com/drawmoon/xrc/pkg/proto"
)

// Config represents the overall configuration for the Xray application.
type Config struct {
	Log       *LogConfig `json:"log,omitempty"`     // logging configuration
	Inbounds  []Inbound  `json:"inbounds"`          // list of inbound configurations
	Outbounds []Outbound `json:"outbounds"`         // list of outbound configurations
	Routing   *Routing   `json:"routing,omitempty"` // routing configuration
}

// Link is a generic interface for different subscription link types.
type Link interface {
	*proto.Trojan | *proto.VMess | *proto.VLess
}

// GenterateConfigFromLink generates an Xray configuration from a given subscription link.
func GenterateConfigFromLink[T Link](link T) (*Config, error) {
	switch t := any(link).(type) {
	case *proto.Trojan:
		return newTrojanConfig(t)
	case *proto.VMess:
		return newVmessConfig(t)
	case *proto.VLess:
		return newVlessConfig(t)
	default:
		return nil, errors.ErrUnsupported
	}
}
