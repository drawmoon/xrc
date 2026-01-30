package netutil_test

import (
	"fmt"
	"net"
	"testing"

	"github.com/drawmoon/xrc/pkg/netutil"
)

func TestPickFreeTCPPort(t *testing.T) {
	port, err := netutil.PickFreeTCPPort()
	if err != nil {
		t.Fatalf("PickFreeTCPPort failed: %v", err)
	}

	if port <= 0 || port > 65535 {
		t.Errorf("Invalid port number: %d", port)
	}

	// Verify the port is free by trying to listen on it
	ln, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	if err != nil {
		t.Errorf("Port %d is not free: %v", port, err)
	} else {
		ln.Close()
	}

	// Test that multiple calls return different ports (with high probability)
	port2, err := netutil.PickFreeTCPPort()
	if err != nil {
		t.Fatalf("Second PickFreeTCPPort failed: %v", err)
	}
	if port == port2 {
		t.Logf("Warning: Same port returned twice (%d), but this is possible", port)
	}
}
