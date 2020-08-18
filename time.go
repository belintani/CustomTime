package ctime

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

/// Time is a custom implementation of time.Time
type Time struct {
	Unix    int64        `json:",omitempty"`
	Year    int          `json:",omitempty"`
	Month   time.Month   `json:",omitempty"`
	Day     int          `json:",omitempty"`
	Weekday time.Weekday `json:",omitempty"`
	Hour    int
	Minute  int
}

// Now returns the current time in UTC
func Now() (value *Time, err error) {
	return NewTime(time.Now().UTC())
}

// Unix converts a unix time to a Time
// unix must be a positive number
func Unix(unix int64) (value *Time, err error) {

	if unix <= 0 {
		return nil, errors.New("unix time is zero")
	}

	tm := time.Unix(unix, 0)

	return NewTime(tm)
}

// NewTime converts a time.Time struct to Time in UTC
// time cannot be zero
func NewTime(time time.Time) (value *Time, err error) {

	if time.IsZero() {
		return nil, errors.New("time is zero")
	}

	t := time.UTC()

	return &Time{
		Unix:    t.Unix(),
		Year:    t.Year(),
		Month:   t.Month(),
		Day:     t.Day(),
		Weekday: t.Weekday(),
		Hour:    t.Hour(),
		Minute:  t.Minute(),
	}, nil
}

// Formats Time to String.
func (ct Time) String() string {
	var b strings.Builder
	fmt.Fprintf(&b, "%d-%02d-%02dT%02d:%02d -> %s", ct.Year, ct.Month, ct.Day, ct.Hour, ct.Minute, ct.Weekday.String())
	return b.String()
}

// Returns a filename as String based on Time.
// It should be helpful for implementation of a datalake
func (ct Time) FileName() string {
	var b strings.Builder
	fmt.Fprintf(&b, "%d-%02d-%02dT%02d:%02d", ct.Year, ct.Month, ct.Day, ct.Hour, ct.Minute)
	return b.String()
}
