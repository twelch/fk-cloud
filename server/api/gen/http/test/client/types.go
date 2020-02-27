// Code generated by goa v3.0.10, DO NOT EDIT.
//
// test HTTP client types
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package client

import (
	test "github.com/fieldkit/cloud/server/api/gen/test"
)

// GetUnauthorizedResponseBody is the type of the "test" service "get" endpoint
// HTTP response body for the "unauthorized" error.
type GetUnauthorizedResponseBody string

// ErrorUnauthorizedResponseBody is the type of the "test" service "error"
// endpoint HTTP response body for the "unauthorized" error.
type ErrorUnauthorizedResponseBody string

// EmailUnauthorizedResponseBody is the type of the "test" service "email"
// endpoint HTTP response body for the "unauthorized" error.
type EmailUnauthorizedResponseBody string

// NewGetUnauthorized builds a test service get endpoint unauthorized error.
func NewGetUnauthorized(body GetUnauthorizedResponseBody) test.Unauthorized {
	v := test.Unauthorized(body)
	return v
}

// NewErrorUnauthorized builds a test service error endpoint unauthorized error.
func NewErrorUnauthorized(body ErrorUnauthorizedResponseBody) test.Unauthorized {
	v := test.Unauthorized(body)
	return v
}

// NewEmailUnauthorized builds a test service email endpoint unauthorized error.
func NewEmailUnauthorized(body EmailUnauthorizedResponseBody) test.Unauthorized {
	v := test.Unauthorized(body)
	return v
}
