package configify

import (
	"strings"
	"time"
)

// Fixed creates a hard-coded map of config values. You can use these as a source on their own
// or you can provide them as the fallback defaults for other sources.
func Fixed(options Options, values Values) Source {
	if options.Defaults == nil {
		options.Defaults = Defaults{}
	}

	// Make sure that all of the keys are properly namespaced. This isn't a big deal when using
	// one of these as a standalone source, but when you use it as the fallback for some other source
	// you want to make sure you're working within the same namespace. And, you know... consistency.
	qualified := make(map[string]interface{}, len(values))
	for key, value := range values {
		qualified[options.QualifyKey(key)] = value
	}

	return &fixedSource{
		options: options,
		values:  qualified,
	}
}

type fixedSource struct {
	options Options
	values  Values
}

func (s fixedSource) Options() Options {
	return s.options
}
func (s fixedSource) String(key string) (string, bool) {
	if val, ok := s.values[s.options.QualifyKey(key)].(string); ok {
		return strings.TrimSpace(val), true
	}
	return s.options.Defaults.String(key)
}

func (s fixedSource) StringSlice(key string) ([]string, bool) {
	if val, ok := s.values[s.options.QualifyKey(key)].([]string); ok {
		return val, true
	}
	return s.options.Defaults.StringSlice(key)
}

func (s fixedSource) Int(key string) (int, bool) {
	if val, ok := s.values[s.options.QualifyKey(key)].(int); ok {
		return val, true
	}
	return s.options.Defaults.Int(key)
}

func (s fixedSource) Uint(key string) (uint, bool) {
	if val, ok := s.values[s.options.QualifyKey(key)].(uint); ok {
		return val, true
	}
	return s.options.Defaults.Uint(key)
}

func (s fixedSource) Duration(key string) (time.Duration, bool) {
	if val, ok := s.values[s.options.QualifyKey(key)].(time.Duration); ok {
		return val, true
	}
	return s.options.Defaults.Duration(key)
}

func (s fixedSource) Time(key string) (time.Time, bool) {
	if val, ok := s.values[s.options.QualifyKey(key)].(time.Time); ok {
		return val, true
	}
	return s.options.Defaults.Time(key)
}
