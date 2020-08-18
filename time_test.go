package ctime

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {

	t.Run("Should Create a Valid Time", func(t *testing.T) {
		now := time.Now().UTC()
		request, err := NewTime(now)
		if err != nil {
			t.Errorf("Fail while creating a valid Time by current time")
		} else {
			assertCustomTime(t, request, now)
		}
	})

	t.Run("Should NOT create a Time from empty time", func(t *testing.T) {
		_, err := NewTime(time.Time{})
		if err == nil {
			t.Errorf("FAILED")
		}
	})

	t.Run("Today should be different from tomorrow", func(t *testing.T) {
		now := time.Now().UTC()
		request, err := NewTime(now)

		if err != nil {
			t.Errorf("Fail while creating a valid Time by current time")
		} else {
			tomorrow := time.Now().UTC().AddDate(0, 0, 1)
			if request.Day == tomorrow.Day() {
				t.Errorf("FAILED")
			}
		}
	})

	t.Run("Should create a Time from unix", func(t *testing.T) {
		request, err := Unix(1587842520)
		if err != nil {
			t.Errorf("Fail while creating a valid Time by UNIX time")
		} else {
			want := time.Date(2020, 04, 25, 19, 22, 0, 0, time.FixedZone("", 0))
			assertCustomTime(t, request, want)
		}
	})

	t.Run("Should NOT create a Time from unix zero value", func(t *testing.T) {
		_, err := Unix(0)
		if err == nil {
			t.Errorf("FAILED")
		}
	})

	t.Run("Should NOT create a Time from negative unix", func(t *testing.T) {
		_, err := Unix(-1)
		if err == nil {
			t.Errorf("FAILED")
		}
	})

	t.Run("Should print Time correctly", func(t *testing.T) {

		request, err := NewTime(time.Date(2004, 10, 5, 16, 30, 44, 44, time.FixedZone("", 0)))

		want := "2004-10-05T16:30 -> Tuesday"

		if err != nil {
			t.Errorf("Fail while creating a valid CustomTime by a fixed time")
		} else if want != request.String() {
			t.Errorf("customTime string is wrong, got '%s' want '%s'", request.String(), want)
		}
	})
}

func assertCustomTime(t *testing.T, got *Time, want time.Time) {
	t.Helper()

	if got.Unix != want.Unix() {
		t.Errorf("property unix is wrong, got '%c' want '%c'", got.Unix, want.Unix())
	}

	if got.Year != want.Year() {
		t.Errorf("property year is wrong, got '%c' want '%c'", got.Year, want.Year())
	}

	if got.Month != want.Month() {
		t.Errorf("property month is wrong, got '%s' want '%s'", got.Month, want.Month())
	}

	if got.Day != want.Day() {
		t.Errorf("property day is wrong, got '%c' want '%c'", got.Day, want.Day())
	}

	if got.Hour != want.Hour() {
		t.Errorf("property hour is wrong, got '%d' want '%d'", got.Hour, want.Hour())
	}

	if got.Minute != want.Minute() {
		t.Errorf("property minute is wrong, got '%c' want '%c'", got.Minute, want.Minute())
	}

	if got.Weekday != want.Weekday() {
		t.Errorf("property weekday is wrong, got '%s' want '%s'", got.Weekday, want.Weekday())
	}
}
