package configify

import (
	"reflect"
	"regexp"
	"strings"
	"time"
)

var typeDuration = reflect.TypeOf(time.Duration(0))
var typeTime = reflect.TypeOf(time.Time{})

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
		Source:      source,
		emptySource: emptySource{},
	}
}

type standardBinder struct {
	Source
	emptySource
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

func (b standardBinder) bindPrefixWithType(_ interface{}, prefix string, outType reflect.Type, outValue reflect.Value) {
	for i := 0; i < outType.NumField(); i++ {
		field := outType.Field(i)
		value := outValue.Field(i)
		key := b.Source.Options().Namespace.Join(prefix, b.resolveName(field))

		b.updateValue(field, value, key)
	}
}

func (b standardBinder) updateValue(field reflect.StructField, value reflect.Value, key string) {
	// There are a couple of common types we support that aren't built-ins, so check those first
	switch field.Type {
	case typeDuration:
		if v, ok := b.Source.Duration(key); ok {
			value.Set(reflect.ValueOf(v))
		}
		return
	case typeTime:
		if v, ok := b.Source.Time(key); ok {
			value.Set(reflect.ValueOf(v))
		}
		return
	}

	switch field.Type.Kind() {
	case reflect.String:
		if v, ok := b.Source.String(key); ok {
			value.SetString(v)
		}
	case reflect.Bool:
		if v, ok := b.Source.Bool(key); ok {
			value.SetBool(v)
		}
	case reflect.Int:
		if v, ok := b.Source.Int(key); ok {
			value.SetInt(int64(v))
		}
	case reflect.Int8:
		if v, ok := b.Source.Int8(key); ok {
			value.SetInt(int64(v))
		}
	case reflect.Int16:
		if v, ok := b.Source.Int16(key); ok {
			value.SetInt(int64(v))
		}
	case reflect.Int32:
		if v, ok := b.Source.Int32(key); ok {
			value.SetInt(int64(v))
		}
	case reflect.Int64:
		if v, ok := b.Source.Int64(key); ok {
			value.SetInt(v)
		}
	case reflect.Uint:
		if v, ok := b.Source.Uint(key); ok {
			value.SetUint(uint64(v))
		}
	case reflect.Uint8:
		if v, ok := b.Source.Uint8(key); ok {
			value.SetUint(uint64(v))
		}
	case reflect.Uint16:
		if v, ok := b.Source.Uint16(key); ok {
			value.SetUint(uint64(v))
		}
	case reflect.Uint32:
		if v, ok := b.Source.Uint32(key); ok {
			value.SetUint(uint64(v))
		}
	case reflect.Uint64:
		if v, ok := b.Source.Uint64(key); ok {
			value.SetUint(v)
		}
	case reflect.Float32:
		if v, ok := b.Source.Float32(key); ok {
			value.SetFloat(float64(v))
		}
	case reflect.Float64:
		if v, ok := b.Source.Float64(key); ok {
			value.SetFloat(v)
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
	// There are a couple of common types we support that aren't built-ins, so check those first
	switch field.Type.Elem() {
	case typeDuration:
		if v, ok := b.Source.Duration(key); ok {
			value.Set(reflect.ValueOf(&v))
		}
		return
	case typeTime:
		if v, ok := b.Source.Time(key); ok {
			value.Set(reflect.ValueOf(&v))
		}
		return
	}

	switch field.Type.Elem().Kind() {
	case reflect.String:
		if v, ok := b.Source.String(key); ok {
			value.Set(reflect.ValueOf(&v))
		}
	case reflect.Bool:
		if v, ok := b.Source.Bool(key); ok {
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
	case reflect.Float64:
		if v, ok := b.Source.Float64(key); ok {
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
