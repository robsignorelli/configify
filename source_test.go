package configify_test

import (
	"testing"

	"github.com/robsignorelli/configify"
	"github.com/stretchr/testify/suite"
)

func TestNamespaceSuite(t *testing.T) {
	suite.Run(t, new(NamespaceSuite))
}

type NamespaceSuite struct {
	suite.Suite
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
