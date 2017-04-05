// Code generated by goagen v1.1.0, command line:
// $ main
//
// API "fieldkit": Application Media Types
//
// The content of this file is auto-generated, DO NOT MODIFY

package client

import (
	"github.com/goadesign/goa"
	"net/http"
	"unicode/utf8"
)

// ProjectAdministrator media type (default view)
//
// Identifier: application/vnd.app.administrator+json; view=default
type ProjectAdministrator struct {
	ProjectID int `form:"project_id" json:"project_id" xml:"project_id"`
	UserID    int `form:"user_id" json:"user_id" xml:"user_id"`
}

// Validate validates the ProjectAdministrator media type instance.
func (mt *ProjectAdministrator) Validate() (err error) {

	return
}

// DecodeProjectAdministrator decodes the ProjectAdministrator instance encoded in resp body.
func (c *Client) DecodeProjectAdministrator(resp *http.Response) (*ProjectAdministrator, error) {
	var decoded ProjectAdministrator
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// ProjectAdministratorCollection is the media type for an array of ProjectAdministrator (default view)
//
// Identifier: application/vnd.app.administrator+json; type=collection; view=default
type ProjectAdministratorCollection []*ProjectAdministrator

// Validate validates the ProjectAdministratorCollection media type instance.
func (mt ProjectAdministratorCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeProjectAdministratorCollection decodes the ProjectAdministratorCollection instance encoded in resp body.
func (c *Client) DecodeProjectAdministratorCollection(resp *http.Response) (ProjectAdministratorCollection, error) {
	var decoded ProjectAdministratorCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// ProjectAdministrators media type (default view)
//
// Identifier: application/vnd.app.administrators+json; view=default
type ProjectAdministrators struct {
	Administrators ProjectAdministratorCollection `form:"administrators" json:"administrators" xml:"administrators"`
}

// Validate validates the ProjectAdministrators media type instance.
func (mt *ProjectAdministrators) Validate() (err error) {
	if mt.Administrators == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "administrators"))
	}
	return
}

// DecodeProjectAdministrators decodes the ProjectAdministrators instance encoded in resp body.
func (c *Client) DecodeProjectAdministrators(resp *http.Response) (*ProjectAdministrators, error) {
	var decoded ProjectAdministrators
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Document media type (default view)
//
// Identifier: application/vnd.app.document+json; view=default
type Document struct {
	Data      interface{} `form:"data" json:"data" xml:"data"`
	ID        string      `form:"id" json:"id" xml:"id"`
	InputID   int         `form:"input_id" json:"input_id" xml:"input_id"`
	Location  *Point      `form:"location" json:"location" xml:"location"`
	TeamID    *int        `form:"team_id,omitempty" json:"team_id,omitempty" xml:"team_id,omitempty"`
	Timestamp int         `form:"timestamp" json:"timestamp" xml:"timestamp"`
	UserID    *int        `form:"user_id,omitempty" json:"user_id,omitempty" xml:"user_id,omitempty"`
}

// Validate validates the Document media type instance.
func (mt *Document) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}

	if mt.Location == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "location"))
	}

	if mt.Location != nil {
		if err2 := mt.Location.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// DecodeDocument decodes the Document instance encoded in resp body.
func (c *Client) DecodeDocument(resp *http.Response) (*Document, error) {
	var decoded Document
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DocumentCollection is the media type for an array of Document (default view)
//
// Identifier: application/vnd.app.document+json; type=collection; view=default
type DocumentCollection []*Document

// Validate validates the DocumentCollection media type instance.
func (mt DocumentCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeDocumentCollection decodes the DocumentCollection instance encoded in resp body.
func (c *Client) DecodeDocumentCollection(resp *http.Response) (DocumentCollection, error) {
	var decoded DocumentCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// Documents media type (default view)
//
// Identifier: application/vnd.app.documents+json; view=default
type Documents struct {
	Documents DocumentCollection `form:"documents" json:"documents" xml:"documents"`
}

// Validate validates the Documents media type instance.
func (mt *Documents) Validate() (err error) {
	if mt.Documents == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "documents"))
	}
	if err2 := mt.Documents.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// DecodeDocuments decodes the Documents instance encoded in resp body.
func (c *Client) DecodeDocuments(resp *http.Response) (*Documents, error) {
	var decoded Documents
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Expedition media type (default view)
//
// Identifier: application/vnd.app.expedition+json; view=default
type Expedition struct {
	Description string `form:"description" json:"description" xml:"description"`
	ID          int    `form:"id" json:"id" xml:"id"`
	Name        string `form:"name" json:"name" xml:"name"`
	Slug        string `form:"slug" json:"slug" xml:"slug"`
}

// Validate validates the Expedition media type instance.
func (mt *Expedition) Validate() (err error) {

	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Slug == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "slug"))
	}
	if mt.Description == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "description"))
	}
	if ok := goa.ValidatePattern(`^[[:alnum:]]+(-[[:alnum:]]+)*$`, mt.Slug); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.slug`, mt.Slug, `^[[:alnum:]]+(-[[:alnum:]]+)*$`))
	}
	if utf8.RuneCountInString(mt.Slug) > 40 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.slug`, mt.Slug, utf8.RuneCountInString(mt.Slug), 40, false))
	}
	return
}

// DecodeExpedition decodes the Expedition instance encoded in resp body.
func (c *Client) DecodeExpedition(resp *http.Response) (*Expedition, error) {
	var decoded Expedition
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// ExpeditionCollection is the media type for an array of Expedition (default view)
//
// Identifier: application/vnd.app.expedition+json; type=collection; view=default
type ExpeditionCollection []*Expedition

// Validate validates the ExpeditionCollection media type instance.
func (mt ExpeditionCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeExpeditionCollection decodes the ExpeditionCollection instance encoded in resp body.
func (c *Client) DecodeExpeditionCollection(resp *http.Response) (ExpeditionCollection, error) {
	var decoded ExpeditionCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// Expeditions media type (default view)
//
// Identifier: application/vnd.app.expeditions+json; view=default
type Expeditions struct {
	Expeditions ExpeditionCollection `form:"expeditions" json:"expeditions" xml:"expeditions"`
}

// Validate validates the Expeditions media type instance.
func (mt *Expeditions) Validate() (err error) {
	if mt.Expeditions == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "expeditions"))
	}
	if err2 := mt.Expeditions.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// DecodeExpeditions decodes the Expeditions instance encoded in resp body.
func (c *Client) DecodeExpeditions(resp *http.Response) (*Expeditions, error) {
	var decoded Expeditions
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// FieldkitInput media type (default view)
//
// Identifier: application/vnd.app.fieldkit_input+json; view=default
type FieldkitInput struct {
	ExpeditionID int    `form:"expedition_id" json:"expedition_id" xml:"expedition_id"`
	ID           int    `form:"id" json:"id" xml:"id"`
	Name         string `form:"name" json:"name" xml:"name"`
	TeamID       *int   `form:"team_id,omitempty" json:"team_id,omitempty" xml:"team_id,omitempty"`
	UserID       *int   `form:"user_id,omitempty" json:"user_id,omitempty" xml:"user_id,omitempty"`
}

// Validate validates the FieldkitInput media type instance.
func (mt *FieldkitInput) Validate() (err error) {

	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	return
}

// DecodeFieldkitInput decodes the FieldkitInput instance encoded in resp body.
func (c *Client) DecodeFieldkitInput(resp *http.Response) (*FieldkitInput, error) {
	var decoded FieldkitInput
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// FieldkitInputCollection is the media type for an array of FieldkitInput (default view)
//
// Identifier: application/vnd.app.fieldkit_input+json; type=collection; view=default
type FieldkitInputCollection []*FieldkitInput

// Validate validates the FieldkitInputCollection media type instance.
func (mt FieldkitInputCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeFieldkitInputCollection decodes the FieldkitInputCollection instance encoded in resp body.
func (c *Client) DecodeFieldkitInputCollection(resp *http.Response) (FieldkitInputCollection, error) {
	var decoded FieldkitInputCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// FieldkitBinary media type (default view)
//
// Identifier: application/vnd.app.fieldkit_input_binary+json; view=default
type FieldkitBinary struct {
	Fields    []string          `form:"fields" json:"fields" xml:"fields"`
	ID        int               `form:"id" json:"id" xml:"id"`
	InputID   int               `form:"input_id" json:"input_id" xml:"input_id"`
	Latitude  *string           `form:"latitude,omitempty" json:"latitude,omitempty" xml:"latitude,omitempty"`
	Longitude *string           `form:"longitude,omitempty" json:"longitude,omitempty" xml:"longitude,omitempty"`
	Mapper    map[string]string `form:"mapper" json:"mapper" xml:"mapper"`
	SchemaID  int               `form:"schema_id" json:"schema_id" xml:"schema_id"`
}

// Validate validates the FieldkitBinary media type instance.
func (mt *FieldkitBinary) Validate() (err error) {

	if mt.Fields == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "fields"))
	}
	if mt.Mapper == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "mapper"))
	}
	for _, e := range mt.Fields {
		if !(e == "varint" || e == "uvarint" || e == "float32" || e == "float64") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.fields[*]`, e, []interface{}{"varint", "uvarint", "float32", "float64"}))
		}
	}
	return
}

// DecodeFieldkitBinary decodes the FieldkitBinary instance encoded in resp body.
func (c *Client) DecodeFieldkitBinary(resp *http.Response) (*FieldkitBinary, error) {
	var decoded FieldkitBinary
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// FieldkitInputs media type (default view)
//
// Identifier: application/vnd.app.fieldkit_inputs+json; view=default
type FieldkitInputs struct {
	FieldkitInputs FieldkitInputCollection `form:"fieldkit_inputs" json:"fieldkit_inputs" xml:"fieldkit_inputs"`
}

// Validate validates the FieldkitInputs media type instance.
func (mt *FieldkitInputs) Validate() (err error) {
	if mt.FieldkitInputs == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "fieldkit_inputs"))
	}
	if err2 := mt.FieldkitInputs.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// DecodeFieldkitInputs decodes the FieldkitInputs instance encoded in resp body.
func (c *Client) DecodeFieldkitInputs(resp *http.Response) (*FieldkitInputs, error) {
	var decoded FieldkitInputs
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Input media type (default view)
//
// Identifier: application/vnd.app.input+json; view=default
type Input struct {
	ExpeditionID int    `form:"expedition_id" json:"expedition_id" xml:"expedition_id"`
	ID           int    `form:"id" json:"id" xml:"id"`
	Name         string `form:"name" json:"name" xml:"name"`
	TeamID       *int   `form:"team_id,omitempty" json:"team_id,omitempty" xml:"team_id,omitempty"`
	UserID       *int   `form:"user_id,omitempty" json:"user_id,omitempty" xml:"user_id,omitempty"`
}

// Validate validates the Input media type instance.
func (mt *Input) Validate() (err error) {

	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	return
}

// DecodeInput decodes the Input instance encoded in resp body.
func (c *Client) DecodeInput(resp *http.Response) (*Input, error) {
	var decoded Input
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// InputToken media type (default view)
//
// Identifier: application/vnd.app.input_token+json; view=default
type InputToken struct {
	ExpeditionID int    `form:"expedition_id" json:"expedition_id" xml:"expedition_id"`
	ID           int    `form:"id" json:"id" xml:"id"`
	Token        string `form:"token" json:"token" xml:"token"`
}

// Validate validates the InputToken media type instance.
func (mt *InputToken) Validate() (err error) {

	if mt.Token == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "token"))
	}

	return
}

// DecodeInputToken decodes the InputToken instance encoded in resp body.
func (c *Client) DecodeInputToken(resp *http.Response) (*InputToken, error) {
	var decoded InputToken
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// InputTokenCollection is the media type for an array of InputToken (default view)
//
// Identifier: application/vnd.app.input_token+json; type=collection; view=default
type InputTokenCollection []*InputToken

// Validate validates the InputTokenCollection media type instance.
func (mt InputTokenCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeInputTokenCollection decodes the InputTokenCollection instance encoded in resp body.
func (c *Client) DecodeInputTokenCollection(resp *http.Response) (InputTokenCollection, error) {
	var decoded InputTokenCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// InputTokens media type (default view)
//
// Identifier: application/vnd.app.input_tokens+json; view=default
type InputTokens struct {
	InputTokens InputTokenCollection `form:"input_tokens" json:"input_tokens" xml:"input_tokens"`
}

// Validate validates the InputTokens media type instance.
func (mt *InputTokens) Validate() (err error) {
	if mt.InputTokens == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "input_tokens"))
	}
	if err2 := mt.InputTokens.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// DecodeInputTokens decodes the InputTokens instance encoded in resp body.
func (c *Client) DecodeInputTokens(resp *http.Response) (*InputTokens, error) {
	var decoded InputTokens
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Inputs media type (default view)
//
// Identifier: application/vnd.app.inputs+json; view=default
type Inputs struct {
	FieldkitInputs       FieldkitInputCollection       `form:"fieldkit_inputs,omitempty" json:"fieldkit_inputs,omitempty" xml:"fieldkit_inputs,omitempty"`
	TwitterAccountInputs TwitterAccountInputCollection `form:"twitter_account_inputs,omitempty" json:"twitter_account_inputs,omitempty" xml:"twitter_account_inputs,omitempty"`
}

// Validate validates the Inputs media type instance.
func (mt *Inputs) Validate() (err error) {
	if err2 := mt.FieldkitInputs.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	if err2 := mt.TwitterAccountInputs.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// DecodeInputs decodes the Inputs instance encoded in resp body.
func (c *Client) DecodeInputs(resp *http.Response) (*Inputs, error) {
	var decoded Inputs
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Location media type (default view)
//
// Identifier: application/vnd.app.location+json; view=default
type Location struct {
	Location string `form:"location" json:"location" xml:"location"`
}

// Validate validates the Location media type instance.
func (mt *Location) Validate() (err error) {
	if mt.Location == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "location"))
	}
	if err2 := goa.ValidateFormat(goa.FormatURI, mt.Location); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`response.location`, mt.Location, goa.FormatURI, err2))
	}
	return
}

// DecodeLocation decodes the Location instance encoded in resp body.
func (c *Client) DecodeLocation(resp *http.Response) (*Location, error) {
	var decoded Location
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// TeamMember media type (default view)
//
// Identifier: application/vnd.app.member+json; view=default
type TeamMember struct {
	Role   string `form:"role" json:"role" xml:"role"`
	TeamID int    `form:"team_id" json:"team_id" xml:"team_id"`
	UserID int    `form:"user_id" json:"user_id" xml:"user_id"`
}

// Validate validates the TeamMember media type instance.
func (mt *TeamMember) Validate() (err error) {

	if mt.Role == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "role"))
	}
	return
}

// DecodeTeamMember decodes the TeamMember instance encoded in resp body.
func (c *Client) DecodeTeamMember(resp *http.Response) (*TeamMember, error) {
	var decoded TeamMember
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// TeamMemberCollection is the media type for an array of TeamMember (default view)
//
// Identifier: application/vnd.app.member+json; type=collection; view=default
type TeamMemberCollection []*TeamMember

// Validate validates the TeamMemberCollection media type instance.
func (mt TeamMemberCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeTeamMemberCollection decodes the TeamMemberCollection instance encoded in resp body.
func (c *Client) DecodeTeamMemberCollection(resp *http.Response) (TeamMemberCollection, error) {
	var decoded TeamMemberCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// TeamMembers media type (default view)
//
// Identifier: application/vnd.app.members+json; view=default
type TeamMembers struct {
	Members TeamMemberCollection `form:"members" json:"members" xml:"members"`
}

// Validate validates the TeamMembers media type instance.
func (mt *TeamMembers) Validate() (err error) {
	if mt.Members == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "members"))
	}
	if err2 := mt.Members.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// DecodeTeamMembers decodes the TeamMembers instance encoded in resp body.
func (c *Client) DecodeTeamMembers(resp *http.Response) (*TeamMembers, error) {
	var decoded TeamMembers
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Point media type (default view)
//
// Identifier: application/vnd.app.point+json; view=default
type Point struct {
	Coordinates []float64 `form:"coordinates" json:"coordinates" xml:"coordinates"`
	Type        string    `form:"type" json:"type" xml:"type"`
}

// Validate validates the Point media type instance.
func (mt *Point) Validate() (err error) {
	if mt.Type == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "type"))
	}
	if mt.Coordinates == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "coordinates"))
	}
	if len(mt.Coordinates) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.coordinates`, mt.Coordinates, len(mt.Coordinates), 2, true))
	}
	if len(mt.Coordinates) > 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.coordinates`, mt.Coordinates, len(mt.Coordinates), 2, false))
	}
	if !(mt.Type == "Point") {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.type`, mt.Type, []interface{}{"Point"}))
	}
	return
}

// DecodePoint decodes the Point instance encoded in resp body.
func (c *Client) DecodePoint(resp *http.Response) (*Point, error) {
	var decoded Point
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Project media type (default view)
//
// Identifier: application/vnd.app.project+json; view=default
type Project struct {
	Description string `form:"description" json:"description" xml:"description"`
	ID          int    `form:"id" json:"id" xml:"id"`
	Name        string `form:"name" json:"name" xml:"name"`
	Slug        string `form:"slug" json:"slug" xml:"slug"`
}

// Validate validates the Project media type instance.
func (mt *Project) Validate() (err error) {

	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Slug == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "slug"))
	}
	if mt.Description == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "description"))
	}
	if ok := goa.ValidatePattern(`^[[:alnum:]]+(-[[:alnum:]]+)*$`, mt.Slug); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.slug`, mt.Slug, `^[[:alnum:]]+(-[[:alnum:]]+)*$`))
	}
	if utf8.RuneCountInString(mt.Slug) > 40 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.slug`, mt.Slug, utf8.RuneCountInString(mt.Slug), 40, false))
	}
	return
}

// DecodeProject decodes the Project instance encoded in resp body.
func (c *Client) DecodeProject(resp *http.Response) (*Project, error) {
	var decoded Project
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// ProjectCollection is the media type for an array of Project (default view)
//
// Identifier: application/vnd.app.project+json; type=collection; view=default
type ProjectCollection []*Project

// Validate validates the ProjectCollection media type instance.
func (mt ProjectCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeProjectCollection decodes the ProjectCollection instance encoded in resp body.
func (c *Client) DecodeProjectCollection(resp *http.Response) (ProjectCollection, error) {
	var decoded ProjectCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// Projects media type (default view)
//
// Identifier: application/vnd.app.projects+json; view=default
type Projects struct {
	Projects ProjectCollection `form:"projects" json:"projects" xml:"projects"`
}

// Validate validates the Projects media type instance.
func (mt *Projects) Validate() (err error) {
	if mt.Projects == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "projects"))
	}
	if err2 := mt.Projects.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// DecodeProjects decodes the Projects instance encoded in resp body.
func (c *Client) DecodeProjects(resp *http.Response) (*Projects, error) {
	var decoded Projects
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Schema media type (default view)
//
// Identifier: application/vnd.app.schema+json; view=default
type Schema struct {
	ID         int         `form:"id" json:"id" xml:"id"`
	JSONSchema interface{} `form:"json_schema" json:"json_schema" xml:"json_schema"`
	ProjectID  *int        `form:"project_id,omitempty" json:"project_id,omitempty" xml:"project_id,omitempty"`
}

// Validate validates the Schema media type instance.
func (mt *Schema) Validate() (err error) {

	return
}

// DecodeSchema decodes the Schema instance encoded in resp body.
func (c *Client) DecodeSchema(resp *http.Response) (*Schema, error) {
	var decoded Schema
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// SchemaCollection is the media type for an array of Schema (default view)
//
// Identifier: application/vnd.app.schema+json; type=collection; view=default
type SchemaCollection []*Schema

// Validate validates the SchemaCollection media type instance.
func (mt SchemaCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeSchemaCollection decodes the SchemaCollection instance encoded in resp body.
func (c *Client) DecodeSchemaCollection(resp *http.Response) (SchemaCollection, error) {
	var decoded SchemaCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// Schemas media type (default view)
//
// Identifier: application/vnd.app.schemas+json; view=default
type Schemas struct {
	Schemas SchemaCollection `form:"schemas,omitempty" json:"schemas,omitempty" xml:"schemas,omitempty"`
}

// DecodeSchemas decodes the Schemas instance encoded in resp body.
func (c *Client) DecodeSchemas(resp *http.Response) (*Schemas, error) {
	var decoded Schemas
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Team media type (default view)
//
// Identifier: application/vnd.app.team+json; view=default
type Team struct {
	Description string `form:"description" json:"description" xml:"description"`
	ID          int    `form:"id" json:"id" xml:"id"`
	Name        string `form:"name" json:"name" xml:"name"`
	Slug        string `form:"slug" json:"slug" xml:"slug"`
}

// Validate validates the Team media type instance.
func (mt *Team) Validate() (err error) {

	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Slug == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "slug"))
	}
	if mt.Description == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "description"))
	}
	if ok := goa.ValidatePattern(`\S`, mt.Name); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.name`, mt.Name, `\S`))
	}
	if utf8.RuneCountInString(mt.Name) > 256 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, mt.Name, utf8.RuneCountInString(mt.Name), 256, false))
	}
	if ok := goa.ValidatePattern(`^[[:alnum:]]+(-[[:alnum:]]+)*$`, mt.Slug); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.slug`, mt.Slug, `^[[:alnum:]]+(-[[:alnum:]]+)*$`))
	}
	if utf8.RuneCountInString(mt.Slug) > 40 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.slug`, mt.Slug, utf8.RuneCountInString(mt.Slug), 40, false))
	}
	return
}

// DecodeTeam decodes the Team instance encoded in resp body.
func (c *Client) DecodeTeam(resp *http.Response) (*Team, error) {
	var decoded Team
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// TeamCollection is the media type for an array of Team (default view)
//
// Identifier: application/vnd.app.team+json; type=collection; view=default
type TeamCollection []*Team

// Validate validates the TeamCollection media type instance.
func (mt TeamCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeTeamCollection decodes the TeamCollection instance encoded in resp body.
func (c *Client) DecodeTeamCollection(resp *http.Response) (TeamCollection, error) {
	var decoded TeamCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// Teams media type (default view)
//
// Identifier: application/vnd.app.teams+json; view=default
type Teams struct {
	Teams TeamCollection `form:"teams" json:"teams" xml:"teams"`
}

// Validate validates the Teams media type instance.
func (mt *Teams) Validate() (err error) {
	if mt.Teams == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "teams"))
	}
	if err2 := mt.Teams.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// DecodeTeams decodes the Teams instance encoded in resp body.
func (c *Client) DecodeTeams(resp *http.Response) (*Teams, error) {
	var decoded Teams
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// TwitterAccountInput media type (default view)
//
// Identifier: application/vnd.app.twitter_account_input+json; view=default
type TwitterAccountInput struct {
	ExpeditionID     int    `form:"expedition_id" json:"expedition_id" xml:"expedition_id"`
	ID               int    `form:"id" json:"id" xml:"id"`
	Name             string `form:"name" json:"name" xml:"name"`
	ScreenName       string `form:"screen_name" json:"screen_name" xml:"screen_name"`
	TeamID           *int   `form:"team_id,omitempty" json:"team_id,omitempty" xml:"team_id,omitempty"`
	TwitterAccountID int    `form:"twitter_account_id" json:"twitter_account_id" xml:"twitter_account_id"`
	UserID           *int   `form:"user_id,omitempty" json:"user_id,omitempty" xml:"user_id,omitempty"`
}

// Validate validates the TwitterAccountInput media type instance.
func (mt *TwitterAccountInput) Validate() (err error) {

	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}

	if mt.ScreenName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "screen_name"))
	}
	return
}

// DecodeTwitterAccountInput decodes the TwitterAccountInput instance encoded in resp body.
func (c *Client) DecodeTwitterAccountInput(resp *http.Response) (*TwitterAccountInput, error) {
	var decoded TwitterAccountInput
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// TwitterAccountInputCollection is the media type for an array of TwitterAccountInput (default view)
//
// Identifier: application/vnd.app.twitter_account_input+json; type=collection; view=default
type TwitterAccountInputCollection []*TwitterAccountInput

// Validate validates the TwitterAccountInputCollection media type instance.
func (mt TwitterAccountInputCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeTwitterAccountInputCollection decodes the TwitterAccountInputCollection instance encoded in resp body.
func (c *Client) DecodeTwitterAccountInputCollection(resp *http.Response) (TwitterAccountInputCollection, error) {
	var decoded TwitterAccountInputCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// TwitterAccountInputs media type (default view)
//
// Identifier: application/vnd.app.twitter_account_intputs+json; view=default
type TwitterAccountInputs struct {
	TwitterAccountInputs TwitterAccountInputCollection `form:"twitter_account_inputs" json:"twitter_account_inputs" xml:"twitter_account_inputs"`
}

// Validate validates the TwitterAccountInputs media type instance.
func (mt *TwitterAccountInputs) Validate() (err error) {
	if mt.TwitterAccountInputs == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "twitter_account_inputs"))
	}
	if err2 := mt.TwitterAccountInputs.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// DecodeTwitterAccountInputs decodes the TwitterAccountInputs instance encoded in resp body.
func (c *Client) DecodeTwitterAccountInputs(resp *http.Response) (*TwitterAccountInputs, error) {
	var decoded TwitterAccountInputs
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// User media type (default view)
//
// Identifier: application/vnd.app.user+json; view=default
type User struct {
	Bio      string `form:"bio" json:"bio" xml:"bio"`
	Email    string `form:"email" json:"email" xml:"email"`
	ID       int    `form:"id" json:"id" xml:"id"`
	Name     string `form:"name" json:"name" xml:"name"`
	Username string `form:"username" json:"username" xml:"username"`
}

// Validate validates the User media type instance.
func (mt *User) Validate() (err error) {

	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Username == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "username"))
	}
	if mt.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "email"))
	}
	if mt.Bio == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "bio"))
	}
	if err2 := goa.ValidateFormat(goa.FormatEmail, mt.Email); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`response.email`, mt.Email, goa.FormatEmail, err2))
	}
	if ok := goa.ValidatePattern(`\S`, mt.Name); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.name`, mt.Name, `\S`))
	}
	if utf8.RuneCountInString(mt.Name) > 256 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, mt.Name, utf8.RuneCountInString(mt.Name), 256, false))
	}
	if ok := goa.ValidatePattern(`^[[:alnum:]]+(-[[:alnum:]]+)*$`, mt.Username); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.username`, mt.Username, `^[[:alnum:]]+(-[[:alnum:]]+)*$`))
	}
	if utf8.RuneCountInString(mt.Username) > 40 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.username`, mt.Username, utf8.RuneCountInString(mt.Username), 40, false))
	}
	return
}

// DecodeUser decodes the User instance encoded in resp body.
func (c *Client) DecodeUser(resp *http.Response) (*User, error) {
	var decoded User
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// UserCollection is the media type for an array of User (default view)
//
// Identifier: application/vnd.app.user+json; type=collection; view=default
type UserCollection []*User

// Validate validates the UserCollection media type instance.
func (mt UserCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeUserCollection decodes the UserCollection instance encoded in resp body.
func (c *Client) DecodeUserCollection(resp *http.Response) (UserCollection, error) {
	var decoded UserCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// Users media type (default view)
//
// Identifier: application/vnd.app.users+json; view=default
type Users struct {
	Users UserCollection `form:"users" json:"users" xml:"users"`
}

// Validate validates the Users media type instance.
func (mt *Users) Validate() (err error) {
	if mt.Users == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "users"))
	}
	if err2 := mt.Users.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// DecodeUsers decodes the Users instance encoded in resp body.
func (c *Client) DecodeUsers(resp *http.Response) (*Users, error) {
	var decoded Users
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeErrorResponse decodes the ErrorResponse instance encoded in resp body.
func (c *Client) DecodeErrorResponse(resp *http.Response) (*goa.ErrorResponse, error) {
	var decoded goa.ErrorResponse
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
