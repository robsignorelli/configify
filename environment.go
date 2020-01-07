package configify

import (
	"os"
	"strings"
	"time"
)

// Environment creates a new config source that pull environment variables to provide configuration
// values. This source will also try to best-guess parse things like numbers since they're all
// natively strings.
func Environment(opts ...Option) (Source, error) {
	options := apply(opts, &Options{
		Defaults: emptySource{},
	})
	return &environmentSource{options: *options, massage: Massage{}}, nil
}

type environmentSource struct {
	options Options
	massage Massage
}

func (e environmentSource) Options() Options {
	return e.options
}

func (e environmentSource) lookup(key string) (string, bool) {
	if value, ok := os.LookupEnv(e.options.Namespace.Qualify(key)); ok {
		return strings.TrimSpace(value), true
	}
	return "", false
}

func (e environmentSource) String(key string) (string, bool) {
	value, ok := e.lookup(key)
	if !ok {
		return e.options.Defaults.String(key)
	}
	return value, true
}

func (e environmentSource) StringSlice(key string) ([]string, bool) {
	value, ok := e.lookup(key)
	if !ok {
		return e.options.Defaults.StringSlice(key)
	}
	return e.massage.StringToSlice(value)
}

func (e environmentSource) Int(key string) (int, bool) {
	value, ok := e.lookup(key)
	if !ok {
		return e.options.Defaults.Int(key)
	}
	number, ok := e.massage.StringToInt64(value)
	return int(number), ok
}

func (e environmentSource) Int8(key string) (int8, bool) {
	value, ok := e.lookup(key)
	if !ok {
		return e.options.Defaults.Int8(key)
	}
	number, ok := e.massage.StringToInt64(value)
	return int8(number), ok
}

func (e environmentSource) Int16(key string) (int16, bool) {
	value, ok := e.lookup(key)
	if !ok {
		return e.options.Defaults.Int16(key)
	}
	number, ok := e.massage.StringToInt64(value)
	return int16(number), ok
}

func (e environmentSource) Int32(key string) (int32, bool) {
	value, ok := e.lookup(key)
	if !ok {
		return e.options.Defaults.Int32(key)
	}
	number, ok := e.massage.StringToInt64(value)
	return int32(number), ok
}

func (e environmentSource) Int64(key string) (int64, bool) {
	value, ok := e.lookup(key)
	if !ok {
		return e.options.Defaults.Int64(key)
	}
	return e.massage.StringToInt64(value)
}

func (e environmentSource) Uint(key string) (uint, bool) {
	value, ok := e.lookup(key)
	if !ok {
		return e.options.Defaults.Uint(key)
	}
	number, ok := e.massage.StringToUint64(value)
	return uint(number), ok
}

func (e environmentSource) Uint8(key string) (uint8, bool) {
	value, ok := e.lookup(key)
	if !ok {
		return e.options.Defaults.Uint8(key)
	}
	number, ok := e.massage.StringToUint64(value)
	return uint8(number), ok
}

func (e environmentSource) Uint16(key string) (uint16, bool) {
	value, ok := e.lookup(key)
	if !ok {
		return e.options.Defaults.Uint16(key)
	}
	number, ok := e.massage.StringToUint64(value)
	return uint16(number), ok
}

func (e environmentSource) Uint32(key string) (uint32, bool) {
	value, ok := e.lookup(key)
	if !ok {
		return e.options.Defaults.Uint32(key)
	}
	number, ok := e.massage.StringToUint64(value)
	return uint32(number), ok
}

func (e environmentSource) Uint64(key string) (uint64, bool) {
	value, ok := e.lookup(key)
	if !ok {
		return e.options.Defaults.Uint64(key)
	}
	return e.massage.StringToUint64(value)
}

func (e environmentSource) Float32(key string) (float32, bool) {
	value, ok := e.lookup(key)
	if !ok {
		return e.options.Defaults.Float32(key)
	}
	number, ok := e.massage.StringToFloat64(value)
	return float32(number), ok
}

func (e environmentSource) Float64(key string) (float64, bool) {
	value, ok := e.lookup(key)
	if !ok {
		return e.options.Defaults.Float64(key)
	}
	return e.massage.StringToFloat64(value)
}

func (e environmentSource) Bool(key string) (bool, bool) {
	value, ok := e.lookup(key)
	if !ok {
		return e.options.Defaults.Bool(key)
	}
	return e.massage.StringToBool(value)
}

func (e environmentSource) Duration(key string) (time.Duration, bool) {
	value, ok := e.lookup(key)
	if !ok {
		return e.options.Defaults.Duration(key)
	}
	return e.massage.StringToDuration(value)
}

func (e environmentSource) Time(key string) (time.Time, bool) {
	value, ok := e.lookup(key)
	if !ok {
		return e.options.Defaults.Time(key)
	}
	return e.massage.StringToTime(value)
}
