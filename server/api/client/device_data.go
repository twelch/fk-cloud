// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "fieldkit": DeviceData Resource Client
//
// Command:
// $ main

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// AllDeviceDataPath computes a request path to the all action of DeviceData.
func AllDeviceDataPath(deviceID string) string {
	param0 := deviceID

	return fmt.Sprintf("/devices/%s/data", param0)
}

// Get all device data
func (c *Client) AllDeviceData(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewAllDeviceDataRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewAllDeviceDataRequest create the request corresponding to the all action endpoint of the DeviceData resource.
func (c *Client) NewAllDeviceDataRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}