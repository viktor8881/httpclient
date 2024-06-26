package simplehttp

import (
	"log"
	"net/http"
	"time"
)

// LoggingRoundTripper is an http.RoundTripper that logs requests and responses
type LoggingRoundTripper struct {
	Proxied   http.RoundTripper
	Logger    *log.Logger
	TurnOnAll bool
}

func NewLoggingRoundTripper(proxied http.RoundTripper, logger *log.Logger, turnOnAll bool) *LoggingRoundTripper {
	return &LoggingRoundTripper{
		Proxied:   proxied,
		Logger:    logger,
		TurnOnAll: turnOnAll,
	}
}

func (lrt *LoggingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	start := time.Now()
	if lrt.TurnOnAll {
		lrt.Logger.Printf("Request: %s %s", req.Method, req.URL)
	}

	resp, err := lrt.Proxied.RoundTrip(req)
	if err != nil {
		lrt.Logger.Printf("Error: %v", err)
		return nil, err
	}

	if lrt.TurnOnAll {
		duration := time.Since(start)
		lrt.Logger.Printf("Response: %s %s in %v", req.Method, req.URL, duration)
		lrt.Logger.Printf("Response Status: %s", resp.Status)
	}

	return resp, nil
}
