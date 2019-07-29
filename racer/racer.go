package racer

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	aDuration := measureRequest(a)
	bDuration :=measureRequest(b)

	if aDuration > bDuration {
		return b
	}

	return a
}

func measureRequest(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	duration := time.Since(start)
	return duration
}
