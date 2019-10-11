package configify

import (
	"strings"
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
	Uint(key string) (uint, bool)
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
	// environments shared with other processes/services that might have their own variables.
	Namespace string

	// NamespaceDelim defines the separator to use when composing a Namespace with a variable
	// name. For instance, if you have the Namespace "FOO" and you're calling `GetString("BAR")`
	// with the NamespaceDelim "." then we'll look up the attribute "FOO.BAR". This defaults to
	// underscore if not supplied explicitly.
	NamespaceDelim string

	// Defaults is an optional "fallback" for values when the source you're creating does not
	// contain the requested value. For instance if your source has values for "FOO" and "BAR"
	// but you request the value for "BAZ", your source will look in the Defaults source to see
	// what we should return as the default value. If you don't supply this, all defaults will
	// simply match sane defaults for the type (e.g. "" for strings, 0 for ints, etc)
	Defaults Source
}

// QualifyKey takes the non-qualified config attribute name (e.g. "PORT") and returns the fully
// qualified attribute name w/ its Namespace prepended (e.g. "MY_APP_PORT").
func (o Options) QualifyKey(key string) string {
	return o.JoinKey(o.NamespaceDelim, o.Namespace, key)
}

// JoinKey constructs a well-formed key by joining the given segments, ignoring any empty "". This
// will ensure that there are no consecutive delimiters or leading/trailing ones.
func (o Options) JoinKey(segments ...string) string {
	delim := strings.TrimSpace(o.NamespaceDelim)
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
