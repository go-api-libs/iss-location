// This file provides a client with methods as well as functions to interact with the HTTP API.
//
// Code generated by github.com/MarkRosemaker DO NOT EDIT.

package isslocation

import (
	"context"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-api-libs/api"
	"github.com/go-json-experiment/json"
)

const (
	userAgent = "IssLocationGoAPILibrary/1.0.0 (https://github.com/go-api-libs/iss-location)"
)

var (
	baseURL = &url.URL{
		Host:   "api.open-notify.org",
		Path:   "/",
		Scheme: "http",
	}

	jsonOpts = json.JoinOptions(
		json.RejectUnknownMembers(true))
)

// Client conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The HTTP client to use for requests.
	cli *http.Client
}

// NewClient creates a new Client.
func NewClient() (*Client, error) {
	return &Client{cli: http.DefaultClient}, nil
}

// GetIssNowJSON defines an operation.
//
//	GET /iss-now.json
func (c *Client) GetIssNowJSON(ctx context.Context) (*GetIssNowJSONOkJSONResponse, error) {
	return GetIssNowJSON[GetIssNowJSONOkJSONResponse](ctx, c)
}

// GetIssNowJSON defines an operation.
// You can define a custom result to unmarshal the response into.
//
//	GET /iss-now.json
func GetIssNowJSON[R any](ctx context.Context, c *Client) (*R, error) {
	u := baseURL.JoinPath("/iss-now.json")
	req := (&http.Request{
		Header:     http.Header{"User-Agent": []string{userAgent}},
		Host:       u.Host,
		Method:     http.MethodGet,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		URL:        u,
	}).WithContext(ctx)

	rsp, err := c.cli.Do(req)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	switch rsp.StatusCode {
	case http.StatusOK:
		// TODO
		switch mt, _, _ := strings.Cut(rsp.Header.Get("Content-Type"), ";"); mt {
		case "application/json":
			var out R
			if err := json.UnmarshalRead(rsp.Body, &out, jsonOpts); err != nil {
				return nil, api.WrapDecodingError(rsp, err)
			}

			return &out, nil
		default:
			return nil, api.NewErrUnknownContentType(rsp)
		}
	default:
		return nil, api.NewErrUnknownStatusCode(rsp)
	}
}
