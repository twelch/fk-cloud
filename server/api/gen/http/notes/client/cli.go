// Code generated by goa v3.1.2, DO NOT EDIT.
//
// notes HTTP client CLI support package
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package client

import (
	"encoding/json"
	"fmt"
	"strconv"

	notes "github.com/fieldkit/cloud/server/api/gen/notes"
	goa "goa.design/goa/v3/pkg"
)

// BuildUpdatePayload builds the payload for the notes update endpoint from CLI
// flags.
func BuildUpdatePayload(notesUpdateBody string, notesUpdateStationID string, notesUpdateAuth string) (*notes.UpdatePayload, error) {
	var err error
	var body UpdateRequestBody
	{
		err = json.Unmarshal([]byte(notesUpdateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"notes\": {\n         \"creating\": [\n            {\n               \"body\": \"Beatae magni aliquam perferendis facilis.\",\n               \"key\": \"Vitae ipsam corrupti ea dolorum consequatur.\",\n               \"mediaId\": 2771950207859890457\n            },\n            {\n               \"body\": \"Beatae magni aliquam perferendis facilis.\",\n               \"key\": \"Vitae ipsam corrupti ea dolorum consequatur.\",\n               \"mediaId\": 2771950207859890457\n            },\n            {\n               \"body\": \"Beatae magni aliquam perferendis facilis.\",\n               \"key\": \"Vitae ipsam corrupti ea dolorum consequatur.\",\n               \"mediaId\": 2771950207859890457\n            }\n         ],\n         \"notes\": [\n            {\n               \"body\": \"Fugit quos sed enim occaecati.\",\n               \"id\": 6139879987527965591,\n               \"key\": \"Dolorum et autem tempora consequatur architecto optio.\",\n               \"mediaId\": 6035690160929153665\n            },\n            {\n               \"body\": \"Fugit quos sed enim occaecati.\",\n               \"id\": 6139879987527965591,\n               \"key\": \"Dolorum et autem tempora consequatur architecto optio.\",\n               \"mediaId\": 6035690160929153665\n            },\n            {\n               \"body\": \"Fugit quos sed enim occaecati.\",\n               \"id\": 6139879987527965591,\n               \"key\": \"Dolorum et autem tempora consequatur architecto optio.\",\n               \"mediaId\": 6035690160929153665\n            }\n         ]\n      }\n   }'")
		}
		if body.Notes == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("notes", "body"))
		}
		if body.Notes != nil {
			if err2 := ValidateFieldNoteUpdateRequestBody(body.Notes); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
		if err != nil {
			return nil, err
		}
	}
	var stationID int32
	{
		var v int64
		v, err = strconv.ParseInt(notesUpdateStationID, 10, 32)
		stationID = int32(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for stationID, must be INT32")
		}
	}
	var auth string
	{
		auth = notesUpdateAuth
	}
	v := &notes.UpdatePayload{}
	if body.Notes != nil {
		v.Notes = marshalFieldNoteUpdateRequestBodyToNotesFieldNoteUpdate(body.Notes)
	}
	v.StationID = stationID
	v.Auth = auth

	return v, nil
}

// BuildGetPayload builds the payload for the notes get endpoint from CLI flags.
func BuildGetPayload(notesGetStationID string, notesGetAuth string) (*notes.GetPayload, error) {
	var err error
	var stationID int32
	{
		var v int64
		v, err = strconv.ParseInt(notesGetStationID, 10, 32)
		stationID = int32(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for stationID, must be INT32")
		}
	}
	var auth string
	{
		auth = notesGetAuth
	}
	v := &notes.GetPayload{}
	v.StationID = stationID
	v.Auth = auth

	return v, nil
}

// BuildMediaPayload builds the payload for the notes media endpoint from CLI
// flags.
func BuildMediaPayload(notesMediaMediaID string, notesMediaAuth string) (*notes.MediaPayload, error) {
	var err error
	var mediaID int32
	{
		var v int64
		v, err = strconv.ParseInt(notesMediaMediaID, 10, 32)
		mediaID = int32(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for mediaID, must be INT32")
		}
	}
	var auth string
	{
		auth = notesMediaAuth
	}
	v := &notes.MediaPayload{}
	v.MediaID = mediaID
	v.Auth = auth

	return v, nil
}

// BuildUploadPayload builds the payload for the notes upload endpoint from CLI
// flags.
func BuildUploadPayload(notesUploadAuth string, notesUploadContentType string, notesUploadContentLength string) (*notes.UploadPayload, error) {
	var err error
	var auth string
	{
		auth = notesUploadAuth
	}
	var contentType string
	{
		contentType = notesUploadContentType
	}
	var contentLength int64
	{
		contentLength, err = strconv.ParseInt(notesUploadContentLength, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value for contentLength, must be INT64")
		}
	}
	v := &notes.UploadPayload{}
	v.Auth = auth
	v.ContentType = contentType
	v.ContentLength = contentLength

	return v, nil
}