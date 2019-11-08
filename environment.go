package configify

import (
	"os"
	"strconv"
	"strings"
	"time"
)

// Environment creates a new config source that pull environment variables to provide configuration
// values. This source will also try to best-guess parse things like numbers since they're all
// natively strings.
func Environment(options Options) (Source, error) {
	if options.Defaults == nil {
		options.Defaults = emptySource{}
	}
	return &environmentSource{options: options}, nil
}

type environmentSource struct {
	options Options
}

func (e environmentSource) Options() Options {
	return e.options
}

func (e environmentSource) lookup(key string) (string, bool) {
	if value, ok := os.LookupEnv(e.options.Namespace.Qualify(key)); ok {
		return strings.TrimSpace(value), true
	}
	return "", false
}

func (e environmentSource) String(key string) (string, bool) {
	value, ok := e.lookup(key)
	if !ok {
		return e.options.Defaults.String(key)
	}
	return value, true
}

func (e environmentSource) StringSlice(key string) ([]string, bool) {
	value, ok := e.lookup(key)
	if !ok {
		return e.options.Defaults.StringSlice(key)
	}
	if value == "" {
		return []string{}, true
	}
	slice := strings.Split(value, ",")
	for i := range slice {
		slice[i] = strings.TrimSpace(slice[i])
	}
	return slice, true
}

func (e environmentSource) Int(key string) (int, bool) {
	value, ok := e.lookup(key)
	if !ok {
		return e.options.Defaults.Int(key)
	}
	number, ok := parseInt64(value, ',', '.')
	return int(number), ok
}

func (e environmentSource) Int8(key string) (int8, bool) {
	value, ok := e.lookup(key)
	if !ok {
		return e.options.Defaults.Int8(key)
	}
	number, ok := parseInt64(value, ',', '.')
	return int8(number), ok
}

func (e environmentSource) Int16(key string) (int16, bool) {
	value, ok := e.lookup(key)
	if !ok {
		return e.options.Defaults.Int16(key)
	}
	number, ok := parseInt64(value, ',', '.')
	return int16(number), ok
}

func (e environmentSource) Int32(key string) (int32, bool) {
	value, ok := e.lookup(key)
	if !ok {
		return e.options.Defaults.Int32(key)
	}
	number, ok := parseInt64(value, ',', '.')
	return int32(number), ok
}

func (e environmentSource) Int64(key string) (int64, bool) {
	value, ok := e.lookup(key)
	if !ok {
		return e.options.Defaults.Int64(key)
	}
	return parseInt64(value, ',', '.')
}

func (e environmentSource) Uint(key string) (uint, bool) {
	value, ok := e.lookup(key)
	if !ok {
		return e.options.Defaults.Uint(key)
	}
	number, ok := parseUint64(value, ',', '.')
	return uint(number), ok
}

func (e environmentSource) Uint8(key string) (uint8, bool) {
	value, ok := e.lookup(key)
	if !ok {
		return e.options.Defaults.Uint8(key)
	}
	number, ok := parseUint64(value, ',', '.')
	return uint8(number), ok
}

func (e environmentSource) Uint16(key string) (uint16, bool) {
	value, ok := e.lookup(key)
	if !ok {
		return e.options.Defaults.Uint16(key)
	}
	number, ok := parseUint64(value, ',', '.')
	return uint16(number), ok
}

func (e environmentSource) Uint32(key string) (uint32, bool) {
	value, ok := e.lookup(key)
	if !ok {
		return e.options.Defaults.Uint32(key)
	}
	number, ok := parseUint64(value, ',', '.')
	return uint32(number), ok
}

func (e environmentSource) Uint64(key string) (uint64, bool) {
	value, ok := e.lookup(key)
	if !ok {
		return e.options.Defaults.Uint64(key)
	}
	return parseUint64(value, ',', '.')
}

func (e environmentSource) Float32(key string) (float32, bool) {
	value, ok := e.lookup(key)
	if !ok {
		return e.options.Defaults.Float32(key)
	}
	number, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return float32(0), false
	}
	return float32(number), true
}

func (e environmentSource) Float64(key string) (float64, bool) {
	value, ok := e.lookup(key)
	if !ok {
		return e.options.Defaults.Float64(key)
	}
	number, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return float64(0), false
	}
	return number, true
}

func (e environmentSource) Bool(key string) (bool, bool) {
	value, ok := e.lookup(key)
	if !ok {
		return e.options.Defaults.Bool(key)
	}
	switch strings.ToLower(value) {
	case "true":
		return true, true
	case "false":
		return false, true
	default:
		return false, false
	}
}

func (e environmentSource) Duration(key string) (time.Duration, bool) {
	value, ok := e.lookup(key)
	if !ok {
		return e.options.Defaults.Duration(key)
	}
	duration, err := time.ParseDuration(value)
	if err != nil {
		return time.Duration(0), false
	}
	return duration, true
}

func (e environmentSource) Time(key string) (time.Time, bool) {
	value, ok := e.lookup(key)
	if !ok {
		return e.options.Defaults.Time(key)
	}
	dateTime, err := parseTime(value)
	if err != nil {
		return time.Time{}, false
	}
	return dateTime, true
}
