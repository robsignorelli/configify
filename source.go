package configify

import (
	"strconv"
	"strings"
	"time"
)

// Source represents a... well... source of configuration values. This can be a file,
// a struct, or your environment variables. The idea is that you extract values from a
// source and put them in your config structures to use in your program. Or you can simply
// fetch individual values on demand if you prefer.
type Source interface {
	Options() Options
	String(key string) (string, bool)
	StringSlice(key string) ([]string, bool)
	Int(key string) (int, bool)
	Int8(key string) (int8, bool)
	Int16(key string) (int16, bool)
	Int32(key string) (int32, bool)
	Int64(key string) (int64, bool)
	Uint(key string) (uint, bool)
	Uint8(key string) (uint8, bool)
	Uint16(key string) (uint16, bool)
	Uint32(key string) (uint32, bool)
	Uint64(key string) (uint64, bool)
	Float32(key string) (float32, bool)
	Float64(key string) (float64, bool)
	Bool(key string) (bool, bool)
	Duration(key string) (time.Duration, bool)
	Time(key string) (time.Time, bool)
}

// Options encapsulate the standard attributes that are shared by all sources in configify.
type Options struct {
	// Namespace is an optional prefix for all of your config values. This is useful for cases
	// like environment variable names that might collide. For instance instead of the variable
	// name "NAME", you should define "MY_APP_NAME" in your environment. In configify, you would
	// use the namespace "MY_APP" for your source and then you can look up "NAME" in your code.
	// Under the hood, configify will look up "MY_APP_NAME" for you.
	//
	// This helps you enforce good variable naming practices especially when running in
	// environments shared with other processes/services/components that might have their own
	// similarly named variables.
	Namespace Namespace

	// Defaults is an optional "fallback" for values when the source you're creating does not
	// contain the requested value. For instance if your source has values for "FOO" and "BAR"
	// but you request the value for "BAZ", your source will look in the Defaults source to see
	// what we should return as the default value. If you don't supply this, all defaults will
	// simply match sane defaults for the type (e.g. "" for strings, 0 for ints, etc)
	Defaults Source
}

// Namespace defines a fixed prefix for keys in your config store. This helps you isolate your
// config values to certain services or components. For instance for all HTTP router configuration
// you can use the namespace "HTTP" or for the configs for your RabbitMQ component, you can use
// the namespace "RABBITMQ", and so on.
type Namespace struct {
	Name      string
	Delimiter string
}

// Qualify takes the raw, unqualified key name (e.g. "PORT") and returns the namespace-qualified
// key name (e.g. "HTTP_PORT").
func (ns Namespace) Qualify(key string) string {
	if ns.Name == "" {
		return key
	}
	return ns.Join(ns.Name, key)
}

// Join constructs a well-formed value by joining the given segments, ignoring any empty "". This
// will ensure that there are no consecutive delimiters or leading/trailing ones. This does NOT
// force the namespace name as a prefix!
func (ns Namespace) Join(segments ...string) string {
	delim := strings.TrimSpace(ns.Delimiter)
	if delim == "" {
		delim = "_"
	}
	var goodSegments []string
	for _, segment := range segments {
		if segment = strings.TrimSpace(segment); segment != "" {
			goodSegments = append(goodSegments, segment)
		}
	}
	return strings.Join(goodSegments, delim)
}

// Values represents a set of key/value pairs as a map.
type Values map[string]interface{}

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

func parseInt64(value string, groupSep rune, decimalSep rune) (int64, bool) {
	number, err := strconv.ParseInt(normalizeInteger(value, groupSep, decimalSep), 10, 64)
	if err != nil {
		return 0, false
	}
	return number, true
}

func parseUint64(value string, groupSep rune, decimalSep rune) (uint64, bool) {
	number, err := strconv.ParseUint(normalizeInteger(value, groupSep, decimalSep), 10, 64)
	if err != nil {
		return 0, false
	}
	return number, true
}
