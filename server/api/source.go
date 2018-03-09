package api

import (
	"github.com/goadesign/goa"

	"github.com/fieldkit/cloud/server/api/app"
	"github.com/fieldkit/cloud/server/backend"
	"github.com/fieldkit/cloud/server/data"
)

func SourceType(source *data.Source) *app.Source {
	sourceType := &app.Source{
		ID:           int(source.ID),
		ExpeditionID: int(source.ExpeditionID),
		Name:         source.Name,
		Active:       &source.Active,
	}

	if source.TeamID != nil {
		teamID := int(*source.TeamID)
		sourceType.TeamID = &teamID
	}

	if source.UserID != nil {
		userID := int(*source.UserID)
		sourceType.UserID = &userID
	}

	return sourceType
}

func GeometryClusterSummaryType(s *backend.GeometryClusterSummary) *app.GeometryClusterSummary {
	return &app.GeometryClusterSummary{
		ID:               s.ID,
		NumberOfFeatures: s.NumberOfFeatures,
		StartTime:        s.StartTime,
		EndTime:          s.EndTime,
		Centroid:         s.Centroid.Coordinates(),
		Radius:           s.Radius,
		Geometry:         s.Geometry.Coordinates(),
	}
}

func GeometryClusterSummariesType(s []*backend.GeometryClusterSummary) []*app.GeometryClusterSummary {
	summaries := make([]*app.GeometryClusterSummary, len(s))
	for i, summary := range s {
		summaries[i] = GeometryClusterSummaryType(summary)
	}
	return summaries
}

func SourceSummaryType(source *data.DeviceSource, spatial, temporal []*backend.GeometryClusterSummary) *app.SourceSummary {
	return &app.SourceSummary{
		ID:       int(source.ID),
		Name:     source.Name,
		Temporal: GeometryClusterSummariesType(temporal),
		Spatial:  GeometryClusterSummariesType(spatial),
	}
}

type SourceControllerOptions struct {
	Backend *backend.Backend
}

type SourceController struct {
	*goa.Controller
	options SourceControllerOptions
}

func NewSourceController(service *goa.Service, options SourceControllerOptions) *SourceController {
	return &SourceController{
		Controller: service.NewController("SourceController"),
		options:    options,
	}
}

func (c *SourceController) Update(ctx *app.UpdateSourceContext) error {
	source, err := c.options.Backend.Source(ctx, int32(ctx.SourceID))
	if err != nil {
		return err
	}

	source.Name = ctx.Payload.Name

	if ctx.Payload.TeamID != nil {
		teamID := int32(*ctx.Payload.TeamID)
		source.TeamID = &teamID
	} else {
		source.TeamID = nil
	}

	if ctx.Payload.UserID != nil {
		userID := int32(*ctx.Payload.UserID)
		source.UserID = &userID
	} else {
		source.UserID = nil
	}

	if err := c.options.Backend.UpdateSource(ctx, source); err != nil {
		return err
	}

	return ctx.OK(SourceType(source))
}

func (c *SourceController) List(ctx *app.ListSourceContext) error {
	twitterAccountSources, err := c.options.Backend.ListTwitterAccountSources(ctx, ctx.Project, ctx.Expedition)
	if err != nil {
		return err
	}

	deviceSources, err := c.options.Backend.ListDeviceSources(ctx, ctx.Project, ctx.Expedition)
	if err != nil {
		return err
	}

	return ctx.OK(&app.Sources{
		TwitterAccountSources: TwitterAccountSourcesType(twitterAccountSources).TwitterAccountSources,
		DeviceSources:         DeviceSourcesType(deviceSources).DeviceSources,
	})
}

func (c *SourceController) ListID(ctx *app.ListIDSourceContext) error {
	deviceSource, err := c.options.Backend.GetDeviceSourceByID(ctx, int32(ctx.SourceID))
	if err != nil {
		return err
	}
	summary, err := c.options.Backend.FeatureSummaryBySourceID(ctx, ctx.SourceID)
	if err != nil {
		return err
	}
	return ctx.OKPublic(DeviceSourcePublicType(deviceSource, summary))
}

func (c *SourceController) SummaryByID(ctx *app.SummaryByIDSourceContext) error {
	deviceSource, err := c.options.Backend.GetDeviceSourceByID(ctx, int32(ctx.SourceID))
	if err != nil {
		return err
	}

	spatial, err := c.options.Backend.SpatialClustersBySourceID(ctx, ctx.SourceID)
	if err != nil {
		return err
	}

	temporal, err := c.options.Backend.TemporalClustersBySourceID(ctx, ctx.SourceID)
	if err != nil {
		return err
	}

	return ctx.OK(SourceSummaryType(deviceSource, spatial, temporal))
}

func (c *SourceController) ListExpeditionID(ctx *app.ListExpeditionIDSourceContext) error {
	twitterAccountSources, err := c.options.Backend.ListTwitterAccountSourcesByExpeditionID(ctx, int32(ctx.ExpeditionID))
	if err != nil {
		return err
	}

	deviceSources, err := c.options.Backend.ListDeviceSourcesByExpeditionID(ctx, int32(ctx.ExpeditionID))
	if err != nil {
		return err
	}

	return ctx.OK(&app.Sources{
		TwitterAccountSources: TwitterAccountSourcesType(twitterAccountSources).TwitterAccountSources,
		DeviceSources:         DeviceSourcesType(deviceSources).DeviceSources,
	})
}
