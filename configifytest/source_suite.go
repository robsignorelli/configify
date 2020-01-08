package configifytest

import (
	"time"

	"github.com/robsignorelli/configify"
	"github.com/stretchr/testify/suite"
)

// SourceSuite is a testify suite that adds some helpers for asserting whether or not the
// source you're testing has a specific value in it.
type SourceSuite struct {
	suite.Suite
	Source configify.Source
}

func (suite SourceSuite) checkOK(key string, expectedOK bool, ok bool) bool {
	if !expectedOK {
		return suite.False(ok, "Value for '%s' should not exist", key)
	}
	return suite.True(ok, "Value for '%s' was not found", key)
}

// ExpectString reads the config value on the suite's source and asserts that exists (or not)
// as you expect and that the value for the config is what you expect.
func (suite SourceSuite) ExpectString(key string, expected string, expectedOK bool) bool {
	output, ok := suite.Source.String(key)
	return suite.checkOK(key, expectedOK, ok) && suite.Equal(expected, output)
}

// ExpectStringSlice reads the config value on the suite's source and asserts that exists (or not)
// as you expect and that the value for the config is what you expect.
func (suite SourceSuite) ExpectStringSlice(key string, expected []string, expectedOK bool) bool {
	output, ok := suite.Source.StringSlice(key)
	return suite.checkOK(key, expectedOK, ok) && suite.ElementsMatch(expected, output)
}

// ExpectInt reads the config value on the suite's source and asserts that exists (or not)
// as you expect and that the value for the config is what you expect.
func (suite SourceSuite) ExpectInt(key string, expected int, expectedOK bool) bool {
	output, ok := suite.Source.Int(key)
	return suite.checkOK(key, expectedOK, ok) && suite.Equal(expected, output)
}

// ExpectInt8 reads the config value on the suite's source and asserts that exists (or not)
// as you expect and that the value for the config is what you expect.
func (suite SourceSuite) ExpectInt8(key string, expected int8, expectedOK bool) bool {
	output, ok := suite.Source.Int8(key)
	return suite.checkOK(key, expectedOK, ok) && suite.Equal(expected, output)
}

// ExpectInt16 reads the config value on the suite's source and asserts that exists (or not)
// as you expect and that the value for the config is what you expect.
func (suite SourceSuite) ExpectInt16(key string, expected int16, expectedOK bool) bool {
	output, ok := suite.Source.Int16(key)
	return suite.checkOK(key, expectedOK, ok) && suite.Equal(expected, output)
}

// ExpectInt32 reads the config value on the suite's source and asserts that exists (or not)
// as you expect and that the value for the config is what you expect.
func (suite SourceSuite) ExpectInt32(key string, expected int32, expectedOK bool) bool {
	output, ok := suite.Source.Int32(key)
	return suite.checkOK(key, expectedOK, ok) && suite.Equal(expected, output)
}

// ExpectInt64 reads the config value on the suite's source and asserts that exists (or not)
// as you expect and that the value for the config is what you expect.
func (suite SourceSuite) ExpectInt64(key string, expected int64, expectedOK bool) bool {
	output, ok := suite.Source.Int64(key)
	return suite.checkOK(key, expectedOK, ok) && suite.Equal(expected, output)
}

// ExpectUint reads the config value on the suite's source and asserts that exists (or not)
// as you expect and that the value for the config is what you expect.
func (suite SourceSuite) ExpectUint(key string, expected uint, expectedOK bool) bool {
	output, ok := suite.Source.Uint(key)
	return suite.checkOK(key, expectedOK, ok) && suite.Equal(expected, output)
}

// ExpectUint8 reads the config value on the suite's source and asserts that exists (or not)
// as you expect and that the value for the config is what you expect.
func (suite SourceSuite) ExpectUint8(key string, expected uint8, expectedOK bool) bool {
	output, ok := suite.Source.Uint8(key)
	return suite.checkOK(key, expectedOK, ok) && suite.Equal(expected, output)
}

// ExpectUint16 reads the config value on the suite's source and asserts that exists (or not)
// as you expect and that the value for the config is what you expect.
func (suite SourceSuite) ExpectUint16(key string, expected uint16, expectedOK bool) bool {
	output, ok := suite.Source.Uint16(key)
	return suite.checkOK(key, expectedOK, ok) && suite.Equal(expected, output)
}

// ExpectUint32 reads the config value on the suite's source and asserts that exists (or not)
// as you expect and that the value for the config is what you expect.
func (suite SourceSuite) ExpectUint32(key string, expected uint32, expectedOK bool) bool {
	output, ok := suite.Source.Uint32(key)
	return suite.checkOK(key, expectedOK, ok) && suite.Equal(expected, output)
}

// ExpectUint64 reads the config value on the suite's source and asserts that exists (or not)
// as you expect and that the value for the config is what you expect.
func (suite SourceSuite) ExpectUint64(key string, expected uint64, expectedOK bool) bool {
	output, ok := suite.Source.Uint64(key)
	return suite.checkOK(key, expectedOK, ok) && suite.Equal(expected, output)
}

// ExpectFloat32 reads the config value on the suite's source and asserts that exists (or not)
// as you expect and that the value for the config is what you expect.
func (suite SourceSuite) ExpectFloat32(key string, expected float32, expectedOK bool) bool {
	output, ok := suite.Source.Float32(key)
	return suite.checkOK(key, expectedOK, ok) && suite.Equal(expected, output)
}

// ExpectFloat64 reads the config value on the suite's source and asserts that exists (or not)
// as you expect and that the value for the config is what you expect.
func (suite SourceSuite) ExpectFloat64(key string, expected float64, expectedOK bool) bool {
	output, ok := suite.Source.Float64(key)
	return suite.checkOK(key, expectedOK, ok) && suite.Equal(expected, output)
}

// ExpectBool reads the config value on the suite's source and asserts that exists (or not)
// as you expect and that the value for the config is what you expect.
func (suite SourceSuite) ExpectBool(key string, expected bool, expectedOK bool) bool {
	output, ok := suite.Source.Bool(key)
	return suite.checkOK(key, expectedOK, ok) && suite.Equal(expected, output)
}

// ExpectDuration reads the config value on the suite's source and asserts that exists (or not)
// as you expect and that the value for the config is what you expect.
func (suite SourceSuite) ExpectDuration(key string, expected time.Duration, expectedOK bool) bool {
	output, ok := suite.Source.Duration(key)
	return suite.checkOK(key, expectedOK, ok) && suite.Equal(expected, output)
}

// ExpectTime reads the config value on the suite's source and asserts that exists (or not)
// as you expect and that the value for the config is what you expect.
func (suite SourceSuite) ExpectTime(key string, expected time.Time, expectedOK bool) bool {
	output, ok := suite.Source.Time(key)
	return suite.checkOK(key, expectedOK, ok) && suite.Equal(expected, output)
}
