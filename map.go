package configify

import (
	"strings"
	"time"
)

// Map creates a hard-coded map of config values. You can use these as a source on their own
// or you can provide them as the fallback defaults for other sources.
func Map(values Values) Source {
	return &mapSource{values: values}
}

type mapSource struct {
	values Values
}

func (s mapSource) Options() Options {
	return Options{}
}

func (s mapSource) String(key string) (string, bool) {
	if val, ok := s.values[key].(string); ok {
		return strings.TrimSpace(val), true
	}
	return "", false
}

func (s mapSource) StringSlice(key string) ([]string, bool) {
	if val, ok := s.values[key].([]string); ok {
		return val, true
	}
	return nil, false
}

func (s mapSource) Int(key string) (int, bool) {
	if val, ok := s.values[key].(int); ok {
		return val, true
	}
	return 0, false
}

func (s mapSource) Int8(key string) (int8, bool) {
	if val, ok := s.values[key].(int8); ok {
		return val, true
	}
	return 0, false
}

func (s mapSource) Int16(key string) (int16, bool) {
	if val, ok := s.values[key].(int16); ok {
		return val, true
	}
	return 0, false
}

func (s mapSource) Int32(key string) (int32, bool) {
	if val, ok := s.values[key].(int32); ok {
		return val, true
	}
	return 0, false
}

func (s mapSource) Int64(key string) (int64, bool) {
	if val, ok := s.values[key].(int64); ok {
		return val, true
	}
	return 0, false
}

func (s mapSource) Uint(key string) (uint, bool) {
	if val, ok := s.values[key].(uint); ok {
		return val, true
	}
	return uint(0), false
}

func (s mapSource) Uint8(key string) (uint8, bool) {
	if val, ok := s.values[key].(uint8); ok {
		return val, true
	}
	return 0, false
}

func (s mapSource) Uint16(key string) (uint16, bool) {
	if val, ok := s.values[key].(uint16); ok {
		return val, true
	}
	return 0, false
}

func (s mapSource) Uint32(key string) (uint32, bool) {
	if val, ok := s.values[key].(uint32); ok {
		return val, true
	}
	return 0, false
}

func (s mapSource) Uint64(key string) (uint64, bool) {
	if val, ok := s.values[key].(uint64); ok {
		return val, true
	}
	return 0, false
}

func (s mapSource) Float64(key string) (float64, bool) {
	if val, ok := s.values[key].(float64); ok {
		return val, true
	}
	return 0.0, false
}

func (s mapSource) Float32(key string) (float32, bool) {
	if val, ok := s.values[key].(float32); ok {
		return float32(val), true
	}
	return 0.0, false
}

func (s mapSource) Bool(key string) (bool, bool) {
	if val, ok := s.values[key].(bool); ok {
		return val, true
	}
	return false, false
}

func (s mapSource) Duration(key string) (time.Duration, bool) {
	if val, ok := s.values[key].(time.Duration); ok {
		return val, true
	}
	return time.Duration(0), false
}

func (s mapSource) Time(key string) (time.Time, bool) {
	if val, ok := s.values[key].(time.Time); ok {
		return val, true
	}
	return time.Time{}, false
}
