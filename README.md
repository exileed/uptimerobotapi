# uptimerobotapi [![Build Status](https://github.com/exileed/uptimerobotapi/workflows/test/badge.svg)](https://github.com/exileed/uptimerobotapi/actions) [![Go Reference](https://pkg.go.dev/badge/github.com/exileed/uptimerobotapi.svg)](https://pkg.go.dev/github.com/exileed/uptimerobotapi)

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
