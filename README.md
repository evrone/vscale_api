# Vscale API [WIP]

## Installation and documentation

To install `Vscale API`, simply run:

```
$ go get github.com/openit-lib/vscale_api
```

## Getting Started

``` go
package main
import(
  "fmt"
  vscale "github.com/openit-lib/vscale_api"
)

func GetAccountInfo() {
  client := api.New("API_SECRET_TOKEN")
  account, _, err := client.Account.Get()
  if err != nil {
    panic(err)
  }
  fmt.Printf("Account info: %v", account)
}
```

## Contribution Guidelines

01. Fork
02. Change
03. PR

*WIP*
