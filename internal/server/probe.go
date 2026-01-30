package server

const (
	TEST_URL = "https://cp.cloudflare.com/generate_204"
)

// Probe measures the delay of the given link by performing network requests.
func Probe(link *any, port int, times int, timeout int) (int, error) {
	return -1, nil
}

// ProbeAll measures the delay of all given links in parallel.
func ProbeAll(links []*any, times int, timeout int) ([]*any, error) {
	return links, nil
}
