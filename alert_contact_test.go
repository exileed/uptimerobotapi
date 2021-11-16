package uptimerobotapi

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetAlertContacts(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/v2/getAlertContacts", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
				{
  					"stat": "ok",
  					"limit": 50,
  					"offset": 0,
  					"total": 1,
  					"alert_contacts": [
    					{
      						"id": "093765",
      						"friendly_name": "John Doe",
      						"type": 2,
      						"status": 1,
      						"value": "johndoe@gmail.com"
    					}
					]
				}
			`)
	})

	params := GetAlertContactsParams{}
	account, err := client.AlertContact.GetAlertContacts(params)

	require.NoError(t, err)

	alertContact := []AlertContact{{Id: "093765", FriendlyName: "John Doe", Type: 2, Status: 1, Value: "johndoe@gmail.com"}}

	want := &AlertContactResp{Stat: StatOk, Limit: 50, Offset: 0, Total: 1, AlertContacts: alertContact}

	require.Equal(t, want, account)
}

func TestNewAlertContact(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/v2/newAlertContact", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
				{
    				"stat": "ok",
    				"alertcontact": {
        				"id": 4561
    				}
				}
			`)
	})

	params := NewAlertContactParams{}
	account, err := client.AlertContact.NewAlertContact(params)

	require.NoError(t, err)

	want := &AlertContactSingleResp{Stat: StatOk, AlertContact: AlertContactSingle{Id: 4561}}

	require.Equal(t, want, account)
}

func TestEditAlertContact(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/v2/editAlertContact", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
				{
    				"stat": "ok",
    				"alertcontact": {
        				"id": 4561
    				}
				}
			`)
	})

	friendlyName := "test"
	params := EditAlertContactParams{
		Id:           4561,
		FriendlyName: &friendlyName,
	}
	account, err := client.AlertContact.EditAlertContact(params)

	require.NoError(t, err)

	want := &AlertContactEditResp{Stat: StatOk, AlertContact: AlertContactSingle{Id: 4561}}

	require.Equal(t, want, account)
}

func TestDeleteAlertContact(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/v2/deleteAlertContact", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
				{
    				"stat": "ok",
    				"alertcontact": {
        				"id": 4561
    				}
				}
			`)
	})

	account, err := client.AlertContact.DeleteAlertContact(4561)

	require.NoError(t, err)

	want := &AlertContactEditResp{Stat: StatOk, AlertContact: AlertContactSingle{Id: 4561}}

	require.Equal(t, want, account)
}
