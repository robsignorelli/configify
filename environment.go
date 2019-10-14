package configify

import (
	"os"
	"strconv"
	"strings"
	"time"
)

func Environment(options Options) (Source, error) {
	if options.Defaults == nil {
		options.Defaults = Defaults{}
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
	if value, ok := os.LookupEnv(e.options.QualifyKey(key)); ok {
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
	number, err := strconv.ParseInt(normalizeInteger(value, ',', '.'), 10, 64)
	if err != nil {
		return 0, false
	}
	return int(number), true
}

func (e environmentSource) Uint(key string) (uint, bool) {
	value, ok := e.lookup(key)
	if !ok {
		return e.options.Defaults.Uint(key)
	}
	number, err := strconv.ParseUint(normalizeInteger(value, ',', '.'), 10, 64)
	if err != nil {
		return uint(0), false
	}
	return uint(number), true
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
