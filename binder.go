package configify

import (
	"reflect"
	"regexp"
	"strings"
)

// Binder defines a component that can overlay config/source values onto an existing struct of yours.
// This lets you have strongly-typed config struct instances in your code that you can pass around
// rather than pulling off individual config values one by one from a source.
type Binder interface {
	Bind(out interface{})
}

// NewBinder creates the standard binder which maps values from your Source to the fields on
// the struct you want to populate.
func NewBinder(source Source) Binder {
	return &standardBinder{
		Source:   source,
		Defaults: Defaults{},
	}
}

type standardBinder struct {
	Source
	Defaults
}

func (b standardBinder) Bind(out interface{}) {
	if b.Source != nil {
		b.bindPrefix(out, "")
	}
}

func (b standardBinder) bindPrefix(out interface{}, prefix string) {
	outType := reflect.TypeOf(out).Elem()
	outValue := reflect.ValueOf(out).Elem()
	b.bindPrefixWithType(out, prefix, outType, outValue)
}

func (b standardBinder) bindPrefixWithType(out interface{}, prefix string, outType reflect.Type, outValue reflect.Value) {
	for i := 0; i < outType.NumField(); i++ {
		field := outType.Field(i)
		value := outValue.Field(i)
		key := b.Source.Options().JoinKey(prefix, b.resolveName(field))
		b.updateValue(field, value, key)
	}
}

func (b standardBinder) updateValue(field reflect.StructField, value reflect.Value, key string) {
	switch field.Type.Kind() {
	case reflect.String:
		if v, ok := b.Source.String(key); ok {
			value.SetString(v)
		}
	case reflect.Int:
		if v, ok := b.Source.Int(key); ok {
			value.SetInt(int64(v))
		}
	case reflect.Uint:
		if v, ok := b.Source.Uint(key); ok {
			value.SetUint(uint64(v))
		}
	case reflect.Struct:
		b.bindPrefix(value.Addr().Interface(), key)
	case reflect.Slice:
		b.updateSlice(field, value, key)
	case reflect.Ptr:
		b.updatePointer(field, value, key)
	}
}

func (b standardBinder) updatePointer(field reflect.StructField, value reflect.Value, key string) {
	switch field.Type.Elem().Kind() {
	case reflect.String:
		if v, ok := b.Source.String(key); ok {
			value.Set(reflect.ValueOf(&v))
		}
	case reflect.Int:
		if v, ok := b.Source.Int(key); ok {
			value.Set(reflect.ValueOf(&v))
		}
	case reflect.Uint:
		if v, ok := b.Source.Uint(key); ok {
			value.Set(reflect.ValueOf(&v))
		}
	case reflect.Struct:
		// Currently, we only support recursion into struct pointers if your input already
		// has a non-nil value for it. I'm not 100% sure on the semantics of how this should
		// work. If you come in null but have no values in your source to fill it, should we
		// leave the pointer nil? If so, how should we detect if anything down the chain was
		// supplied? Or should we supply a non-nil value for the struct with all default
		// values for its fields? I don't have a strong pull in either direction, so we can
		// leave this as a future enhancement.
		if value.Pointer() != 0 {
			b.bindPrefix(value.Interface(), key)
		}
	}
}

func (b standardBinder) updateSlice(field reflect.StructField, value reflect.Value, key string) {
	// Determine what this is a slice of and invoke the appropriate slice getter on the source.
	switch field.Type.Elem().Kind() {
	case reflect.String:
		if v, ok := b.Source.StringSlice(key); ok {
			value.Set(reflect.ValueOf(v))
		}
	}
}

// resolveName looks at a struct field/attribute and determines the config key we should use to
// try and look up its value. We'll first attempt to locate the 'conf' tag in case you defined
// a specific name. Otherwise, we'll just use the upper-snake-cased version of the attribute name.
func (b standardBinder) resolveName(field reflect.StructField) string {
	if name := field.Tag.Get("conf"); name != "" {
		return name
	}

	// Convert camel-cased field names to upper snake case by default.
	// For example "FooBar" will become "FOO_BAR". If you don't like it,
	// use the 'conf' tags.
	snake := matchFirstCap.ReplaceAllString(field.Name, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToUpper(snake)
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
