# 🛰️ ISS Current Location
[![Go Reference](https://pkg.go.dev/badge/github.com/go-api-libs/iss-location.svg)](https://pkg.go.dev/github.com/go-api-libs/iss-location/pkg/isslocation)
[![Official Documentation](https://img.shields.io/badge/docs-API-blue)](http://open-notify.org/Open-Notify-API/ISS-Location-Now/)
[![OpenAPI](https://img.shields.io/badge/OpenAPI-3.1-blue)](/api/openapi.json)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-api-libs/iss-location)](https://goreportcard.com/report/github.com/go-api-libs/iss-location)
![Code Coverage](https://img.shields.io/badge/coverage-100%25-brightgreen)
![API Health](https://img.shields.io/badge/API_health-85%25-green)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](./LICENSE)

Provides the current location of the International Space Station including latitude, longitude, and timestamp. Offers both JSON and JSONP formats for easy access. Users are advised to poll the API no more than once every 5 seconds. ([Source](https://freepublicapis.com/iss-current-location))

## Installation

To install the library, use the following command:

```shell
go get github.com/go-api-libs/iss-location/pkg/isslocation
```

## Usage

### Example: 

```go
package main

import (
	"context"

	"github.com/go-api-libs/iss-location/pkg/isslocation"
)

func main() {
	c, err := isslocation.NewClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	issLocation, err := c.GetIssLocation(ctx)
	if err != nil {
		panic(err)
	}

	// Use issLocation object
}

```

## Additional Information

- [**Go Reference**](https://pkg.go.dev/github.com/go-api-libs/iss-location/pkg/isslocation): The Go reference documentation for the client package.
- [**Official Documentation**](http://open-notify.org/Open-Notify-API/ISS-Location-Now/): The official API documentation.
- [**OpenAPI Specification**](./api/openapi.json): The OpenAPI 3.1.0 specification.
- [**Go Report Card**](https://goreportcard.com/report/github.com/go-api-libs/iss-location): Check the code quality report.

## Contributing

If you have any contributions to make, please submit a pull request or open an issue on the [GitHub repository](https://github.com/go-api-libs/iss-location).

## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.
