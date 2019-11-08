package configify_test

import (
	"time"

	"github.com/robsignorelli/configify"
	"github.com/stretchr/testify/mock"
)

// NewMockSource creates a mock source that lets you instruct it exactly how to
// respond for specific inputs.
func NewMockSource(setup func(*MockSource)) configify.Source {
	// We call setup() before using `mock.Anything` so that testify will try to apply those
	// matchers first on invocation. So the stuff you set up takes precedence over these defaults.
	s := new(MockSource)
	setup(s)
	s.On("Options", mock.Anything).Return(configify.Options{})
	s.On("String", mock.Anything).Return("", false)
	s.On("StringSlice", mock.Anything).Return([]string{}, false)
	s.On("Int", mock.Anything).Return(0, false)
	s.On("Int8", mock.Anything).Return(int8(0), false)
	s.On("Int16", mock.Anything).Return(int16(0), false)
	s.On("Int32", mock.Anything).Return(int32(0), false)
	s.On("Int64", mock.Anything).Return(int64(0), false)
	s.On("Uint", mock.Anything).Return(uint(0), false)
	s.On("Uint8", mock.Anything).Return(uint8(0), false)
	s.On("Uint16", mock.Anything).Return(uint16(0), false)
	s.On("Uint32", mock.Anything).Return(uint32(0), false)
	s.On("Uint64", mock.Anything).Return(uint64(0), false)
	s.On("Bool", mock.Anything).Return(false, false)
	s.On("Float32", mock.Anything).Return(float32(0), false)
	s.On("Float64", mock.Anything).Return(float64(0), false)
	s.On("Duration", mock.Anything).Return(time.Duration(0), false)
	s.On("Time", mock.Anything).Return(time.Time{}, false)
	return s
}

type MockSource struct {
	mock.Mock
}

func (s MockSource) Options() configify.Options {
	return s.Called().Get(0).(configify.Options)
}

func (s MockSource) String(key string) (string, bool) {
	args := s.Called(key)
	return args.Get(0).(string), args.Get(1).(bool)
}

func (s MockSource) StringSlice(key string) ([]string, bool) {
	args := s.Called(key)
	return args.Get(0).([]string), args.Get(1).(bool)
}

func (s MockSource) Int(key string) (int, bool) {
	args := s.Called(key)
	return args.Get(0).(int), args.Get(1).(bool)
}

func (s MockSource) Int8(key string) (int8, bool) {
	args := s.Called(key)
	return args.Get(0).(int8), args.Get(1).(bool)
}

func (s MockSource) Int16(key string) (int16, bool) {
	args := s.Called(key)
	return args.Get(0).(int16), args.Get(1).(bool)
}

func (s MockSource) Int32(key string) (int32, bool) {
	args := s.Called(key)
	return args.Get(0).(int32), args.Get(1).(bool)
}

func (s MockSource) Int64(key string) (int64, bool) {
	args := s.Called(key)
	return args.Get(0).(int64), args.Get(1).(bool)
}

func (s MockSource) Uint(key string) (uint, bool) {
	args := s.Called(key)
	return args.Get(0).(uint), args.Get(1).(bool)
}

func (s MockSource) Uint8(key string) (uint8, bool) {
	args := s.Called(key)
	return args.Get(0).(uint8), args.Get(1).(bool)
}

func (s MockSource) Uint16(key string) (uint16, bool) {
	args := s.Called(key)
	return args.Get(0).(uint16), args.Get(1).(bool)
}

func (s MockSource) Uint32(key string) (uint32, bool) {
	args := s.Called(key)
	return args.Get(0).(uint32), args.Get(1).(bool)
}

func (s MockSource) Uint64(key string) (uint64, bool) {
	args := s.Called(key)
	return args.Get(0).(uint64), args.Get(1).(bool)
}

func (s MockSource) Float32(key string) (float32, bool) {
	args := s.Called(key)
	return args.Get(0).(float32), args.Get(1).(bool)
}

func (s MockSource) Float64(key string) (float64, bool) {
	args := s.Called(key)
	return args.Get(0).(float64), args.Get(1).(bool)
}

func (s MockSource) Bool(key string) (bool, bool) {
	args := s.Called(key)
	return args.Get(0).(bool), args.Get(1).(bool)
}

func (s MockSource) Duration(key string) (time.Duration, bool) {
	args := s.Called(key)
	return args.Get(0).(time.Duration), args.Get(1).(bool)
}

func (s MockSource) Time(key string) (time.Time, bool) {
	args := s.Called(key)
	return args.Get(0).(time.Time), args.Get(1).(bool)
}
