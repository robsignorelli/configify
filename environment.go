package configify

import (
	"os"
	"strconv"
	"strings"
)

func EnvironmentSource(prefix string) SourceBinder {
	return &environmentSource{prefix: prefix}
}

type environmentSource struct {
	prefix string
}

func (e environmentSource) lookup(key string) string {
	return strings.TrimSpace(os.Getenv(joinKey(e.prefix, key)))
}

func (e environmentSource) GetString(key string) string {
	return e.lookup(key)
}

func (e environmentSource) GetStringSlice(key string) []string {
	value := e.lookup(key)
	if value == "" {
		return emptyStringSlice
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
		return 0
	}
	number, err := strconv.ParseInt(normalizeInteger(value, ',', '.'), 10, 64)
	if err != nil {
		return 0
	}
	return int(number)
}

func (e environmentSource) GetUint(key string) uint {
	value := e.lookup(key)
	if value == "" {
		return 0
	}
	number, err := strconv.ParseUint(normalizeInteger(value, ',', '.'), 10, 64)
	if err != nil {
		return 0
	}
	return uint(number)
}
