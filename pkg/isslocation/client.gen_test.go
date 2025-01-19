// This file provides tests for the client package.
//
// Code generated by github.com/MarkRosemaker DO NOT EDIT.

package isslocation_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"maps"
	"net/http"
	"net/url"
	"path"
	"slices"
	"strings"
	"testing"

	"github.com/go-api-libs/api"
	"github.com/go-api-libs/iss-location/pkg/isslocation"
	"gopkg.in/dnaeon/go-vcr.v3/cassette"
	"gopkg.in/dnaeon/go-vcr.v3/recorder"
)

type testRoundTripper struct {
	rsp *http.Response
	err error
}

func (t *testRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return t.rsp, t.err
}

func TestClient_Error(t *testing.T) {
	ctx := context.Background()

	c, err := isslocation.NewClient()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Do", func(t *testing.T) {
		testErr := errors.New("test error")
		http.DefaultClient.Transport = &testRoundTripper{err: testErr}

		if _, err := c.GetIssLocation(ctx); err == nil {
			t.Fatal("expected error")
		} else if !errors.Is(err, testErr) {
			t.Fatalf("want: %v, got: %v", testErr, err)
		}

		if _, err := c.GetPeopleInSpace(ctx); err == nil {
			t.Fatal("expected error")
		} else if !errors.Is(err, testErr) {
			t.Fatalf("want: %v, got: %v", testErr, err)
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		errDecode := &api.DecodingError{}

		t.Run("GetIssLocation", func(t *testing.T) {
			// unknown status code
			http.DefaultClient.Transport = &testRoundTripper{rsp: &http.Response{StatusCode: http.StatusTeapot}}

			if _, err := c.GetIssLocation(ctx); err == nil {
				t.Fatal("expected error")
			} else if !errors.Is(err, api.ErrUnknownStatusCode) {
				t.Fatalf("want: %v, got: %v", api.ErrUnknownStatusCode, err)
			}

			// unknown content type for 200 OK
			http.DefaultClient.Transport = &testRoundTripper{rsp: &http.Response{
				Header:     http.Header{"Content-Type": []string{"foo"}},
				StatusCode: http.StatusOK,
			}}

			if _, err := c.GetIssLocation(ctx); err == nil {
				t.Fatal("expected error")
			} else if !errors.Is(err, api.ErrUnknownContentType) {
				t.Fatalf("want: %v, got: %v", api.ErrUnknownContentType, err)
			}

			// decoding error for known content type "application/json"
			http.DefaultClient.Transport = &testRoundTripper{rsp: &http.Response{
				Body:       io.NopCloser(strings.NewReader("{")),
				Header:     http.Header{"Content-Type": []string{"application/json"}},
				StatusCode: http.StatusOK,
			}}

			if _, err := c.GetIssLocation(ctx); err == nil {
				t.Fatal("expected error")
			} else if !errors.As(err, &errDecode) {
				t.Fatalf("want: %v, got: %v", errDecode, err)
			}
		})

		t.Run("GetPeopleInSpace", func(t *testing.T) {
			// unknown status code
			http.DefaultClient.Transport = &testRoundTripper{rsp: &http.Response{StatusCode: http.StatusTeapot}}

			if _, err := c.GetPeopleInSpace(ctx); err == nil {
				t.Fatal("expected error")
			} else if !errors.Is(err, api.ErrUnknownStatusCode) {
				t.Fatalf("want: %v, got: %v", api.ErrUnknownStatusCode, err)
			}

			// unknown content type for 200 OK
			http.DefaultClient.Transport = &testRoundTripper{rsp: &http.Response{
				Header:     http.Header{"Content-Type": []string{"foo"}},
				StatusCode: http.StatusOK,
			}}

			if _, err := c.GetPeopleInSpace(ctx); err == nil {
				t.Fatal("expected error")
			} else if !errors.Is(err, api.ErrUnknownContentType) {
				t.Fatalf("want: %v, got: %v", api.ErrUnknownContentType, err)
			}

			// decoding error for known content type "application/json"
			http.DefaultClient.Transport = &testRoundTripper{rsp: &http.Response{
				Body:       io.NopCloser(strings.NewReader("{")),
				Header:     http.Header{"Content-Type": []string{"application/json"}},
				StatusCode: http.StatusOK,
			}}

			if _, err := c.GetPeopleInSpace(ctx); err == nil {
				t.Fatal("expected error")
			} else if !errors.As(err, &errDecode) {
				t.Fatalf("want: %v, got: %v", errDecode, err)
			}
		})
	})
}

func replay(t *testing.T, cassette string) {
	t.Helper()

	r, err := recorder.NewWithOptions(&recorder.Options{
		CassetteName:       cassette,
		Mode:               recorder.ModeReplayOnly,
		RealTransport:      http.DefaultTransport,
		SkipRequestLatency: true,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		_ = r.Stop()
	})

	r.SetMatcher(matcher)
	http.DefaultClient.Transport = r
}

func matcher(r *http.Request, i cassette.Request) bool {
	u, err := url.Parse(i.URL)
	if err != nil {
		panic(err)
	}

	return r.Method == i.Method &&
		r.URL.Scheme == u.Scheme &&
		r.URL.Opaque == u.Opaque &&
		r.URL.Host == u.Host &&
		path.Clean("/"+r.URL.Path) == path.Clean("/"+u.Path) &&
		r.URL.Fragment == u.Fragment &&
		maps.EqualFunc(r.URL.Query(), u.Query(), slices.Equal) &&
		getBody(r) == i.Body
}

func getBody(r *http.Request) string {
	if r.Body == nil {
		return ""
	}

	if r.GetBody == nil {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		r.Body.Close()
		r.Body = io.NopCloser(bytes.NewReader(b))
		return string(b)
	}

	body, err := r.GetBody()
	if err != nil {
		panic(err)
	}
	defer body.Close()

	b, err := io.ReadAll(body)
	if err != nil {
		panic(err)
	}

	return string(b)
}

func TestClient_VCR(t *testing.T) {
	ctx := context.Background()

	c, err := isslocation.NewClient()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("2024-12-22", func(t *testing.T) {
		replay(t, "vcr/2024-12-22")

		{
			res, err := c.GetIssLocation(ctx)
			if err != nil {
				t.Fatal(err)
			} else if res == nil {
				t.Fatal("result is nil")
			}
		}

		{
			res, err := c.GetPeopleInSpace(ctx)
			if err != nil {
				t.Fatal(err)
			} else if res == nil {
				t.Fatal("result is nil")
			}
		}
	})
}
