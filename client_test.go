package uptimerobotapi

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

// setup sets up a test HTTP server along with a uptimerobotapi.Client that is
// configured to talk to that test server.  Tests should register handlers on
// mux which provide mock responses for the API method being tested.
func setup(t *testing.T) (*http.ServeMux, *httptest.Server, *Client) {
	// mux is the HTTP request multiplexer used with the test server.
	mux := http.NewServeMux()

	// server is a test HTTP server used to provide mock API responses.
	server := httptest.NewServer(mux)

	// client is the UptimeRobot API client being tests
	client := NewClient("")

	url, _ := url.Parse(server.URL)
	client.baseUrl = url

	return mux, server, client
}

// teardown closes the test HTTP server.
func teardown(server *httptest.Server) {
	server.Close()
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %s, want %s", got, want)
	}
}
