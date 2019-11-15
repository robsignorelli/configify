package configify_test

import (
	"testing"
	"time"

	"github.com/robsignorelli/configify"
	"github.com/robsignorelli/configify/configifytest"
	"github.com/stretchr/testify/suite"
)

func TestDefaultSuite(t *testing.T) {
	suite.Run(t, new(DefaultSuite))
}

type DefaultSuite struct {
	configifytest.SourceSuite
}

func (suite DefaultSuite) TestAll() {
	suite.Source = configify.Empty()
	options := suite.Source.Options()
	suite.Equal("", options.Namespace.Name)
	suite.Equal("", options.Namespace.Delimiter)
	suite.Nil(options.Defaults)

	suite.ExpectString("", "", false)
	suite.ExpectString("anything", "", false)

	suite.ExpectStringSlice("", []string{}, false)
	suite.ExpectStringSlice("anything", []string{}, false)

	suite.ExpectInt("", 0, false)
	suite.ExpectInt("anything", 0, false)

	suite.ExpectInt8("", int8(0), false)
	suite.ExpectInt8("anything", int8(0), false)

	suite.ExpectInt16("", int16(0), false)
	suite.ExpectInt16("anything", int16(0), false)

	suite.ExpectInt32("", int32(0), false)
	suite.ExpectInt32("anything", int32(0), false)

	suite.ExpectInt64("", int64(0), false)
	suite.ExpectInt64("anything", int64(0), false)

	suite.ExpectUint("", uint(0), false)
	suite.ExpectUint("anything", uint(0), false)

	suite.ExpectUint8("", uint8(0), false)
	suite.ExpectUint8("anything", uint8(0), false)

	suite.ExpectUint16("", uint16(0), false)
	suite.ExpectUint16("anything", uint16(0), false)

	suite.ExpectUint32("", uint32(0), false)
	suite.ExpectUint32("anything", uint32(0), false)

	suite.ExpectUint64("", uint64(0), false)
	suite.ExpectUint64("anything", uint64(0), false)

	suite.ExpectFloat32("", float32(0), false)
	suite.ExpectFloat32("anything", float32(0), false)

	suite.ExpectFloat64("", float64(0), false)
	suite.ExpectFloat64("anything", float64(0), false)

	suite.ExpectBool("", false, false)
	suite.ExpectBool("anything", false, false)

	suite.ExpectDuration("", time.Duration(0), false)
	suite.ExpectDuration("anything", time.Duration(0), false)

	suite.ExpectTime("", time.Time{}, false)
	suite.ExpectTime("anything", time.Time{}, false)
}
