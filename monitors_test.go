package uptimerobotapi

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetMonitors(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/v2/getMonitors", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
				{
  					"stat": "ok",
  					"pagination": {
    					"offset": 0,
    					"limit": 50,
    					"total": 1
  					},
  					"monitors": [
    					{
      						"id": 111,
      						"friendly_name": "Google",
      						"url": "http://www.google.com",
      						"type": 1,
							"sub_type": "",
            				"keyword_type": null,
            				"keyword_case_type": 0,
            				"port": "",
            				"interval": 600,
            				"timeout": 30,
            				"status": 0,
            				"create_datetime": 1484160022
						}
  					]
				}
			`)
	})

	params := GetMonitorsParams{}
	account, err := client.Monitor.GetMonitors(params)

	require.NoError(t, err)

	timeout := 30
	keywordCaseType := 0
	monitors := []Monitor{{Id: 111, FriendlyName: "Google", Url: "http://www.google.com", Type: 1, Status: 0, Interval: 600, Timeout: timeout, KeywordCaseType: &keywordCaseType, CreateDatetime: 1484160022}}

	want := &MonitorsResp{Stat: StatOk, Pagination: Pagination{Offset: 0, Limit: 50, Total: 1}, Monitors: monitors}

	require.Equal(t, want, account)
}

func TestNewMonitor(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/v2/newMonitor", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
				{
      				"stat": "ok",
      				"monitor": {
        				"id": 111,
        				"status": 1
      				}
				}
			`)
	})

	params := NewMonitorsParams{}
	account, err := client.Monitor.NewMonitor(params)

	require.NoError(t, err)

	status := 1
	want := &MonitorsSingResp{Stat: StatOk, Monitor: MonitorSingle{Id: 111, Status: &status}}

	require.Equal(t, want, account)
}

func TestEditMonitor(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/v2/editMonitor", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
				{
      				"stat": "ok",
      				"monitor": {
        				"id": 111
      				}
				}
			`)
	})

	params := EditMonitorsParams{}
	account, err := client.Monitor.EditMonitor(111, params)

	require.NoError(t, err)

	want := &MonitorsSingResp{Stat: StatOk, Monitor: MonitorSingle{Id: 111}}

	require.Equal(t, want, account)
}

func TestDeleteMonitor(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/v2/deleteMonitor", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
				{
      				"stat": "ok",
      				"monitor": {
        				"id": 111
      				}
				}
			`)
	})

	account, err := client.Monitor.DeleteMonitor(111)

	require.NoError(t, err)

	want := &MonitorsSingResp{Stat: StatOk, Monitor: MonitorSingle{Id: 111}}

	require.Equal(t, want, account)
}

func TestResetMonitor(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/v2/resetMonitor", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
				{
      				"stat": "ok",
      				"monitor": {
        				"id": 111
      				}
				}
			`)
	})

	account, err := client.Monitor.ResetMonitor(111)

	require.NoError(t, err)

	want := &MonitorsSingResp{Stat: StatOk, Monitor: MonitorSingle{Id: 111}}

	require.Equal(t, want, account)
}
