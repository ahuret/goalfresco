package time

import (
	"strconv"
	"time"
	gotime "time"
)

// custom unmarshaller purpose

var timeformats = []string{
	"2006-01-02T15:04:05.999Z0700",
	time.RFC3339,
}

type Time struct {
	gotime.Time
}

func (t *Time) UnmarshalJSON(data []byte) error {
	d, err := strconv.Unquote(string(data))
	if err != nil {
		return err
	}
	return t.Parse(d)
}

func (t *Time) Parse(data string) error {
	var err error
	var ti gotime.Time
	for i := 0; (ti == gotime.Time{}) && i < len(timeformats); i++ {
		ti, err = gotime.Parse(timeformats[i], data)
	}
	if err != nil {
		return err
	}
	*t = Time{ti}
	return nil
}

// map some functions, methods and types of golang builtin time.Time library used in openapi generated code

const (
	RFC3339 = gotime.RFC3339
	RFC1123 = gotime.RFC1123
)

func (t Time) Add(d gotime.Duration) Time {
	t.Time = t.Time.Add(d)
	return t
}

func Parse(layout, value string) (Time, error) {
	t, err := gotime.Parse(layout, value)
	if err != nil {
		return Time{}, err
	}
	return Time{t}, nil
}

func Now() Time {
	return Time{gotime.Now()}
}

func ParseDuration(s string) (gotime.Duration, error) {
	return gotime.ParseDuration(s)
}
