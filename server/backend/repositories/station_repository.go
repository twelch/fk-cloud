package repositories

import (
	"context"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/conservify/sqlxcache"

	"github.com/jmoiron/sqlx"

	pbapp "github.com/fieldkit/app-protocol"

	"github.com/fieldkit/cloud/server/data"
)

var (
	StatusReplySourceID = int32(0)
)

type StationRepository struct {
	db *sqlxcache.DB
}

func NewStationRepository(db *sqlxcache.DB) (rr *StationRepository, err error) {
	return &StationRepository{db: db}, nil
}

func (r *StationRepository) Add(ctx context.Context, adding *data.Station) (station *data.Station, err error) {
	if err := r.db.NamedGetContext(ctx, adding, `
		INSERT INTO fieldkit.station
		(name, device_id, owner_id, created_at, updated_at, location, location_name, place_other, place_native,
		 battery, memory_used, memory_available, firmware_number, firmware_time, recording_started_at) VALUES
		(:name, :device_id, :owner_id, :created_at, :updated_at, ST_SetSRID(ST_GeomFromText(:location), 4326), :location_name, :place_other, :place_native,
         :battery, :memory_used, :memory_available, :firmware_number, :firmware_time, :recording_started_at)
		RETURNING id
		`, adding); err != nil {
		return nil, err
	}

	return adding, nil
}

func (r *StationRepository) Update(ctx context.Context, station *data.Station) (err error) {
	if _, err := r.db.NamedExecContext(ctx, `
		UPDATE fieldkit.station SET
			   name = :name,
			   battery = :battery,
			   location = ST_SetSRID(ST_GeomFromText(:location), 4326),
			   recording_started_at = :recording_started_at,
			   location_name = :location_name,
               place_other = :place_other,
               place_native = :place_native,
			   memory_available = :memory_available,
			   memory_used = :memory_used,
			   firmware_number = :firmware_number,
			   firmware_time = :firmware_time,
			   updated_at = :updated_at
		WHERE id = :id
		`, station); err != nil {
		return err
	}

	return nil
}

func (r *StationRepository) QueryStationByID(ctx context.Context, id int32) (station *data.Station, err error) {
	station = &data.Station{}
	if err := r.db.GetContext(ctx, station, `
		SELECT
			id, name, device_id, owner_id, created_at, updated_at, battery, location_name, place_other, place_native,
			recording_started_at, memory_used, memory_available, firmware_number, firmware_time, ST_AsBinary(location) AS location
		FROM fieldkit.station WHERE id = $1
		`, id); err != nil {
		return nil, err
	}
	return station, nil
}

func (r *StationRepository) QueryStationsByDeviceID(ctx context.Context, deviceIdBytes []byte) (stations []*data.Station, err error) {
	stations = []*data.Station{}
	if err := r.db.SelectContext(ctx, &stations, `
		SELECT
			id, name, device_id, owner_id, created_at, updated_at, battery, location_name, place_other, place_native,
			recording_started_at, memory_used, memory_available, firmware_number, firmware_time, ST_AsBinary(location) AS location
		FROM fieldkit.station WHERE device_id = $1
		`, deviceIdBytes); err != nil {
		return nil, err
	}
	return stations, nil
}

func (r *StationRepository) QueryStationByDeviceID(ctx context.Context, deviceIdBytes []byte) (station *data.Station, err error) {
	station = &data.Station{}
	if err := r.db.GetContext(ctx, station, `
		SELECT
			id, name, device_id, owner_id, created_at, updated_at, battery, location_name, place_other, place_native,
			recording_started_at, memory_used, memory_available, firmware_number, firmware_time, ST_AsBinary(location) AS location
		FROM fieldkit.station WHERE device_id = $1
		`, deviceIdBytes); err != nil {
		return nil, err
	}
	return station, nil
}

func (r *StationRepository) TryQueryStationByDeviceID(ctx context.Context, deviceIdBytes []byte) (station *data.Station, err error) {
	stations := []*data.Station{}
	if err := r.db.SelectContext(ctx, &stations, `
		SELECT
			id, name, device_id, owner_id, created_at, updated_at, battery, location_name, place_other, place_native,
			recording_started_at, memory_used, memory_available, firmware_number, firmware_time, ST_AsBinary(location) AS location
		FROM fieldkit.station WHERE device_id = $1
		`, deviceIdBytes); err != nil {
		return nil, err
	}
	if len(stations) != 1 {
		return nil, nil
	}
	return stations[0], nil
}

func (r *StationRepository) UpsertConfiguration(ctx context.Context, configuration *data.StationConfiguration) (*data.StationConfiguration, error) {
	if configuration.SourceID != nil && configuration.MetaRecordID == nil {
		if err := r.db.NamedGetContext(ctx, configuration, `
			INSERT INTO fieldkit.station_configuration
				(provision_id, source_id, updated_at) VALUES
				(:provision_id, :source_id, :updated_at)
			ON CONFLICT (provision_id, source_id) DO UPDATE SET updated_at = EXCLUDED.updated_at
			RETURNING *
			`, configuration); err != nil {
			return nil, err
		}
		return configuration, nil
	}
	if configuration.SourceID == nil && configuration.MetaRecordID != nil {
		if err := r.db.NamedGetContext(ctx, configuration, `
			INSERT INTO fieldkit.station_configuration
				(provision_id, meta_record_id, updated_at) VALUES
				(:provision_id, :meta_record_id, :updated_at)
			ON CONFLICT (provision_id, meta_record_id) DO UPDATE SET updated_at = EXCLUDED.updated_at
			RETURNING *
			`, configuration); err != nil {
			return nil, err
		}
		return configuration, nil
	}
	return nil, fmt.Errorf("invalid StationConfiguration")
}

func (r *StationRepository) UpsertStationModule(ctx context.Context, module *data.StationModule) (*data.StationModule, error) {
	if err := r.db.NamedGetContext(ctx, module, `
		INSERT INTO fieldkit.station_module
			(configuration_id, hardware_id, module_index, position, flags, name, manufacturer, kind, version) VALUES
			(:configuration_id, :hardware_id, :module_index, :position, :flags, :name, :manufacturer, :kind, :version)
		ON CONFLICT (configuration_id, hardware_id)
			DO UPDATE SET module_index = EXCLUDED.module_index,
                          position = EXCLUDED.position,
                          flags = EXCLUDED.flags,
                          name = EXCLUDED.name,
                          manufacturer = EXCLUDED.manufacturer,
                          kind = EXCLUDED.kind,
						  version = EXCLUDED.version
		RETURNING *
		`, module); err != nil {
		return nil, err
	}
	return module, nil
}

func (r *StationRepository) UpsertModuleSensor(ctx context.Context, sensor *data.ModuleSensor) (*data.ModuleSensor, error) {
	if err := r.db.NamedGetContext(ctx, sensor, `
		INSERT INTO fieldkit.module_sensor AS s
			(module_id, configuration_id, sensor_index, unit_of_measure, name, reading_last, reading_time) VALUES
			(:module_id, :configuration_id, :sensor_index, :unit_of_measure, :name, :reading_last, :reading_time)
		ON CONFLICT (module_id, sensor_index)
			DO UPDATE SET name = EXCLUDED.name,
                          unit_of_measure = EXCLUDED.unit_of_measure,
						  reading_last = COALESCE(s.reading_last, EXCLUDED.reading_last),
						  reading_time = COALESCE(s.reading_time, EXCLUDED.reading_time)
		RETURNING *
		`, sensor); err != nil {
		return nil, err
	}
	return sensor, nil
}

func (r *StationRepository) UpdateStationModelFromStatus(ctx context.Context, s *data.Station, rawStatus string) error {
	statusReply, err := s.ParseHttpReply(rawStatus)
	if err != nil {
		return err
	}

	if statusReply.Status == nil || statusReply.Status.Identity == nil || statusReply.Status.Identity.Generation == nil {
		return fmt.Errorf("incomplete status, no identity or generation")
	}

	if err := r.updateStationConfigurationFromStatus(ctx, s, statusReply); err != nil {
		return err
	}

	if err := r.updateDeployedActivityFromStatus(ctx, s); err != nil {
		return err
	}

	return nil
}

func (r *StationRepository) updateStationConfigurationFromStatus(ctx context.Context, station *data.Station, statusReply *pbapp.HttpReply) error {
	pr, err := NewProvisionRepository(r.db)
	if err != nil {
		return err
	}

	p, err := pr.QueryOrCreateProvision(ctx, station.DeviceID, statusReply.Status.Identity.Generation)
	if err != nil {
		return err
	}

	configuration := &data.StationConfiguration{
		ProvisionID: p.ID,
		SourceID:    &StatusReplySourceID,
		UpdatedAt:   time.Now(),
	}
	if _, err := r.UpsertConfiguration(ctx, configuration); err != nil {
		return err
	}

	keepingModules := make([]int64, 0)

	if statusReply.LiveReadings != nil {
		for _, lr := range statusReply.LiveReadings {
			time := time.Unix(int64(lr.Time), 0)

			for moduleIndex, lrm := range lr.Modules {
				m := lrm.Module

				module := newStationModule(m, configuration, uint32(moduleIndex))
				if _, err := r.UpsertStationModule(ctx, module); err != nil {
					return err
				}

				keepingModules = append(keepingModules, module.ID)

				keepingSensors := make([]int64, 0)

				for sensorIndex, lrs := range lrm.Readings {
					s := lrs.Sensor
					value := float64(lrs.Value)

					sensor := newModuleSensor(s, module, configuration, uint32(sensorIndex), &time, &value)
					if _, err := r.UpsertModuleSensor(ctx, sensor); err != nil {
						return err
					}

					keepingSensors = append(keepingSensors, sensor.ID)
				}

				if err := r.deleteModuleSensorsExcept(ctx, module.ID, keepingSensors); err != nil {
					return err
				}
			}
		}
	} else {
		for moduleIndex, m := range statusReply.Modules {
			module := newStationModule(m, configuration, uint32(moduleIndex))
			if _, err := r.UpsertStationModule(ctx, module); err != nil {
				return err
			}

			keepingModules = append(keepingModules, module.ID)

			keepingSensors := make([]int64, 0)

			for sensorIndex, s := range m.Sensors {
				sensor := newModuleSensor(s, module, configuration, uint32(sensorIndex), nil, nil)
				if _, err := r.UpsertModuleSensor(ctx, sensor); err != nil {
					return err
				}

				keepingSensors = append(keepingSensors, sensor.ID)
			}

			if err := r.deleteModuleSensorsExcept(ctx, module.ID, keepingSensors); err != nil {
				return err
			}
		}
	}

	if err := r.deleteStationModulesExcept(ctx, configuration.ID, keepingModules); err != nil {
		return err
	}

	log := Logger(ctx).Sugar()

	log.Infow("configuration", "station_id", station.ID, "configuration_id", configuration.ID, "provision_id", p.ID)

	if _, err := r.db.ExecContext(ctx, `
		INSERT INTO fieldkit.visible_configuration (station_id, configuration_id)
        SELECT $1 AS station_id, $2 AS configuration_id
		ON CONFLICT ON CONSTRAINT visible_configuration_pkey
		DO UPDATE SET configuration_id = EXCLUDED.configuration_id
		`, station.ID, configuration.ID); err != nil {
		return err
	}

	return nil
}

func (r *StationRepository) updateDeployedActivityFromStatus(ctx context.Context, station *data.Station) error {
	if station.RecordingStartedAt == nil {
		return nil
	}

	if station.Location == nil {
		log := Logger(ctx).Sugar()
		log.Warnw("deployed station w/o location", "station_id", station.ID)
		return nil
	}

	if _, err := r.db.ExecContext(ctx, `
		INSERT INTO fieldkit.station_deployed AS sd
			(created_at, station_id, deployed_at, location) VALUES
			(NOW(), $1, $2, ST_SetSRID(ST_GeomFromText($3), 4326))
		ON CONFLICT (station_id, deployed_at)
		DO UPDATE SET location = EXCLUDED.location
		RETURNING id
		`, station.ID, station.RecordingStartedAt, station.Location); err != nil {
		return err
	}

	return nil
}

func (r *StationRepository) deleteModuleSensorsExcept(ctx context.Context, moduleID int64, keeping []int64) error {
	if query, args, err := sqlx.In(`
		DELETE FROM fieldkit.module_sensor WHERE module_id = ? AND id NOT IN (?)
		`, moduleID, keeping); err != nil {
		return err
	} else {
		if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), args...); err != nil {
			return err
		}
	}

	return nil
}

func (r *StationRepository) deleteStationModulesExcept(ctx context.Context, configurationID int64, keeping []int64) error {
	if query, args, err := sqlx.In(`
		DELETE FROM fieldkit.module_sensor WHERE module_id IN (SELECT id FROM fieldkit.station_module WHERE configuration_id = ? AND id NOT IN (?))
		`, configurationID, keeping); err != nil {
		return err
	} else {
		if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), args...); err != nil {
			return err
		}
	}

	if query, args, err := sqlx.In(`
		DELETE FROM fieldkit.station_module WHERE configuration_id = ? AND id NOT IN (?)
		`, configurationID, keeping); err != nil {
		return err
	} else {
		if _, err := r.db.ExecContext(ctx, r.db.Rebind(query), args...); err != nil {
			return err
		}
	}

	return nil
}

func (r *StationRepository) QueryStationFull(ctx context.Context, id int32) (*data.StationFull, error) {
	stations := []*data.Station{}
	if err := r.db.SelectContext(ctx, &stations, `
		SELECT
			id, name, device_id, owner_id, created_at, updated_at, battery, location_name, place_other, place_native, recording_started_at,
			memory_used, memory_available, firmware_number, firmware_time, ST_AsBinary(location) AS location
		FROM fieldkit.station WHERE id = $1
		`, id); err != nil {
		return nil, err
	}

	if len(stations) != 1 {
		return nil, fmt.Errorf("no such station: %v", id)
	}

	owners := []*data.User{}
	if err := r.db.SelectContext(ctx, &owners, `
		SELECT * FROM fieldkit.user WHERE id = $1
		`, stations[0].OwnerID); err != nil {
		return nil, err
	}

	areas := []*data.StationArea{}
	if err := r.db.SelectContext(ctx, &areas, `
		SELECT
			s.id,
			c.name AS name,
			ST_AsBinary(c.geom) AS geometry
		FROM fieldkit.station AS s
		JOIN fieldkit.counties AS c ON (ST_Contains(c.geom, s.location))
		WHERE s.id = $1
		  AND s.location IS NOT NULL
		`, id); err != nil {
		return nil, err
	}

	dataSummaries := []*data.AggregatedDataSummary{}
	if err := r.db.SelectContext(ctx, &dataSummaries, `
		SELECT
			a.station_id, MIN(a.time) AS start, MAX(a.time) AS end, SUM(a.nsamples) AS number_samples
		FROM fieldkit.aggregated_24h AS a
		WHERE station_id IN ($1)
		GROUP BY a.station_id
	`, stations[0].ID); err != nil {
		return nil, err
	}

	ingestions := []*data.Ingestion{}
	if err := r.db.SelectContext(ctx, &ingestions, `
		SELECT id, time, upload_id, user_id, device_id, generation, size, url, type, blocks, flags
		FROM fieldkit.ingestion WHERE device_id = $1 ORDER BY time DESC LIMIT 10
		`, stations[0].DeviceID); err != nil {
		return nil, err
	}

	provisions := []*data.Provision{}
	if err := r.db.SelectContext(ctx, &provisions, `
		SELECT * FROM fieldkit.provision WHERE device_id = $1
		`, stations[0].DeviceID); err != nil {
		return nil, err
	}

	configurations := []*data.StationConfiguration{}
	if err := r.db.SelectContext(ctx, &configurations, `
		SELECT * FROM fieldkit.station_configuration WHERE provision_id IN (
			SELECT id FROM fieldkit.provision WHERE device_id = $1
		) ORDER BY updated_at DESC
		`, stations[0].DeviceID); err != nil {
		return nil, err
	}

	modules := []*data.StationModule{}
	if err := r.db.SelectContext(ctx, &modules, `
		SELECT
			sm.*
		FROM fieldkit.station_module AS sm
		WHERE sm.configuration_id IN (
			SELECT id FROM fieldkit.station_configuration WHERE provision_id IN (
				SELECT id FROM fieldkit.provision WHERE device_id = $1
			)
		)
		ORDER BY sm.module_index
		`, stations[0].DeviceID); err != nil {
		return nil, err
	}

	sensors := []*data.ModuleSensor{}
	if err := r.db.SelectContext(ctx, &sensors, `
		SELECT
			ms.*
		FROM fieldkit.module_sensor AS ms
		WHERE ms.configuration_id IN (
			SELECT id FROM fieldkit.station_configuration WHERE provision_id IN (
				SELECT id FROM fieldkit.provision WHERE device_id = $1
			)
		)
		ORDER BY ms.sensor_index
		`, stations[0].DeviceID); err != nil {
		return nil, err
	}

	all, err := r.toStationFull(stations, owners, areas, dataSummaries, ingestions, provisions, configurations, modules, sensors)
	if err != nil {
		return nil, err
	}

	return all[0], nil
}

func (r *StationRepository) QueryStationFullByOwnerID(ctx context.Context, id int32) ([]*data.StationFull, error) {
	stations := []*data.Station{}
	if err := r.db.SelectContext(ctx, &stations, `
		SELECT
			id, name, device_id, owner_id, created_at, updated_at, battery, location_name, place_other, place_native, recording_started_at,
			memory_used, memory_available, firmware_number, firmware_time, ST_AsBinary(location) AS location
		FROM fieldkit.station WHERE owner_id = $1
		`, id); err != nil {
		return nil, err
	}

	owners := []*data.User{}
	if err := r.db.SelectContext(ctx, &owners, `
		SELECT * FROM fieldkit.user WHERE id = $1
		`, id); err != nil {
		return nil, err
	}

	areas := []*data.StationArea{}
	if err := r.db.SelectContext(ctx, &areas, `
		SELECT
			s.id,
			c.name AS name,
			ST_AsBinary(c.geom) AS geometry
		FROM fieldkit.station AS s
		JOIN fieldkit.counties AS c ON (ST_Contains(c.geom, s.location))
		WHERE s.owner_id = $1
		  AND s.location IS NOT NULL
		`, id); err != nil {
		return nil, err
	}

	dataSummaries := []*data.AggregatedDataSummary{}
	if err := r.db.SelectContext(ctx, &dataSummaries, `
		SELECT
			a.station_id, MIN(a.time) AS start, MAX(a.time) AS end, SUM(a.nsamples) AS number_samples
		FROM fieldkit.aggregated_24h AS a
		WHERE station_id IN (SELECT id FROM fieldkit.station WHERE owner_id = $1)
		GROUP BY a.station_id
	`, id); err != nil {
		return nil, err
	}

	ingestions := []*data.Ingestion{}
	if err := r.db.SelectContext(ctx, &ingestions, `
		SELECT
			id, time, upload_id, user_id, device_id, generation, size, url, type, blocks, flags
		FROM (
			SELECT
				*,
				rank() OVER (PARTITION BY device_id ORDER BY time DESC)
			FROM
			fieldkit.ingestion
			WHERE device_id IN (SELECT device_id FROM fieldkit.station WHERE owner_id = $1) ORDER BY time DESC
		) AS ranked WHERE rank < 10;
		`, id); err != nil {
		return nil, err
	}

	provisions := []*data.Provision{}
	if err := r.db.SelectContext(ctx, &provisions, `
		SELECT * FROM fieldkit.provision WHERE device_id IN (SELECT device_id FROM fieldkit.station WHERE owner_id = $1)
		`, id); err != nil {
		return nil, err
	}

	configurations := []*data.StationConfiguration{}
	if err := r.db.SelectContext(ctx, &configurations, `
		SELECT id, provision_id, meta_record_id, source_id, updated_at FROM (
			SELECT
				sc.*,
				rank() OVER (PARTITION BY provision_id ORDER BY updated_at DESC) AS rank
			FROM fieldkit.station_configuration AS sc
			WHERE sc.provision_id IN (
				SELECT id FROM fieldkit.provision WHERE device_id IN (
					SELECT device_id FROM fieldkit.station WHERE owner_id = $1
				)
			)
		) AS q
		WHERE rank <= 1
		ORDER BY updated_at DESC
		`, id); err != nil {
		return nil, err
	}

	modules := []*data.StationModule{}
	if err := r.db.SelectContext(ctx, &modules, `
		SELECT
			sm.*
		FROM fieldkit.station_module AS sm
		WHERE sm.configuration_id IN (
			SELECT id FROM fieldkit.station_configuration WHERE provision_id IN (
				SELECT id FROM fieldkit.provision WHERE device_id IN (SELECT device_id FROM fieldkit.station WHERE owner_id = $1)
			)
		)
		ORDER BY sm.configuration_id, sm.module_index
		`, id); err != nil {
		return nil, err
	}

	sensors := []*data.ModuleSensor{}
	if err := r.db.SelectContext(ctx, &sensors, `
		SELECT
			ms.*
		FROM fieldkit.module_sensor AS ms
		WHERE ms.configuration_id IN (
			SELECT id FROM fieldkit.station_configuration WHERE provision_id IN (
				SELECT id FROM fieldkit.provision WHERE device_id IN (SELECT device_id FROM fieldkit.station WHERE owner_id = $1)
			)
		)
		ORDER BY ms.configuration_id, ms.sensor_index
		`, id); err != nil {
		return nil, err
	}

	return r.toStationFull(stations, owners, areas, dataSummaries, ingestions, provisions, configurations, modules, sensors)
}

func (r *StationRepository) QueryStationFullByProjectID(ctx context.Context, id int32) ([]*data.StationFull, error) {
	stations := []*data.Station{}
	if err := r.db.SelectContext(ctx, &stations, `
		SELECT
			id, name, device_id, owner_id, created_at, updated_at, battery, location_name, place_other, place_native, recording_started_at,
			memory_used, memory_available, firmware_number, firmware_time, ST_AsBinary(location) AS location
		FROM fieldkit.station WHERE id IN (SELECT station_id FROM fieldkit.project_station WHERE project_id = $1)
		`, id); err != nil {
		return nil, err
	}

	owners := []*data.User{}
	if err := r.db.SelectContext(ctx, &owners, `
		SELECT *
		FROM fieldkit.user
		WHERE id IN (SELECT owner_id FROM fieldkit.station WHERE id IN (SELECT station_id FROM fieldkit.project_station WHERE project_id = $1))
		`, id); err != nil {
		return nil, err
	}

	areas := []*data.StationArea{}
	if err := r.db.SelectContext(ctx, &areas, `
		SELECT
			s.id,
			c.name AS name,
			ST_AsBinary(c.geom) AS geometry
		FROM fieldkit.station AS s
		JOIN fieldkit.counties AS c ON (ST_Contains(c.geom, s.location))
		WHERE s.id IN (SELECT station_id FROM fieldkit.project_station WHERE project_id = $1)
		  AND s.location IS NOT NULL
		`, id); err != nil {
		return nil, err
	}

	dataSummaries := []*data.AggregatedDataSummary{}
	if err := r.db.SelectContext(ctx, &dataSummaries, `
		SELECT
			a.station_id, MIN(a.time) AS start, MAX(a.time) AS end, SUM(a.nsamples) AS number_samples
		FROM fieldkit.aggregated_24h AS a
		WHERE station_id IN (SELECT station_id FROM fieldkit.project_station WHERE project_id = $1)
		GROUP BY a.station_id
	`, id); err != nil {
		return nil, err
	}

	ingestions := []*data.Ingestion{}
	if err := r.db.SelectContext(ctx, &ingestions, `
		SELECT
			id, time, upload_id, user_id, device_id, generation, size, url, type, blocks, flags
		FROM (
			SELECT
				*,
				rank() OVER (PARTITION BY device_id ORDER BY time DESC)
			FROM
			fieldkit.ingestion
			WHERE device_id IN (SELECT s.device_id FROM fieldkit.station AS s JOIN fieldkit.project_station AS ps ON (s.id = ps.station_id) WHERE project_id = $1)
		) AS ranked WHERE rank < 10;
		`, id); err != nil {
		return nil, err
	}

	provisions := []*data.Provision{}
	if err := r.db.SelectContext(ctx, &provisions, `
		SELECT
			p.*
		FROM fieldkit.provision AS p WHERE device_id IN (
			SELECT device_id FROM fieldkit.station WHERE id IN (
				SELECT station_id FROM fieldkit.project_station WHERE project_id = $1
			)
		)
		ORDER BY p.updated DESC
		`, id); err != nil {
		return nil, err
	}

	configurations := []*data.StationConfiguration{}
	if err := r.db.SelectContext(ctx, &configurations, `
		SELECT id, provision_id, meta_record_id, source_id, updated_at FROM (
			SELECT
				sc.*,
				rank() OVER (PARTITION BY provision_id ORDER BY updated_at DESC) AS rank
			FROM fieldkit.station_configuration AS sc
			WHERE sc.provision_id IN (
				SELECT id FROM fieldkit.provision WHERE device_id IN (
					SELECT device_id FROM fieldkit.station WHERE id IN (
						SELECT station_id FROM fieldkit.project_station WHERE project_id = $1
					)
				)
			)
		) AS q
		WHERE rank <= 1
		ORDER BY updated_at DESC
		`, id); err != nil {
		return nil, err
	}

	modules := []*data.StationModule{}
	if err := r.db.SelectContext(ctx, &modules, `
		SELECT
			sm.*
		FROM fieldkit.station_module AS sm
        WHERE sm.configuration_id IN (
			SELECT id FROM fieldkit.station_configuration WHERE provision_id IN (
				SELECT id FROM fieldkit.provision WHERE device_id IN (
					SELECT device_id FROM fieldkit.station WHERE id IN (
						SELECT station_id FROM fieldkit.project_station WHERE project_id = $1
					)
				)
			)
		)
		ORDER BY sm.configuration_id, sm.module_index
		`, id); err != nil {
		return nil, err
	}

	sensors := []*data.ModuleSensor{}
	if err := r.db.SelectContext(ctx, &sensors, `
		SELECT
			ms.*
		FROM fieldkit.module_sensor AS ms
		WHERE ms.configuration_id IN (
			SELECT id FROM fieldkit.station_configuration WHERE provision_id IN (
				SELECT id FROM fieldkit.provision WHERE device_id IN (
					SELECT device_id FROM fieldkit.station WHERE id IN (
						SELECT station_id FROM fieldkit.project_station WHERE project_id = $1
					)
				)
			)
		)
		ORDER BY ms.configuration_id, ms.sensor_index
		`, id); err != nil {
		return nil, err
	}

	return r.toStationFull(stations, owners, areas, dataSummaries, ingestions, provisions, configurations, modules, sensors)
}

func (r *StationRepository) toStationFull(stations []*data.Station, owners []*data.User, areas []*data.StationArea, dataSummaries []*data.AggregatedDataSummary, ingestions []*data.Ingestion,
	provisions []*data.Provision, configurations []*data.StationConfiguration,
	modules []*data.StationModule, sensors []*data.ModuleSensor) ([]*data.StationFull, error) {
	ownersByID := make(map[int32]*data.User)
	ingestionsByDeviceID := make(map[string][]*data.Ingestion)
	summariesByStationID := make(map[int32]*data.AggregatedDataSummary)
	mediaByStationID := make(map[int32][]*data.FieldNoteMedia)
	modulesByStationID := make(map[int32][]*data.StationModule)
	sensorsByStationID := make(map[int32][]*data.ModuleSensor)
	stationIDsByDeviceID := make(map[string]int32)
	areasByStationID := make(map[int32][]*data.StationArea)

	for _, station := range stations {
		key := hex.EncodeToString(station.DeviceID)
		ingestionsByDeviceID[key] = make([]*data.Ingestion, 0)
		mediaByStationID[station.ID] = make([]*data.FieldNoteMedia, 0)
		modulesByStationID[station.ID] = make([]*data.StationModule, 0)
		sensorsByStationID[station.ID] = make([]*data.ModuleSensor, 0)
		areasByStationID[station.ID] = make([]*data.StationArea, 0)
		stationIDsByDeviceID[key] = station.ID
	}

	for _, v := range owners {
		ownersByID[v.ID] = v
	}

	for _, v := range areas {
		areasByStationID[v.ID] = append(areasByStationID[v.ID], v)
	}

	for _, v := range dataSummaries {
		summariesByStationID[v.StationID] = v
	}

	for _, v := range ingestions {
		key := hex.EncodeToString(v.DeviceID)
		ingestionsByDeviceID[key] = append(ingestionsByDeviceID[key], v)
	}

	stationIDsByProvisionID := make(map[int64]int32)
	for _, p := range provisions {
		deviceID := hex.EncodeToString(p.DeviceID)
		stationIDsByProvisionID[p.ID] = stationIDsByDeviceID[deviceID]
	}

	configurationsByStationID := make(map[int32][]*data.StationConfiguration)
	modulesByConfigurationID := make(map[int64][]*data.StationModule)
	for _, v := range configurations {
		stationID := stationIDsByProvisionID[v.ProvisionID]
		configurationsByStationID[stationID] = append(configurationsByStationID[stationID], v)
		modulesByConfigurationID[v.ID] = make([]*data.StationModule, 0)
	}

	for _, v := range modules {
		modulesByConfigurationID[v.ConfigurationID] = append(modulesByConfigurationID[v.ConfigurationID], v)
	}

	stationIDByModuleID := make(map[int64]int32)
	for _, v := range configurations {
		modules := modulesByConfigurationID[v.ID]
		stationID := stationIDsByProvisionID[v.ProvisionID]
		modulesByStationID[stationID] = append(modulesByStationID[stationID], modules...)
		for _, m := range modules {
			stationIDByModuleID[m.ID] = stationID
		}
	}

	for _, v := range sensors {
		stationID := stationIDByModuleID[v.ModuleID]
		sensorsByStationID[stationID] = append(sensorsByStationID[stationID], v)
	}

	all := make([]*data.StationFull, 0, len(stations))
	for _, station := range stations {
		all = append(all, &data.StationFull{
			Station:        station,
			Owner:          ownersByID[station.OwnerID],
			Areas:          areasByStationID[station.ID],
			Ingestions:     ingestionsByDeviceID[station.DeviceIDHex()],
			Configurations: configurationsByStationID[station.ID],
			Modules:        modulesByStationID[station.ID],
			Sensors:        sensorsByStationID[station.ID],
			DataSummary:    summariesByStationID[station.ID],
		})
	}

	return all, nil
}

func newStationModule(m *pbapp.ModuleCapabilities, c *data.StationConfiguration, moduleIndex uint32) *data.StationModule {
	if m.Header == nil {
		return &data.StationModule{
			ConfigurationID: c.ID,
			HardwareID:      m.Id,
			Index:           moduleIndex,
			Position:        m.Position,
			Flags:           m.Flags,
			Name:            m.Name,
		}
	}
	return &data.StationModule{
		ConfigurationID: c.ID,
		HardwareID:      m.Id,
		Index:           moduleIndex,
		Position:        m.Position,
		Flags:           m.Flags,
		Name:            m.Name,
		Manufacturer:    m.Header.Manufacturer,
		Kind:            m.Header.Kind,
		Version:         m.Header.Version,
	}
}

func newModuleSensor(s *pbapp.SensorCapabilities, m *data.StationModule, c *data.StationConfiguration, sensorIndex uint32, time *time.Time, value *float64) *data.ModuleSensor {
	return &data.ModuleSensor{
		ModuleID:        m.ID,
		ConfigurationID: c.ID,
		Index:           sensorIndex,
		UnitOfMeasure:   s.UnitOfMeasure,
		Name:            s.Name,
		ReadingTime:     time,
		ReadingValue:    value,
	}
}

type EssentialQueryParams struct {
	Page     int32
	PageSize int32
}

type QueriedEssential struct {
	Stations []*data.EssentialStation
	Total    int32
}

func (sr *StationRepository) QueryEssentialStations(ctx context.Context, qp *EssentialQueryParams) (*QueriedEssential, error) {
	total := int32(0)
	if err := sr.db.GetContext(ctx, &total, `SELECT COUNT(*) FROM fieldkit.station AS s`); err != nil {
		return nil, err
	}

	stations := []*data.EssentialStation{}
	if err := sr.db.SelectContext(ctx, &stations, `
		SELECT q.* FROM
		(
			SELECT
				s.id, s.device_id, s.name, u.id AS owner_id, u.name AS owner_name,
				s.created_at, s.updated_at,
				s.memory_used, s.memory_available,
				s.firmware_time, s.firmware_number,
				s.recording_started_at,
				ST_AsBinary(location) AS location,
				(SELECT MAX(i.time) AS last_ingestion_at FROM fieldkit.ingestion AS i WHERE i.device_id = s.device_id)
			FROM fieldkit.station AS s
			JOIN fieldkit.user AS u ON (s.owner_id = u.id)
        ) AS q
		ORDER BY CASE WHEN q.last_ingestion_at IS NULL THEN q.updated_at ELSE q.last_ingestion_at END DESC, name
		LIMIT $1 OFFSET $2
		`, qp.PageSize, qp.PageSize*qp.Page); err != nil {
		return nil, err
	}

	return &QueriedEssential{
		Stations: stations,
		Total:    total,
	}, nil
}

func (sr *StationRepository) Delete(ctx context.Context, stationID int32) error {
	queries := []string{
		`DELETE FROM fieldkit.aggregated_24h WHERE station_id IN ($1)`,
		`DELETE FROM fieldkit.aggregated_12h WHERE station_id IN ($1)`,
		`DELETE FROM fieldkit.aggregated_6h WHERE station_id IN ($1)`,
		`DELETE FROM fieldkit.aggregated_1h WHERE station_id IN ($1)`,
		`DELETE FROM fieldkit.aggregated_30m WHERE station_id IN ($1)`,
		`DELETE FROM fieldkit.aggregated_10m WHERE station_id IN ($1)`,
		`DELETE FROM fieldkit.aggregated_1m WHERE station_id IN ($1)`,
		`DELETE FROM fieldkit.visible_configuration WHERE station_id IN ($1)`,
		`DELETE FROM fieldkit.notes_media WHERE station_id IN ($1)`,
		`DELETE FROM fieldkit.notes WHERE station_id IN ($1)`,
		`DELETE FROM fieldkit.station_activity WHERE station_id IN ($1)`,
		`DELETE FROM fieldkit.project_station WHERE station_id IN ($1)`,
		`DELETE FROM fieldkit.station WHERE id IN ($1)`,
	}

	return sr.db.WithNewTransaction(ctx, func(txCtx context.Context) error {
		for _, query := range queries {
			if _, err := sr.db.ExecContext(ctx, query, stationID); err != nil {
				return err
			}
		}
		return nil
	})
}
