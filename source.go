package configify

import (
	"strings"
)

var emptyStringSlice = make([]string, 0)

// Source represents a... well... source of configuration values. This can be a file,
// a struct, or your environment variables. The idea is that you extract values from a
// source and put them in your config structures to use in your program. Or you can simply
// fetch individual values on demand if you prefer.
type Source interface {
	GetString(key string) string
	GetStringSlice(key string) []string
	GetInt(key string) int
	GetUint(key string) uint
}

// Binder defines the ability to overlay config values onto some existing structure. This way
// in one fell swoop you can apply all of your environment variables to your config structures.
type Binder interface {
}

type SourceBinder interface {
	Source
	Binder
}

func joinKey(components... string) string {
	return strings.Join(components, "_")
}

func normalizeInteger(value string, groupSep rune, decimalSep rune) string {
	decimalPos := strings.IndexRune(value, decimalSep)
	if decimalPos == 0 {
		return ""
	}
	if decimalPos > 0 {
		value = value[:decimalPos]
	}
	return strings.ReplaceAll(value, string(groupSep), "")
}