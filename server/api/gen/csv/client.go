// Code generated by goa v3.1.2, DO NOT EDIT.
//
// csv client
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package csv

import (
	"context"
	"io"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "csv" service client.
type Client struct {
	ExportEndpoint   goa.Endpoint
	DownloadEndpoint goa.Endpoint
}

// NewClient initializes a "csv" service client given the endpoints.
func NewClient(export, download goa.Endpoint) *Client {
	return &Client{
		ExportEndpoint:   export,
		DownloadEndpoint: download,
	}
}

// Export calls the "export" endpoint of the "csv" service.
func (c *Client) Export(ctx context.Context, p *ExportPayload) (res *ExportResult, err error) {
	var ires interface{}
	ires, err = c.ExportEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*ExportResult), nil
}

// Download calls the "download" endpoint of the "csv" service.
// Download may return the following errors:
//	- "busy" (type *goa.ServiceError)
//	- error: internal error
func (c *Client) Download(ctx context.Context, p *DownloadPayload) (res *DownloadResult, resp io.ReadCloser, err error) {
	var ires interface{}
	ires, err = c.DownloadEndpoint(ctx, p)
	if err != nil {
		return
	}
	o := ires.(*DownloadResponseData)
	return o.Result, o.Body, nil
}
