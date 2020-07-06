// Code generated by goa v3.1.2, DO NOT EDIT.
//
// notes HTTP client types
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package client

import (
	notes "github.com/fieldkit/cloud/server/api/gen/notes"
	notesviews "github.com/fieldkit/cloud/server/api/gen/notes/views"
	goa "goa.design/goa/v3/pkg"
)

// UpdateRequestBody is the type of the "notes" service "update" endpoint HTTP
// request body.
type UpdateRequestBody struct {
	Notes *FieldNoteUpdateRequestBody `form:"notes" json:"notes" xml:"notes"`
}

// UpdateResponseBody is the type of the "notes" service "update" endpoint HTTP
// response body.
type UpdateResponseBody struct {
	Notes []*FieldNoteResponseBody `form:"notes,omitempty" json:"notes,omitempty" xml:"notes,omitempty"`
}

// GetResponseBody is the type of the "notes" service "get" endpoint HTTP
// response body.
type GetResponseBody struct {
	Notes []*FieldNoteResponseBody `form:"notes,omitempty" json:"notes,omitempty" xml:"notes,omitempty"`
}

// UploadResponseBody is the type of the "notes" service "upload" endpoint HTTP
// response body.
type UploadResponseBody struct {
	ID  *int64  `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	URL *string `form:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
}

// UpdateBadRequestResponseBody is the type of the "notes" service "update"
// endpoint HTTP response body for the "bad-request" error.
type UpdateBadRequestResponseBody string

// UpdateForbiddenResponseBody is the type of the "notes" service "update"
// endpoint HTTP response body for the "forbidden" error.
type UpdateForbiddenResponseBody string

// UpdateNotFoundResponseBody is the type of the "notes" service "update"
// endpoint HTTP response body for the "not-found" error.
type UpdateNotFoundResponseBody string

// UpdateUnauthorizedResponseBody is the type of the "notes" service "update"
// endpoint HTTP response body for the "unauthorized" error.
type UpdateUnauthorizedResponseBody string

// GetBadRequestResponseBody is the type of the "notes" service "get" endpoint
// HTTP response body for the "bad-request" error.
type GetBadRequestResponseBody string

// GetForbiddenResponseBody is the type of the "notes" service "get" endpoint
// HTTP response body for the "forbidden" error.
type GetForbiddenResponseBody string

// GetNotFoundResponseBody is the type of the "notes" service "get" endpoint
// HTTP response body for the "not-found" error.
type GetNotFoundResponseBody string

// GetUnauthorizedResponseBody is the type of the "notes" service "get"
// endpoint HTTP response body for the "unauthorized" error.
type GetUnauthorizedResponseBody string

// MediaBadRequestResponseBody is the type of the "notes" service "media"
// endpoint HTTP response body for the "bad-request" error.
type MediaBadRequestResponseBody string

// MediaForbiddenResponseBody is the type of the "notes" service "media"
// endpoint HTTP response body for the "forbidden" error.
type MediaForbiddenResponseBody string

// MediaNotFoundResponseBody is the type of the "notes" service "media"
// endpoint HTTP response body for the "not-found" error.
type MediaNotFoundResponseBody string

// MediaUnauthorizedResponseBody is the type of the "notes" service "media"
// endpoint HTTP response body for the "unauthorized" error.
type MediaUnauthorizedResponseBody string

// UploadBadRequestResponseBody is the type of the "notes" service "upload"
// endpoint HTTP response body for the "bad-request" error.
type UploadBadRequestResponseBody string

// UploadForbiddenResponseBody is the type of the "notes" service "upload"
// endpoint HTTP response body for the "forbidden" error.
type UploadForbiddenResponseBody string

// UploadNotFoundResponseBody is the type of the "notes" service "upload"
// endpoint HTTP response body for the "not-found" error.
type UploadNotFoundResponseBody string

// UploadUnauthorizedResponseBody is the type of the "notes" service "upload"
// endpoint HTTP response body for the "unauthorized" error.
type UploadUnauthorizedResponseBody string

// FieldNoteUpdateRequestBody is used to define fields on request body types.
type FieldNoteUpdateRequestBody struct {
	Notes    []*ExistingFieldNoteRequestBody `form:"notes" json:"notes" xml:"notes"`
	Creating []*NewFieldNoteRequestBody      `form:"creating" json:"creating" xml:"creating"`
}

// ExistingFieldNoteRequestBody is used to define fields on request body types.
type ExistingFieldNoteRequestBody struct {
	ID      int64   `form:"id" json:"id" xml:"id"`
	Key     *string `form:"key,omitempty" json:"key,omitempty" xml:"key,omitempty"`
	Body    *string `form:"body,omitempty" json:"body,omitempty" xml:"body,omitempty"`
	MediaID *int64  `form:"mediaId,omitempty" json:"mediaId,omitempty" xml:"mediaId,omitempty"`
}

// NewFieldNoteRequestBody is used to define fields on request body types.
type NewFieldNoteRequestBody struct {
	Key     *string `form:"key,omitempty" json:"key,omitempty" xml:"key,omitempty"`
	Body    *string `form:"body,omitempty" json:"body,omitempty" xml:"body,omitempty"`
	MediaID *int64  `form:"mediaId,omitempty" json:"mediaId,omitempty" xml:"mediaId,omitempty"`
}

// FieldNoteResponseBody is used to define fields on response body types.
type FieldNoteResponseBody struct {
	ID        *int64                       `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	CreatedAt *int64                       `form:"createdAt,omitempty" json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Author    *FieldNoteAuthorResponseBody `form:"author,omitempty" json:"author,omitempty" xml:"author,omitempty"`
	Key       *string                      `form:"key,omitempty" json:"key,omitempty" xml:"key,omitempty"`
	Body      *string                      `form:"body,omitempty" json:"body,omitempty" xml:"body,omitempty"`
	MediaID   *int64                       `form:"mediaId,omitempty" json:"mediaId,omitempty" xml:"mediaId,omitempty"`
}

// FieldNoteAuthorResponseBody is used to define fields on response body types.
type FieldNoteAuthorResponseBody struct {
	ID       *int32  `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Name     *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	MediaURL *string `form:"mediaUrl,omitempty" json:"mediaUrl,omitempty" xml:"mediaUrl,omitempty"`
}

// NewUpdateRequestBody builds the HTTP request body from the payload of the
// "update" endpoint of the "notes" service.
func NewUpdateRequestBody(p *notes.UpdatePayload) *UpdateRequestBody {
	body := &UpdateRequestBody{}
	if p.Notes != nil {
		body.Notes = marshalNotesFieldNoteUpdateToFieldNoteUpdateRequestBody(p.Notes)
	}
	return body
}

// NewUpdateFieldNotesOK builds a "notes" service "update" endpoint result from
// a HTTP "OK" response.
func NewUpdateFieldNotesOK(body *UpdateResponseBody) *notesviews.FieldNotesView {
	v := &notesviews.FieldNotesView{}
	v.Notes = make([]*notesviews.FieldNoteView, len(body.Notes))
	for i, val := range body.Notes {
		v.Notes[i] = unmarshalFieldNoteResponseBodyToNotesviewsFieldNoteView(val)
	}

	return v
}

// NewUpdateBadRequest builds a notes service update endpoint bad-request error.
func NewUpdateBadRequest(body UpdateBadRequestResponseBody) notes.BadRequest {
	v := notes.BadRequest(body)
	return v
}

// NewUpdateForbidden builds a notes service update endpoint forbidden error.
func NewUpdateForbidden(body UpdateForbiddenResponseBody) notes.Forbidden {
	v := notes.Forbidden(body)
	return v
}

// NewUpdateNotFound builds a notes service update endpoint not-found error.
func NewUpdateNotFound(body UpdateNotFoundResponseBody) notes.NotFound {
	v := notes.NotFound(body)
	return v
}

// NewUpdateUnauthorized builds a notes service update endpoint unauthorized
// error.
func NewUpdateUnauthorized(body UpdateUnauthorizedResponseBody) notes.Unauthorized {
	v := notes.Unauthorized(body)
	return v
}

// NewGetFieldNotesOK builds a "notes" service "get" endpoint result from a
// HTTP "OK" response.
func NewGetFieldNotesOK(body *GetResponseBody) *notesviews.FieldNotesView {
	v := &notesviews.FieldNotesView{}
	v.Notes = make([]*notesviews.FieldNoteView, len(body.Notes))
	for i, val := range body.Notes {
		v.Notes[i] = unmarshalFieldNoteResponseBodyToNotesviewsFieldNoteView(val)
	}

	return v
}

// NewGetBadRequest builds a notes service get endpoint bad-request error.
func NewGetBadRequest(body GetBadRequestResponseBody) notes.BadRequest {
	v := notes.BadRequest(body)
	return v
}

// NewGetForbidden builds a notes service get endpoint forbidden error.
func NewGetForbidden(body GetForbiddenResponseBody) notes.Forbidden {
	v := notes.Forbidden(body)
	return v
}

// NewGetNotFound builds a notes service get endpoint not-found error.
func NewGetNotFound(body GetNotFoundResponseBody) notes.NotFound {
	v := notes.NotFound(body)
	return v
}

// NewGetUnauthorized builds a notes service get endpoint unauthorized error.
func NewGetUnauthorized(body GetUnauthorizedResponseBody) notes.Unauthorized {
	v := notes.Unauthorized(body)
	return v
}

// NewMediaResultOK builds a "notes" service "media" endpoint result from a
// HTTP "OK" response.
func NewMediaResultOK(length int64, contentType string) *notes.MediaResult {
	v := &notes.MediaResult{}
	v.Length = length
	v.ContentType = contentType

	return v
}

// NewMediaBadRequest builds a notes service media endpoint bad-request error.
func NewMediaBadRequest(body MediaBadRequestResponseBody) notes.BadRequest {
	v := notes.BadRequest(body)
	return v
}

// NewMediaForbidden builds a notes service media endpoint forbidden error.
func NewMediaForbidden(body MediaForbiddenResponseBody) notes.Forbidden {
	v := notes.Forbidden(body)
	return v
}

// NewMediaNotFound builds a notes service media endpoint not-found error.
func NewMediaNotFound(body MediaNotFoundResponseBody) notes.NotFound {
	v := notes.NotFound(body)
	return v
}

// NewMediaUnauthorized builds a notes service media endpoint unauthorized
// error.
func NewMediaUnauthorized(body MediaUnauthorizedResponseBody) notes.Unauthorized {
	v := notes.Unauthorized(body)
	return v
}

// NewUploadNoteMediaOK builds a "notes" service "upload" endpoint result from
// a HTTP "OK" response.
func NewUploadNoteMediaOK(body *UploadResponseBody) *notesviews.NoteMediaView {
	v := &notesviews.NoteMediaView{
		ID:  body.ID,
		URL: body.URL,
	}

	return v
}

// NewUploadBadRequest builds a notes service upload endpoint bad-request error.
func NewUploadBadRequest(body UploadBadRequestResponseBody) notes.BadRequest {
	v := notes.BadRequest(body)
	return v
}

// NewUploadForbidden builds a notes service upload endpoint forbidden error.
func NewUploadForbidden(body UploadForbiddenResponseBody) notes.Forbidden {
	v := notes.Forbidden(body)
	return v
}

// NewUploadNotFound builds a notes service upload endpoint not-found error.
func NewUploadNotFound(body UploadNotFoundResponseBody) notes.NotFound {
	v := notes.NotFound(body)
	return v
}

// NewUploadUnauthorized builds a notes service upload endpoint unauthorized
// error.
func NewUploadUnauthorized(body UploadUnauthorizedResponseBody) notes.Unauthorized {
	v := notes.Unauthorized(body)
	return v
}

// ValidateFieldNoteUpdateRequestBody runs the validations defined on
// FieldNoteUpdateRequestBody
func ValidateFieldNoteUpdateRequestBody(body *FieldNoteUpdateRequestBody) (err error) {
	if body.Notes == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("notes", "body"))
	}
	if body.Creating == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("creating", "body"))
	}
	return
}

// ValidateFieldNoteResponseBody runs the validations defined on
// FieldNoteResponseBody
func ValidateFieldNoteResponseBody(body *FieldNoteResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.CreatedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("createdAt", "body"))
	}
	if body.Author == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("author", "body"))
	}
	if body.Author != nil {
		if err2 := ValidateFieldNoteAuthorResponseBody(body.Author); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateFieldNoteAuthorResponseBody runs the validations defined on
// FieldNoteAuthorResponseBody
func ValidateFieldNoteAuthorResponseBody(body *FieldNoteAuthorResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.MediaURL == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("mediaUrl", "body"))
	}
	return
}