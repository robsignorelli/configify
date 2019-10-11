package configify_test

import (
	"testing"

	"github.com/robsignorelli/configify"
	"github.com/stretchr/testify/suite"
)

func TestDefaultSuite(t *testing.T) {
	suite.Run(t, new(DefaultSuite))
}

type DefaultSuite struct {
	SourceSuite
}

func (suite DefaultSuite) TestAll() {
	suite.source = configify.Defaults{}
	options := suite.source.Options()
	suite.Equal("", options.Namespace)
	suite.Equal("", options.NamespaceDelim)
	suite.Nil(options.Defaults)

	suite.ExpectString("", "", false)
	suite.ExpectString("anything", "", false)

	suite.ExpectStringSlice("", []string{}, false)
	suite.ExpectStringSlice("anything", []string{}, false)

	suite.ExpectInt("", 0, false)
	suite.ExpectInt("anything", 0, false)

	suite.ExpectUint("", uint(0), false)
	suite.ExpectUint("anything", uint(0), false)
}
