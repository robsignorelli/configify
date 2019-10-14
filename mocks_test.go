package configify_test

import (
	"time"

	"github.com/robsignorelli/configify"
	"github.com/stretchr/testify/mock"
)

func NewMockSource(setup func(*MockSource)) configify.Source {
	// We call setup() before using `mock.Anything` so that testify will try to apply those
	// matchers first on invocation. So the stuff you set up takes precedence over these defaults.
	s := new(MockSource)
	setup(s)
	s.On("Options", mock.Anything).Return(configify.Options{})
	s.On("String", mock.Anything).Return("", false)
	s.On("StringSlice", mock.Anything).Return([]string{}, false)
	s.On("Int", mock.Anything).Return(0, false)
	s.On("Uint", mock.Anything).Return(uint(0), false)
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

func (s MockSource) Uint(key string) (uint, bool) {
	args := s.Called(key)
	return args.Get(0).(uint), args.Get(1).(bool)
}

func (s MockSource) Duration(key string) (time.Duration, bool) {
	args := s.Called(key)
	return args.Get(0).(time.Duration), args.Get(1).(bool)
}

func (s MockSource) Time(key string) (time.Time, bool) {
	args := s.Called(key)
	return args.Get(0).(time.Time), args.Get(1).(bool)
}
