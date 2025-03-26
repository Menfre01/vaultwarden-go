# vaultwarden-go

这个仓库目是一个 RESTful API to the Bitwarden CLI 的 Client。

## Requirements

使用前请先确保你已经使用 bw CLI host your own RESTful API Server.

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


