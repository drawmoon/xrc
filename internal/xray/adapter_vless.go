package xray

import (
	"errors"

	"github.com/drawmoon/xrc/pkg/proto"
)

// newVlessConfig generates an Xray configuration for a VLess link.
func newVlessConfig(_ *proto.VLess) (*Config, error) {
	return nil, errors.ErrUnsupported
}
