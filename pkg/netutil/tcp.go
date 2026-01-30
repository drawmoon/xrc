package netutil

import (
	"fmt"
	"log/slog"
	"net"
	"time"
)

// PickFreeTCPPort finds and returns a free TCP port on the local machine.
func PickFreeTCPPort() (int, error) {
	const maxRetries = 5
	const waitInterval = 10 * time.Millisecond

	var port int
	var lastErr error

	for i := 0; i < maxRetries; i++ {
		ln, cvtErr := net.Listen("tcp", "127.0.0.1:0")
		if cvtErr != nil {
			time.Sleep(waitInterval)
			lastErr = cvtErr
			continue
		}

		cvtErr = nil

		if addr, ok := ln.Addr().(*net.TCPAddr); ok {
			port = addr.Port
		} else {
			msg := "could not cast terminal address to *net.TCPAddr"
			slog.Error(msg)
			cvtErr = fmt.Errorf(msg)
		}

		// Ensure the listener is closed after retrieving the port.
		closeErr := ln.Close()
		if cvtErr != nil {
			return 0, cvtErr
		} else if closeErr != nil {
			// Here we prioritize returning any previous error over the close error.
			if lastErr == nil {
				lastErr = closeErr
			}
			return 0, lastErr
		}

		if port != 0 {
			return port, nil
		}
	}

	return 0, fmt.Errorf("failed to pick a free port after %d attempts: %w", maxRetries, lastErr)
}
