// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "wit": workitems Resource Client
//
// Command:
// $ goagen
// --design=github.com/fabric8-services/fabric8-wit/design
// --out=$(GOPATH)/src/github.com/fabric8-services/fabric8-wit
// --version=v1.3.0

package client

import (
	"bytes"
	"context"
	"fmt"
	uuid "github.com/goadesign/goa/uuid"
	"net/http"
	"net/url"
	"strconv"
)

// CreateWorkitemsPayload is the workitems create action payload.
type CreateWorkitemsPayload struct {
	Data *WorkItem `form:"data" json:"data" xml:"data"`
	// An array of mixed types
	Included []interface{}  `form:"included,omitempty" json:"included,omitempty" xml:"included,omitempty"`
	Links    *WorkItemLinks `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
}

// CreateWorkitemsPath computes a request path to the create action of workitems.
func CreateWorkitemsPath(spaceID uuid.UUID) string {
	param0 := spaceID.String()

	return fmt.Sprintf("/api/spaces/%s/workitems", param0)
}

// create work item with type and id.
func (c *Client) CreateWorkitems(ctx context.Context, path string, payload *CreateWorkitemsPayload) (*http.Response, error) {
	req, err := c.NewCreateWorkitemsRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateWorkitemsRequest create the request corresponding to the create action endpoint of the workitems resource.
func (c *Client) NewCreateWorkitemsRequest(ctx context.Context, path string, payload *CreateWorkitemsPayload) (*http.Request, error) {
	var body bytes.Buffer
	err := c.Encoder.Encode(payload, &body, "*/*")
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	if c.JWTSigner != nil {
		c.JWTSigner.Sign(req)
	}
	return req, nil
}

// ListWorkitemsPath computes a request path to the list action of workitems.
func ListWorkitemsPath(spaceID uuid.UUID) string {
	param0 := spaceID.String()

	return fmt.Sprintf("/api/spaces/%s/workitems", param0)
}

// List work items.
func (c *Client) ListWorkitems(ctx context.Context, path string, filter *string, filterArea *string, filterAssignee *string, filterExpression *string, filterIteration *string, filterParentexists *bool, filterWorkitemstate *string, filterWorkitemtype *uuid.UUID, pageLimit *int, pageOffset *string, sort *string, ifModifiedSince *string, ifNoneMatch *string) (*http.Response, error) {
	req, err := c.NewListWorkitemsRequest(ctx, path, filter, filterArea, filterAssignee, filterExpression, filterIteration, filterParentexists, filterWorkitemstate, filterWorkitemtype, pageLimit, pageOffset, sort, ifModifiedSince, ifNoneMatch)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListWorkitemsRequest create the request corresponding to the list action endpoint of the workitems resource.
func (c *Client) NewListWorkitemsRequest(ctx context.Context, path string, filter *string, filterArea *string, filterAssignee *string, filterExpression *string, filterIteration *string, filterParentexists *bool, filterWorkitemstate *string, filterWorkitemtype *uuid.UUID, pageLimit *int, pageOffset *string, sort *string, ifModifiedSince *string, ifNoneMatch *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if filter != nil {
		values.Set("filter", *filter)
	}
	if filterArea != nil {
		values.Set("filter[area]", *filterArea)
	}
	if filterAssignee != nil {
		values.Set("filter[assignee]", *filterAssignee)
	}
	if filterExpression != nil {
		values.Set("filter[expression]", *filterExpression)
	}
	if filterIteration != nil {
		values.Set("filter[iteration]", *filterIteration)
	}
	if filterParentexists != nil {
		tmp148 := strconv.FormatBool(*filterParentexists)
		values.Set("filter[parentexists]", tmp148)
	}
	if filterWorkitemstate != nil {
		values.Set("filter[workitemstate]", *filterWorkitemstate)
	}
	if filterWorkitemtype != nil {
		tmp149 := filterWorkitemtype.String()
		values.Set("filter[workitemtype]", tmp149)
	}
	if pageLimit != nil {
		tmp150 := strconv.Itoa(*pageLimit)
		values.Set("page[limit]", tmp150)
	}
	if pageOffset != nil {
		values.Set("page[offset]", *pageOffset)
	}
	if sort != nil {
		values.Set("sort", *sort)
	}
	u.RawQuery = values.Encode()
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

// ReorderWorkitemsPayload is the workitems reorder action payload.
type ReorderWorkitemsPayload struct {
	Data     []*WorkItem              `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
	Position *WorkItemReorderPosition `form:"position,omitempty" json:"position,omitempty" xml:"position,omitempty"`
}

// ReorderWorkitemsPath computes a request path to the reorder action of workitems.
func ReorderWorkitemsPath(spaceID uuid.UUID) string {
	param0 := spaceID.String()

	return fmt.Sprintf("/api/spaces/%s/workitems/reorder", param0)
}

// reorder the work items
func (c *Client) ReorderWorkitems(ctx context.Context, path string, payload *ReorderWorkitemsPayload) (*http.Response, error) {
	req, err := c.NewReorderWorkitemsRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewReorderWorkitemsRequest create the request corresponding to the reorder action endpoint of the workitems resource.
func (c *Client) NewReorderWorkitemsRequest(ctx context.Context, path string, payload *ReorderWorkitemsPayload) (*http.Request, error) {
	var body bytes.Buffer
	err := c.Encoder.Encode(payload, &body, "*/*")
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("PATCH", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	if c.JWTSigner != nil {
		c.JWTSigner.Sign(req)
	}
	return req, nil
}