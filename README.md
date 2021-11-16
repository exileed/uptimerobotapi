# uptimerobotapi [![Build Status](https://github.com/exileed/uptimerobotapi/workflows/test/badge.svg)](https://github.com/exileed/uptimerobotapi/actions) [![Go Reference](https://pkg.go.dev/badge/github.com/exileed/uptimerobotapi.svg)](https://pkg.go.dev/github.com/exileed/uptimerobotapi) [![Go Report Card](https://goreportcard.com/badge/github.com/exileed/uptimerobotapi)](https://goreportcard.com/report/github.com/exileed/uptimerobotapi)


A Go client for [UptimeRobot API](https://uptimerobot.com/api/).

## Usage

See the [full API reference on Go.dev](https://pkg.go.dev/github.com/exileed/uptimerobotapi).

### Client initialization

All API requests are made through [`uptimerobotapi.Client`](https://pkg.go.dev/github.com/exileed/uptimerobotapi#Client). Make sure to include an API token:

``` go
package main
import (
	"os"
	
	"github.com/exileed/uptimerobotapi"
)
func main() {
	client := uptimerobotapi.NewClient(os.Getenv("UPTIMEROBOT_API_TOKEN"))
	...
}
```

### Making API requests

Use an initialized client to make API requests:

``` go
package main
import (
	"os"
	
	"github.com/exileed/uptimerobotapi"
)
func main() {
	client := uptimerobotapi.NewClient(os.Getenv("UPTIMEROBOT_API_TOKEN"))
	
	accounts, err := client.Account.GetAccountDetails()
	if err != nil {
		panic(err)
	}
	...
}
```


#### API Error Responses

For cases where your request results in an error from the API, you can use the
`errors.As()` function from the standard library to extract the
`uptimerobotapi.APIError` error value and inspect more details about the error,
including the HTTP response code and UptimeRobot API Error Code.

```go
package main

import (
	"fmt"
	
	"github.com/exileed/uptimerobotapi"
)


func main() {
	client := uptimerobotapi.NewClient("")
	account, err := client.Account.GetAccountDetails()
	var apiErr uptimerobotapi.APIError
		
	
		if errors.As(err, &apiErr){
			if apiErr.RateLimited() {
				fmt.Println("rate limited")
				return
			}

			fmt.Println("unknown status code:", apiErr.StatusCode)
		}

		panic(err)
	}
	fmt.Println(account)
}
