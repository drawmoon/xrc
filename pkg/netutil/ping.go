package netutil

import (
	"net/http"
	"time"
)

// Ping performs HTTP GET requests to the specified test_url for a given number of times.
// It returns the average response time in milliseconds. If any request fails or
// does not return a 204 status with zero content length, it returns -1.
func Ping(client *http.Client, test_url string, times int) int {
	req, _ := http.NewRequest("GET", test_url, nil)

	var totalElapsedMillis int
	var skip bool

	for i := 0; i < times; i++ {
		elapsedMillis := func() int {
			now := time.Now()

			resp, err := client.Do(req)
			if err != nil {
				return -1
			}
			defer resp.Body.Close()

			millis := int(time.Since(now).Milliseconds())
			if resp.StatusCode == 204 && resp.ContentLength == 0 {
				return millis
			}
			return -1
		}()
		if elapsedMillis == -1 {
			skip = true
			break
		}
		totalElapsedMillis += elapsedMillis
	}

	if skip {
		return -1
	}
	return totalElapsedMillis / times
}
