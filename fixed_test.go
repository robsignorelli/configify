package configify_test

import (
	"testing"

	"github.com/robsignorelli/configify"
	"github.com/stretchr/testify/suite"
)

func TestFixedSuite(t *testing.T) {
	suite.Run(t, new(FixedSuite))
}

type FixedSuite struct {
	suite.Suite
	source configify.Source
}

func (suite *FixedSuite) SetupSuite() {
	suite.source = configify.FixedSource("TEST", map[string]interface{}{
		"EMPTY":              "",
		"STRING":             "foo",
		"STRING_SPACE":       "  foo bar ",
		"STRING_SLICE":       []string{"foo", "bar", "baz", "5"},
		"STRING_SLICE_EMPTY": []string{},
		"STRING_SLICE_NIL":   nil,
		"INT":                5,
		"INT8":               int8(6),
		"INT16":              int16(7),
		"INT32":              int32(8),
		"INT64":              int64(9),
		"LARGE_INT":          5300123,
		"NEGATIVE":           -3,
		"UINT":               uint(5),
		"UINT8":              uint(6),
		"UINT16":             uint(7),
		"UINT32":             uint(8),
		"UINT64":             uint(9),
		"FLOAT":              5.430,
		"LARGE_FLOAT":        5300123.430,
		"NEGATIVE_FLOAT":     -5.1,
	})
}

func (suite FixedSuite) TestGetString() {
	get := func(key string) string {
		return suite.source.GetString(key)
	}
	suite.Equal("", get("NOT_FOUND"))
	suite.Equal("", get("EMPTY"))
	suite.Equal("foo", get("STRING"))
	suite.Equal("foo bar", get("STRING_SPACE"))

	// Only values strongly typed as strings will resolve properly.
	suite.Equal("", get("STRING_SLICE"))
	suite.Equal("", get("INT"))
	suite.Equal("", get("UINT"))
	suite.Equal("", get("NEGATIVE"))
	suite.Equal("", get("FLOAT"))
	suite.Equal("", get("NEGATIVE_FLOAT"))
}

func (suite FixedSuite) TestGetStringSlice() {
	get := func(key string) []string {
		return suite.source.GetStringSlice(key)
	}

	suite.ElementsMatch([]string{"foo", "bar", "baz", "5"}, get("STRING_SLICE"))
	suite.ElementsMatch([]string{}, get("STRING_SLICE_EMPTY"))
	suite.ElementsMatch([]string{}, get("STRING_SLICE_NIL"))

	// Only values strongly typed as []string will resolve properly.
	suite.Len(get("NOT_FOUND"), 0)
	suite.Len(get("EMPTY"), 0)
	suite.Len(get("STRING"), 0)
	suite.Len(get("STRING_SPACE"), 0)
	suite.Len(get("INT"), 0)
	suite.Len(get("UINT"), 0)
	suite.Len(get("FLOAT"), 0)
}

func (suite FixedSuite) TestGetInt() {
	get := func(key string) int {
		return suite.source.GetInt(key)
	}

	suite.Equal(5, get("INT"))
	suite.Equal(5300123, get("LARGE_INT"))
	suite.Equal(-3, get("NEGATIVE"))

	// Only values strongly typed as strings will resolve properly.
	suite.Equal(0, get("NOT_FOUND"))
	suite.Equal(0, get("EMPTY"))
	suite.Equal(0, get("STRING"))
	suite.Equal(0, get("STRING_SPACE"))
	suite.Equal(0, get("STRING_SLICE"))
	suite.Equal(0, get("UINT"))
	suite.Equal(0, get("FLOAT"))
	suite.Equal(0, get("NEGATIVE_FLOAT"))
}

func (suite FixedSuite) TestGetUint() {
	get := func(key string) uint {
		return suite.source.GetUint(key)
	}

	suite.Equal(uint(5), get("UINT"))

	// Only values strongly typed as strings will resolve properly.
	suite.Equal(uint(0), get("NOT_FOUND"))
	suite.Equal(uint(0), get("EMPTY"))
	suite.Equal(uint(0), get("STRING"))
	suite.Equal(uint(0), get("STRING_SPACE"))
	suite.Equal(uint(0), get("STRING_SLICE"))
	suite.Equal(uint(0), get("FLOAT"))
	suite.Equal(uint(0), get("NEGATIVE_FLOAT"))
	suite.Equal(uint(0), get("INT"))
	suite.Equal(uint(0), get("LARGE_INT"))
	suite.Equal(uint(0), get("NEGATIVE"))
}
