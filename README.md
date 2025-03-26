# vaultwarden-go

This repository is a RESTful API client for the Bitwarden CLI own RESTful API Server.

## Requirements

Before using, please ensure you have hosted your own RESTful API Server using the bw CLI.

```sh
bw serve --port <number> --hostname <hostname>
```

## Getting Started

```sh
go get github.com/Menfre01/vaultwarden-go
```

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/Menfre01/vaultwarden-go/vaultwarden"
)

func main() {
	cli := vaultwarden.NewClient("http://localhost:8087", &http.Client{})
	orgResp, err := cli.ListOrganizations("your org name")
	if err != nil {
		panic(err)
	}
	for _, org := range orgResp.Data.Data {
		fmt.Println(org)
	}
}
```


