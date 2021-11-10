package uptimerobotapi

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (

	// defaultBaseURL is the base URL of the UptimeRobot API.
	defaultBaseURL = "https://api.uptimerobot.com/"

	// apiVersion is the revision of the UptimeRobot API.
	apiVersion = "v2"

	// apiVersion default user agent of the  UptimeRobot API.
	defaultUserAgent = "go-uptimerobotapi" //@todo

	headerRateLimit = "X-RateLimit-Limit" //@todo
	headerRateReset = "X-RateLimit-Reset" //@todo
)

// ClientConfig specifies configuration with which to initialize a UptimeRobot API
// client.
type ClientConfig struct {
	// APIToken is the UptimeRobot API token to use for authentication.
	APIToken string

	// HTTPClient is your own HTTP client. The library will otherwise use a
	// parameter-less `&http.Client{}`, resulting in default everything.
	HTTPClient *http.Client

	// userAgent User agent used when communicating with the UptimeRobot API.
	userAgent *string
}

// A Client manages communication with the UptimeRobot API.
type Client struct {

	// HTTP client used to communicate with the API.
	httpClient *http.Client
	// Base URL for API requests.
	baseUrl *url.URL

	apiVersion string
	// User agent used when communicating with the Uptime Robot API.
	userAgent string

	// Token used to make authenticated API calls.
	Token string

	// Services used for talking to different parts of the UptimeRobot API.
	Account      AccountService
	AlertContact AlertContactService
	Monitor      MonitorService
	MWindow      MWindowService
}

// NewClient returns a new UptimeRobot API client.
func NewClient(token string) *Client {
	client := newClient(nil)

	client.Token = token
	return client
}

// NewClientWithConfig returns a new UptimeRobot API client.
func NewClientWithConfig(config *ClientConfig) *Client {
	client := newClient(config)
	return client
}

// newClient returns a new UptimeRobot API client.
func newClient(config *ClientConfig) *Client {

	if config == nil {
		config = &ClientConfig{}
	}

	var httpClient *http.Client

	if config.HTTPClient == nil {
		httpClient = &http.Client{}
	} else {
		httpClient = config.HTTPClient
	}

	apiUrl, err := url.Parse(defaultBaseURL)
	if err != nil {
		panic(err)
	}

	c := &Client{
		httpClient: httpClient,
		baseUrl:    apiUrl,
		apiVersion: apiVersion,
		Token:      config.APIToken,
		userAgent:  defaultUserAgent,
	}

	c.Account = AccountService{apiClient: c}
	c.AlertContact = AlertContactService{apiClient: c}
	c.Monitor = MonitorService{apiClient: c}
	c.MWindow = MWindowService{apiClient: c}
	return c
}

func (c *Client) request(method string, urlStr string, opt interface{}, responseModel interface{}) error {
	u, err := c.baseUrl.Parse(fmt.Sprintf("%s/%s", c.apiVersion, urlStr))
	if err != nil {
		return err
	}

	q, err := query.Values(opt)
	if err != nil {
		return err
	}

	q.Set("api_key", c.Token)
	q.Set("format", "json")

	req, err := http.NewRequest(method, u.String(), strings.NewReader(q.Encode()))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return err
	}

	var apiResp map[string]interface{}

	_, err = c.decodeAPIResponse(res.Body, &apiResp)
	if err != nil {
		return err
	}

	if apiResp["stat"] != StatOk {
		var apiErrResp ErrorResponse

		err := createFromMap(apiResp, &apiErrResp)

		if err != nil {
			return err
		}

		_, err = c.decodeAPIResponse(res.Body, &apiErrResp)
		if err != nil {
			return err
		}

		var apiErrResp2 ErrorResponse

		return &apiErrResp2.Error
	}

	err = createFromMap(apiResp, &responseModel)

	return err
}

// decodeAPIResponse decode response
func (c *Client) decodeAPIResponse(responseBody io.Reader, resp interface{}) (_ io.Reader, err error) {
	data, err := ioutil.ReadAll(responseBody)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, resp)
	if err != nil {
		return
	}

	return responseBody, nil
}

func createFromMap(m map[string]interface{}, result interface{}) error {
	data, _ := json.Marshal(m)
	err := json.Unmarshal(data, &result)
	return err
}
