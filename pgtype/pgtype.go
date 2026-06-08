package pgtype

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	pgxpgtype "github.com/jackc/pgx/v5/pgtype"
	commonpgtype "github.com/webdevelop-pro/go-common/orm/pgtype"
)

type Status = commonpgtype.Status
type InfinityModifier = commonpgtype.InfinityModifier

type Timestamptz commonpgtype.Timestamptz

const (
	Undefined        = commonpgtype.Undefined
	Null             = commonpgtype.Null
	Present          = commonpgtype.Present
	Infinity         = commonpgtype.Infinity
	None             = commonpgtype.None
	NegativeInfinity = commonpgtype.NegativeInfinity
)

func (t *Timestamptz) Set(src any) error {
	return (*commonpgtype.Timestamptz)(t).Set(src)
}

func (t *Timestamptz) ScanTimestamptz(src pgxpgtype.Timestamptz) error {
	return (*commonpgtype.Timestamptz)(t).ScanTimestamptz(src)
}

func (t Timestamptz) TimestamptzValue() (pgxpgtype.Timestamptz, error) {
	return (commonpgtype.Timestamptz)(t).TimestamptzValue()
}

func (t *Timestamptz) Scan(src any) error {
	if src == nil {
		*t = Timestamptz{Status: Null}
		return nil
	}

	switch value := src.(type) {
	case string:
		return t.setString(value)
	case []byte:
		return t.setString(string(value))
	default:
		return t.Set(value)
	}
}

func (t Timestamptz) Value() (driver.Value, error) {
	switch t.Status {
	case Present:
		switch t.InfinityModifier {
		case Infinity:
			return "infinity", nil
		case NegativeInfinity:
			return "-infinity", nil
		default:
			return t.Time, nil
		}
	case Null:
		return nil, nil
	default:
		return nil, fmt.Errorf("cannot encode status %d", t.Status)
	}
}

func (t Timestamptz) MarshalJSON() ([]byte, error) {
	switch t.Status {
	case Null:
		return []byte("null"), nil
	case Present:
		switch t.InfinityModifier {
		case Infinity:
			return json.Marshal("infinity")
		case NegativeInfinity:
			return json.Marshal("-infinity")
		default:
			return json.Marshal(t.Time.Format(time.RFC3339Nano))
		}
	default:
		return nil, fmt.Errorf("cannot encode status %d", t.Status)
	}
}

func (t *Timestamptz) UnmarshalJSON(data []byte) error {
	var value *string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	if value == nil {
		*t = Timestamptz{Status: Null}
		return nil
	}

	return t.setString(*value)
}

func (t *Timestamptz) setString(value string) error {
	switch value {
	case "infinity":
		*t = Timestamptz{Status: Present, InfinityModifier: Infinity}
		return nil
	case "-infinity":
		*t = Timestamptz{Status: Present, InfinityModifier: NegativeInfinity}
		return nil
	}

	parsed, err := time.Parse(time.RFC3339Nano, value)
	if err != nil {
		return err
	}
	*t = Timestamptz{Time: parsed, Status: Present}
	return nil
}
