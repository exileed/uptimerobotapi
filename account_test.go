package uptimerobotapi

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetAccountDetails(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/v2/getAccountDetails", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
				{
  					"stat": "ok",
					"account": {
    					"email": "test@domain.com",
    					"monitor_limit": 50,
    					"monitor_interval": 1,
    					"up_monitors": 1,
    					"down_monitors": 0,
    					"paused_monitors": 2
  					}
				}
			`)
	})

	account, err := client.Account.GetAccountDetails()

	require.NoError(t, err)

	want := &AccountResp{Object: Object{Stat: "ok"}, Account: Account{Email: "test@domain.com", MonitorLimit: 50, MonitorInterval: 1, UpMonitors: 1, DownMonitors: 0, PausedMonitors: 2}}

	require.Equal(t, want, account)
}
