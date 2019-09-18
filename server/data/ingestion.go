package data

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jmoiron/sqlx/types"

	"github.com/lib/pq"
)

type Ingestion struct {
	ID        int64         `db:"id"`
	Time      time.Time     `db:"time"`
	UploadID  string        `db:"upload_id"`
	UserID    int32         `db:"user_id"`
	DeviceID  []byte        `db:"device_id"`
	Size      int64         `db:"size"`
	URL       string        `db:"url"`
	Type      *string       `db:"type"`
	Blocks    Int64Range    `db:"blocks"`
	Flags     pq.Int64Array `db:"flags"`
	Errors    *bool         `db:"errors"`
	Completed *time.Time    `db:"processing_completed"`
	Attempted *time.Time    `db:"processing_attempted"`
}

type Provision struct {
	ID          int64     `db:"id"`
	Time        time.Time `db:"time"`
	DeviceID    []byte    `db:"device_id"`
	Generation  []byte    `db:"generation"`
	IngestionID int64     `db:"ingestion_id"`
}

type Int64Range []int64

// Scan implements the sql.Scanner interface.
func (a *Int64Range) Scan(src interface{}) error {
	switch src := src.(type) {
	case []byte:
		return a.parseString(string(src))
	case string:
		return a.parseString(src)
	case nil:
		*a = nil
		return nil
	}

	return fmt.Errorf("pq: cannot convert %T to Int64Range", src)
}

func (a *Int64Range) parseString(s string) error {
	if s[0] != '[' || s[len(s)-1] != ')' {
		return fmt.Errorf("Unexpected range boundaries. I was lazy.")
	}

	values := s[1 : len(s)-1]
	b, err := ParseBlocks(values)
	if err != nil {
		return err
	}

	b[1] = b[1] - 1

	*a = b

	return nil
}

// Value implements the driver.Valuer interface.
func (a Int64Range) Value() (driver.Value, error) {
	if a == nil {
		return nil, nil
	}

	if n := len(a); n > 0 {
		// There will be at least two curly brackets, N bytes of values,
		// and N-1 bytes of delimiters.
		b := make([]byte, 1, 1+2*n)
		b[0] = '['

		b = strconv.AppendInt(b, a[0], 10)
		for i := 1; i < n; i++ {
			b = append(b, ',')
			b = strconv.AppendInt(b, a[i], 10)
		}

		return string(append(b, ']')), nil
	}

	return "{}", nil
}

func ParseBlocks(s string) ([]int64, error) {
	parts := strings.Split(s, ",")

	if len(parts) != 2 {
		return nil, fmt.Errorf("Malformed block range")
	}

	blocks := make([]int64, 2)
	for i, p := range parts {
		b, err := strconv.Atoi(strings.TrimSpace(p))
		if err != nil {
			return nil, err
		}
		blocks[i] = int64(b)
	}

	return blocks, nil
}

type MetaRecord struct {
	ID          int64          `db:"id"`
	ProvisionID int64          `db:"provision_id"`
	Time        time.Time      `db:"time"`
	Number      int64          `db:"number"`
	Data        types.JSONText `db:"raw"`
}

func (d *DataRecord) SetData(data interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	d.Data = jsonData
	return nil
}

func (d *DataRecord) GetData() (fields map[string]interface{}, err error) {
	err = json.Unmarshal(d.Data, &fields)
	if err != nil {
		return nil, err
	}
	return
}

type DataRecord struct {
	ID          int64          `db:"id"`
	ProvisionID int64          `db:"provision_id"`
	Time        time.Time      `db:"time"`
	Number      int64          `db:"number"`
	Location    *Location      `db:"location"`
	Data        types.JSONText `db:"raw"`
}

func (d *MetaRecord) SetData(data interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	d.Data = jsonData
	return nil
}

func (d *MetaRecord) GetData() (fields map[string]interface{}, err error) {
	err = json.Unmarshal(d.Data, &fields)
	if err != nil {
		return nil, err
	}
	return
}
