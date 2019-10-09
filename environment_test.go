package configify_test

import (
	"os"
	"testing"

	"github.com/robsignorelli/configify"
	"github.com/stretchr/testify/suite"
)

func TestEnvironmentSuite(t *testing.T) {
	suite.Run(t, new(EnvironmentSuite))
}

type EnvironmentSuite struct {
	suite.Suite
	source configify.Source
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

func (suite EnvironmentSuite) TestGetString() {
	get := func(key string) string {
		return suite.source.GetString(key)
	}
	suite.Equal("", get("NOT_FOUND"))
	suite.Equal("", get("EMPTY"))
	suite.Equal("foo", get("STRING"))
	suite.Equal("foo bar", get("STRING_SPACE"))
	suite.Equal("foo, bar, baz ,5", get("STRING_SLICE"))
	suite.Equal("5", get("INT"))
	suite.Equal("-3", get("NEGATIVE"))
	suite.Equal("5.430", get("FLOAT"))
	suite.Equal("-5.1", get("NEGATIVE_FLOAT"))

	// Does not fetch values from other namespaces
	suite.Equal("", get("FOO_EMPTY"))
	suite.Equal("", get("FOO_STRING"))
	suite.Equal("", get("FOO_INT"))
}

func (suite EnvironmentSuite) TestGetStringSlice() {
	get := func(key string) []string {
		return suite.source.GetStringSlice(key)
	}
	suite.Len(get("NOT_FOUND"), 0)
	suite.Len(get("EMPTY"), 0)
	suite.EqualValues([]string{"foo"}, get("STRING"))
	suite.EqualValues([]string{"foo bar"}, get("STRING_SPACE"))
	suite.EqualValues([]string{"foo", "bar", "baz", "5"}, get("STRING_SLICE"))
	suite.EqualValues([]string{"5"}, get("INT"))
	suite.EqualValues([]string{"5", "300", "123"}, get("LARGE_INT"))
	suite.EqualValues([]string{"-3"}, get("NEGATIVE"))
	suite.EqualValues([]string{"5.430"}, get("FLOAT"))
	suite.EqualValues([]string{"-5.1"}, get("NEGATIVE_FLOAT"))
	suite.EqualValues([]string{"5", "300", "123.430"}, get("LARGE_FLOAT"))

	// Does not fetch values from other namespaces
	suite.Len(get("FOO_EMPTY"), 0)
	suite.Len(get("FOO_STRING"), 0)
	suite.Len(get("FOO_INT"), 0)
}

func (suite EnvironmentSuite) TestGetInt() {
	get := func(key string) int {
		return suite.source.GetInt(key)
	}
	suite.Equal(5300123, get("LARGE_INT"))
	suite.Equal(0, get("EMPTY"))
	suite.Equal(0, get("STRING"))
	suite.Equal(0, get("STRING_SPACE"))
	suite.Equal(0, get("STRING_SLICE"))
	suite.Equal(5, get("INT"))
	suite.Equal(5300123, get("LARGE_INT"))
	suite.Equal(-3, get("NEGATIVE"))
	suite.Equal(5, get("FLOAT"))
	suite.Equal(-5, get("NEGATIVE_FLOAT"))
	suite.Equal(5300123, get("LARGE_FLOAT"))
	suite.Equal(0, get("JUST_FLOAT"))

	// Does not fetch values from other namespaces
	suite.Equal(0, get("FOO_EMPTY"))
	suite.Equal(0, get("FOO_STRING"))
	suite.Equal(0, get("FOO_INT"))
}

func (suite EnvironmentSuite) TestGetUint() {
	get := func(key string) uint {
		return suite.source.GetUint(key)
	}
	suite.Equal(uint(5300123), get("LARGE_INT"))
	suite.Equal(uint(0), get("EMPTY"))
	suite.Equal(uint(0), get("STRING"))
	suite.Equal(uint(0), get("STRING_SPACE"))
	suite.Equal(uint(0), get("STRING_SLICE"))
	suite.Equal(uint(5), get("INT"))
	suite.Equal(uint(5300123), get("LARGE_INT"))
	suite.Equal(uint(5), get("FLOAT"))
	suite.Equal(uint(5300123), get("LARGE_FLOAT"))
	suite.Equal(uint(0), get("JUST_FLOAT"))

	// Negatives resolve to zero, not the value w/o the minus sign.
	suite.Equal(uint(0), get("NEGATIVE"))
	suite.Equal(uint(0), get("NEGATIVE_FLOAT"))

	// Does not fetch values from other namespaces
	suite.Equal(uint(0), get("FOO_EMPTY"))
	suite.Equal(uint(0), get("FOO_STRING"))
	suite.Equal(uint(0), get("FOO_INT"))
}

func (suite EnvironmentSuite) TestDefaults() {
	def := NewMockSource(func(s *MockSource) {
		s.On("GetString", "STRING_MOCK").Return("asdf")
		s.On("GetStringSlice", "STRING_SLICE_MOCK").Return([]string{"a", "b"})
		s.On("GetInt", "INT_MOCK").Return(8)
		s.On("GetUint", "UINT_MOCK").Return(uint(9))
	})

	env, err := configify.Environment(configify.Options{Namespace: "TEST", Defaults: def})
	suite.Require().NoError(err)

	// For each, make sure that (A) a valid env value resolves, (B) a value in the fixed fallback
	// resolves, and (C) a value that doesn't exist in either uses the hard-coded defaults.
	suite.Equal("foo", env.GetString("STRING"))
	suite.Equal("asdf", env.GetString("STRING_MOCK"))
	suite.Equal("", env.GetString("STRING_XXX"))

	suite.ElementsMatch([]string{"foo", "bar", "baz", "5"}, env.GetStringSlice("STRING_SLICE"))
	suite.ElementsMatch([]string{"a", "b"}, env.GetStringSlice("STRING_SLICE_MOCK"))
	suite.ElementsMatch([]string{}, env.GetStringSlice("STRING_SLICE_XXX"))

	suite.Equal(5, env.GetInt("INT"))
	suite.Equal(8, env.GetInt("INT_MOCK"))
	suite.Equal(0, env.GetInt("INT_XXX"))

	suite.Equal(uint(90), env.GetUint("UINT"))
	suite.Equal(uint(9), env.GetUint("UINT_MOCK"))
	suite.Equal(uint(0), env.GetUint("UINT_XXX"))
}
