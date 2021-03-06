// Code generated by goa v3.1.2, DO NOT EDIT.
//
// data views
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// DeviceDataSummaryResponse is the viewed result type that is projected based
// on a view.
type DeviceDataSummaryResponse struct {
	// Type to project
	Projected *DeviceDataSummaryResponseView
	// View to render
	View string
}

// DeviceDataSummaryResponseView is a type that runs validations on a projected
// type.
type DeviceDataSummaryResponseView struct {
	Provisions DeviceProvisionSummaryCollectionView
}

// DeviceProvisionSummaryCollectionView is a type that runs validations on a
// projected type.
type DeviceProvisionSummaryCollectionView []*DeviceProvisionSummaryView

// DeviceProvisionSummaryView is a type that runs validations on a projected
// type.
type DeviceProvisionSummaryView struct {
	Generation *string
	Created    *int64
	Updated    *int64
	Meta       *DeviceMetaSummaryView
	Data       *DeviceDataSummaryView
}

// DeviceMetaSummaryView is a type that runs validations on a projected type.
type DeviceMetaSummaryView struct {
	Size  *int64
	First *int64
	Last  *int64
}

// DeviceDataSummaryView is a type that runs validations on a projected type.
type DeviceDataSummaryView struct {
	Size  *int64
	First *int64
	Last  *int64
}

var (
	// DeviceDataSummaryResponseMap is a map of attribute names in result type
	// DeviceDataSummaryResponse indexed by view name.
	DeviceDataSummaryResponseMap = map[string][]string{
		"default": []string{
			"provisions",
		},
	}
	// DeviceProvisionSummaryCollectionMap is a map of attribute names in result
	// type DeviceProvisionSummaryCollection indexed by view name.
	DeviceProvisionSummaryCollectionMap = map[string][]string{
		"default": []string{
			"generation",
			"created",
			"updated",
			"meta",
			"data",
		},
	}
	// DeviceProvisionSummaryMap is a map of attribute names in result type
	// DeviceProvisionSummary indexed by view name.
	DeviceProvisionSummaryMap = map[string][]string{
		"default": []string{
			"generation",
			"created",
			"updated",
			"meta",
			"data",
		},
	}
	// DeviceMetaSummaryMap is a map of attribute names in result type
	// DeviceMetaSummary indexed by view name.
	DeviceMetaSummaryMap = map[string][]string{
		"default": []string{
			"size",
			"first",
			"last",
		},
	}
	// DeviceDataSummaryMap is a map of attribute names in result type
	// DeviceDataSummary indexed by view name.
	DeviceDataSummaryMap = map[string][]string{
		"default": []string{
			"size",
			"first",
			"last",
		},
	}
)

// ValidateDeviceDataSummaryResponse runs the validations defined on the viewed
// result type DeviceDataSummaryResponse.
func ValidateDeviceDataSummaryResponse(result *DeviceDataSummaryResponse) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateDeviceDataSummaryResponseView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateDeviceDataSummaryResponseView runs the validations defined on
// DeviceDataSummaryResponseView using the "default" view.
func ValidateDeviceDataSummaryResponseView(result *DeviceDataSummaryResponseView) (err error) {

	if result.Provisions != nil {
		if err2 := ValidateDeviceProvisionSummaryCollectionView(result.Provisions); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateDeviceProvisionSummaryCollectionView runs the validations defined on
// DeviceProvisionSummaryCollectionView using the "default" view.
func ValidateDeviceProvisionSummaryCollectionView(result DeviceProvisionSummaryCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateDeviceProvisionSummaryView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateDeviceProvisionSummaryView runs the validations defined on
// DeviceProvisionSummaryView using the "default" view.
func ValidateDeviceProvisionSummaryView(result *DeviceProvisionSummaryView) (err error) {
	if result.Generation == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("generation", "result"))
	}
	if result.Created == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("created", "result"))
	}
	if result.Updated == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("updated", "result"))
	}
	if result.Meta != nil {
		if err2 := ValidateDeviceMetaSummaryView(result.Meta); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if result.Data != nil {
		if err2 := ValidateDeviceDataSummaryView(result.Data); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateDeviceMetaSummaryView runs the validations defined on
// DeviceMetaSummaryView using the "default" view.
func ValidateDeviceMetaSummaryView(result *DeviceMetaSummaryView) (err error) {
	if result.Size == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("size", "result"))
	}
	if result.First == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("first", "result"))
	}
	if result.Last == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("last", "result"))
	}
	return
}

// ValidateDeviceDataSummaryView runs the validations defined on
// DeviceDataSummaryView using the "default" view.
func ValidateDeviceDataSummaryView(result *DeviceDataSummaryView) (err error) {
	if result.Size == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("size", "result"))
	}
	if result.First == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("first", "result"))
	}
	if result.Last == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("last", "result"))
	}
	return
}
