package api

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"strings"
	"time"

	"github.com/iancoleman/strcase"

	"goa.design/goa/v3/security"

	station "github.com/fieldkit/cloud/server/api/gen/station"

	"github.com/fieldkit/cloud/server/backend/repositories"
	"github.com/fieldkit/cloud/server/data"
)

type StationService struct {
	options *ControllerOptions
}

func NewStationService(ctx context.Context, options *ControllerOptions) *StationService {
	return &StationService{
		options: options,
	}
}

func (c *StationService) updateStation(ctx context.Context, station *data.Station, rawStatusPb *string) error {
	log := Logger(ctx).Sugar()

	sr, err := repositories.NewStationRepository(c.options.Database)
	if err != nil {
		return err
	}

	if rawStatusPb != nil {
		log.Infow("updating station from status", "station_id", station.ID)

		if err := station.UpdateFromStatus(ctx, *rawStatusPb); err != nil {
			log.Errorw("error updating from status", "error", err)
		}

		if err := sr.UpdateStationModelFromStatus(ctx, station, *rawStatusPb); err != nil {
			log.Errorw("error updating model status", "error", err)
		}

		if station.Location != nil && c.options.locations != nil {
			log.Infow("describing location", "station_id", station.ID)
			names, err := c.options.locations.Describe(ctx, station.Location)
			if err != nil {
				log.Errorw("error updating from location", "error", err)
			} else if names != nil {
				station.PlaceOther = names.OtherLandName
				station.PlaceNative = names.NativeLandName
			}
		}
	}

	station.UpdatedAt = time.Now()

	if err := sr.Update(ctx, station); err != nil {
		return err
	}

	return nil
}

func (c *StationService) Add(ctx context.Context, payload *station.AddPayload) (response *station.StationFull, err error) {
	log := Logger(ctx).Sugar()

	p, err := NewPermissions(ctx, c.options).Unwrap()
	if err != nil {
		return nil, err
	}

	deviceId, err := hex.DecodeString(payload.DeviceID)
	if err != nil {
		return nil, err
	}

	log.Infow("adding station", "device_id", payload.DeviceID)

	sr, err := repositories.NewStationRepository(c.options.Database)
	if err != nil {
		return nil, err
	}

	stations, err := sr.QueryStationsByDeviceID(ctx, deviceId)
	if err != nil {
		return nil, err
	}

	if len(stations) > 0 {
		existing := stations[0]

		if existing.OwnerID != p.UserID() {
			log.Infow("station conflict", "device_id", deviceId, "user_id", p.UserID(), "owner_id", existing.OwnerID, "station_id", existing.ID)
			return nil, station.MakeStationOwnerConflict(errors.New("station already registered"))
		}

		if payload.LocationName != nil {
			existing.LocationName = payload.LocationName
		}

		if err := c.updateStation(ctx, existing, payload.StatusPb); err != nil {
			return nil, err
		}

		return c.Get(ctx, &station.GetPayload{
			Auth: payload.Auth,
			ID:   existing.ID,
		})
	}

	adding := &data.Station{
		OwnerID:      p.UserID(),
		Name:         payload.Name,
		DeviceID:     deviceId,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		LocationName: payload.LocationName,
	}

	added, err := sr.Add(ctx, adding)
	if err != nil {
		return nil, err
	}

	if err := c.updateStation(ctx, added, payload.StatusPb); err != nil {
		return nil, err
	}

	pr := repositories.NewProjectRepository(c.options.Database)
	if err != nil {
		return nil, err
	}

	if err := pr.AddStationToDefaultProjectMaybe(ctx, adding); err != nil {
		return nil, err
	}

	return c.Get(ctx, &station.GetPayload{
		Auth: payload.Auth,
		ID:   adding.ID,
	})
}

func (c *StationService) Get(ctx context.Context, payload *station.GetPayload) (response *station.StationFull, err error) {
	p, err := NewPermissions(ctx, c.options).ForStationByID(int(payload.ID))
	if err != nil {
		return nil, err
	}

	if err := p.CanView(); err != nil {
		return nil, err
	}

	r, err := repositories.NewStationRepository(c.options.Database)
	if err != nil {
		return nil, err
	}

	sf, err := r.QueryStationFull(ctx, payload.ID)
	if err != nil {
		return nil, err
	}

	preciseLocation := p.UserID() == sf.Owner.ID

	return transformStationFull(c.options.signer, p, sf, preciseLocation)
}

func (c *StationService) Update(ctx context.Context, payload *station.UpdatePayload) (response *station.StationFull, err error) {
	sr, err := repositories.NewStationRepository(c.options.Database)
	if err != nil {
		return nil, err
	}

	updating, err := sr.QueryStationByID(ctx, payload.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, station.MakeNotFound(errors.New("station not found"))
		}
		return nil, err
	}

	p, err := NewPermissions(ctx, c.options).ForStation(updating)
	if err != nil {
		return nil, err
	}

	if err := p.CanModify(); err != nil {
		return nil, err
	}

	updating.Name = payload.Name
	if payload.LocationName != nil {
		updating.LocationName = payload.LocationName
	}

	if err := c.updateStation(ctx, updating, payload.StatusPb); err != nil {
		return nil, err
	}

	return c.Get(ctx, &station.GetPayload{
		Auth: payload.Auth,
		ID:   payload.ID,
	})
}

func (c *StationService) ListMine(ctx context.Context, payload *station.ListMinePayload) (response *station.StationsFull, err error) {
	p, err := NewPermissions(ctx, c.options).Unwrap()
	if err != nil {
		return nil, err
	}

	r, err := repositories.NewStationRepository(c.options.Database)
	if err != nil {
		return nil, err
	}

	sfs, err := r.QueryStationFullByOwnerID(ctx, p.UserID())
	if err != nil {
		return nil, err
	}

	stations, err := transformAllStationFull(c.options.signer, p, sfs, true)
	if err != nil {
		return nil, err
	}

	response = &station.StationsFull{
		Stations: stations,
	}

	return
}

func (c *StationService) ListProject(ctx context.Context, payload *station.ListProjectPayload) (response *station.StationsFull, err error) {
	pr := repositories.NewProjectRepository(c.options.Database)

	project, err := pr.QueryByID(ctx, payload.ID)
	if err != nil {
		return nil, err
	}

	preciseLocation := project.Privacy == data.Public
	p, err := NewPermissions(ctx, c.options).Unwrap()
	if err == nil {
		relationships, err := pr.QueryUserProjectRelationships(ctx, p.UserID())
		if err != nil {
			return nil, err
		}
		if _, ok := relationships[payload.ID]; ok {
			preciseLocation = true
		}
	}

	r, err := repositories.NewStationRepository(c.options.Database)
	if err != nil {
		return nil, err
	}

	sfs, err := r.QueryStationFullByProjectID(ctx, payload.ID)
	if err != nil {
		return nil, err
	}

	stations, err := transformAllStationFull(c.options.signer, p, sfs, preciseLocation)
	if err != nil {
		return nil, err
	}

	response = &station.StationsFull{
		Stations: stations,
	}

	return
}

func (c *StationService) ListAll(ctx context.Context, payload *station.ListAllPayload) (*station.PageOfStations, error) {
	sr, err := repositories.NewStationRepository(c.options.Database)
	if err != nil {
		return nil, err
	}

	qp := &repositories.EssentialQueryParams{
		Page:     0,
		PageSize: 20,
	}

	if payload.Page != nil {
		qp.Page = *payload.Page
	}
	if payload.PageSize != nil {
		qp.PageSize = *payload.PageSize
	}

	queried, err := sr.QueryEssentialStations(ctx, qp)
	if err != nil {
		return nil, err
	}

	stationsWm := make([]*station.EssentialStation, 0)

	for _, es := range queried.Stations {
		var lastIngestionAt *int64

		if es.LastIngestionAt != nil {
			jsEpoch := es.LastIngestionAt.Unix() * 1000
			lastIngestionAt = &jsEpoch
		}

		var location *station.StationLocation = nil
		if es.Location != nil {
			location = &station.StationLocation{
				Precise: []float64{es.Location.Longitude(), es.Location.Latitude()},
			}
		}

		var recordingStartedAt *int64
		if es.RecordingStartedAt != nil {
			unix := es.RecordingStartedAt.Unix() * 1000
			recordingStartedAt = &unix
		}

		stationsWm = append(stationsWm, &station.EssentialStation{
			ID:       es.ID,
			Name:     es.Name,
			DeviceID: hex.EncodeToString(es.DeviceID),
			Owner: &station.StationOwner{
				ID:   es.OwnerID,
				Name: es.OwnerName,
			},
			CreatedAt:          es.CreatedAt.Unix() * 1000,
			UpdatedAt:          es.UpdatedAt.Unix() * 1000,
			RecordingStartedAt: recordingStartedAt,
			MemoryUsed:         es.MemoryUsed,
			MemoryAvailable:    es.MemoryAvailable,
			FirmwareNumber:     es.FirmwareNumber,
			FirmwareTime:       es.FirmwareTime,
			Location:           location,
			LastIngestionAt:    lastIngestionAt,
		})
	}

	wm := &station.PageOfStations{
		Stations: stationsWm,
		Total:    queried.Total,
	}

	return wm, nil
}

func (c *StationService) Delete(ctx context.Context, payload *station.DeletePayload) error {
	sr, err := repositories.NewStationRepository(c.options.Database)
	if err != nil {
		return err
	}

	return sr.Delete(ctx, payload.StationID)
}

func (c *StationService) DownloadPhoto(ctx context.Context, payload *station.DownloadPhotoPayload) (*station.DownloadedPhoto, error) {
	x := uint(124)
	y := uint(100)

	allMedia := []*data.FieldNoteMedia{}
	if err := c.options.Database.SelectContext(ctx, &allMedia, `
		SELECT * FROM fieldkit.notes_media WHERE station_id = $1 ORDER BY created_at DESC
		`, payload.StationID); err != nil {
		return nil, err
	}

	if len(allMedia) == 0 {
		return defaultPhoto(payload.StationID)
	}

	media := allMedia[0]

	etag := quickHash(media.URL)
	if payload.Size != nil {
		etag += fmt.Sprintf(":%d", *payload.Size)
	}

	if payload.IfNoneMatch != nil {
		if *payload.IfNoneMatch == fmt.Sprintf(`"%s"`, etag) {
			return &station.DownloadedPhoto{
				ContentType: "image/jpeg",
				Etag:        etag,
				Body:        []byte{},
			}, nil
		}
	}

	mr := repositories.NewMediaRepository(c.options.MediaFiles)
	lm, err := mr.LoadByURL(ctx, media.URL)
	if err != nil {
		return nil, err
	}

	original, _, err := image.Decode(lm.Reader)
	if err != nil {
		log := Logger(ctx).Sugar()
		log.Warnw("image-error", "error", err, "station_id", payload.StationID)
		return defaultPhoto(payload.StationID)
	}

	data := []byte{}

	if payload.Size != nil {
		if media.ContentType == "image/jpeg" || media.ContentType == "image/png" {
			resized, err := resizeLoadedMedia(ctx, lm, uint(*payload.Size), 0)
			if err != nil {
				return nil, err
			}
			return &station.DownloadedPhoto{
				Length:      resized.Size,
				ContentType: resized.ContentType,
				Etag:        etag,
				Body:        resized.Data,
			}, nil
		}
	} else {
		cropped, err := smartCrop(original, x, y)
		if err != nil {
			return nil, err
		}

		options := jpeg.Options{
			Quality: 80,
		}

		buf := new(bytes.Buffer)
		if err := jpeg.Encode(buf, cropped, &options); err != nil {
			return nil, err
		}

		data = buf.Bytes()
	}

	if len(data) == 0 {
		log := Logger(ctx).Sugar()
		log.Warnw("empty-image", "station_id", payload.StationID)
		return defaultPhoto(payload.StationID)
	}

	return &station.DownloadedPhoto{
		Length:      int64(len(data)),
		ContentType: "image/jpeg",
		Etag:        etag,
		Body:        data,
	}, nil
}

func (s *StationService) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	return Authenticate(ctx, AuthAttempt{
		Token:        token,
		Scheme:       scheme,
		Key:          s.options.JWTHMACKey,
		NotFound:     func(m string) error { return station.MakeNotFound(errors.New(m)) },
		Unauthorized: func(m string) error { return station.MakeUnauthorized(errors.New(m)) },
		Forbidden:    func(m string) error { return station.MakeForbidden(errors.New(m)) },
	})
}

func transformUploads(from []*data.Ingestion) (to []*station.StationUpload) {
	to = make([]*station.StationUpload, 0, len(from))
	for _, v := range from {
		to = append(to, &station.StationUpload{
			ID:       v.ID,
			Time:     v.Time.Unix() * 1000,
			UploadID: v.UploadID,
			Size:     v.Size,
			Type:     v.Type,
			URL:      v.URL,
			Blocks:   v.Blocks.ToInt64Array(),
		})
	}
	return to
}

func transformReading(s *data.ModuleSensor) *station.SensorReading {
	if s.ReadingValue == nil {
		return nil
	}

	time := int64(0)
	if s.ReadingTime != nil {
		time = s.ReadingTime.Unix() * 1000
	}

	return &station.SensorReading{
		Last: float32(*s.ReadingValue),
		Time: time,
	}
}

func transformConfigurations(from *data.StationFull) (to []*station.StationConfiguration) {
	to = make([]*station.StationConfiguration, 0)
	for _, v := range from.Configurations {
		to = append(to, &station.StationConfiguration{
			ID:      v.ID,
			Time:    v.UpdatedAt.Unix() * 1000,
			Modules: transformModules(from, v.ID),
		})
	}
	return
}

func transformModules(from *data.StationFull, configurationID int64) (to []*station.StationModule) {
	to = make([]*station.StationModule, 0)
	for _, v := range from.Modules {
		if v.ConfigurationID != configurationID {
			continue
		}

		sensors := make([]*data.ModuleSensor, 0)
		for _, s := range from.Sensors {
			if s.ModuleID == v.ID {
				sensors = append(sensors, s)
			}
		}

		sensorsWm := make([]*station.StationSensor, 0)
		translatedName := translateModuleName(v.Name, sensors)
		moduleKey := translateModuleKey(translatedName)

		for _, s := range sensors {
			key := strcase.ToLowerCamel(s.Name)

			sensorsWm = append(sensorsWm, &station.StationSensor{
				Key:           key,
				FullKey:       moduleKey + "." + key,
				Name:          s.Name,
				UnitOfMeasure: s.UnitOfMeasure,
				Reading:       transformReading(s),
			})
		}

		hardwareID := hex.EncodeToString(v.HardwareID)

		to = append(to, &station.StationModule{
			ID:         v.ID,
			FullKey:    moduleKey,
			HardwareID: &hardwareID,
			Name:       translatedName,
			Position:   int32(v.Position),
			Flags:      int32(v.Flags),
			Internal:   v.Flags > 0 || v.Position == 255,
			Sensors:    sensorsWm,
		})

	}
	return
}

var (
	NameMap = map[string]string{
		"distance":    "modules.distance",
		"weather":     "modules.weather",
		"diagnostics": "modules.diagnostics",
		"ultrasonic":  "modules.distance",
	}
)

// This is going to go away eventually once no modules with old names
// are coming in. Then we can do a database migration and get rid of
// them completely. I'm deciding to leave them in so we can see them
// disappear over time.
func translateModuleName(old string, sensors []*data.ModuleSensor) string {
	if newName, ok := NameMap[old]; ok {
		return newName
	}

	if old == "water" {
		if len(sensors) == 1 {
			return "modules.water." + sensors[0].Name
		} else {
			return "modules.water.ec"
		}
	}

	return old
}

func translateModuleKey(name string) string {
	return strings.Replace(name, "modules.", "fk.", 1)
}

func transformLocation(sf *data.StationFull, preciseLocation bool) *station.StationLocation {
	if l := sf.Station.Location; l != nil {
		sl := &station.StationLocation{}

		if preciseLocation {
			sl.Precise = []float64{l.Longitude(), l.Latitude()}
		}

		regions := make([]*station.StationRegion, 0)
		for _, area := range sf.Areas {
			regions = append(regions, &station.StationRegion{
				Name:  area.Name,
				Shape: area.Geometry.Coordinates(),
			})
		}
		if false {
			sl.Regions = regions
		}

		return sl
	}
	return nil
}

func transformStationFull(signer *Signer, p Permissions, sf *data.StationFull, preciseLocation bool) (*station.StationFull, error) {
	readOnly := true
	if p != nil {
		sp, err := p.ForStation(sf.Station)
		if err != nil {
			return nil, err
		}

		readOnly = sp.IsReadOnly()
	}

	configurations := transformConfigurations(sf)

	dataSummary := transformDataSummary(sf.DataSummary)

	location := transformLocation(sf, preciseLocation)

	var recordingStartedAt *int64
	if sf.Station.RecordingStartedAt != nil {
		unix := sf.Station.RecordingStartedAt.Unix() * 1000
		recordingStartedAt = &unix
	}

	return &station.StationFull{
		ID:       sf.Station.ID,
		Name:     sf.Station.Name,
		ReadOnly: readOnly,
		DeviceID: hex.EncodeToString(sf.Station.DeviceID),
		Uploads:  transformUploads(sf.Ingestions),
		Configurations: &station.StationConfigurations{
			All: configurations,
		},
		RecordingStartedAt: recordingStartedAt,
		Battery:            sf.Station.Battery,
		MemoryUsed:         sf.Station.MemoryUsed,
		MemoryAvailable:    sf.Station.MemoryAvailable,
		FirmwareNumber:     sf.Station.FirmwareNumber,
		FirmwareTime:       sf.Station.FirmwareTime,
		UpdatedAt:          sf.Station.UpdatedAt.Unix() * 1000,
		LocationName:       sf.Station.LocationName,
		PlaceNameOther:     sf.Station.PlaceOther,
		PlaceNameNative:    sf.Station.PlaceNative,
		Location:           location,
		Data:               dataSummary,
		Owner: &station.StationOwner{
			ID:   sf.Owner.ID,
			Name: sf.Owner.Name,
		},
		Photos: &station.StationPhotos{
			Small: fmt.Sprintf("/stations/%d/photo", sf.Station.ID),
		},
	}, nil
}

func transformDataSummary(ads *data.AggregatedDataSummary) *station.StationDataSummary {
	if ads == nil {
		return nil
	}
	if ads.Start == nil || ads.End == nil || ads.NumberSamples == nil {
		return nil
	}
	return &station.StationDataSummary{
		Start:           (*ads.Start).Unix() * 1000,
		End:             (*ads.End).Unix() * 1000,
		NumberOfSamples: *ads.NumberSamples,
	}
}

func transformAllStationFull(signer *Signer, p Permissions, sfs []*data.StationFull, preciseLocation bool) ([]*station.StationFull, error) {
	stations := make([]*station.StationFull, 0)

	for _, sf := range sfs {
		after, err := transformStationFull(signer, p, sf, preciseLocation)
		if err != nil {
			return nil, err
		}

		stations = append(stations, after)
	}

	return stations, nil
}

func defaultPhoto(id int32) (*station.DownloadedPhoto, error) {
	defaultPhotoContentType := "image/png"
	defaultPhoto, err := StationDefaultPicture(int64(id))
	if err != nil {
		// NOTE This, hopefully never happens because we've got no image to send back.
		return nil, err
	}

	return &station.DownloadedPhoto{
		ContentType: defaultPhotoContentType,
		Length:      int64(len(defaultPhoto)),
		Body:        defaultPhoto,
	}, nil
}
