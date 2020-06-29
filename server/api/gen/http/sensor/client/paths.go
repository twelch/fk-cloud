// Code generated by goa v3.1.2, DO NOT EDIT.
//
// HTTP request path constructors for the sensor service.
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package client

// MetaSensorPath returns the URL path to the sensor service meta HTTP endpoint.
func MetaSensorPath() string {
	return "/sensors"
}

// DataSensorPath returns the URL path to the sensor service data HTTP endpoint.
func DataSensorPath() string {
	return "/sensors/data"
}
