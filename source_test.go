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

func (suite NamespaceSuite) TestNamespace_Qualify() {
	ns := configify.Namespace{}
	suite.Equal("BAR", ns.Qualify("BAR"))

	ns = configify.Namespace{Name: "FOO"}
	suite.Equal("FOO_BAR", ns.Qualify("BAR"))

	ns = configify.Namespace{Name: "FOO", Delimiter: "."}
	suite.Equal("FOO.BAR", ns.Qualify("BAR"))

	ns = configify.Namespace{Name: "FOO  ", Delimiter: "  .  "}
	suite.Equal("FOO.BAR", ns.Qualify("BAR"))
}
