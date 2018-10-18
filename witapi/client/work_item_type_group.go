// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "wit": work_item_type_group Resource Client
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

// ShowWorkItemTypeGroupPath computes a request path to the show action of work_item_type_group.
func ShowWorkItemTypeGroupPath(groupID uuid.UUID) string {
	param0 := groupID.String()

	return fmt.Sprintf("/api/workitemtypegroups/%s", param0)
}

// Show work item type group for given ID
func (c *Client) ShowWorkItemTypeGroup(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowWorkItemTypeGroupRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowWorkItemTypeGroupRequest create the request corresponding to the show action endpoint of the work_item_type_group resource.
func (c *Client) NewShowWorkItemTypeGroupRequest(ctx context.Context, path string) (*http.Request, error) {
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