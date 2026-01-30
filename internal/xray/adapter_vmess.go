package xray

import (
	"errors"

	"github.com/drawmoon/xrc/pkg/proto"
)

// newVmessConfig generates an Xray configuration for a VMess link.
func newVmessConfig(_ *proto.VMess) (*Config, error) {
	return nil, errors.ErrUnsupported
}
