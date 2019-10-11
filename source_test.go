package configify_test

import (
	"github.com/robsignorelli/configify"
	"github.com/stretchr/testify/suite"
)

type SourceSuite struct {
	suite.Suite
	source configify.Source
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

func (suite SourceSuite) ExpectUint(key string, expected uint, expectedOK bool) bool {
	output, ok := suite.source.Uint(key)
	return suite.checkOK(key, expectedOK, ok) && suite.Equal(expected, output)
}
