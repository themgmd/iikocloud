package http

import (
	"net/http"
	"strconv"
	"time"
)

// Option .
type Option func(*http.Request)

// WithCustomTimeout .
func WithCustomTimeout(t time.Duration) Option {
	return func(r *http.Request) {
		r.Header.Set("Timeout", strconv.Itoa(int(t.Seconds())))
	}
}

// WithAuthorization .
func WithAuthorization(token string) Option {
	return func(r *http.Request) {
		r.Header.Set("Authorization", "Bearer "+token)
	}
}
