# go-lighthouse

[![Go Reference](https://pkg.go.dev/badge/github.com/guardian360/go-lighthouse.svg)](https://pkg.go.dev/github.com/guardian360/go-lighthouse)

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
`client.New` function. You can configure the client with the base URL.

```go
package main

import (
    "fmt"
    "os"

    "github.com/guardian360/go-lighthouse/api"
    "github.com/guardian360/go-lighthouse/client"
)

func main() {
    config := client.Config{
        BaseURL:  "https://lighthouse.guardian360.nl",
    }
    lighthouse := client.New(config)

    response, err := api.V1(lighthouse).Heartbeat().Get()
    if err != nil {
        fmt.Println("Error fetching heartbeat:", err)
        os.Exit(1)
    }

    fmt.Println("Lighthouse API heartbeat response:", response)
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

    "github.com/guardian360/go-lighthouse/api"
    "github.com/guardian360/go-lighthouse/client"
)

func main() {
    config := client.Config{
        BaseURL:  "https://lighthouse.guardian360.nl",
        Insecure: false,
    }
    lighthouse := client.New(config).WithOAuth(client.ClientCredentialsGrant{
        TokenURL: "https://lighthouse.guardian360.nl/oauth/token",
        ClientID: "client_id",
        ClientSecret: "client_secret",
    })

    response, err := api.V1(lighthouse).Probes().Get()
    if err != nil {
        fmt.Println("Error fetching probes:", err)
        os.Exit(1)
    }

    fmt.Println("Lighthouse API probes response:", response)
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
