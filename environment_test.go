package configify_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/robsignorelli/configify"
	"github.com/robsignorelli/configify/configifytest"
	"github.com/stretchr/testify/suite"
)

func TestEnvironmentSuite(t *testing.T) {
	suite.Run(t, new(EnvironmentSuite))
}

type EnvironmentSuite struct {
	configifytest.SourceSuite
}

func (suite *EnvironmentSuite) SetupSuite() {
	os.Clearenv()

	// All of the different data types we're likely to encounter/test
	suite.set("TEST_EMPTY", "")
	suite.set("TEST_STRING", "foo")
	suite.set("TEST_STRING_SPACE", "  foo bar ")
	suite.set("TEST_STRING_SLICE", "foo, bar, baz ,5 ")
	suite.set("TEST_INT", "5")
	suite.set("TEST_INT8", "8")
	suite.set("TEST_INT16", "16")
	suite.set("TEST_INT32", "32")
	suite.set("TEST_INT64", "64")
	suite.set("TEST_UINT", "90")
	suite.set("TEST_UINT8", "80")
	suite.set("TEST_UINT16", "160")
	suite.set("TEST_UINT32", "320")
	suite.set("TEST_UINT64", "640")
	suite.set("TEST_BOOL_TRUE", "true")
	suite.set("TEST_BOOL_TRUE_UPPER", "TRUE")
	suite.set("TEST_BOOL_FALSE", "false")
	suite.set("TEST_DURATION", "5m3s")
	suite.set("TEST_TIME_YYYYMMDD", "2019-12-25")
	suite.set("TEST_TIME_RFC3339", "2019-12-25T12:00:05.0Z")
	suite.set("TEST_LARGE_INT", "5,300,123")
	suite.set("TEST_NEGATIVE", "-3")
	suite.set("TEST_FLOAT", "5.430")
	suite.set("TEST_FLOAT32", "2.89")
	suite.set("TEST_LARGE_FLOAT", "5,300,123.430")
	suite.set("TEST_NEGATIVE_FLOAT", "-5.1")
	suite.set("TEST_JUST_FLOAT", ".1")

	// Not part of the "test namespace"
	suite.set("FOO_EMPTY", "")
	suite.set("FOO_STRING", "foo")
	suite.set("FOO_INT", "5")

	suite.Source, _ = configify.Environment(configify.Options{
		Namespace: configify.Namespace{Name: "TEST"},
	})
}

func (suite EnvironmentSuite) set(key string, value string) {
	_ = os.Setenv(key, value)
}

func (suite EnvironmentSuite) TestFactory() {
	_, err := configify.Environment(configify.Options{})
	suite.NoError(err)

	_, err = configify.Environment(configify.Options{Namespace: configify.Namespace{}})
	suite.NoError(err)

	_, err = configify.Environment(configify.Options{
		Namespace: configify.Namespace{Name: "", Delimiter: ""},
	})
	suite.NoError(err)

	_, err = configify.Environment(configify.Options{
		Namespace: configify.Namespace{Name: "FOO", Delimiter: "."},
	})
	suite.NoError(err)
}

func (suite EnvironmentSuite) TestOptions() {
	source, _ := configify.Environment(configify.Options{
		Namespace: configify.Namespace{Name: "FOO", Delimiter: "."},
	})
	suite.Equal("FOO", source.Options().Namespace.Name)
	suite.Equal(".", source.Options().Namespace.Delimiter)
}

func (suite EnvironmentSuite) TestString() {
	suite.ExpectString("NOT_FOUND", "", false)
	suite.ExpectString("EMPTY", "", true)
	suite.ExpectString("STRING", "foo", true)
	suite.ExpectString("STRING_SPACE", "foo bar", true)
	suite.ExpectString("STRING_SLICE", "foo, bar, baz ,5", true)
	suite.ExpectString("INT", "5", true)
	suite.ExpectString("NEGATIVE", "-3", true)
	suite.ExpectString("FLOAT", "5.430", true)
	suite.ExpectString("NEGATIVE_FLOAT", "-5.1", true)

	// Does not fetch values from other namespaces
	suite.ExpectString("FOO_EMPTY", "", false)
	suite.ExpectString("FOO_STRING", "", false)
	suite.ExpectString("FOO_INT", "", false)
}

func (suite EnvironmentSuite) TestStringSlice() {
	suite.ExpectStringSlice("NOT_FOUND", []string{}, false)
	suite.ExpectStringSlice("EMPTY", []string{}, true)
	suite.ExpectStringSlice("STRING", []string{"foo"}, true)
	suite.ExpectStringSlice("STRING_SPACE", []string{"foo bar"}, true)
	suite.ExpectStringSlice("STRING_SLICE", []string{"foo", "bar", "baz", "5"}, true)
	suite.ExpectStringSlice("INT", []string{"5"}, true)
	suite.ExpectStringSlice("LARGE_INT", []string{"5", "300", "123"}, true)
	suite.ExpectStringSlice("NEGATIVE", []string{"-3"}, true)
	suite.ExpectStringSlice("NEGATIVE_FLOAT", []string{"-5.1"}, true)
	suite.ExpectStringSlice("LARGE_FLOAT", []string{"5", "300", "123.430"}, true)

	// Does not fetch values from other namespaces
	suite.ExpectStringSlice("FOO_EMPTY", []string{}, false)
	suite.ExpectStringSlice("FOO_STRING", []string{}, false)
	suite.ExpectStringSlice("FOO_INT", []string{}, false)
}

func (suite EnvironmentSuite) TestInt() {
	// Only values that properly parse to integers are "ok"
	suite.ExpectInt("NOT_FOUND", 0, false)
	suite.ExpectInt("EMPTY", 0, false)
	suite.ExpectInt("STRING", 0, false)
	suite.ExpectInt("STRING_SPACE", 0, false)
	suite.ExpectInt("STRING_SLICE", 0, false)
	suite.ExpectInt("INT", 5, true)
	suite.ExpectInt("LARGE_INT", 5300123, true)
	suite.ExpectInt("NEGATIVE", -3, true)
	suite.ExpectInt("FLOAT", 5, true)
	suite.ExpectInt("NEGATIVE_FLOAT", -5, true)
	suite.ExpectInt("LARGE_FLOAT", 5300123, true)
	suite.ExpectInt("JUST_FLOAT", 0, false)

	// Does not fetch values from other namespaces
	suite.ExpectInt("FOO_EMPTY", 0, false)
	suite.ExpectInt("FOO_STRING", 0, false)
	suite.ExpectInt("FOO_INT", 0, false)
}

func (suite EnvironmentSuite) TestInt8() {
	// Only values that properly parse to integers are "ok"
	suite.ExpectInt8("NOT_FOUND", int8(0), false)
	suite.ExpectInt8("EMPTY", int8(0), false)
	suite.ExpectInt8("STRING", int8(0), false)
	suite.ExpectInt8("INT", int8(5), true)
	suite.ExpectInt8("INT8", int8(8), true)
	suite.ExpectInt8("INT16", int8(16), true)

	// Does not fetch values from other namespaces
	suite.ExpectInt("FOO_INT", 0, false)
}

func (suite EnvironmentSuite) TestInt16() {
	// Only values that properly parse to integers are "ok"
	suite.ExpectInt16("NOT_FOUND", int16(0), false)
	suite.ExpectInt16("EMPTY", int16(0), false)
	suite.ExpectInt16("STRING", int16(0), false)
	suite.ExpectInt16("INT", int16(5), true)
	suite.ExpectInt16("INT16", int16(16), true)
	suite.ExpectInt16("INT16", int16(16), true)

	// Does not fetch values from other namespaces
	suite.ExpectInt("FOO_INT", 0, false)
}

func (suite EnvironmentSuite) TestInt32() {
	// Only values that properly parse to integers are "ok"
	suite.ExpectInt32("NOT_FOUND", int32(0), false)
	suite.ExpectInt32("EMPTY", int32(0), false)
	suite.ExpectInt32("STRING", int32(0), false)
	suite.ExpectInt32("INT", int32(5), true)
	suite.ExpectInt32("INT32", int32(32), true)
	suite.ExpectInt32("INT16", int32(16), true)

	// Does not fetch values from other namespaces
	suite.ExpectInt("FOO_INT", 0, false)
}

func (suite EnvironmentSuite) TestInt64() {
	// Only values that properly parse to integers are "ok"
	suite.ExpectInt64("NOT_FOUND", int64(0), false)
	suite.ExpectInt64("EMPTY", int64(0), false)
	suite.ExpectInt64("STRING", int64(0), false)
	suite.ExpectInt64("INT", int64(5), true)
	suite.ExpectInt64("INT64", int64(64), true)
	suite.ExpectInt64("INT16", int64(16), true)

	// Does not fetch values from other namespaces
	suite.ExpectInt("FOO_INT", 0, false)
}

func (suite EnvironmentSuite) TestUint() {
	// Only values that properly parse to integers are "ok"
	suite.ExpectUint("NOT_FOUND", uint(0), false)
	suite.ExpectUint("EMPTY", uint(0), false)
	suite.ExpectUint("STRING", uint(0), false)
	suite.ExpectUint("STRING_SPACE", uint(0), false)
	suite.ExpectUint("STRING_SLICE", uint(0), false)
	suite.ExpectUint("INT", uint(5), true)
	suite.ExpectUint("LARGE_INT", uint(5300123), true)
	suite.ExpectUint("FLOAT", uint(5), true)
	suite.ExpectUint("LARGE_FLOAT", uint(5300123), true)
	suite.ExpectUint("JUST_FLOAT", uint(0), false)

	// Negatives resolve to zero, not the value w/o the minus sign.
	suite.ExpectUint("NEGATIVE", uint(0), false)
	suite.ExpectUint("NEGATIVE_FLOAT", uint(0), false)

	// Does not fetch values from other namespaces
	suite.ExpectUint("FOO_EMPTY", 0, false)
	suite.ExpectUint("FOO_STRING", 0, false)
	suite.ExpectUint("FOO_INT", 0, false)
}

func (suite EnvironmentSuite) TestUint8() {
	// Only values that properly parse to uintegers are "ok"
	suite.ExpectUint8("NOT_FOUND", uint8(0), false)
	suite.ExpectUint8("EMPTY", uint8(0), false)
	suite.ExpectUint8("STRING", uint8(0), false)
	suite.ExpectUint8("UINT", uint8(90), true)
	suite.ExpectUint8("UINT8", uint8(80), true)
	suite.ExpectUint8("UINT16", uint8(160), true)

	// Does not fetch values from other namespaces
	suite.ExpectUint("FOO_UINT", 0, false)
}

func (suite EnvironmentSuite) TestUint16() {
	// Only values that properly parse to uintegers are "ok"
	suite.ExpectUint16("NOT_FOUND", uint16(0), false)
	suite.ExpectUint16("EMPTY", uint16(0), false)
	suite.ExpectUint16("STRING", uint16(0), false)
	suite.ExpectUint16("UINT", uint16(90), true)
	suite.ExpectUint16("UINT16", uint16(160), true)
	suite.ExpectUint16("UINT16", uint16(160), true)

	// Does not fetch values from other namespaces
	suite.ExpectUint("FOO_UINT", 0, false)
}

func (suite EnvironmentSuite) TestUint32() {
	// Only values that properly parse to uintegers are "ok"
	suite.ExpectUint32("NOT_FOUND", uint32(0), false)
	suite.ExpectUint32("EMPTY", uint32(0), false)
	suite.ExpectUint32("STRING", uint32(0), false)
	suite.ExpectUint32("UINT", uint32(90), true)
	suite.ExpectUint32("UINT32", uint32(320), true)
	suite.ExpectUint32("UINT16", uint32(160), true)

	// Does not fetch values from other namespaces
	suite.ExpectUint("FOO_UINT", 0, false)
}

func (suite EnvironmentSuite) TestUint64() {
	// Only values that properly parse to uintegers are "ok"
	suite.ExpectUint64("NOT_FOUND", uint64(0), false)
	suite.ExpectUint64("EMPTY", uint64(0), false)
	suite.ExpectUint64("STRING", uint64(0), false)
	suite.ExpectUint64("UINT", uint64(90), true)
	suite.ExpectUint64("UINT64", uint64(640), true)
	suite.ExpectUint64("UINT16", uint64(160), true)

	// Does not fetch values from other namespaces
	suite.ExpectUint("FOO_UINT", 0, false)
}

func (suite EnvironmentSuite) TestFloat64() {
	suite.ExpectFloat64("NOT_FOUND", float64(0), false)
	suite.ExpectFloat64("FLOAT", 5.43, true)
	suite.ExpectFloat64("NEGATIVE_FLOAT", -5.1, true)
	suite.ExpectFloat64("JUST_FLOAT", 0.1, true)

	suite.ExpectFloat64("EMPTY", float64(0), false)
	suite.ExpectFloat64("STRING", float64(0), false)
	suite.ExpectFloat64("STRING_SLICE", float64(0), false)
	suite.ExpectFloat64("LARGE_INT", float64(0), false)
}

func (suite EnvironmentSuite) TestFloat32() {
	suite.ExpectFloat32("NOT_FOUND", float32(0), false)
	suite.ExpectFloat32("FLOAT", 5.43, true)
	suite.ExpectFloat32("NEGATIVE_FLOAT", -5.1, true)
	suite.ExpectFloat32("JUST_FLOAT", 0.1, true)

	suite.ExpectFloat32("EMPTY", float32(0), false)
	suite.ExpectFloat32("STRING", float32(0), false)
	suite.ExpectFloat32("STRING_SLICE", float32(0), false)
	suite.ExpectFloat32("LARGE_INT", float32(0), false)
}

func (suite EnvironmentSuite) TestBool() {
	suite.ExpectBool("NOT_FOUND", false, false)
	suite.ExpectBool("BOOL_FALSE", false, true)
	suite.ExpectBool("BOOL_TRUE", true, true)
	suite.ExpectBool("BOOL_TRUE_UPPER", true, true)

	suite.ExpectBool("EMPTY", false, false)
	suite.ExpectBool("STRING", false, false)
	suite.ExpectBool("STRING_SLICE", false, false)
	suite.ExpectBool("LARGE_INT", false, false)
}

func (suite EnvironmentSuite) TestDuration() {
	suite.ExpectDuration("NOT_FOUND", time.Duration(0), false)
	suite.ExpectDuration("DURATION", 5*time.Minute+3*time.Second, true)

	suite.ExpectDuration("EMPTY", time.Duration(0), false)
	suite.ExpectDuration("STRING", time.Duration(0), false)
	suite.ExpectDuration("STRING_SLICE", time.Duration(0), false)
	suite.ExpectDuration("LARGE_INT", time.Duration(0), false)
}

func (suite EnvironmentSuite) TestTime() {
	suite.ExpectTime("NOT_FOUND", time.Time{}, false)
	suite.ExpectTime("TIME_YYYYMMDD", time.Date(2019, 12, 25, 0, 0, 0, 0, time.UTC), true)
	suite.ExpectTime("TIME_RFC3339", time.Date(2019, 12, 25, 12, 0, 5, 0, time.UTC), true)

	suite.ExpectTime("EMPTY", time.Time{}, false)
	suite.ExpectTime("STRING", time.Time{}, false)
	suite.ExpectTime("STRING_SLICE", time.Time{}, false)
	suite.ExpectTime("LARGE_INT", time.Time{}, false)
}

func (suite EnvironmentSuite) TestDefaults() {
	def := configifytest.NewMockSource(func(s *configifytest.MockSource) {
		s.On("String", "STRING_MOCK").Return("asdf", true)
		s.On("StringSlice", "STRING_SLICE_MOCK").Return([]string{"a", "b"}, true)
		s.On("Int", "INT_MOCK").Return(8, true)
		s.On("Uint", "UINT_MOCK").Return(uint(9), true)
	})

	var err error
	suite.Source, err = configify.Environment(configify.Options{
		Namespace: configify.Namespace{Name: "TEST"},
		Defaults:  def,
	})
	suite.Require().NoError(err)

	// For each, make sure that (A) a valid env value resolves, (B) a value in the fixed fallback
	// resolves, and (C) a value that doesn't exist in either uses the hard-coded defaults.
	suite.ExpectString("STRING", "foo", true)
	suite.ExpectString("STRING_MOCK", "asdf", true)
	suite.ExpectString("STRING_XXX", "", false)

	suite.ExpectStringSlice("STRING_SLICE", []string{"foo", "bar", "baz", "5"}, true)
	suite.ExpectStringSlice("STRING_SLICE_MOCK", []string{"a", "b"}, true)
	suite.ExpectStringSlice("STRING_SLICE_XXX", []string{}, false)

	suite.ExpectInt("INT", 5, true)
	suite.ExpectInt("INT_MOCK", 8, true)
	suite.ExpectInt("INT_XXX", 0, false)

	suite.ExpectUint("UINT", uint(90), true)
	suite.ExpectUint("UINT_MOCK", uint(9), true)
	suite.ExpectUint("UINT_XXX", uint(0), false)
}

func ExampleEnvironment() {
	// Obviously, these would be normally applied by whatever you're using
	// for orchestration.
	_ = os.Setenv("HELLO_HOST", "localhost")
	_ = os.Setenv("HELLO_PORT", "8080")
	_ = os.Setenv("HELLO_TIMEOUT", "20s")

	// Use the namespace "HELLO" because it's the common prefix for all
	// of our environment variables. By default, we'll use "_" as the delimiter.
	config, err := configify.Environment(configify.Options{
		Namespace: configify.Namespace{Name: "HELLO"},
	})
	if err != nil {
		panic("Aww nuts...")
	}

	// Each value fetch gives you the parsed value as well as an 'ok'
	// as to whether the value actually existed in the source or not.
	host, ok := config.String("HOST")
	fmt.Printf("Host:    [%s] (%v)\n", host, ok)

	port, ok := config.Uint16("PORT")
	fmt.Printf("Port:    [%d] (%v)\n", port, ok)

	timeout, ok := config.Duration("TIMEOUT")
	fmt.Printf("Timeout: [%d] (%v)\n", timeout, ok)

	// The ok value is false for things not in your environment.
	foo, ok := config.String("FOO")
	fmt.Printf("Foo:     [%s] (%v)\n", foo, ok)

	// Output: Host:    [localhost] (true)
	// Port:    [8080] (true)
	// Timeout: [20000000000] (true)
	// Foo:     [] (false)
}

func ExampleEnvironmentDefaults() {
	// Obviously, these would be normally applied by whatever you're using
	// for orchestration.
	_ = os.Setenv("HELLO_HOST", "localhost")
	_ = os.Setenv("HELLO_PORT", "8080")
	_ = os.Setenv("HELLO_TIMEOUT", "20s")

	// Define default values to fall back to if the environment doesn't define them.
	defaults := configify.Map(configify.Values{
		"PORT": 9999,
		"FOO":  "foo value",
		"BAR":  "bar value",
	})

	// Use that fixed map source as the defaults for the environment source.
	config, err := configify.Environment(configify.Options{
		Namespace: configify.Namespace{Name: "HELLO"},
		Defaults:  defaults,
	})
	if err != nil {
		panic("Aww nuts...")
	}

	host, ok := config.String("HOST")
	fmt.Printf("Host:    [%s] (%v)\n", host, ok)

	// We'll use the environment value not the default (i.e. 8080, not 9999)
	port, ok := config.Uint16("PORT")
	fmt.Printf("Port:    [%d] (%v)\n", port, ok)

	timeout, ok := config.Duration("TIMEOUT")
	fmt.Printf("Timeout: [%d] (%v)\n", timeout, ok)

	// These two will use the fallback defaults. Note that 'ok' is true, not false!
	foo, ok := config.String("FOO")
	fmt.Printf("Foo:     [%s] (%v)\n", foo, ok)
	bar, ok := config.String("BAR")
	fmt.Printf("Bar:     [%s] (%v)\n", bar, ok)

	// But neither has "BAZ", so this will still be the natural default for a string
	baz, ok := config.String("BAZ")
	fmt.Printf("Baz:     [%s] (%v)\n", baz, ok)

	// Output: Host:    [localhost] (true)
	// Port:    [8080] (true)
	// Timeout: [20000000000] (true)
	// Foo:     [foo value] (true)
	// Bar:     [bar value] (true)
	// Baz:     [] (false)
}
