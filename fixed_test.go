package configify_test

import (
	"testing"
	"time"

	"github.com/robsignorelli/configify"
	"github.com/stretchr/testify/suite"
)

func TestFixedSuite(t *testing.T) {
	suite.Run(t, new(FixedSuite))
}

type FixedSuite struct {
	SourceSuite
}

func (suite *FixedSuite) SetupSuite() {
	suite.source = configify.Fixed(configify.Values{
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
		"BOOL_TRUE":          true,
		"BOOL_FALSE":         false,
		"LARGE_FLOAT":        5300123.430,
		"NEGATIVE_FLOAT":     -5.1,
		"DURATION":           5 * time.Minute,
		"TIME":               time.Date(2019, time.December, 25, 8, 33, 40, 0, time.UTC),
	})
}

func (suite FixedSuite) TestString() {
	suite.ExpectString("NOT_FOUND", "", false)
	suite.ExpectString("EMPTY", "", true)
	suite.ExpectString("STRING", "foo", true)
	suite.ExpectString("STRING_SPACE", "foo bar", true)

	// Only values strongly typed as strings will resolve properly.
	suite.ExpectString("STRING_SLICE", "", false)
	suite.ExpectString("INT", "", false)
	suite.ExpectString("FLOAT", "", false)
}

func (suite FixedSuite) TestStringSlice() {
	suite.ExpectStringSlice("NOT_FOUND", []string{}, false)
	suite.ExpectStringSlice("STRING_SLICE", []string{"foo", "bar", "baz", "5"}, true)
	suite.ExpectStringSlice("STRING_SLICE_EMPTY", []string{}, true)

	// We can't distinguish between nil you explicitly put in and nil that never existed.
	suite.ExpectStringSlice("STRING_SLICE_NIL", []string{}, false)

	// Only values strongly typed as []string will resolve properly.
	suite.ExpectStringSlice("EMPTY", []string{}, false)
	suite.ExpectStringSlice("STRING", []string{}, false)
	suite.ExpectStringSlice("INT", []string{}, false)
	suite.ExpectStringSlice("FLOAT", []string{}, false)
}

func (suite FixedSuite) TestInt() {
	suite.ExpectInt("NOT_FOUND", 0, false)
	suite.ExpectInt("INT", 5, true)
	suite.ExpectInt("LARGE_INT", 5300123, true)
	suite.ExpectInt("NEGATIVE", -3, true)

	// Only values strongly typed as strings will resolve properly.
	suite.ExpectInt("EMPTY", 0, false)
	suite.ExpectInt("STRING", 0, false)
	suite.ExpectInt("STRING_SLICE", 0, false)
	suite.ExpectInt("UINT", 0, false)
}

func (suite FixedSuite) TestUint() {
	suite.ExpectUint("NOT_FOUND", uint(0), false)
	suite.ExpectUint("UINT", uint(5), true)

	// Only values strongly typed as strings will resolve properly.
	suite.ExpectUint("EMPTY", uint(0), false)
	suite.ExpectUint("STRING", uint(0), false)
	suite.ExpectUint("STRING_SLICE", uint(0), false)
	suite.ExpectUint("INT", uint(0), false)
	suite.ExpectUint("LARGE_INT", uint(0), false)
	suite.ExpectUint("NEGATIVE", uint(0), false)
	suite.ExpectUint("FLOAT", uint(0), false)
}

func (suite FixedSuite) TestBool() {
	suite.ExpectBool("NOT_FOUND", false, false)
	suite.ExpectBool("BOOL_TRUE", true, true)
	suite.ExpectBool("BOOL_FALSE", false, true)

	// Only values strongly typed as strings will resolve properly.
	suite.ExpectBool("EMPTY", false, false)
	suite.ExpectBool("STRING", false, false)
	suite.ExpectBool("STRING_SLICE", false, false)
	suite.ExpectBool("UINT", false, false)
}

func (suite FixedSuite) TestFloat() {
	suite.ExpectFloat("NOT_FOUND", float64(0), false)
	suite.ExpectFloat("FLOAT", 5.430, true)

	// Only values strongly typed as strings will resolve properly.
	suite.ExpectFloat("EMPTY", float64(0), false)
	suite.ExpectFloat("STRING", float64(0), false)
	suite.ExpectFloat("STRING_SLICE", float64(0), false)
	suite.ExpectFloat("UINT", float64(0), false)
}

func (suite FixedSuite) TestDuration() {
	suite.ExpectDuration("NOT_FOUND", time.Duration(0), false)
	suite.ExpectDuration("DURATION", 5*time.Minute, true)

	// Only values strongly typed as strings will resolve properly.
	suite.ExpectDuration("EMPTY", time.Duration(0), false)
	suite.ExpectDuration("STRING", time.Duration(0), false)
	suite.ExpectDuration("STRING_SLICE", time.Duration(0), false)
	suite.ExpectDuration("INT", time.Duration(0), false)
	suite.ExpectDuration("LARGE_INT", time.Duration(0), false)
	suite.ExpectDuration("NEGATIVE", time.Duration(0), false)
	suite.ExpectDuration("FLOAT", time.Duration(0), false)
}

func (suite FixedSuite) TestTime() {
	suite.ExpectTime("NOT_FOUND", time.Time{}, false)
	suite.ExpectTime("TIME", time.Date(2019, time.December, 25, 8, 33, 40, 0, time.UTC), true)

	// Only values strongly typed as strings will resolve properly.
	suite.ExpectTime("EMPTY", time.Time{}, false)
	suite.ExpectTime("STRING", time.Time{}, false)
	suite.ExpectTime("STRING_SLICE", time.Time{}, false)
	suite.ExpectTime("INT", time.Time{}, false)
	suite.ExpectTime("LARGE_INT", time.Time{}, false)
	suite.ExpectTime("NEGATIVE", time.Time{}, false)
	suite.ExpectTime("FLOAT", time.Time{}, false)
	suite.ExpectTime("DURATION", time.Time{}, false)
}
