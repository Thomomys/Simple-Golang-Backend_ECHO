package types

import (
	"database/sql/driver"
	"errors"
	"time"
)

// DateTime represents a custom datetime type based on time.Time.
type DateTime struct {
	time.Time
}

// NewDateTime creates a new instance of DateTime with the given year, month, and day.
func NewDateTime(year int, month time.Month, day int) DateTime {
	return DateTime{time.Date(year, month, day, 0, 0, 0, 0, time.UTC)}
}

// Value implements the driver.Valuer interface for DateTime.
func (cd DateTime) Value() (driver.Value, error) {
	return cd.Time, nil
}

// Scan implements the sql.Scanner interface for DateTime.
func (cd *DateTime) Scan(value interface{}) error {
	var t []byte                  // Use var declaration for clarity and to avoid shadowing the imported "time" package.
	if len(value.([]byte)) == 0 { // Check if the byte slice is empty to avoid unnecessary parsing.
		return errors.New("empty time value")
	}

	parsedTime, err := time.Parse(time.RFC3339, string(t)) // Use time.RFC3339 for consistent parsing.
	if err != nil {
		return err
	}

	cd.Time = parsedTime
	return nil
}
