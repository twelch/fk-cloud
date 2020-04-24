// Code generated by goa v3.1.1, DO NOT EDIT.
//
// activity views
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// StationActivityPage is the viewed result type that is projected based on a
// view.
type StationActivityPage struct {
	// Type to project
	Projected *StationActivityPageView
	// View to render
	View string
}

// ProjectActivityPage is the viewed result type that is projected based on a
// view.
type ProjectActivityPage struct {
	// Type to project
	Projected *ProjectActivityPageView
	// View to render
	View string
}

// StationActivityPageView is a type that runs validations on a projected type.
type StationActivityPageView struct {
	Activities ActivityEntryCollectionView
	Total      *int32
	Page       *int32
}

// ActivityEntryCollectionView is a type that runs validations on a projected
// type.
type ActivityEntryCollectionView []*ActivityEntryView

// ActivityEntryView is a type that runs validations on a projected type.
type ActivityEntryView struct {
	ID        *int64
	Key       *string
	Project   *ProjectSummaryView
	Station   *StationSummaryView
	CreatedAt *int64
	Type      *string
	Meta      interface{}
}

// ProjectSummaryView is a type that runs validations on a projected type.
type ProjectSummaryView struct {
	ID   *int64
	Name *string
}

// StationSummaryView is a type that runs validations on a projected type.
type StationSummaryView struct {
	ID   *int64
	Name *string
}

// ProjectActivityPageView is a type that runs validations on a projected type.
type ProjectActivityPageView struct {
	Activities ActivityEntryCollectionView
	Total      *int32
	Page       *int32
}

var (
	// StationActivityPageMap is a map of attribute names in result type
	// StationActivityPage indexed by view name.
	StationActivityPageMap = map[string][]string{
		"default": []string{
			"activities",
			"total",
			"page",
		},
	}
	// ProjectActivityPageMap is a map of attribute names in result type
	// ProjectActivityPage indexed by view name.
	ProjectActivityPageMap = map[string][]string{
		"default": []string{
			"activities",
			"total",
			"page",
		},
	}
	// ActivityEntryCollectionMap is a map of attribute names in result type
	// ActivityEntryCollection indexed by view name.
	ActivityEntryCollectionMap = map[string][]string{
		"default": []string{
			"id",
			"key",
			"project",
			"station",
			"created_at",
			"type",
			"meta",
		},
	}
	// ActivityEntryMap is a map of attribute names in result type ActivityEntry
	// indexed by view name.
	ActivityEntryMap = map[string][]string{
		"default": []string{
			"id",
			"key",
			"project",
			"station",
			"created_at",
			"type",
			"meta",
		},
	}
)

// ValidateStationActivityPage runs the validations defined on the viewed
// result type StationActivityPage.
func ValidateStationActivityPage(result *StationActivityPage) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateStationActivityPageView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateProjectActivityPage runs the validations defined on the viewed
// result type ProjectActivityPage.
func ValidateProjectActivityPage(result *ProjectActivityPage) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateProjectActivityPageView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateStationActivityPageView runs the validations defined on
// StationActivityPageView using the "default" view.
func ValidateStationActivityPageView(result *StationActivityPageView) (err error) {
	if result.Total == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("total", "result"))
	}
	if result.Page == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("page", "result"))
	}
	if result.Activities != nil {
		if err2 := ValidateActivityEntryCollectionView(result.Activities); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateActivityEntryCollectionView runs the validations defined on
// ActivityEntryCollectionView using the "default" view.
func ValidateActivityEntryCollectionView(result ActivityEntryCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateActivityEntryView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateActivityEntryView runs the validations defined on ActivityEntryView
// using the "default" view.
func ValidateActivityEntryView(result *ActivityEntryView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Key == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("key", "result"))
	}
	if result.Project == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("project", "result"))
	}
	if result.Station == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("station", "result"))
	}
	if result.CreatedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("created_at", "result"))
	}
	if result.Type == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("type", "result"))
	}
	if result.Meta == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("meta", "result"))
	}
	if result.Project != nil {
		if err2 := ValidateProjectSummaryView(result.Project); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if result.Station != nil {
		if err2 := ValidateStationSummaryView(result.Station); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateProjectSummaryView runs the validations defined on
// ProjectSummaryView.
func ValidateProjectSummaryView(result *ProjectSummaryView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	return
}

// ValidateStationSummaryView runs the validations defined on
// StationSummaryView.
func ValidateStationSummaryView(result *StationSummaryView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	return
}

// ValidateProjectActivityPageView runs the validations defined on
// ProjectActivityPageView using the "default" view.
func ValidateProjectActivityPageView(result *ProjectActivityPageView) (err error) {
	if result.Total == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("total", "result"))
	}
	if result.Page == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("page", "result"))
	}
	if result.Activities != nil {
		if err2 := ValidateActivityEntryCollectionView(result.Activities); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}
