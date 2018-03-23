// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// archiver HTTP server types
//
// Command:
// $ goa gen goa.design/plugins/goakit/examples/fetcher/archiver/design

package server

import (
	goa "goa.design/goa"
	archiversvc "goa.design/plugins/goakit/examples/fetcher/archiver/gen/archiver"
)

// ArchiveRequestBody is the type of the "archiver" service "archive" endpoint
// HTTP request body.
type ArchiveRequestBody struct {
	// HTTP status
	Status *int `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// HTTP response body content
	Body *string `form:"body,omitempty" json:"body,omitempty" xml:"body,omitempty"`
}

// ArchiveResponseBody is the type of the "archiver" service "archive" endpoint
// HTTP response body.
type ArchiveResponseBody struct {
	// The archive resouce href
	Href string `form:"href" json:"href" xml:"href"`
	// HTTP status
	Status int `form:"status" json:"status" xml:"status"`
	// HTTP response body content
	Body string `form:"body" json:"body" xml:"body"`
}

// ReadResponseBody is the type of the "archiver" service "read" endpoint HTTP
// response body.
type ReadResponseBody struct {
	// The archive resouce href
	Href string `form:"href" json:"href" xml:"href"`
	// HTTP status
	Status int `form:"status" json:"status" xml:"status"`
	// HTTP response body content
	Body string `form:"body" json:"body" xml:"body"`
}

// ReadNotFoundResponseBody is the type of the "archiver" service "read"
// endpoint HTTP response body for the "not_found" error.
type ReadNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
}

// ReadBadRequestResponseBody is the type of the "archiver" service "read"
// endpoint HTTP response body for the "bad_request" error.
type ReadBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
}

// NewArchiveResponseBody builds the HTTP response body from the result of the
// "archive" endpoint of the "archiver" service.
func NewArchiveResponseBody(res *archiversvc.ArchiveMedia) *ArchiveResponseBody {
	body := &ArchiveResponseBody{
		Href:   res.Href,
		Status: res.Status,
		Body:   res.Body,
	}
	return body
}

// NewReadResponseBody builds the HTTP response body from the result of the
// "read" endpoint of the "archiver" service.
func NewReadResponseBody(res *archiversvc.ArchiveMedia) *ReadResponseBody {
	body := &ReadResponseBody{
		Href:   res.Href,
		Status: res.Status,
		Body:   res.Body,
	}
	return body
}

// NewReadNotFoundResponseBody builds the HTTP response body from the result of
// the "read" endpoint of the "archiver" service.
func NewReadNotFoundResponseBody(res *archiversvc.Error) *ReadNotFoundResponseBody {
	body := &ReadNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
	}
	return body
}

// NewReadBadRequestResponseBody builds the HTTP response body from the result
// of the "read" endpoint of the "archiver" service.
func NewReadBadRequestResponseBody(res *archiversvc.Error) *ReadBadRequestResponseBody {
	body := &ReadBadRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
	}
	return body
}

// NewArchiveArchivePayload builds a archiver service archive endpoint payload.
func NewArchiveArchivePayload(body *ArchiveRequestBody) *archiversvc.ArchivePayload {
	v := &archiversvc.ArchivePayload{
		Status: *body.Status,
		Body:   *body.Body,
	}
	return v
}

// NewReadReadPayload builds a archiver service read endpoint payload.
func NewReadReadPayload(id int) *archiversvc.ReadPayload {
	return &archiversvc.ReadPayload{
		ID: id,
	}
}

// Validate runs the validations defined on ArchiveRequestBody
func (body *ArchiveRequestBody) Validate() (err error) {
	if body.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "body"))
	}
	if body.Body == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("body", "body"))
	}
	if body.Status != nil {
		if *body.Status < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.status", *body.Status, 0, true))
		}
	}
	return
}
