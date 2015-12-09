# Vscale API [WIP]

[![Vexor status](https://ci.vexor.io/projects/0fccbce3-edc7-49f4-9677-6996e94e3fa5/status.svg)](https://ci.vexor.io/ui/projects/0fccbce3-edc7-49f4-9677-6996e94e3fa5/builds)

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
  client := vscale.New("API_SECRET_TOKEN")
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
