package configify_test

import (
	"testing"

	"github.com/robsignorelli/configify"
	"github.com/stretchr/testify/suite"
)

func TestDefaultsSuite(t *testing.T) {
	suite.Run(t, new(DefaultsSuite))
}

type DefaultsSuite struct {
	suite.Suite
	source configify.Source
}

func (s *DefaultsSuite) SetupSuite() {
	s.source = configify.Defaults("TEST", map[string]interface{}{
		suite.set("TEST_EMPTY", "")
		suite.set("TEST_STRING", "foo")
		suite.set("TEST_STRING_SPACE", "  foo bar ")
		suite.set("TEST_STRING_SLICE", "foo, bar, baz ,5 ")
		suite.set("TEST_INT", "5")
		suite.set("TEST_LARGE_INT", "5,300,123")
		suite.set("TEST_NEGATIVE", "-3")
		suite.set("TEST_FLOAT", "5.430")
		suite.set("TEST_LARGE_FLOAT", "5,300,123.430")
		suite.set("TEST_NEGATIVE_FLOAT", "-5.1")
		suite.set("TEST_JUST_FLOAT", ".1")

	})
}
