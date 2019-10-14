package configify_test

import (
	"testing"
	"time"

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

	suite.ExpectFloat("", float64(0), false)
	suite.ExpectFloat("anything", float64(0), false)

	suite.ExpectBool("", false, false)
	suite.ExpectBool("anything", false, false)

	suite.ExpectDuration("", time.Duration(0), false)
	suite.ExpectDuration("anything", time.Duration(0), false)

	suite.ExpectTime("", time.Time{}, false)
	suite.ExpectTime("anything", time.Time{}, false)
}
