package simplehttp

import (
	"net/http"
	"time"
)

type RetryRoundTripper struct {
	Proxied  http.RoundTripper
	Attempts int
	Delay    time.Duration
}

func (rrt *RetryRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	var resp *http.Response
	var err error
	for i := 0; i < rrt.Attempts; i++ {
		resp, err = rrt.Proxied.RoundTrip(req)
		if err == nil {
			return resp, nil
		}
		time.Sleep(rrt.Delay)
	}
	return resp, err
}
