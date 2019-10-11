package configify

import (
	"reflect"
	"regexp"
	"strings"
)

type Binder interface {
	Bind(out interface{})
}

func NewBinder(source Source) Binder {
	return &binder{
		Source:   source,
		Defaults: Defaults{},
	}
}

type binder struct {
	Source
	Defaults
}

func (b binder) Bind(out interface{}) {
	b.BindPrefix(out, "")
}

func (b binder) BindPrefix(out interface{}, prefix string) {
	outType := reflect.TypeOf(out).Elem()
	outValue := reflect.ValueOf(out).Elem()

	for i := 0; i < outType.NumField(); i++ {
		field := outType.Field(i)
		value := outValue.Field(i)
		key := b.Source.Options().JoinKey(prefix, b.resolveName(field))
		b.updateField(field, value, key)
	}
}

func (b binder) updateField(field reflect.StructField, value reflect.Value, key string) {
	switch field.Type.Kind() {
	case reflect.Slice:
		b.updateSlice(field, value, key)
	case reflect.Struct:
		b.BindPrefix(b.pointerTo(value), key)
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
	}
}

func (b binder) updateSlice(field reflect.StructField, value reflect.Value, key string) {
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
func (b binder) resolveName(field reflect.StructField) string {
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

// pointerTo is utilized when we perform a recursive binding on an attribute which is another
// struct. This ensures that we're passing a pointer to the struct so modifications actually
// update the original instance, not a copy.
func (b binder) pointerTo(value reflect.Value) interface{} {
	if value.Elem().Kind() == reflect.Ptr {
		return value.Interface()
	}
	return value.Addr().Interface()
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
