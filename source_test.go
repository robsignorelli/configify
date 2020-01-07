package configify_test

import (
	"time"

	"github.com/robsignorelli/configify"
	"github.com/stretchr/testify/suite"
)

type SourceSuite struct {
	suite.Suite
	source configify.Source
}

func (suite SourceSuite) TestNamespace_Qualify() {
	options := configify.Options{}
	suite.Equal("BAR", options.Namespace.Qualify("BAR"))

	options = configify.Options{}
	configify.Namespace("FOO")(&options)
	suite.Equal("FOO_BAR", options.Namespace.Qualify("BAR"))

	options = configify.Options{}
	configify.Namespace("FOO")(&options)
	configify.NamespaceDelim(".")(&options)
	suite.Equal("FOO.BAR", options.Namespace.Qualify("BAR"))

	options = configify.Options{}
	configify.Namespace("FOO")(&options)
	configify.NamespaceDelim("  .  ")(&options)
	suite.Equal("FOO.BAR", options.Namespace.Qualify("BAR"))
}

func (suite SourceSuite) checkOK(key string, expectedOK bool, ok bool) bool {
	if !expectedOK {
		return suite.False(ok, "Value for '%s' should not exist", key)
	}
	return suite.True(ok, "Value for '%s' was not found", key)
}

func (suite SourceSuite) ExpectString(key string, expected string, expectedOK bool) bool {
	output, ok := suite.source.String(key)
	return suite.checkOK(key, expectedOK, ok) && suite.Equal(expected, output)
}

func (suite SourceSuite) ExpectStringSlice(key string, expected []string, expectedOK bool) bool {
	output, ok := suite.source.StringSlice(key)
	return suite.checkOK(key, expectedOK, ok) && suite.ElementsMatch(expected, output)
}

func (suite SourceSuite) ExpectInt(key string, expected int, expectedOK bool) bool {
	output, ok := suite.source.Int(key)
	return suite.checkOK(key, expectedOK, ok) && suite.Equal(expected, output)
}

func (suite SourceSuite) ExpectInt8(key string, expected int8, expectedOK bool) bool {
	output, ok := suite.source.Int8(key)
	return suite.checkOK(key, expectedOK, ok) && suite.Equal(expected, output)
}

func (suite SourceSuite) ExpectInt16(key string, expected int16, expectedOK bool) bool {
	output, ok := suite.source.Int16(key)
	return suite.checkOK(key, expectedOK, ok) && suite.Equal(expected, output)
}

func (suite SourceSuite) ExpectInt32(key string, expected int32, expectedOK bool) bool {
	output, ok := suite.source.Int32(key)
	return suite.checkOK(key, expectedOK, ok) && suite.Equal(expected, output)
}

func (suite SourceSuite) ExpectInt64(key string, expected int64, expectedOK bool) bool {
	output, ok := suite.source.Int64(key)
	return suite.checkOK(key, expectedOK, ok) && suite.Equal(expected, output)
}

func (suite SourceSuite) ExpectUint(key string, expected uint, expectedOK bool) bool {
	output, ok := suite.source.Uint(key)
	return suite.checkOK(key, expectedOK, ok) && suite.Equal(expected, output)
}

func (suite SourceSuite) ExpectUint8(key string, expected uint8, expectedOK bool) bool {
	output, ok := suite.source.Uint8(key)
	return suite.checkOK(key, expectedOK, ok) && suite.Equal(expected, output)
}

func (suite SourceSuite) ExpectUint16(key string, expected uint16, expectedOK bool) bool {
	output, ok := suite.source.Uint16(key)
	return suite.checkOK(key, expectedOK, ok) && suite.Equal(expected, output)
}

func (suite SourceSuite) ExpectUint32(key string, expected uint32, expectedOK bool) bool {
	output, ok := suite.source.Uint32(key)
	return suite.checkOK(key, expectedOK, ok) && suite.Equal(expected, output)
}

func (suite SourceSuite) ExpectUint64(key string, expected uint64, expectedOK bool) bool {
	output, ok := suite.source.Uint64(key)
	return suite.checkOK(key, expectedOK, ok) && suite.Equal(expected, output)
}

func (suite SourceSuite) ExpectFloat32(key string, expected float32, expectedOK bool) bool {
	output, ok := suite.source.Float32(key)
	return suite.checkOK(key, expectedOK, ok) && suite.Equal(expected, output)
}

func (suite SourceSuite) ExpectFloat64(key string, expected float64, expectedOK bool) bool {
	output, ok := suite.source.Float64(key)
	return suite.checkOK(key, expectedOK, ok) && suite.Equal(expected, output)
}

func (suite SourceSuite) ExpectBool(key string, expected bool, expectedOK bool) bool {
	output, ok := suite.source.Bool(key)
	return suite.checkOK(key, expectedOK, ok) && suite.Equal(expected, output)
}

func (suite SourceSuite) ExpectDuration(key string, expected time.Duration, expectedOK bool) bool {
	output, ok := suite.source.Duration(key)
	return suite.checkOK(key, expectedOK, ok) && suite.Equal(expected, output)
}

func (suite SourceSuite) ExpectTime(key string, expected time.Time, expectedOK bool) bool {
	output, ok := suite.source.Time(key)
	return suite.checkOK(key, expectedOK, ok) && suite.Equal(expected, output)
}
