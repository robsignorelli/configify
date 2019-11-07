package configify

import (
	"errors"
	"time"
)

var emptyStringSlice = make([]string, 0)

// Defaults lets us "dog food" our own Source representation so that when some "real" source
// fails to resolve a value, you can always fall back to this source which is nothing but hard-coded
// values such as "all ints default to 0" and "all strings default to empty".
type Defaults struct{}

func (Defaults) Options() Options {
	return Options{}
}

func (Defaults) String(string) (string, bool) {
	return "", false
}

func (Defaults) StringSlice(string) ([]string, bool) {
	return emptyStringSlice, false
}

func (Defaults) Int(string) (int, bool) {
	return 0, false
}

func (Defaults) Int8(string) (int8, bool) {
	return int8(0), false
}

func (Defaults) Int16(string) (int16, bool) {
	return int16(0), false
}

func (Defaults) Int32(string) (int32, bool) {
	return int32(0), false
}

func (Defaults) Int64(string) (int64, bool) {
	return int64(0), false
}

func (Defaults) Uint(string) (uint, bool) {
	return uint(0), false
}

func (Defaults) Uint8(string) (uint8, bool) {
	return uint8(0), false
}

func (Defaults) Uint16(string) (uint16, bool) {
	return uint16(0), false
}

func (Defaults) Uint32(string) (uint32, bool) {
	return uint32(0), false
}

func (Defaults) Uint64(string) (uint64, bool) {
	return uint64(0), false
}

func (Defaults) Float32(string) (float32, bool) {
	return float32(0), false
}

func (Defaults) Float64(string) (float64, bool) {
	return float64(0), false
}

func (Defaults) Bool(string) (bool, bool) {
	return false, false
}

func (Defaults) Duration(string) (time.Duration, bool) {
	return time.Duration(0), false
}

func (Defaults) Time(string) (time.Time, bool) {
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
