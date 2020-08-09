// Code generated by goa v3.1.2, DO NOT EDIT.
//
// HTTP request path constructors for the csv service.
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package server

import (
	"fmt"
)

// ExportCsvPath returns the URL path to the csv service export HTTP endpoint.
func ExportCsvPath() string {
	return "/sensors/data/export/csv"
}

// DownloadCsvPath returns the URL path to the csv service download HTTP endpoint.
func DownloadCsvPath(id string) string {
	return fmt.Sprintf("/sensors/data/export/csv/%v", id)
}
