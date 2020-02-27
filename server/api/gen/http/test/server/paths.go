// Code generated by goa v3.0.10, DO NOT EDIT.
//
// HTTP request path constructors for the test service.
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package server

import (
	"fmt"
)

// GetTestPath returns the URL path to the test service get HTTP endpoint.
func GetTestPath(id int64) string {
	return fmt.Sprintf("/test/%v", id)
}

// ErrorTestPath returns the URL path to the test service error HTTP endpoint.
func ErrorTestPath() string {
	return "/test/error"
}

// EmailTestPath returns the URL path to the test service email HTTP endpoint.
func EmailTestPath() string {
	return "/test/email"
}
