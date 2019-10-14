package configify_test

import (
	"os"
	"testing"
	"time"

	"github.com/robsignorelli/configify"
	"github.com/stretchr/testify/suite"
)

func TestEnvironmentSuite(t *testing.T) {
	suite.Run(t, new(EnvironmentSuite))
}

type EnvironmentSuite struct {
	SourceSuite
}

func (suite *EnvironmentSuite) SetupSuite() {
	os.Clearenv()

	// All of the different data types we're likely to encounter/test
	suite.set("TEST_EMPTY", "")
	suite.set("TEST_STRING", "foo")
	suite.set("TEST_STRING_SPACE", "  foo bar ")
	suite.set("TEST_STRING_SLICE", "foo, bar, baz ,5 ")
	suite.set("TEST_INT", "5")
	suite.set("TEST_UINT", "90")
	suite.set("TEST_DURATION", "5m3s")
	suite.set("TEST_TIME_YYYYMMDD", "2019-12-25")
	suite.set("TEST_TIME_RFC3339", "2019-12-25T12:00:05.0Z")
	suite.set("TEST_LARGE_INT", "5,300,123")
	suite.set("TEST_NEGATIVE", "-3")
	suite.set("TEST_FLOAT", "5.430")
	suite.set("TEST_LARGE_FLOAT", "5,300,123.430")
	suite.set("TEST_NEGATIVE_FLOAT", "-5.1")
	suite.set("TEST_JUST_FLOAT", ".1")

	// Not part of the "test namespace"
	suite.set("FOO_EMPTY", "")
	suite.set("FOO_STRING", "foo")
	suite.set("FOO_INT", "5")

	suite.source, _ = configify.Environment(configify.Options{
		Namespace: "TEST",
	})
}

func (suite EnvironmentSuite) set(key string, value string) {
	_ = os.Setenv(key, value)
}

func (suite EnvironmentSuite) TestFactory() {
	_, err := configify.Environment(configify.Options{Namespace: ""})
	suite.NoError(err)

	_, err = configify.Environment(configify.Options{Namespace: "FOO"})
	suite.NoError(err)
}

func (suite EnvironmentSuite) TestOptions() {
	source, _ := configify.Environment(configify.Options{
		Namespace:      "FOO",
		NamespaceDelim: ".",
	})
	suite.Equal("FOO", source.Options().Namespace)
	suite.Equal(".", source.Options().NamespaceDelim)
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
	def := NewMockSource(func(s *MockSource) {
		s.On("String", "STRING_MOCK").Return("asdf", true)
		s.On("StringSlice", "STRING_SLICE_MOCK").Return([]string{"a", "b"}, true)
		s.On("Int", "INT_MOCK").Return(8, true)
		s.On("Uint", "UINT_MOCK").Return(uint(9), true)
	})

	var err error
	suite.source, err = configify.Environment(configify.Options{Namespace: "TEST", Defaults: def})
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
