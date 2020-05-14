// Code generated by goa v3.1.2, DO NOT EDIT.
//
// project HTTP server types
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package server

import (
	project "github.com/fieldkit/cloud/server/api/gen/project"
	projectviews "github.com/fieldkit/cloud/server/api/gen/project/views"
	goa "goa.design/goa/v3/pkg"
)

// UpdateRequestBody is the type of the "project" service "update" endpoint
// HTTP request body.
type UpdateRequestBody struct {
	Body *string `form:"body,omitempty" json:"body,omitempty" xml:"body,omitempty"`
}

// InvitesResponseBody is the type of the "project" service "invites" endpoint
// HTTP response body.
type InvitesResponseBody struct {
	Pending []*PendingInviteResponseBody `form:"pending" json:"pending" xml:"pending"`
}

// LookupInviteResponseBody is the type of the "project" service "lookup
// invite" endpoint HTTP response body.
type LookupInviteResponseBody struct {
	Pending []*PendingInviteResponseBody `form:"pending" json:"pending" xml:"pending"`
}

// UpdateBadRequestResponseBody is the type of the "project" service "update"
// endpoint HTTP response body for the "bad-request" error.
type UpdateBadRequestResponseBody string

// UpdateForbiddenResponseBody is the type of the "project" service "update"
// endpoint HTTP response body for the "forbidden" error.
type UpdateForbiddenResponseBody string

// UpdateNotFoundResponseBody is the type of the "project" service "update"
// endpoint HTTP response body for the "not-found" error.
type UpdateNotFoundResponseBody string

// UpdateUnauthorizedResponseBody is the type of the "project" service "update"
// endpoint HTTP response body for the "unauthorized" error.
type UpdateUnauthorizedResponseBody string

// InvitesBadRequestResponseBody is the type of the "project" service "invites"
// endpoint HTTP response body for the "bad-request" error.
type InvitesBadRequestResponseBody string

// InvitesForbiddenResponseBody is the type of the "project" service "invites"
// endpoint HTTP response body for the "forbidden" error.
type InvitesForbiddenResponseBody string

// InvitesNotFoundResponseBody is the type of the "project" service "invites"
// endpoint HTTP response body for the "not-found" error.
type InvitesNotFoundResponseBody string

// InvitesUnauthorizedResponseBody is the type of the "project" service
// "invites" endpoint HTTP response body for the "unauthorized" error.
type InvitesUnauthorizedResponseBody string

// LookupInviteBadRequestResponseBody is the type of the "project" service
// "lookup invite" endpoint HTTP response body for the "bad-request" error.
type LookupInviteBadRequestResponseBody string

// LookupInviteForbiddenResponseBody is the type of the "project" service
// "lookup invite" endpoint HTTP response body for the "forbidden" error.
type LookupInviteForbiddenResponseBody string

// LookupInviteNotFoundResponseBody is the type of the "project" service
// "lookup invite" endpoint HTTP response body for the "not-found" error.
type LookupInviteNotFoundResponseBody string

// LookupInviteUnauthorizedResponseBody is the type of the "project" service
// "lookup invite" endpoint HTTP response body for the "unauthorized" error.
type LookupInviteUnauthorizedResponseBody string

// AcceptInviteBadRequestResponseBody is the type of the "project" service
// "accept invite" endpoint HTTP response body for the "bad-request" error.
type AcceptInviteBadRequestResponseBody string

// AcceptInviteForbiddenResponseBody is the type of the "project" service
// "accept invite" endpoint HTTP response body for the "forbidden" error.
type AcceptInviteForbiddenResponseBody string

// AcceptInviteNotFoundResponseBody is the type of the "project" service
// "accept invite" endpoint HTTP response body for the "not-found" error.
type AcceptInviteNotFoundResponseBody string

// AcceptInviteUnauthorizedResponseBody is the type of the "project" service
// "accept invite" endpoint HTTP response body for the "unauthorized" error.
type AcceptInviteUnauthorizedResponseBody string

// RejectInviteBadRequestResponseBody is the type of the "project" service
// "reject invite" endpoint HTTP response body for the "bad-request" error.
type RejectInviteBadRequestResponseBody string

// RejectInviteForbiddenResponseBody is the type of the "project" service
// "reject invite" endpoint HTTP response body for the "forbidden" error.
type RejectInviteForbiddenResponseBody string

// RejectInviteNotFoundResponseBody is the type of the "project" service
// "reject invite" endpoint HTTP response body for the "not-found" error.
type RejectInviteNotFoundResponseBody string

// RejectInviteUnauthorizedResponseBody is the type of the "project" service
// "reject invite" endpoint HTTP response body for the "unauthorized" error.
type RejectInviteUnauthorizedResponseBody string

// PendingInviteResponseBody is used to define fields on response body types.
type PendingInviteResponseBody struct {
	ID      int64                       `form:"id" json:"id" xml:"id"`
	Project *ProjectSummaryResponseBody `form:"project" json:"project" xml:"project"`
	Time    int64                       `form:"time" json:"time" xml:"time"`
	Role    int32                       `form:"role" json:"role" xml:"role"`
}

// ProjectSummaryResponseBody is used to define fields on response body types.
type ProjectSummaryResponseBody struct {
	ID   int64  `form:"id" json:"id" xml:"id"`
	Name string `form:"name" json:"name" xml:"name"`
}

// NewInvitesResponseBody builds the HTTP response body from the result of the
// "invites" endpoint of the "project" service.
func NewInvitesResponseBody(res *projectviews.PendingInvitesView) *InvitesResponseBody {
	body := &InvitesResponseBody{}
	if res.Pending != nil {
		body.Pending = make([]*PendingInviteResponseBody, len(res.Pending))
		for i, val := range res.Pending {
			body.Pending[i] = marshalProjectviewsPendingInviteViewToPendingInviteResponseBody(val)
		}
	}
	return body
}

// NewLookupInviteResponseBody builds the HTTP response body from the result of
// the "lookup invite" endpoint of the "project" service.
func NewLookupInviteResponseBody(res *projectviews.PendingInvitesView) *LookupInviteResponseBody {
	body := &LookupInviteResponseBody{}
	if res.Pending != nil {
		body.Pending = make([]*PendingInviteResponseBody, len(res.Pending))
		for i, val := range res.Pending {
			body.Pending[i] = marshalProjectviewsPendingInviteViewToPendingInviteResponseBody(val)
		}
	}
	return body
}

// NewUpdateBadRequestResponseBody builds the HTTP response body from the
// result of the "update" endpoint of the "project" service.
func NewUpdateBadRequestResponseBody(res project.BadRequest) UpdateBadRequestResponseBody {
	body := UpdateBadRequestResponseBody(res)
	return body
}

// NewUpdateForbiddenResponseBody builds the HTTP response body from the result
// of the "update" endpoint of the "project" service.
func NewUpdateForbiddenResponseBody(res project.Forbidden) UpdateForbiddenResponseBody {
	body := UpdateForbiddenResponseBody(res)
	return body
}

// NewUpdateNotFoundResponseBody builds the HTTP response body from the result
// of the "update" endpoint of the "project" service.
func NewUpdateNotFoundResponseBody(res project.NotFound) UpdateNotFoundResponseBody {
	body := UpdateNotFoundResponseBody(res)
	return body
}

// NewUpdateUnauthorizedResponseBody builds the HTTP response body from the
// result of the "update" endpoint of the "project" service.
func NewUpdateUnauthorizedResponseBody(res project.Unauthorized) UpdateUnauthorizedResponseBody {
	body := UpdateUnauthorizedResponseBody(res)
	return body
}

// NewInvitesBadRequestResponseBody builds the HTTP response body from the
// result of the "invites" endpoint of the "project" service.
func NewInvitesBadRequestResponseBody(res project.BadRequest) InvitesBadRequestResponseBody {
	body := InvitesBadRequestResponseBody(res)
	return body
}

// NewInvitesForbiddenResponseBody builds the HTTP response body from the
// result of the "invites" endpoint of the "project" service.
func NewInvitesForbiddenResponseBody(res project.Forbidden) InvitesForbiddenResponseBody {
	body := InvitesForbiddenResponseBody(res)
	return body
}

// NewInvitesNotFoundResponseBody builds the HTTP response body from the result
// of the "invites" endpoint of the "project" service.
func NewInvitesNotFoundResponseBody(res project.NotFound) InvitesNotFoundResponseBody {
	body := InvitesNotFoundResponseBody(res)
	return body
}

// NewInvitesUnauthorizedResponseBody builds the HTTP response body from the
// result of the "invites" endpoint of the "project" service.
func NewInvitesUnauthorizedResponseBody(res project.Unauthorized) InvitesUnauthorizedResponseBody {
	body := InvitesUnauthorizedResponseBody(res)
	return body
}

// NewLookupInviteBadRequestResponseBody builds the HTTP response body from the
// result of the "lookup invite" endpoint of the "project" service.
func NewLookupInviteBadRequestResponseBody(res project.BadRequest) LookupInviteBadRequestResponseBody {
	body := LookupInviteBadRequestResponseBody(res)
	return body
}

// NewLookupInviteForbiddenResponseBody builds the HTTP response body from the
// result of the "lookup invite" endpoint of the "project" service.
func NewLookupInviteForbiddenResponseBody(res project.Forbidden) LookupInviteForbiddenResponseBody {
	body := LookupInviteForbiddenResponseBody(res)
	return body
}

// NewLookupInviteNotFoundResponseBody builds the HTTP response body from the
// result of the "lookup invite" endpoint of the "project" service.
func NewLookupInviteNotFoundResponseBody(res project.NotFound) LookupInviteNotFoundResponseBody {
	body := LookupInviteNotFoundResponseBody(res)
	return body
}

// NewLookupInviteUnauthorizedResponseBody builds the HTTP response body from
// the result of the "lookup invite" endpoint of the "project" service.
func NewLookupInviteUnauthorizedResponseBody(res project.Unauthorized) LookupInviteUnauthorizedResponseBody {
	body := LookupInviteUnauthorizedResponseBody(res)
	return body
}

// NewAcceptInviteBadRequestResponseBody builds the HTTP response body from the
// result of the "accept invite" endpoint of the "project" service.
func NewAcceptInviteBadRequestResponseBody(res project.BadRequest) AcceptInviteBadRequestResponseBody {
	body := AcceptInviteBadRequestResponseBody(res)
	return body
}

// NewAcceptInviteForbiddenResponseBody builds the HTTP response body from the
// result of the "accept invite" endpoint of the "project" service.
func NewAcceptInviteForbiddenResponseBody(res project.Forbidden) AcceptInviteForbiddenResponseBody {
	body := AcceptInviteForbiddenResponseBody(res)
	return body
}

// NewAcceptInviteNotFoundResponseBody builds the HTTP response body from the
// result of the "accept invite" endpoint of the "project" service.
func NewAcceptInviteNotFoundResponseBody(res project.NotFound) AcceptInviteNotFoundResponseBody {
	body := AcceptInviteNotFoundResponseBody(res)
	return body
}

// NewAcceptInviteUnauthorizedResponseBody builds the HTTP response body from
// the result of the "accept invite" endpoint of the "project" service.
func NewAcceptInviteUnauthorizedResponseBody(res project.Unauthorized) AcceptInviteUnauthorizedResponseBody {
	body := AcceptInviteUnauthorizedResponseBody(res)
	return body
}

// NewRejectInviteBadRequestResponseBody builds the HTTP response body from the
// result of the "reject invite" endpoint of the "project" service.
func NewRejectInviteBadRequestResponseBody(res project.BadRequest) RejectInviteBadRequestResponseBody {
	body := RejectInviteBadRequestResponseBody(res)
	return body
}

// NewRejectInviteForbiddenResponseBody builds the HTTP response body from the
// result of the "reject invite" endpoint of the "project" service.
func NewRejectInviteForbiddenResponseBody(res project.Forbidden) RejectInviteForbiddenResponseBody {
	body := RejectInviteForbiddenResponseBody(res)
	return body
}

// NewRejectInviteNotFoundResponseBody builds the HTTP response body from the
// result of the "reject invite" endpoint of the "project" service.
func NewRejectInviteNotFoundResponseBody(res project.NotFound) RejectInviteNotFoundResponseBody {
	body := RejectInviteNotFoundResponseBody(res)
	return body
}

// NewRejectInviteUnauthorizedResponseBody builds the HTTP response body from
// the result of the "reject invite" endpoint of the "project" service.
func NewRejectInviteUnauthorizedResponseBody(res project.Unauthorized) RejectInviteUnauthorizedResponseBody {
	body := RejectInviteUnauthorizedResponseBody(res)
	return body
}

// NewUpdatePayload builds a project service update endpoint payload.
func NewUpdatePayload(body *UpdateRequestBody, id int64, auth string) *project.UpdatePayload {
	v := &project.UpdatePayload{
		Body: *body.Body,
	}
	v.ID = id
	v.Auth = auth

	return v
}

// NewInvitesPayload builds a project service invites endpoint payload.
func NewInvitesPayload(auth string) *project.InvitesPayload {
	v := &project.InvitesPayload{}
	v.Auth = auth

	return v
}

// NewLookupInvitePayload builds a project service lookup invite endpoint
// payload.
func NewLookupInvitePayload(token string, auth string) *project.LookupInvitePayload {
	v := &project.LookupInvitePayload{}
	v.Token = token
	v.Auth = auth

	return v
}

// NewAcceptInvitePayload builds a project service accept invite endpoint
// payload.
func NewAcceptInvitePayload(id int64, token *string, auth string) *project.AcceptInvitePayload {
	v := &project.AcceptInvitePayload{}
	v.ID = id
	v.Token = token
	v.Auth = auth

	return v
}

// NewRejectInvitePayload builds a project service reject invite endpoint
// payload.
func NewRejectInvitePayload(id int64, token *string, auth string) *project.RejectInvitePayload {
	v := &project.RejectInvitePayload{}
	v.ID = id
	v.Token = token
	v.Auth = auth

	return v
}

// ValidateUpdateRequestBody runs the validations defined on UpdateRequestBody
func ValidateUpdateRequestBody(body *UpdateRequestBody) (err error) {
	if body.Body == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("body", "body"))
	}
	return
}
