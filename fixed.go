package configify

import "strings"

// Fixed creates a hard-coded map of config values. You can use these as a source on their own
// or you can provide them as the fallback defaults for other sources.
func Fixed(options Options, values map[string]interface{}) Source {
	if options.Defaults == nil {
		options.Defaults = defaults{}
	}

	// Make sure that all of the keys are properly namespaced. This isn't a big deal when using
	// one of these as a standalone source, but when you use it as the fallback for some other source
	// you want to make sure you're working within the same namespace. And, you know... consistency.
	qualified := make(map[string]interface{}, len(values))
	for key, value := range values {
		qualified[options.ResolveKey(key)] = value
	}

	return &fixedSource{
		options: options,
		values:  qualified,
	}
}

type fixedSource struct {
	options Options
	values  map[string]interface{}
}

func (s fixedSource) GetString(key string) string {
	if val, ok := s.values[s.options.ResolveKey(key)].(string); ok {
		return strings.TrimSpace(val)
	}
	return s.options.Defaults.GetString(key)
}

func (s fixedSource) GetStringSlice(key string) []string {
	if val, ok := s.values[s.options.ResolveKey(key)].([]string); ok {
		return val
	}
	return s.options.Defaults.GetStringSlice(key)
}

func (s fixedSource) GetInt(key string) int {
	if val, ok := s.values[s.options.ResolveKey(key)].(int); ok {
		return val
	}
	return s.options.Defaults.GetInt(key)
}

func (s fixedSource) GetUint(key string) uint {
	if val, ok := s.values[s.options.ResolveKey(key)].(uint); ok {
		return val
	}
	return s.options.Defaults.GetUint(key)
}
