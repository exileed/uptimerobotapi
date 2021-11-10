package uptimerobotapi

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetMWindows(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/v2/getMWindows", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
				{
  					"stat": "ok",
  					"pagination": {
    					"limit": 10,
    					"offset": 0,
    					"total": 1
  					},
					"mwindows": [
    					{
      						"id": 111,
      						"user": 1,
      						"type": 1,
      						"friendly_name": "Once Backup",
      						"start_time": 1461024000,
      						"duration": 12,
      						"value": "",
      						"status": 1
    					}
  					]
				}
			`)
	})

	params := GetMWindowParams{}
	windows, err := client.MWindow.GetMWindows(params)

	require.NoError(t, err)

	mwindows := []MWindow{{Id: 111, User: 1, Type: 1, FriendlyName: "Once Backup", StartTime: 1461024000, Duration: 12, Value: "", Status: 1}}

	want := &MWindowsResp{Stat: StatOk, Pagination: Pagination{Total: 1, Limit: 10, Offset: 0}, MWindows: mwindows}

	require.Equal(t, want, windows)
}
