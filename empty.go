package configify

import (
	"errors"
	"time"
)

var emptyStringSlice = make([]string, 0)

// Empty is a nop source that always returns the sane default for any type; "" for
// string values, 0 for ints, and so on.
func Empty() Source {
	return &emptySource{}
}

type emptySource struct{}

func (emptySource) Options() Options {
	return Options{}
}

func (emptySource) String(string) (string, bool) {
	return "", false
}

func (emptySource) StringSlice(string) ([]string, bool) {
	return emptyStringSlice, false
}

func (emptySource) Int(string) (int, bool) {
	return 0, false
}

func (emptySource) Int8(string) (int8, bool) {
	return int8(0), false
}

func (emptySource) Int16(string) (int16, bool) {
	return int16(0), false
}

func (emptySource) Int32(string) (int32, bool) {
	return int32(0), false
}

func (emptySource) Int64(string) (int64, bool) {
	return int64(0), false
}

func (emptySource) Uint(string) (uint, bool) {
	return uint(0), false
}

func (emptySource) Uint8(string) (uint8, bool) {
	return uint8(0), false
}

func (emptySource) Uint16(string) (uint16, bool) {
	return uint16(0), false
}

func (emptySource) Uint32(string) (uint32, bool) {
	return uint32(0), false
}

func (emptySource) Uint64(string) (uint64, bool) {
	return uint64(0), false
}

func (emptySource) Float32(string) (float32, bool) {
	return float32(0), false
}

func (emptySource) Float64(string) (float64, bool) {
	return float64(0), false
}

func (emptySource) Bool(string) (bool, bool) {
	return false, false
}

func (emptySource) Duration(string) (time.Duration, bool) {
	return time.Duration(0), false
}

func (emptySource) Time(string) (time.Time, bool) {
	return time.Time{}, false
}

func parseTime(input string) (time.Time, error) {
	switch len(input) {
	case 0:
		return time.Time{}, errors.New("invalid time: empty")
	case 10:
		return time.Parse("2006-01-02", input)
	default:
		return time.Parse(time.RFC3339, input)
	}
}
