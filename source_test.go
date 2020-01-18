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

func (suite OptionsSuite) TestCredentials() {
	options := configify.Options{}
	suite.Empty(options.Username)
	suite.Empty(options.Password)

	options = configify.Options{}
	configify.Username("")(&options)
	suite.Empty(options.Username)

	options = configify.Options{}
	configify.Username(" foo ! bar#.")(&options)
	suite.Equal(" foo ! bar#.", options.Username)

	options = configify.Options{}
	configify.Password("")(&options)
	suite.Empty(options.Password)

	options = configify.Options{}
	configify.Password(" foo ! bar#.")(&options)
	suite.Equal(" foo ! bar#.", options.Password)

	options = configify.Options{}
	configify.Username("hello")(&options)
	configify.Password("world")(&options)
	suite.Equal("hello", options.Username)
	suite.Equal("world", options.Password)
}

func (suite OptionsSuite) TestAddress() {
	options := configify.Options{}
	suite.Empty(options.Address)

	options = configify.Options{}
	configify.Address("")(&options)
	suite.Empty(options.Address)

	options = configify.Options{}
	configify.Address("1.1.1.1:9023")(&options)
	suite.Equal("1.1.1.1:9023", options.Address)

	options = configify.Options{}
	configify.Address(" foo ! bar#.")(&options)
	suite.Equal(" foo ! bar#.", options.Address) // doesn't require a "host:port" format
}
