package uptimerobotapi

import (
	"net/http"
)

type MWindowService struct {
	apiClient *Client
}

type GetMWindowParams struct {
	FriendlyName string `url:"friendly_name"`
	Type         string `url:"type"`
	Value        string `url:"value"`
	StartTime    string `url:"start_time"`
	Duration     string `url:"duration"`
	Offset        int   `url:"offset,omitempty"`
	Limit        *int   `url:"limit,omitempty"`
}

type NewMWindowParams struct {
	FriendlyName string `url:"friendly_name"`
	Type         string `url:"type"`
	Value        string `url:"value"`
	StartTime    string `url:"start_time"`
	Duration     string `url:"duration"`
}

type EditMWindowParams struct {
	FriendlyName    string  `url:"friendly_name"`
	Url             string  `url:"url"`
	SubType         *int    `url:"sub_type,omitempty"`
	Port            *int    `url:"port,omitempty"`
	KeywordType     *int    `url:"keyword_type,omitempty"`
	KeywordCaseType *int    `json:"keyword_case_type,omitempty"`
	KeywordValue    *string `url:"keyword_value,omitempty"`
	Interval        *int    `url:"interval,omitempty"`
	Timeout         *int    `url:"timeout,omitempty"`
	HttpUsername    *string `url:"http_username,omitempty"`
	HttpPassword    *string `url:"http_username,omitempty"`
	HttpAuthType    *int    `url:"http_username,omitempty"`
	HttpMethod      *int    `url:"http_method,omitempty"`
	AlertContacts   *string `url:"alert_contacts,omitempty"`
	MWindows        *string `url:"mwindows,omitempty"`
	//CustomHttpHeaders  *string `url:"custom_http_headers,omitempty"`
	//CustomHttpStatuses *string `url:"custom_http_statuses,omitempty"`
	IgnoreSSLErrors *bool `url:"ignore_ssl_errors,omitempty"`
}

type DeleteMWindowWrapper struct {
	Id int `url:"id"`
}

type MWindowsResp struct {
	Stat       string     `json:"stat"`
	Pagination Pagination `json:"pagination"`
	MWindows   []MWindow  `json:"mwindows"`
}

type MWindowResp struct {
	Stat    string `json:"stat"`
	MWindow struct {
		Id     int  `json:"id"`
		Status *int `json:"status"`
	} `json:"mwindow"`
}

type MWindow struct {
	Id           int    `json:"id"`
	User         int    `json:"user"`
	Type         int    `json:"type"`
	FriendlyName string `json:"friendly_name"`
	// StartTime comes from API as a string value with a format like “18:20.”
	StartTime string `json:"start_time"`
	Duration  int    `json:"duration"`
	Value     string `json:"value"`
	Status    int    `json:"status"`
}

// GetMWindows Get https://uptimerobot.com/#getMWindows
func (ms *MWindowService) GetMWindows(params GetMWindowParams) (*MWindowsResp, error) {
	obj := &MWindowsResp{}

	err := ms.apiClient.request(http.MethodPost, "getMWindows", params, &obj)

	return obj, err
}

// NewMWindow Get https://uptimerobot.com/#newMWindow
func (ms *MWindowService) NewMWindow(params NewMWindowParams) (*MWindowsResp, error) {
	obj := &MWindowsResp{}

	err := ms.apiClient.request(http.MethodPost, "newMWindow", params, &obj)

	return obj, err
}

// EditMWindow Get https://uptimerobot.com/#editMWindow
func (ms *MWindowService) EditMWindow(params EditMWindowParams) (*MWindowResp, error) {
	obj := &MWindowResp{}

	err := ms.apiClient.request(http.MethodPost, "editMWindow", params, &obj)

	return obj, err
}

// DeleteMWindow Get https://uptimerobot.com/#deleteMwindowWrap
func (ms *MWindowService) DeleteMWindow(id int) (*MWindowResp, error) {
	obj := &MWindowResp{}

	params := DeleteMWindowWrapper{Id: id}

	err := ms.apiClient.request(http.MethodPost, "deleteMWindow", params, &obj)

	return obj, err
}
