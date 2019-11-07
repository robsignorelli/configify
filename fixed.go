package configify

import (
	"strings"
	"time"
)

// Fixed creates a hard-coded map of config values. You can use these as a source on their own
// or you can provide them as the fallback defaults for other sources.
func Fixed(values Values) Source {
	return &fixedSource{values: values}
}

type fixedSource struct {
	values Values
}

func (s fixedSource) Options() Options {
	return Options{}
}

func (s fixedSource) String(key string) (string, bool) {
	if val, ok := s.values[key].(string); ok {
		return strings.TrimSpace(val), true
	}
	return "", false
}

func (s fixedSource) StringSlice(key string) ([]string, bool) {
	if val, ok := s.values[key].([]string); ok {
		return val, true
	}
	return nil, false
}

func (s fixedSource) Int(key string) (int, bool) {
	if val, ok := s.values[key].(int); ok {
		return val, true
	}
	return 0, false
}

func (s fixedSource) Uint(key string) (uint, bool) {
	if val, ok := s.values[key].(uint); ok {
		return val, true
	}
	return uint(0), false
}

func (s fixedSource) Float(key string) (float64, bool) {
	if val, ok := s.values[key].(float64); ok {
		return val, true
	}
	return 0.0, false
}

func (s fixedSource) Bool(key string) (bool, bool) {
	if val, ok := s.values[key].(bool); ok {
		return val, true
	}
	return false, false
}

func (s fixedSource) Duration(key string) (time.Duration, bool) {
	if val, ok := s.values[key].(time.Duration); ok {
		return val, true
	}
	return time.Duration(0), false
}

func (s fixedSource) Time(key string) (time.Time, bool) {
	if val, ok := s.values[key].(time.Time); ok {
		return val, true
	}
	return time.Time{}, false
}
