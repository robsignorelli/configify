package configify

import "strings"

// FixedSource creates a hard-coded map of config values. You can use these as a source on their own
// or you can provide them as the fallback defaults for other sources.
func FixedSource(namespace string, values map[string]interface{}) Source {
	// Make sure that all of the keys are properly namespaced. This isn't a big deal when using
	// one of these as a standalone source, but when you use it as the fallback for some other source
	// you want to make sure you're working within the same namespace. And, you know... consistency.
	qualified := make(map[string]interface{}, len(values))
	for key, value := range values {
		qualified[joinKey(namespace, key)] = value
	}

	return &fixedSource{
		namespace: namespace,
		defaults:  qualified,
	}
}

type fixedSource struct {
	namespace string
	defaults  map[string]interface{}
}

func (s fixedSource) GetString(key string) string {
	if val, ok := s.defaults[joinKey(s.namespace, key)].(string); ok {
		return strings.TrimSpace(val)
	}
	return ""
}

func (s fixedSource) GetStringSlice(key string) []string {
	if val, ok := s.defaults[joinKey(s.namespace, key)].([]string); ok {
		return val
	}
	return emptyStringSlice
}

func (s fixedSource) GetInt(key string) int {
	if val, ok := s.defaults[joinKey(s.namespace, key)].(int); ok {
		return val
	}
	return 0
}

func (s fixedSource) GetUint(key string) uint {
	if val, ok := s.defaults[joinKey(s.namespace, key)].(uint); ok {
		return val
	}
	return uint(0)
}
