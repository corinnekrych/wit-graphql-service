// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "wit": filter Resource Client
//
// Command:
// $ goagen
// --design=github.com/fabric8-services/fabric8-wit/design
// --out=$(GOPATH)/src/github.com/fabric8-services/fabric8-wit
// --version=v1.3.0

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// ListFilterPath computes a request path to the list action of filter.
func ListFilterPath() string {

	return fmt.Sprintf("/api/filters")
}

// List filters.
func (c *Client) ListFilter(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListFilterRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListFilterRequest create the request corresponding to the list action endpoint of the filter resource.
func (c *Client) NewListFilterRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
