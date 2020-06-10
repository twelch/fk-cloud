// Code generated by goa v3.1.2, DO NOT EDIT.
//
// station HTTP client CLI support package
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package client

import (
	"encoding/json"
	"fmt"
	"strconv"

	station "github.com/fieldkit/cloud/server/api/gen/station"
	goa "goa.design/goa/v3/pkg"
)

// BuildAddPayload builds the payload for the station add endpoint from CLI
// flags.
func BuildAddPayload(stationAddBody string, stationAddAuth string) (*station.AddPayload, error) {
	var err error
	var body AddRequestBody
	{
		err = json.Unmarshal([]byte(stationAddBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"device_id\": \"Harum dicta aut at suscipit.\",\n      \"location_name\": \"Sed et animi.\",\n      \"name\": \"Ut placeat.\",\n      \"status_json\": {\n         \"Non nulla.\": \"Sequi libero.\",\n         \"Qui aut iste.\": \"Aperiam labore nemo corrupti et non suscipit.\",\n         \"Voluptate fugit autem.\": \"Omnis magnam velit id quam nulla.\"\n      },\n      \"status_pb\": \"Aut laborum aliquam tempora ullam.\"\n   }'")
		}
		if body.StatusJSON == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("status_json", "body"))
		}
		if err != nil {
			return nil, err
		}
	}
	var auth string
	{
		auth = stationAddAuth
	}
	v := &station.AddPayload{
		Name:         body.Name,
		DeviceID:     body.DeviceID,
		LocationName: body.LocationName,
		StatusPb:     body.StatusPb,
	}
	if body.StatusJSON != nil {
		v.StatusJSON = make(map[string]interface{}, len(body.StatusJSON))
		for key, val := range body.StatusJSON {
			tk := key
			tv := val
			v.StatusJSON[tk] = tv
		}
	}
	v.Auth = auth

	return v, nil
}

// BuildGetPayload builds the payload for the station get endpoint from CLI
// flags.
func BuildGetPayload(stationGetID string, stationGetAuth string) (*station.GetPayload, error) {
	var err error
	var id int32
	{
		var v int64
		v, err = strconv.ParseInt(stationGetID, 10, 32)
		id = int32(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be INT32")
		}
	}
	var auth string
	{
		auth = stationGetAuth
	}
	v := &station.GetPayload{}
	v.ID = id
	v.Auth = auth

	return v, nil
}

// BuildUpdatePayload builds the payload for the station update endpoint from
// CLI flags.
func BuildUpdatePayload(stationUpdateBody string, stationUpdateID string, stationUpdateAuth string) (*station.UpdatePayload, error) {
	var err error
	var body UpdateRequestBody
	{
		err = json.Unmarshal([]byte(stationUpdateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"location_name\": \"Similique similique vel.\",\n      \"name\": \"Quae sunt in officia perspiciatis maiores.\",\n      \"status_json\": {\n         \"Et vero aut qui.\": \"Dolor eveniet ipsum aperiam et eaque.\",\n         \"Impedit ipsam enim minima recusandae modi aliquid.\": \"Qui dolores sit.\",\n         \"Ut odit.\": \"Animi est rerum similique architecto.\"\n      },\n      \"status_pb\": \"Quidem qui voluptatem nulla ipsa.\"\n   }'")
		}
		if body.StatusJSON == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("status_json", "body"))
		}
		if err != nil {
			return nil, err
		}
	}
	var id int32
	{
		var v int64
		v, err = strconv.ParseInt(stationUpdateID, 10, 32)
		id = int32(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be INT32")
		}
	}
	var auth string
	{
		auth = stationUpdateAuth
	}
	v := &station.UpdatePayload{
		Name:         body.Name,
		LocationName: body.LocationName,
		StatusPb:     body.StatusPb,
	}
	if body.StatusJSON != nil {
		v.StatusJSON = make(map[string]interface{}, len(body.StatusJSON))
		for key, val := range body.StatusJSON {
			tk := key
			tv := val
			v.StatusJSON[tk] = tv
		}
	}
	v.ID = id
	v.Auth = auth

	return v, nil
}

// BuildListMinePayload builds the payload for the station list mine endpoint
// from CLI flags.
func BuildListMinePayload(stationListMineAuth string) (*station.ListMinePayload, error) {
	var auth string
	{
		auth = stationListMineAuth
	}
	v := &station.ListMinePayload{}
	v.Auth = auth

	return v, nil
}

// BuildListProjectPayload builds the payload for the station list project
// endpoint from CLI flags.
func BuildListProjectPayload(stationListProjectID string, stationListProjectAuth string) (*station.ListProjectPayload, error) {
	var err error
	var id int32
	{
		var v int64
		v, err = strconv.ParseInt(stationListProjectID, 10, 32)
		id = int32(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be INT32")
		}
	}
	var auth string
	{
		auth = stationListProjectAuth
	}
	v := &station.ListProjectPayload{}
	v.ID = id
	v.Auth = auth

	return v, nil
}

// BuildPhotoPayload builds the payload for the station photo endpoint from CLI
// flags.
func BuildPhotoPayload(stationPhotoID string, stationPhotoAuth string) (*station.PhotoPayload, error) {
	var err error
	var id int32
	{
		var v int64
		v, err = strconv.ParseInt(stationPhotoID, 10, 32)
		id = int32(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be INT32")
		}
	}
	var auth string
	{
		auth = stationPhotoAuth
	}
	v := &station.PhotoPayload{}
	v.ID = id
	v.Auth = auth

	return v, nil
}
