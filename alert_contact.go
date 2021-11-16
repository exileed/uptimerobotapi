package uptimerobotapi

import (
	"net/http"
)

type AlertContactService struct {
	apiClient *Client
}

type AlertContactSingleResp struct {
	Stat         string             `json:"stat"`
	AlertContact AlertContactSingle `json:"alertcontact"`
}

type AlertContactEditResp struct {
	Stat         string             `json:"stat"`
	AlertContact AlertContactSingle `json:"alertcontact"`
}

type AlertContactSingle struct {
	Id int `json:"id"`
}

type AlertContactResp struct {
	Stat          string         `json:"stat"`
	Limit         int            `json:"limit"`
	Offset        int            `json:"offset"`
	Total         int            `json:"total"`
	AlertContacts []AlertContact `json:"alert_contacts"`
}

type AlertContact struct {
	Id           string `json:"id"`
	FriendlyName string `json:"friendly_name"`
	Type         int    `json:"type"`
	Status       int    `json:"status"`
	Value        string `json:"value"`
}

// GetAlertContactsParams are parameters for GetAlertContacts.
type GetAlertContactsParams struct {
	Offset        *int    `url:"offset,omitempty"`
	Limit         *int    `url:"limit,omitempty"`
	AlertContacts *string `url:"alert_contacts,omitempty"`
}

// NewAlertContactParams are parameters for NewAlertContact.
type NewAlertContactParams struct {
	TypeContact  string `url:"type"`
	Value        string `url:"value"`
	FriendlyName string `url:"friendly_name"`
}

// EditAlertContactParams are parameters for EditAlertContact.
type EditAlertContactParams struct {
	Id           int     `url:"id"`
	Value        *string `url:"value,omitempty"`
	FriendlyName *string `url:"friendly_name,omitempty"`
}

// DeleteAlertContactWrapper are wrapper for DeleteAlertContact.
type DeleteAlertContactWrapper struct {
	Id int `url:"id"`
}

// GetAlertContacts Get https://uptimerobot.com/#getAccountDetailsWrap
// The list of alert contacts
func (ac *AlertContactService) GetAlertContacts(params GetAlertContactsParams) (*AlertContactResp, error) {
	obj := &AlertContactResp{}

	err := ac.apiClient.request(http.MethodPost, "getAlertContacts", params, &obj)

	return obj, err
}

// NewAlertContact Get https://uptimerobot.com/#newAlertContactWrap
func (ac *AlertContactService) NewAlertContact(params NewAlertContactParams) (*AlertContactSingleResp, error) {
	obj := &AlertContactSingleResp{}

	err := ac.apiClient.request(http.MethodPost, "newAlertContact", params, &obj)

	return obj, err
}

// EditAlertContact Get https://uptimerobot.com/#editAlertContactWrap
func (ac *AlertContactService) EditAlertContact(params EditAlertContactParams) (*AlertContactEditResp, error) {
	obj := &AlertContactEditResp{}

	err := ac.apiClient.request(http.MethodPost, "editAlertContact", params, &obj)

	return obj, err
}

// DeleteAlertContact Get https://uptimerobot.com/#deleteAlertContactWrap
func (ac *AlertContactService) DeleteAlertContact(id int) (*AlertContactEditResp, error) {
	obj := &AlertContactEditResp{}

	params := DeleteAlertContactWrapper{Id: id}

	err := ac.apiClient.request(http.MethodPost, "deleteAlertContact", params, &obj)

	return obj, err
}
