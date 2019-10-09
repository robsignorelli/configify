package configify

import (
	"os"
	"strconv"
	"strings"
)

func Environment(options Options) (SourceBinder, error) {
	if options.Defaults == nil {
		options.Defaults = defaults{}
	}
	return &environmentSource{options: options}, nil
}

type environmentSource struct {
	options Options
}

func (e environmentSource) lookup(key string) string {
	return strings.TrimSpace(os.Getenv(e.options.ResolveKey(key)))
}

func (e environmentSource) GetString(key string) string {
	if value := e.lookup(key); value != "" {
		return value
	}
	return e.options.Defaults.GetString(key)
}

func (e environmentSource) GetStringSlice(key string) []string {
	value := e.lookup(key)
	if value == "" {
		return e.options.Defaults.GetStringSlice(key)
	}
	slice := strings.Split(value, ",")
	for i := range slice {
		slice[i] = strings.TrimSpace(slice[i])
	}
	return slice
}

func (e environmentSource) GetInt(key string) int {
	value := e.lookup(key)
	if value == "" {
		return e.options.Defaults.GetInt(key)
	}
	number, err := strconv.ParseInt(normalizeInteger(value, ',', '.'), 10, 64)
	if err != nil {
		return e.options.Defaults.GetInt(key)
	}
	return int(number)
}

func (e environmentSource) GetUint(key string) uint {
	value := e.lookup(key)
	if value == "" {
		return e.options.Defaults.GetUint(key)
	}
	number, err := strconv.ParseUint(normalizeInteger(value, ',', '.'), 10, 64)
	if err != nil {
		return e.options.Defaults.GetUint(key)
	}
	return uint(number)
}
