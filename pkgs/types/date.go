package types

import (
	"database/sql/driver"
	"errors"
	"time"
)

type Date struct {
	time.Time
}

func NewDate(year int, month time.Month, day int) Date {
	return Date{
		time.Date(year, month, day, 0, 0, 0, 0, time.UTC),
	}
}

func (ct *Date) Value() (driver.Value, error) {
	return ct.Time, nil
}

func (ct *Date) Scan(value interface{}) error {
	t, ok := value.([]uint8)
	if !ok {
		return errors.New("invalid time format")
	}

	parsedTime, err := time.Parse("2006-01-02 15:04:05", string(t))
	if err != nil {
		return err
	}

	ct.Time = parsedTime
	return nil
}
