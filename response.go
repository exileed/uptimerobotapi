package uptimerobotapi

const (
	StatFail = "fail"
	StatOk   = "ok"
)

// APIResponse is a response from the Uptime robot API
type APIResponse struct {
	Stat string `json:"stat"`
}

type ErrorResponse struct {
	Stat  string `json:"stat"`
	Error Error  `json:"error"`
}

type Error struct {
	Stat    string            `json:"stat"`
	Value   map[string]string `json:"value"`
	Message string
}

// Error returns the error message that came back with the API error.
func (e *Error) Error() string {
	return e.Message
}

type Pagination struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

// Object contains the common fields of every resource in the Uptimerobot API.
type Object struct {
	Stat string `json:"stat"`
}
