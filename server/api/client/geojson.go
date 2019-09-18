// Code generated by goagen v1.4.0, DO NOT EDIT.
//
// API "fieldkit": GeoJSON Resource Client
//
// Command:
// $ main

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// GeographicalQueryGeoJSONPath computes a request path to the geographical query action of GeoJSON.
func GeographicalQueryGeoJSONPath() string {

	return fmt.Sprintf("/features")
}

// List features in a geographical area.
func (c *Client) GeographicalQueryGeoJSON(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewGeographicalQueryGeoJSONRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGeographicalQueryGeoJSONRequest create the request corresponding to the geographical query action endpoint of the GeoJSON resource.
func (c *Client) NewGeographicalQueryGeoJSONRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListByIDGeoJSONPath computes a request path to the list by id action of GeoJSON.
func ListByIDGeoJSONPath(featureID int) string {
	param0 := strconv.Itoa(featureID)

	return fmt.Sprintf("/features/%s/geojson", param0)
}

// List a feature's GeoJSON by id.
func (c *Client) ListByIDGeoJSON(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListByIDGeoJSONRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListByIDGeoJSONRequest create the request corresponding to the list by id action endpoint of the GeoJSON resource.
func (c *Client) NewListByIDGeoJSONRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListBySourceGeoJSONPath computes a request path to the list by source action of GeoJSON.
func ListBySourceGeoJSONPath(sourceID int) string {
	param0 := strconv.Itoa(sourceID)

	return fmt.Sprintf("/sources/%s/geojson", param0)
}

// List an source's features in a GeoJSON.
func (c *Client) ListBySourceGeoJSON(ctx context.Context, path string, descending *bool) (*http.Response, error) {
	req, err := c.NewListBySourceGeoJSONRequest(ctx, path, descending)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListBySourceGeoJSONRequest create the request corresponding to the list by source action endpoint of the GeoJSON resource.
func (c *Client) NewListBySourceGeoJSONRequest(ctx context.Context, path string, descending *bool) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if descending != nil {
		tmp180 := strconv.FormatBool(*descending)
		values.Set("descending", tmp180)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
