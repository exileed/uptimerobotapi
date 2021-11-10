package uptimerobotapi

import (
	"net/http"
)

type AccountService struct {
	apiClient *Client
}

type AccountResp struct {
	Object
	Account Account `json:"account"`
}

// Account returns basic information for the account making the API request,
type Account struct {
	Email              string `json:"email"`
	UserId             int    `json:"user_id"`
	FirstName          string `json:"firstname"`
	SmsCredits         int    `json:"sms_credits"`
	MonitorLimit       int    `json:"monitor_limit"`
	MonitorInterval    int    `json:"monitor_interval"`
	UpMonitors         int    `json:"up_monitors"`
	DownMonitors       int    `json:"down_monitors"`
	PausedMonitors     int    `json:"paused_monitors"`
	TotalMonitorsCount int    `json:"total_monitors_count"`
}

// GetAccountDetails Get https://uptimerobot.com/#getAccountDetailsWrap
func (ac *AccountService) GetAccountDetails() (*AccountResp, error) {
	obj := &AccountResp{}

	err := ac.apiClient.request(http.MethodPost, "getAccountDetails", nil, &obj)

	return obj, err
}
