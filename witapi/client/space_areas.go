// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "wit": space_areas Resource Client
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
	uuid "github.com/goadesign/goa/uuid"
	"net/http"
	"net/url"
)

// ListSpaceAreasPath computes a request path to the list action of space_areas.
func ListSpaceAreasPath(spaceID uuid.UUID) string {
	param0 := spaceID.String()

	return fmt.Sprintf("/api/spaces/%s/areas", param0)
}

// List Areas.
func (c *Client) ListSpaceAreas(ctx context.Context, path string, ifModifiedSince *string, ifNoneMatch *string) (*http.Response, error) {
	req, err := c.NewListSpaceAreasRequest(ctx, path, ifModifiedSince, ifNoneMatch)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListSpaceAreasRequest create the request corresponding to the list action endpoint of the space_areas resource.
func (c *Client) NewListSpaceAreasRequest(ctx context.Context, path string, ifModifiedSince *string, ifNoneMatch *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if ifModifiedSince != nil {

		header.Set("If-Modified-Since", *ifModifiedSince)
	}
	if ifNoneMatch != nil {

		header.Set("If-None-Match", *ifNoneMatch)
	}
	return req, nil
}
