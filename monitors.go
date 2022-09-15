package uptimerobotapi

import (
	"net/http"
)

type MonitorService struct {
	apiClient *Client
}

type GetMonitorsParams struct {
	Monitors               *string `url:"monitors,omitempty"`
	Types                  *string `url:"types,omitempty"`
	Statuses               *string `url:"statuses,omitempty"`
	CustomUptimeRatios     *string `url:"custom_uptime_ratios,omitempty"`
	CustomUptimeDurations  int     `url:"custom_down_durations,omitempty"`
	AllTimeUptimeRation    int     `url:"all_time_uptime_ratio,omitempty"`
	AllTimeUptimeDurations int     `url:"all_time_uptime_durations,omitempty"`
	Logs                   int     `url:"logs,omitempty"`
	AlertContacts          int     `url:"alert_contacts,omitempty"`
	MWindows               int     `url:"mwindows,omitempty"`
	SSL                    int     `url:"ssl,omitempty"`
	CustomHttpHeaders      int     `url:"custom_http_headers,omitempty"`
	CustomHttpStatuses     int     `url:"custom_http_statuses,omitempty"`
	Timezone               int     `url:"timezone,omitempty"`
	Offset                 int     `url:"offset,omitempty"`
	Limit                  *int    `url:"limit,omitempty"`
	Search                 *string `url:"search,omitempty"`
}

type NewMonitorsParams struct {
	Type               int     `url:"type"`
	FriendlyName       string  `url:"friendly_name"`
	Url                string  `url:"url"`
	SubType            *int    `url:"sub_type,omitempty"`
	Port               *int    `url:"port,omitempty"`
	KeywordType        *int    `url:"keyword_type,omitempty"`
	KeywordCaseType    *int    `json:"keyword_case_type,omitempty"`
	KeywordValue       *string `url:"keyword_value,omitempty"`
	Interval           *int    `url:"interval,omitempty"`
	Timeout            *int    `url:"timeout,omitempty"`
	HttpUsername       *string `url:"http_username,omitempty"`
	HttpPassword       *string `url:"http_username,omitempty"`
	HttpAuthType       *int    `url:"http_username,omitempty"`
	HttpMethod         *int    `url:"http_method,omitempty"`
	AlertContacts      *string `url:"alert_contacts,omitempty"`
	MWindows           *string `url:"mwindows,omitempty"`
	CustomHttpHeaders  *string `url:"custom_http_headers,omitempty"`
	CustomHttpStatuses *string `url:"custom_http_statuses,omitempty"`
	IgnoreSSLErrors    *bool   `url:"ignore_ssl_errors,omitempty"`
}

type EditMonitorsParams struct {
	FriendlyName       string  `url:"friendly_name"`
	Url                string  `url:"url"`
	SubType            *int    `url:"sub_type,omitempty"`
	Port               *int    `url:"port,omitempty"`
	KeywordType        *int    `url:"keyword_type,omitempty"`
	KeywordCaseType    *int    `json:"keyword_case_type,omitempty"`
	KeywordValue       *string `url:"keyword_value,omitempty"`
	Interval           *int    `url:"interval,omitempty"`
	Timeout            *int    `url:"timeout,omitempty"`
	HttpUsername       *string `url:"http_username,omitempty"`
	HttpPassword       *string `url:"http_username,omitempty"`
	HttpAuthType       *int    `url:"http_username,omitempty"`
	HttpMethod         *int    `url:"http_method,omitempty"`
	AlertContacts      *string `url:"alert_contacts,omitempty"`
	MWindows           *string `url:"mwindows,omitempty"`
	CustomHttpHeaders  *string `url:"custom_http_headers,omitempty"`
	CustomHttpStatuses *string `url:"custom_http_statuses,omitempty"`
	IgnoreSSLErrors    *bool   `url:"ignore_ssl_errors,omitempty"`
}

type EditMonitorsWrappers struct {
	Id int `url:"id"`
	EditMonitorsParams
}

type DeleteOrResetMonitorsWrappers struct {
	Id int `url:"id"`
}

type MonitorsResp struct {
	Stat       string     `json:"stat"`
	Pagination Pagination `json:"pagination"`
	Monitors   []Monitor  `json:"monitors"`
}

type MonitorsSingResp struct {
	Stat    string        `json:"stat"`
	Monitor MonitorSingle `json:"monitor"`
}

type MonitorSingle struct {
	Id     int  `json:"id"`
	Status *int `json:"status"`
}

type AlertContactMonitor struct {
	Id         string `json:"id"`
	Value      string `json:"value"`
	Type       int    `json:"type"`
	Threshold  int    `json:"threshold"`
	Recurrence int    `json:"recurrence"`
}

type Monitor struct {
	Id              int                    `json:"id"`
	FriendlyName    string                 `json:"friendly_name"`
	Url             string                 `json:"url"`
	Type            int                    `json:"type"`
	SubType         string                 `json:"sub_type"`
	Port            int                    `json:"port"`
	KeywordType     *int                   `json:"keyword_type"`
	KeywordCaseType *int                   `json:"keyword_case_type"`
	KeywordValue    string                 `json:"keyword_value"`
	HttpUsername    string                 `json:"http_username"`
	HttpPassword    string                 `json:"http_password"`
	Interval        int                    `json:"interval"`
	Timeout         int                    `json:"timeout"`
	Status          int                    `json:"status"`
	CreateDatetime  int                    `json:"create_datetime"`
	MonitorGroup    int                    `json:"monitor_group"`
	IsGroupMain     int                    `json:"is_group_main"`
	Logs            []MonitorLog           `json:"logs"`
	AlertContacts   *[]AlertContactMonitor `json:"alert_contacts"`
	SSL             *MonitorSSL            `json:"ssl"`
}

type MonitorSSL struct {
	Brand                string      `json:"brand"`
	Product              string      `json:"product"`
	Expires              int         `json:"expires"`
	LastCheck            interface{} `json:"last_check"`
	IgnoreErrors         int         `json:"ignore_errors"`
	DisableNotifications int         `json:"disable_notifications"`
}
type MonitorLog struct {
	Type     int `json:"type"`
	Datetime int `json:"datetime"`
	Duration int `json:"duration"`
}

// GetMonitors Get https://uptimerobot.com/#getMonitorsWrap
func (ms *MonitorService) GetMonitors(params GetMonitorsParams) (*MonitorsResp, error) {
	obj := &MonitorsResp{}

	err := ms.apiClient.request(http.MethodPost, "getMonitors", params, &obj)

	return obj, err
}

// NewMonitor Get https://uptimerobot.com/#newMonitorWrap
func (ms *MonitorService) NewMonitor(params NewMonitorsParams) (*MonitorsSingResp, error) {
	obj := &MonitorsSingResp{}

	err := ms.apiClient.request(http.MethodPost, "newMonitor", params, &obj)

	return obj, err
}

// EditMonitor Get https://uptimerobot.com/#editMonitorWrap
func (ms *MonitorService) EditMonitor(id int, request EditMonitorsParams) (*MonitorsSingResp, error) {
	obj := &MonitorsSingResp{}

	params := EditMonitorsWrappers{Id: id, EditMonitorsParams: request}

	err := ms.apiClient.request(http.MethodPost, "editMonitor", params, &obj)

	return obj, err
}

// DeleteMonitor Get https://uptimerobot.com/#deleteMonitorWrap
func (ms *MonitorService) DeleteMonitor(id int) (*MonitorsSingResp, error) {
	obj := &MonitorsSingResp{}

	params := DeleteOrResetMonitorsWrappers{Id: id}

	err := ms.apiClient.request(http.MethodPost, "deleteMonitor", params, &obj)

	return obj, err
}

// ResetMonitor Get https://uptimerobot.com/#resetMonitorWrap
func (ms *MonitorService) ResetMonitor(id int) (*MonitorsSingResp, error) {
	obj := &MonitorsSingResp{}

	params := DeleteOrResetMonitorsWrappers{Id: id}

	err := ms.apiClient.request(http.MethodPost, "resetMonitor", params, &obj)

	return obj, err
}
