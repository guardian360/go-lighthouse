# go-lighthouse

[![Go Reference](https://pkg.go.dev/badge/github.com/guardian360/go-lighthouse.svg)](https://pkg.go.dev/github.com/guardian360/go-lighthouse)

---

## ⚠️🚧 Work in Progress 🚧⚠️

**This project is still under active development.**
APIs and functionality may change, and certain features may not yet be fully
implemented.

❗ **Not recommended for production use.**

---

A Go client library for interacting with the Lighthouse API.

## 📖 Overview

The `go-lighthouse` library provides a Go client for interacting with
Lighthouse API endpoints. It supports OAuth authentication, request handling,
and API versioning.

## 🚀 Installation

```sh
go get github.com/guardian360/go-lighthouse
```

## 🎯 Usage

### Creating a Client

To interact with the Lighthouse API, you need to create a client using the
`client.New` function. You can configure the client with the base URL. Next,
you can create an API instance using the `api.New` function.

```go
package main

import (
    "fmt"
    "os"

    api "github.com/guardian360/go-lighthouse/api/v2"
    "github.com/guardian360/go-lighthouse/client"
)

func main() {
    cfg := client.Config{
        BaseURL: "https://lighthouse.guardian360.nl",
    }
    clt := client.New(cfg)
    lighthouse := api.New(clt)

    response, err := lighthouse.Health().Get()
    if err != nil {
        fmt.Println("Error fetching health:", err)
        os.Exit(1)
    }

    fmt.Printf("Lighthouse API response: %+v\n", response)
}
```

### Authenticating with OAuth

Most Lighthouse API endpoints require OAuth authentication. The `go-lighthouse`
library supports OAuth authentication using the client credentials grant. You
can create an OAuth client using the [Lighthouse
portal](https://lighthouse.guardian360.nl) to obtain a client ID and client
secret.

```go
package main

import (
    "fmt"
    "os"

    api "github.com/guardian360/go-lighthouse/api/v2"
    "github.com/guardian360/go-lighthouse/client"
)

func main() {
    baseURL := "https://lighthouse.guardian360.nl"
    clt := client.New(baseURL).WithClientCredentials(
        baseURL+"/oauth/token",
        "client_id",
        "client_secret",
    )
    lighthouse := api.New(clt)

    resp, err := lighthouse.Probes().Get()
    if err != nil {
        fmt.Println("Error fetching probes:", err)
        os.Exit(1)
    }

    fmt.Printf("Lighthouse API response: %+v\n", resp)
}
```

## 🛠 Contributing

Contributions are welcome! For feature requests, bug reports, or questions,
please open an issue. For code contributions, please open a pull request.

## 📚 Documentation

For more information, please refer to the [Go reference
documentation](https://pkg.go.dev/github.com/guardian360/go-lighthouse).

## 📜 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE)
file for details.
