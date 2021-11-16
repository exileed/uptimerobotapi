package uptimerobotapi

import (
	"fmt"
	"net/http"
)

const (
	StatFail = "fail"
	StatOk   = "ok"
)

// APIResponse is a response from the Uptime robot API
type APIResponse struct {
	Stat string `json:"stat"`
}

type ErrorResponse struct {
	Stat  string   `json:"stat"`
	Error APIError `json:"error"`
}

// APIError represents the error response received when an API call fails. The
// HTTP response code is set inside of the StatusCode field, with the Message
type APIError struct {
	// StatusCode is the HTTP response status code
	StatusCode int `json:"-"`

	Message string `json:"message,omitempty"`
}

// RateLimited returns whether the response had a status of 429, and as such the
// client is rate limited. The PagerDuty rate limits should reset once per
// minute, and for the API they are an account-wide rate limit (not per
// API key or IP).
func (e APIError) RateLimited() bool {
	return e.StatusCode == http.StatusTooManyRequests
}

// Temporary returns whether it was a temporary error, one of which is a
// RateLimited error.
func (e APIError) Temporary() bool {
	return e.RateLimited() || (e.StatusCode >= 500 && e.StatusCode < 600)
}

// Error returns the error message that came back with the API error.
func (e APIError) Error() string {
	return fmt.Sprintf(
		"HTTP response failed with status code %d, message: %s",
		e.StatusCode,
		e.Message,
	)
}

type Pagination struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

// Object contains the common fields of every resource in the UptimeRobot API.
type Object struct {
	Stat string `json:"stat"`
}
