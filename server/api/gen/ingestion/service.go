// Code generated by goa v3.1.2, DO NOT EDIT.
//
// ingestion service
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package ingestion

import (
	"context"

	"goa.design/goa/v3/security"
)

// Service is the ingestion service interface.
type Service interface {
	// ProcessPending implements process pending.
	ProcessPending(context.Context, *ProcessPendingPayload) (err error)
	// ProcessStation implements process station.
	ProcessStation(context.Context, *ProcessStationPayload) (err error)
	// ProcessIngestion implements process ingestion.
	ProcessIngestion(context.Context, *ProcessIngestionPayload) (err error)
	// Delete implements delete.
	Delete(context.Context, *DeletePayload) (err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "ingestion"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [4]string{"process pending", "process station", "process ingestion", "delete"}

// ProcessPendingPayload is the payload type of the ingestion service process
// pending method.
type ProcessPendingPayload struct {
	Auth string
}

// ProcessStationPayload is the payload type of the ingestion service process
// station method.
type ProcessStationPayload struct {
	Auth      string
	StationID int32
}

// ProcessIngestionPayload is the payload type of the ingestion service process
// ingestion method.
type ProcessIngestionPayload struct {
	Auth        string
	IngestionID int64
}

// DeletePayload is the payload type of the ingestion service delete method.
type DeletePayload struct {
	Auth        string
	IngestionID int64
}

// unauthorized
type Unauthorized string

// forbidden
type Forbidden string

// not-found
type NotFound string

// bad-request
type BadRequest string

// Error returns an error description.
func (e Unauthorized) Error() string {
	return "unauthorized"
}

// ErrorName returns "unauthorized".
func (e Unauthorized) ErrorName() string {
	return "unauthorized"
}

// Error returns an error description.
func (e Forbidden) Error() string {
	return "forbidden"
}

// ErrorName returns "forbidden".
func (e Forbidden) ErrorName() string {
	return "forbidden"
}

// Error returns an error description.
func (e NotFound) Error() string {
	return "not-found"
}

// ErrorName returns "not-found".
func (e NotFound) ErrorName() string {
	return "not-found"
}

// Error returns an error description.
func (e BadRequest) Error() string {
	return "bad-request"
}

// ErrorName returns "bad-request".
func (e BadRequest) ErrorName() string {
	return "bad-request"
}