package configify_test

import "github.com/stretchr/testify/mock"

func NewMockSource(setup func(*MockSource)) *MockSource {
	// We call setup() before using `mock.Anything` so that testify will try to apply those
	// matchers first on invocation. So the stuff you set up takes precedence over these defaults.
	s := new(MockSource)
	setup(s)
	s.On("GetString", mock.Anything).Return("")
	s.On("GetStringSlice", mock.Anything).Return([]string{})
	s.On("GetInt", mock.Anything).Return(0)
	s.On("GetUint", mock.Anything).Return(uint(0))
	return s
}

type MockSource struct {
	mock.Mock
}

func (s MockSource) GetString(key string) string {
	return s.Called(key).Get(0).(string)
}

func (s MockSource) GetStringSlice(key string) []string {
	return s.Called(key).Get(0).([]string)
}

func (s MockSource) GetInt(key string) int {
	return s.Called(key).Get(0).(int)
}

func (s MockSource) GetUint(key string) uint {
	return s.Called(key).Get(0).(uint)
}
