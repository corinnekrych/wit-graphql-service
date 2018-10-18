// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "wit": comments Resource Client
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
)

// DeleteCommentsPath computes a request path to the delete action of comments.
func DeleteCommentsPath(commentID uuid.UUID) string {
	param0 := commentID.String()

	return fmt.Sprintf("/api/comments/%s", param0)
}

// Delete work item with given id.
func (c *Client) DeleteComments(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeleteCommentsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteCommentsRequest create the request corresponding to the delete action endpoint of the comments resource.
func (c *Client) NewDeleteCommentsRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		c.JWTSigner.Sign(req)
	}
	return req, nil
}

// ShowCommentsPath computes a request path to the show action of comments.
func ShowCommentsPath(commentID uuid.UUID) string {
	param0 := commentID.String()

	return fmt.Sprintf("/api/comments/%s", param0)
}

// Retrieve comment with given commentId.
func (c *Client) ShowComments(ctx context.Context, path string, ifModifiedSince *string, ifNoneMatch *string) (*http.Response, error) {
	req, err := c.NewShowCommentsRequest(ctx, path, ifModifiedSince, ifNoneMatch)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowCommentsRequest create the request corresponding to the show action endpoint of the comments resource.
func (c *Client) NewShowCommentsRequest(ctx context.Context, path string, ifModifiedSince *string, ifNoneMatch *string) (*http.Request, error) {
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

// UpdateCommentsPayload is the comments update action payload.
type UpdateCommentsPayload struct {
	Data *Comment `form:"data" json:"data" xml:"data"`
	// An array of mixed types
	Included []interface{} `form:"included,omitempty" json:"included,omitempty" xml:"included,omitempty"`
}

// UpdateCommentsPath computes a request path to the update action of comments.
func UpdateCommentsPath(commentID uuid.UUID) string {
	param0 := commentID.String()

	return fmt.Sprintf("/api/comments/%s", param0)
}

// update the comment with the given commentId.
func (c *Client) UpdateComments(ctx context.Context, path string, payload *UpdateCommentsPayload) (*http.Response, error) {
	req, err := c.NewUpdateCommentsRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateCommentsRequest create the request corresponding to the update action endpoint of the comments resource.
func (c *Client) NewUpdateCommentsRequest(ctx context.Context, path string, payload *UpdateCommentsPayload) (*http.Request, error) {
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
