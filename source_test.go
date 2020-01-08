package configify_test

import (
	"context"
	"testing"

	"github.com/robsignorelli/configify"
	"github.com/stretchr/testify/suite"
)

func TestOptionsSuite(t *testing.T) {
	suite.Run(t, new(OptionsSuite))
}

type OptionsSuite struct {
	suite.Suite
}

func (suite OptionsSuite) TestNamespace() {
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

func (suite OptionsSuite) TestContext() {
	options := configify.Options{}
	suite.Nil(options.Context)

	ctx := context.TODO()
	options = configify.Options{}
	configify.Context(ctx)(&options)
	suite.Equal(ctx, options.Context)

	ctx = context.WithValue(context.TODO(), "foo", "bar")
	options = configify.Options{}
	configify.Context(ctx)(&options)
	suite.Equal("bar", options.Context.Value("foo").(string))
}
