package database

import (
	"database/sql/driver"
	"encoding/json"
	"strings"
	"time"
)

// Date is a custom type that handles date-only strings (YYYY-MM-DD)
type Date struct {
	time.Time
}

// UnmarshalJSON implements json.Unmarshaler
func (d *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	if s == "" || s == "null" {
		d.Time = time.Time{}
		return nil
	}

	// Try parsing as date-only first (YYYY-MM-DD)
	if t, err := time.Parse("2006-01-02", s); err == nil {
		d.Time = t
		return nil
	}

	// Try parsing as datetime (RFC3339)
	if t, err := time.Parse(time.RFC3339, s); err == nil {
		d.Time = t
		return nil
	}

	// Try parsing as datetime without timezone
	if t, err := time.Parse("2006-01-02T15:04:05", s); err == nil {
		d.Time = t
		return nil
	}

	// Try parsing as datetime-local format
	if t, err := time.Parse("2006-01-02T15:04:05Z07:00", s); err == nil {
		d.Time = t
		return nil
	}

	// If all else fails, try standard formats
	return json.Unmarshal(b, &d.Time)
}

// MarshalJSON implements json.Marshaler
func (d Date) MarshalJSON() ([]byte, error) {
	if d.Time.IsZero() {
		return []byte("null"), nil
	}
	return json.Marshal(d.Time.Format("2006-01-02"))
}

// Value implements driver.Valuer for database storage
func (d Date) Value() (driver.Value, error) {
	if d.Time.IsZero() {
		return nil, nil
	}
	return d.Time, nil
}

// Scan implements sql.Scanner for database retrieval
func (d *Date) Scan(value interface{}) error {
	if value == nil {
		d.Time = time.Time{}
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		d.Time = v
		return nil
	case []byte:
		return d.UnmarshalJSON(v)
	case string:
		return d.UnmarshalJSON([]byte(`"` + v + `"`))
	}
	return nil
}

// DateTime is a custom type that handles datetime strings
type DateTime struct {
	time.Time
}

// UnmarshalJSON implements json.Unmarshaler
func (dt *DateTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	if s == "" || s == "null" {
		dt.Time = time.Time{}
		return nil
	}

	// Try parsing as datetime-local format (from HTML datetime-local input)
	if t, err := time.Parse("2006-01-02T15:04", s); err == nil {
		dt.Time = t
		return nil
	}

	// Try parsing as RFC3339
	if t, err := time.Parse(time.RFC3339, s); err == nil {
		dt.Time = t
		return nil
	}

	// Try parsing as date-only and set to midnight
	if t, err := time.Parse("2006-01-02", s); err == nil {
		dt.Time = t
		return nil
	}

	// If all else fails, try standard formats
	return json.Unmarshal(b, &dt.Time)
}

// MarshalJSON implements json.Marshaler
func (dt DateTime) MarshalJSON() ([]byte, error) {
	if dt.Time.IsZero() {
		return []byte("null"), nil
	}
	return json.Marshal(dt.Time.Format(time.RFC3339))
}

// Value implements driver.Valuer for database storage
func (dt DateTime) Value() (driver.Value, error) {
	if dt.Time.IsZero() {
		return nil, nil
	}
	return dt.Time, nil
}

// Scan implements sql.Scanner for database retrieval
func (dt *DateTime) Scan(value interface{}) error {
	if value == nil {
		dt.Time = time.Time{}
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		dt.Time = v
		return nil
	case []byte:
		return dt.UnmarshalJSON(v)
	case string:
		return dt.UnmarshalJSON([]byte(`"` + v + `"`))
	}
	return nil
}

