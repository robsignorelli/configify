package configify_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/robsignorelli/configify"
	"github.com/robsignorelli/configify/configifytest"
	"github.com/stretchr/testify/suite"
)

func TestMapSuite(t *testing.T) {
	suite.Run(t, new(MapSuite))
}

type MapSuite struct {
	configifytest.SourceSuite
}

func (suite *MapSuite) SetupSuite() {
	suite.Source = configify.Map(configify.Values{
		"EMPTY":              "",
		"STRING":             "foo",
		"STRING_SPACE":       "  foo bar ",
		"STRING_SLICE":       []string{"foo", "bar", "baz", "5"},
		"STRING_SLICE_EMPTY": []string{},
		"STRING_SLICE_NIL":   nil,
		"INT":                5,
		"INT8":               int8(8),
		"INT16":              int16(16),
		"INT32":              int32(32),
		"INT64":              int64(64),
		"LARGE_INT":          5300123,
		"NEGATIVE":           -3,
		"UINT":               uint(5),
		"UINT8":              uint8(80),
		"UINT16":             uint16(160),
		"UINT32":             uint32(320),
		"UINT64":             uint64(640),
		"FLOAT32":            float32(2.89),
		"FLOAT64":            5.430,
		"BOOL_TRUE":          true,
		"BOOL_FALSE":         false,
		"LARGE_FLOAT":        5300123.430,
		"NEGATIVE_FLOAT":     -5.1,
		"DURATION":           5 * time.Minute,
		"TIME":               time.Date(2019, time.December, 25, 8, 33, 40, 0, time.UTC),
	})
}

func (suite MapSuite) TestString() {
	suite.ExpectString("NOT_FOUND", "", false)
	suite.ExpectString("EMPTY", "", true)
	suite.ExpectString("STRING", "foo", true)
	suite.ExpectString("STRING_SPACE", "foo bar", true)

	// Only values strongly typed as strings will resolve properly.
	suite.ExpectString("STRING_SLICE", "", false)
	suite.ExpectString("INT", "", false)
	suite.ExpectString("FLOAT64", "", false)
}

func (suite MapSuite) TestStringSlice() {
	suite.ExpectStringSlice("NOT_FOUND", []string{}, false)
	suite.ExpectStringSlice("STRING_SLICE", []string{"foo", "bar", "baz", "5"}, true)
	suite.ExpectStringSlice("STRING_SLICE_EMPTY", []string{}, true)

	// We can't distinguish between nil you explicitly put in and nil that never existed.
	suite.ExpectStringSlice("STRING_SLICE_NIL", []string{}, false)

	// Only values strongly typed as []string will resolve properly.
	suite.ExpectStringSlice("EMPTY", []string{}, false)
	suite.ExpectStringSlice("STRING", []string{}, false)
	suite.ExpectStringSlice("INT", []string{}, false)
	suite.ExpectStringSlice("FLOAT64", []string{}, false)
}

func (suite MapSuite) TestInt() {
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

func (suite MapSuite) TestInt8() {
	suite.ExpectInt8("NOT_FOUND", int8(0), false)
	suite.ExpectInt8("INT8", int8(8), true)

	// Only values strongly typed as strings will resolve properly.
	suite.ExpectInt8("EMPTY", int8(0), false)
	suite.ExpectInt8("STRING", int8(0), false)
	suite.ExpectInt8("STRING_SLICE", int8(0), false)
	suite.ExpectInt8("INT", int8(0), false)
	suite.ExpectInt8("LARGE_INT", int8(0), false)
	suite.ExpectInt8("NEGATIVE", int8(0), false)
	suite.ExpectInt8("FLOAT64", int8(0), false)
}

func (suite MapSuite) TestInt16() {
	suite.ExpectInt16("NOT_FOUND", int16(0), false)
	suite.ExpectInt16("INT16", int16(16), true)

	// Only values strongly typed as strings will resolve properly.
	suite.ExpectInt16("EMPTY", int16(0), false)
	suite.ExpectInt16("STRING", int16(0), false)
	suite.ExpectInt16("STRING_SLICE", int16(0), false)
	suite.ExpectInt16("INT", int16(0), false)
	suite.ExpectInt16("LARGE_INT", int16(0), false)
	suite.ExpectInt16("NEGATIVE", int16(0), false)
	suite.ExpectInt16("FLOAT64", int16(0), false)
}

func (suite MapSuite) TestInt32() {
	suite.ExpectInt32("NOT_FOUND", int32(0), false)
	suite.ExpectInt32("INT32", int32(32), true)

	// Only values strongly typed as strings will resolve properly.
	suite.ExpectInt32("EMPTY", int32(0), false)
	suite.ExpectInt32("STRING", int32(0), false)
	suite.ExpectInt32("STRING_SLICE", int32(0), false)
	suite.ExpectInt32("INT", int32(0), false)
	suite.ExpectInt32("LARGE_INT", int32(0), false)
	suite.ExpectInt32("NEGATIVE", int32(0), false)
	suite.ExpectInt32("FLOAT64", int32(0), false)
}

func (suite MapSuite) TestInt64() {
	suite.ExpectInt64("NOT_FOUND", int64(0), false)
	suite.ExpectInt64("INT64", int64(64), true)

	// Only values strongly typed as strings will resolve properly.
	suite.ExpectInt64("EMPTY", int64(0), false)
	suite.ExpectInt64("STRING", int64(0), false)
	suite.ExpectInt64("STRING_SLICE", int64(0), false)
	suite.ExpectInt64("INT", int64(0), false)
	suite.ExpectInt64("LARGE_INT", int64(0), false)
	suite.ExpectInt64("NEGATIVE", int64(0), false)
	suite.ExpectInt64("FLOAT64", int64(0), false)
}

func (suite MapSuite) TestUint() {
	suite.ExpectUint("NOT_FOUND", uint(0), false)
	suite.ExpectUint("UINT", uint(5), true)

	// Only values strongly typed as strings will resolve properly.
	suite.ExpectUint("EMPTY", uint(0), false)
	suite.ExpectUint("STRING", uint(0), false)
	suite.ExpectUint("STRING_SLICE", uint(0), false)
	suite.ExpectUint("INT", uint(0), false)
	suite.ExpectUint("LARGE_INT", uint(0), false)
	suite.ExpectUint("NEGATIVE", uint(0), false)
	suite.ExpectUint("FLOAT64", uint(0), false)
}

func (suite MapSuite) TestUint8() {
	suite.ExpectUint8("NOT_FOUND", uint8(0), false)
	suite.ExpectUint8("UINT8", uint8(80), true)

	// Only values strongly typed as strings will resolve properly.
	suite.ExpectUint8("EMPTY", uint8(0), false)
	suite.ExpectUint8("STRING", uint8(0), false)
	suite.ExpectUint8("STRING_SLICE", uint8(0), false)
	suite.ExpectUint8("INT", uint8(0), false)
	suite.ExpectUint8("LARGE_INT", uint8(0), false)
	suite.ExpectUint8("NEGATIVE", uint8(0), false)
	suite.ExpectUint8("FLOAT64", uint8(0), false)
}

func (suite MapSuite) TestUint16() {
	suite.ExpectUint16("NOT_FOUND", uint16(0), false)
	suite.ExpectUint16("UINT16", uint16(160), true)

	// Only values strongly typed as strings will resolve properly.
	suite.ExpectUint16("EMPTY", uint16(0), false)
	suite.ExpectUint16("STRING", uint16(0), false)
	suite.ExpectUint16("STRING_SLICE", uint16(0), false)
	suite.ExpectUint16("INT", uint16(0), false)
	suite.ExpectUint16("LARGE_INT", uint16(0), false)
	suite.ExpectUint16("NEGATIVE", uint16(0), false)
	suite.ExpectUint16("FLOAT64", uint16(0), false)
}

func (suite MapSuite) TestUint32() {
	suite.ExpectUint32("NOT_FOUND", uint32(0), false)
	suite.ExpectUint32("UINT32", uint32(320), true)

	// Only values strongly typed as strings will resolve properly.
	suite.ExpectUint32("EMPTY", uint32(0), false)
	suite.ExpectUint32("STRING", uint32(0), false)
	suite.ExpectUint32("STRING_SLICE", uint32(0), false)
	suite.ExpectUint32("INT", uint32(0), false)
	suite.ExpectUint32("LARGE_INT", uint32(0), false)
	suite.ExpectUint32("NEGATIVE", uint32(0), false)
	suite.ExpectUint32("FLOAT64", uint32(0), false)
}

func (suite MapSuite) TestUint64() {
	suite.ExpectUint64("NOT_FOUND", uint64(0), false)
	suite.ExpectUint64("UINT64", uint64(640), true)

	// Only values strongly typed as strings will resolve properly.
	suite.ExpectUint64("EMPTY", uint64(0), false)
	suite.ExpectUint64("STRING", uint64(0), false)
	suite.ExpectUint64("STRING_SLICE", uint64(0), false)
	suite.ExpectUint64("INT", uint64(0), false)
	suite.ExpectUint64("LARGE_INT", uint64(0), false)
	suite.ExpectUint64("NEGATIVE", uint64(0), false)
	suite.ExpectUint64("FLOAT64", uint64(0), false)
}

func (suite MapSuite) TestBool() {
	suite.ExpectBool("NOT_FOUND", false, false)
	suite.ExpectBool("BOOL_TRUE", true, true)
	suite.ExpectBool("BOOL_FALSE", false, true)

	// Only values strongly typed as strings will resolve properly.
	suite.ExpectBool("EMPTY", false, false)
	suite.ExpectBool("STRING", false, false)
	suite.ExpectBool("STRING_SLICE", false, false)
	suite.ExpectBool("UINT8", false, false)
}

func (suite MapSuite) TestFloat32() {
	suite.ExpectFloat32("NOT_FOUND", float32(0), false)
	suite.ExpectFloat32("FLOAT32", float32(2.89), true)

	// Only values strongly typed as strings will resolve properly.
	suite.ExpectFloat32("EMPTY", float32(0), false)
	suite.ExpectFloat32("STRING", float32(0), false)
	suite.ExpectFloat32("STRING_SLICE", float32(0), false)
	suite.ExpectFloat32("UINT8", float32(0), false)
	suite.ExpectFloat32("FLOAT64", float32(0), false)
}

func (suite MapSuite) TestFloat64() {
	suite.ExpectFloat64("NOT_FOUND", float64(0), false)
	suite.ExpectFloat64("FLOAT64", 5.430, true)

	// Only values strongly typed as strings will resolve properly.
	suite.ExpectFloat64("EMPTY", float64(0), false)
	suite.ExpectFloat64("STRING", float64(0), false)
	suite.ExpectFloat64("STRING_SLICE", float64(0), false)
	suite.ExpectFloat64("UINT8", float64(0), false)
}

func (suite MapSuite) TestDuration() {
	suite.ExpectDuration("NOT_FOUND", time.Duration(0), false)
	suite.ExpectDuration("DURATION", 5*time.Minute, true)

	// Only values strongly typed as strings will resolve properly.
	suite.ExpectDuration("EMPTY", time.Duration(0), false)
	suite.ExpectDuration("STRING", time.Duration(0), false)
	suite.ExpectDuration("STRING_SLICE", time.Duration(0), false)
	suite.ExpectDuration("INT", time.Duration(0), false)
	suite.ExpectDuration("LARGE_INT", time.Duration(0), false)
	suite.ExpectDuration("NEGATIVE", time.Duration(0), false)
	suite.ExpectDuration("FLOAT64", time.Duration(0), false)
}

func (suite MapSuite) TestTime() {
	suite.ExpectTime("NOT_FOUND", time.Time{}, false)
	suite.ExpectTime("TIME", time.Date(2019, time.December, 25, 8, 33, 40, 0, time.UTC), true)

	// Only values strongly typed as strings will resolve properly.
	suite.ExpectTime("EMPTY", time.Time{}, false)
	suite.ExpectTime("STRING", time.Time{}, false)
	suite.ExpectTime("STRING_SLICE", time.Time{}, false)
	suite.ExpectTime("INT", time.Time{}, false)
	suite.ExpectTime("LARGE_INT", time.Time{}, false)
	suite.ExpectTime("NEGATIVE", time.Time{}, false)
	suite.ExpectTime("FLOAT64", time.Time{}, false)
	suite.ExpectTime("DURATION", time.Time{}, false)
}

func ExampleMap() {
	// Since you have full control over the values, be sure to strongly-type them
	// to the types you expect to get out. For instance if you plan to grab the
	// key "FOO" as a uint8, then put it in the values map as a uint8, not a plain
	// old int.
	config := configify.Map(configify.Values{
		"HOST":    "localhost",
		"PORT":    uint16(1234),
		"TIMEOUT": 20 * time.Second,
		"THINGS":  []string{"foo", "bar", "baz"},
	})

	host, ok := config.String("HOST")
	fmt.Printf("Host:    [%s] (%v)\n", host, ok)

	port, ok := config.Uint16("PORT")
	fmt.Printf("Port:    [%d] (%v)\n", port, ok)

	timeout, ok := config.Duration("TIMEOUT")
	fmt.Printf("Timeout: [%d] (%v)\n", timeout, ok)

	things, ok := config.StringSlice("THINGS")
	fmt.Printf("Things:  [%v] (%v)\n", things, ok)

	// The ok value is false for things not in your value map.
	foo, ok := config.String("FOO")
	fmt.Printf("Foo:     [%s] (%v)\n", foo, ok)

	// Output: Host:    [localhost] (true)
	// Port:    [1234] (true)
	// Timeout: [20000000000] (true)
	// Things:  [[foo bar baz]] (true)
	// Foo:     [] (false)
}
