package backend

import (
	"github.com/conservify/sqlxcache"

	"github.com/fieldkit/cloud/server/backend/handlers"
)

func NewAllHandlers(db *sqlxcache.DB) (RecordHandler, error) {
	return handlers.NewStationModelRecordHandler(db), nil
}
