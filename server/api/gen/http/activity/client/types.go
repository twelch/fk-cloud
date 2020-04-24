// Code generated by goa v3.1.1, DO NOT EDIT.
//
// activity HTTP client types
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package client

import (
	activityviews "github.com/fieldkit/cloud/server/api/gen/activity/views"
	goa "goa.design/goa/v3/pkg"
)

// StationResponseBody is the type of the "activity" service "station" endpoint
// HTTP response body.
type StationResponseBody struct {
	Activities ActivityEntryCollectionResponseBody `form:"activities,omitempty" json:"activities,omitempty" xml:"activities,omitempty"`
	Total      *int32                              `form:"total,omitempty" json:"total,omitempty" xml:"total,omitempty"`
	Page       *int32                              `form:"page,omitempty" json:"page,omitempty" xml:"page,omitempty"`
}

// ProjectResponseBody is the type of the "activity" service "project" endpoint
// HTTP response body.
type ProjectResponseBody struct {
	Activities ActivityEntryCollectionResponseBody `form:"activities,omitempty" json:"activities,omitempty" xml:"activities,omitempty"`
	Total      *int32                              `form:"total,omitempty" json:"total,omitempty" xml:"total,omitempty"`
	Page       *int32                              `form:"page,omitempty" json:"page,omitempty" xml:"page,omitempty"`
}

// ActivityEntryCollectionResponseBody is used to define fields on response
// body types.
type ActivityEntryCollectionResponseBody []*ActivityEntryResponseBody

// ActivityEntryResponseBody is used to define fields on response body types.
type ActivityEntryResponseBody struct {
	ID        *int64                      `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Project   *ProjectSummaryResponseBody `form:"project,omitempty" json:"project,omitempty" xml:"project,omitempty"`
	Station   *StationSummaryResponseBody `form:"station,omitempty" json:"station,omitempty" xml:"station,omitempty"`
	CreatedAt *int64                      `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	Type      *string                     `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	Meta      interface{}                 `form:"meta,omitempty" json:"meta,omitempty" xml:"meta,omitempty"`
}

// ProjectSummaryResponseBody is used to define fields on response body types.
type ProjectSummaryResponseBody struct {
	ID   *int64  `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// StationSummaryResponseBody is used to define fields on response body types.
type StationSummaryResponseBody struct {
	ID   *int64  `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// NewStationActivityPageViewOK builds a "activity" service "station" endpoint
// result from a HTTP "OK" response.
func NewStationActivityPageViewOK(body *StationResponseBody) *activityviews.StationActivityPageView {
	v := &activityviews.StationActivityPageView{
		Total: body.Total,
		Page:  body.Page,
	}
	v.Activities = make([]*activityviews.ActivityEntryView, len(body.Activities))
	for i, val := range body.Activities {
		v.Activities[i] = unmarshalActivityEntryResponseBodyToActivityviewsActivityEntryView(val)
	}

	return v
}

// NewProjectActivityPageViewOK builds a "activity" service "project" endpoint
// result from a HTTP "OK" response.
func NewProjectActivityPageViewOK(body *ProjectResponseBody) *activityviews.ProjectActivityPageView {
	v := &activityviews.ProjectActivityPageView{
		Total: body.Total,
		Page:  body.Page,
	}
	v.Activities = make([]*activityviews.ActivityEntryView, len(body.Activities))
	for i, val := range body.Activities {
		v.Activities[i] = unmarshalActivityEntryResponseBodyToActivityviewsActivityEntryView(val)
	}

	return v
}

// ValidateActivityEntryCollectionResponseBody runs the validations defined on
// ActivityEntryCollectionResponseBody
func ValidateActivityEntryCollectionResponseBody(body ActivityEntryCollectionResponseBody) (err error) {
	for _, e := range body {
		if e != nil {
			if err2 := ValidateActivityEntryResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateActivityEntryResponseBody runs the validations defined on
// ActivityEntryResponseBody
func ValidateActivityEntryResponseBody(body *ActivityEntryResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Project == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("project", "body"))
	}
	if body.Station == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("station", "body"))
	}
	if body.CreatedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("created_at", "body"))
	}
	if body.Type == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("type", "body"))
	}
	if body.Meta == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("meta", "body"))
	}
	if body.Project != nil {
		if err2 := ValidateProjectSummaryResponseBody(body.Project); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if body.Station != nil {
		if err2 := ValidateStationSummaryResponseBody(body.Station); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateProjectSummaryResponseBody runs the validations defined on
// ProjectSummaryResponseBody
func ValidateProjectSummaryResponseBody(body *ProjectSummaryResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	return
}

// ValidateStationSummaryResponseBody runs the validations defined on
// StationSummaryResponseBody
func ValidateStationSummaryResponseBody(body *StationSummaryResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	return
}
