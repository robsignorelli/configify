package configify_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/robsignorelli/configify"
	"github.com/robsignorelli/configify/configifytest"
	"github.com/stretchr/testify/suite"
)

func TestBinderSuite(t *testing.T) {
	suite.Run(t, new(BinderSuite))
}

type BinderSuite struct {
	configifytest.SourceSuite
}

type TestStruct struct {
	String        string
	String2       string
	StringValue   string
	StringRenamed string `conf:"some_string"`
	StringPointer *string

	StringSlice        []string
	StringSlice2       []string
	StringSliceRenamed []string `conf:"some_string_slice"`

	Int        int
	Int2       int
	IntValue   int
	IntRenamed int `conf:"some_int"`
	IntPointer *int

	Int8  int8
	Int16 int16
	Int32 int32
	Int64 int64

	Uint        uint
	Uint2       uint
	UintValue   uint
	UintRenamed uint `conf:"some_uint"`
	UintPointer *uint

	Uint8  uint8
	Uint16 uint16
	Uint32 uint32
	Uint64 uint64

	Bool        bool
	BoolPointer *bool

	Float32        float32
	Float32Pointer *float32

	Float64        float64
	Float64Pointer *float64

	Duration        time.Duration
	DurationPointer *time.Duration

	Time        time.Time
	TimePointer *time.Time

	Nested           Nested
	NestedRenamed    Nested `conf:"DUDE"`
	NestedPointer    *Nested
	NestedPointerNil *Nested
	Anonymous        struct {
		InnerString        string
		InnerStringRenamed string `conf:"inside_string"`
		InnerStringPointer *string
	}
}

type Nested struct {
	InnerString        string
	InnerStringRenamed string `conf:"inside_string"`
	InnerStringPointer *string
	InnerInt           int
	InnerUint          uint
}

func (suite BinderSuite) populateTestStruct() TestStruct {
	stringPointerVal := "E"
	intPointerVal := 5
	uintPointerVal := uint(15)
	nestedStringVal := "NestedInnerC"
	nestedRenamedStringVal := "NestedRenamedInnerC"
	nestedPointerStringVal := "NestedPointerInnerC"
	durationPointerVal := 10 * time.Minute
	timePointerVal := time.Date(2019, 12, 25, 11, 59, 59, 0, time.UTC)
	boolPointerVal := true
	float32PointerVal := float32(1.57)
	float64PointerVal := 3.14

	value := TestStruct{
		String:        "A",
		String2:       "B",
		StringValue:   "C",
		StringRenamed: "D",
		StringPointer: &stringPointerVal,

		Int:        1,
		Int2:       2,
		IntValue:   3,
		IntRenamed: 4,
		IntPointer: &intPointerVal,

		Uint:        uint(11),
		Uint2:       uint(12),
		UintValue:   uint(13),
		UintRenamed: uint(14),
		UintPointer: &uintPointerVal,

		Uint8:  72,
		Uint16: 1072,
		Uint32: 1000072,
		Uint64: 1000000072,

		Int8:  72,
		Int16: 1072,
		Int32: 1000072,
		Int64: 1000000072,

		StringSlice:        []string{"Foo1", "Bar1"},
		StringSlice2:       []string{"Foo2", "Bar2"},
		StringSliceRenamed: []string{"Foo3", "Bar3"},

		Bool:        true,
		BoolPointer: &boolPointerVal,

		Float32:        999.888,
		Float32Pointer: &float32PointerVal,

		Float64:        123.456,
		Float64Pointer: &float64PointerVal,

		Duration:        5 * time.Minute,
		DurationPointer: &durationPointerVal,

		Time:        time.Date(2019, 9, 1, 12, 0, 0, 0, time.UTC),
		TimePointer: &timePointerVal,

		Nested: Nested{
			InnerString:        "NestedInnerA",
			InnerStringRenamed: "NestedInnerB",
			InnerStringPointer: &nestedStringVal,
			InnerInt:           50,
			InnerUint:          uint(60),
		},

		NestedRenamed: Nested{
			InnerString:        "NestedRenamedInnerA",
			InnerStringRenamed: "NestedRenamedInnerB",
			InnerStringPointer: &nestedRenamedStringVal,
			InnerInt:           51,
			InnerUint:          uint(61),
		},

		NestedPointer: &Nested{
			InnerString:        "NestedPointerInnerA",
			InnerStringRenamed: "NestedPointerInnerB",
			InnerStringPointer: &nestedPointerStringVal,
			InnerInt:           52,
			InnerUint:          uint(62),
		},
	}

	anonymousStringVal := "AnonymousC"
	value.Anonymous.InnerString = "AnonymousA"
	value.Anonymous.InnerStringRenamed = "AnonymousB"
	value.Anonymous.InnerStringPointer = &anonymousStringVal
	return value
}

// TestModelBinder_NilSource ensures that we don't panic when binding objects.
func (suite BinderSuite) TestModelBinder_NilSource() {
	input := suite.populateTestStruct()
	configify.NewBinder(nil).Bind(&input)

	// Other tests will make sure that all defaults are kosher. Just make sure we didn't
	// hose some the original values.
	suite.Equal("A", input.String)
	suite.Equal("B", input.String2)
	suite.Equal("C", input.StringValue)
	suite.Equal("D", input.StringRenamed)
	suite.Equal("E", *input.StringPointer)
}

// TestModelBinder_NoDefaults ensures that binding a struct w/ no values set
// up and a source that will never return valid values leaves the input alone.
func (suite BinderSuite) TestModelBinder_NoDefaults() {
	input := TestStruct{}
	source := configifytest.NewMockSource(func(source *configifytest.MockSource) {})
	configify.NewBinder(source).Bind(&input)

	suite.Equal("", input.String)
	suite.Equal("", input.String2)
	suite.Equal("", input.StringValue)
	suite.Equal("", input.StringRenamed)
	suite.Nil(input.StringPointer)

	suite.Equal(0, input.Int)
	suite.Equal(0, input.Int2)
	suite.Equal(0, input.IntValue)
	suite.Equal(0, input.IntRenamed)
	suite.Nil(input.IntPointer)

	suite.Equal(uint(0), input.Uint)
	suite.Equal(uint(0), input.Uint2)
	suite.Equal(uint(0), input.UintValue)
	suite.Equal(uint(0), input.UintRenamed)
	suite.Nil(input.UintPointer)

	suite.Equal(int8(0), input.Int8)
	suite.Equal(int16(0), input.Int16)
	suite.Equal(int32(0), input.Int32)
	suite.Equal(int64(0), input.Int64)

	suite.Equal(int8(0), input.Int8)
	suite.Equal(int16(0), input.Int16)
	suite.Equal(int32(0), input.Int32)
	suite.Equal(int64(0), input.Int64)

	suite.Nil(input.StringSlice)
	suite.Nil(input.StringSlice2)
	suite.Nil(input.StringSliceRenamed)

	suite.Equal(false, input.Bool)
	suite.Nil(input.BoolPointer)

	suite.Equal(float32(0), input.Float32)
	suite.Nil(input.Float32Pointer)

	suite.Equal(float64(0), input.Float64)
	suite.Nil(input.Float64Pointer)

	suite.Equal(time.Duration(0), input.Duration)
	suite.Nil(input.DurationPointer)

	suite.Equal(time.Time{}, input.Time)
	suite.Nil(input.TimePointer)

	suite.Equal("", input.Nested.InnerString)
	suite.Equal(0, input.Nested.InnerInt)
	suite.Nil(input.NestedPointer)
	suite.Nil(input.NestedPointerNil)
	suite.Equal("", input.Anonymous.InnerString)
}

// TestModelBinder_KeepDefaults ensures that binding a struct w/ initial values set and a source
// with no overrides leaves the input alone. It should have all the same values it came in with.
func (suite BinderSuite) TestModelBinder_KeepDefaults() {
	input := suite.populateTestStruct()
	source := configifytest.NewMockSource(func(source *configifytest.MockSource) {})
	configify.NewBinder(source).Bind(&input)

	suite.Equal("A", input.String)
	suite.Equal("B", input.String2)
	suite.Equal("C", input.StringValue)
	suite.Equal("D", input.StringRenamed)
	suite.Equal("E", *input.StringPointer)

	suite.Equal(1, input.Int)
	suite.Equal(2, input.Int2)
	suite.Equal(3, input.IntValue)
	suite.Equal(4, input.IntRenamed)
	suite.Equal(5, *input.IntPointer)

	suite.Equal(uint(11), input.Uint)
	suite.Equal(uint(12), input.Uint2)
	suite.Equal(uint(13), input.UintValue)
	suite.Equal(uint(14), input.UintRenamed)
	suite.Equal(uint(15), *input.UintPointer)

	suite.Equal(uint8(72), input.Uint8)
	suite.Equal(uint16(1072), input.Uint16)
	suite.Equal(uint32(1000072), input.Uint32)
	suite.Equal(uint64(1000000072), input.Uint64)

	suite.Equal(int8(72), input.Int8)
	suite.Equal(int16(1072), input.Int16)
	suite.Equal(int32(1000072), input.Int32)
	suite.Equal(int64(1000000072), input.Int64)

	suite.ElementsMatch([]string{"Foo1", "Bar1"}, input.StringSlice)
	suite.ElementsMatch([]string{"Foo2", "Bar2"}, input.StringSlice2)
	suite.ElementsMatch([]string{"Foo3", "Bar3"}, input.StringSliceRenamed)

	suite.Equal(5*time.Minute, input.Duration)
	suite.Equal(10*time.Minute, *input.DurationPointer)

	suite.Equal(true, input.Bool)
	suite.Equal(true, *input.BoolPointer)

	suite.Equal(float32(999.888), input.Float32)
	suite.Equal(float32(1.57), *input.Float32Pointer)

	suite.Equal(123.456, input.Float64)
	suite.Equal(3.14, *input.Float64Pointer)

	suite.Equal(time.Date(2019, 9, 1, 12, 0, 0, 0, time.UTC), input.Time)
	suite.Equal(time.Date(2019, 12, 25, 11, 59, 59, 0, time.UTC), *input.TimePointer)

	suite.Equal("AnonymousA", input.Anonymous.InnerString)
	suite.Equal("AnonymousB", input.Anonymous.InnerStringRenamed)
	suite.Equal("AnonymousC", *input.Anonymous.InnerStringPointer)

	suite.Equal("NestedInnerA", input.Nested.InnerString)
	suite.Equal("NestedInnerB", input.Nested.InnerStringRenamed)
	suite.Equal("NestedInnerC", *input.Nested.InnerStringPointer)
	suite.Equal(50, input.Nested.InnerInt)
	suite.Equal(uint(60), input.Nested.InnerUint)

	suite.Equal("NestedRenamedInnerA", input.NestedRenamed.InnerString)
	suite.Equal("NestedRenamedInnerB", input.NestedRenamed.InnerStringRenamed)
	suite.Equal("NestedRenamedInnerC", *input.NestedRenamed.InnerStringPointer)
	suite.Equal(51, input.NestedRenamed.InnerInt)
	suite.Equal(uint(61), input.NestedRenamed.InnerUint)

	suite.Nil(input.NestedPointerNil)

	suite.Require().NotNil(input.NestedPointer)
	suite.Equal("NestedPointerInnerA", input.NestedPointer.InnerString)
	suite.Equal("NestedPointerInnerB", input.NestedPointer.InnerStringRenamed)
	suite.Equal("NestedPointerInnerC", *input.NestedPointer.InnerStringPointer)
	suite.Equal(52, input.NestedPointer.InnerInt)
	suite.Equal(uint(62), input.NestedPointer.InnerUint)
}

// TestModelBinder_OverrideEverything ensures that ALL of our supported types can be bound when they
// are found within the source. This will also ensure that all of the necessary
func (suite BinderSuite) TestModelBinder_OverrideEverything() {
	input := suite.populateTestStruct()

	source := configifytest.NewMockSource(func(source *configifytest.MockSource) {
		source.On("String", "STRING").Return("New-A", true)
		source.On("String", "STRING2").Return("New-B", true)
		source.On("String", "STRING_VALUE").Return("New-C", true)
		source.On("String", "some_string").Return("New-D", true)
		source.On("String", "STRING_POINTER").Return("New-E", true)

		source.On("StringSlice", "STRING_SLICE").Return([]string{"New-Foo1", "New-Bar1"}, true)
		source.On("StringSlice", "STRING_SLICE2").Return([]string{"New-Foo2", "New-Bar2"}, true)
		source.On("StringSlice", "some_string_slice").Return([]string{"New-Foo3", "New-Bar3"}, true)
		source.On("StringSlice", "STRING_SLICE_RENAMED").Return([]string{"NOT THIS ONE"}, true)

		source.On("Int", "INT").Return(101, true)
		source.On("Int", "INT2").Return(102, true)
		source.On("Int", "INT_VALUE").Return(103, true)
		source.On("Int", "some_int").Return(104, true)
		source.On("Int", "INT_POINTER").Return(105, true)

		source.On("Uint", "UINT").Return(uint(111), true)
		source.On("Uint", "UINT2").Return(uint(112), true)
		source.On("Uint", "UINT_VALUE").Return(uint(113), true)
		source.On("Uint", "some_uint").Return(uint(114), true)
		source.On("Uint", "UINT_POINTER").Return(uint(115), true)

		source.On("Uint8", "UINT8").Return(uint8(120), true)
		source.On("Uint16", "UINT16").Return(uint16(121), true)
		source.On("Uint32", "UINT32").Return(uint32(122), true)
		source.On("Uint64", "UINT64").Return(uint64(123), true)

		source.On("Int8", "INT8").Return(int8(124), true)
		source.On("Int16", "INT16").Return(int16(125), true)
		source.On("Int32", "INT32").Return(int32(126), true)
		source.On("Int64", "INT64").Return(int64(127), true)

		// The defaults on the struct are true, so flip them back to false.
		source.On("Bool", "BOOL").Return(false, true)
		source.On("Bool", "BOOL_POINTER").Return(false, true)

		source.On("Float32", "FLOAT32").Return(float32(12345.5), true)
		source.On("Float32", "FLOAT32_POINTER").Return(float32(8080.90), true)

		source.On("Float64", "FLOAT64").Return(99.123, true)
		source.On("Float64", "FLOAT64_POINTER").Return(98765.4321, true)

		source.On("Duration", "DURATION").Return(1*time.Hour, true)
		source.On("Duration", "DURATION_POINTER").Return(2*time.Hour, true)

		source.On("Time", "TIME").Return(time.Date(2000, 1, 1, 3, 4, 5, 6, time.UTC), true)
		source.On("Time", "TIME_POINTER").Return(time.Date(2100, 2, 3, 4, 5, 6, 7, time.UTC), true)

		source.On("String", "ANONYMOUS_INNER_STRING").Return("New-AnonymousA", true)
		source.On("String", "ANONYMOUS_inside_string").Return("New-AnonymousB", true)
		source.On("String", "ANONYMOUS_INNER_STRING_RENAMED").Return("SHOULD BE RENAMED!", true)
		source.On("String", "ANONYMOUS_INNER_STRING_POINTER").Return("New-AnonymousC", true)

		source.On("String", "NESTED_INNER_STRING").Return("New-NestedInnerA", true)
		source.On("String", "NESTED_inside_string").Return("New-NestedInnerB", true)
		source.On("String", "NESTED_INNER_STRING_RENAMED").Return("SHOULD BE RENAMED", true)
		source.On("String", "NESTED_INNER_STRING_POINTER").Return("New-NestedInnerC", true)
		source.On("Int", "NESTED_INNER_INT").Return(150, true)
		source.On("Uint", "NESTED_INNER_UINT").Return(uint(160), true)

		// We have the 'conf' tag on the struct field to rename the prefix from "NESTED_RENAMED" to "DUDE"
		source.On("String", "DUDE_INNER_STRING").Return("New-NestedRenamedInnerA", true)
		source.On("String", "DUDE_inside_string").Return("New-NestedRenamedInnerB", true)
		source.On("String", "DUDE_INNER_STRING_RENAMED").Return("SHOULD BE RENAMED", true)
		source.On("String", "DUDE_INNER_STRING_POINTER").Return("New-NestedRenamedInnerC", true)
		source.On("Int", "DUDE_INNER_INT").Return(151, true)
		source.On("Uint", "DUDE_INNER_UINT").Return(uint(161), true)

		source.On("String", "NESTED_POINTER_INNER_STRING").Return("New-NestedPointerInnerA", true)
		source.On("String", "NESTED_POINTER_inside_string").Return("New-NestedPointerInnerB", true)
		source.On("String", "NESTED_POINTER_INNER_STRING_RENAMED").Return("SHOULD BE RENAMED", true)
		source.On("String", "NESTED_POINTER_INNER_STRING_POINTER").Return("New-NestedPointerInnerC", true)
		source.On("Int", "NESTED_POINTER_INNER_INT").Return(152, true)
		source.On("Uint", "NESTED_POINTER_INNER_UINT").Return(uint(162), true)

		source.On("String", "NESTED_POINTER_NIL_INNER_STRING").Return("SHOULD STILL BE NIL", true)
	})
	configify.NewBinder(source).Bind(&input)

	suite.Equal("New-A", input.String)
	suite.Equal("New-B", input.String2)
	suite.Equal("New-C", input.StringValue)
	suite.Equal("New-D", input.StringRenamed)
	suite.Equal("New-E", *input.StringPointer)

	suite.Equal(101, input.Int)
	suite.Equal(102, input.Int2)
	suite.Equal(103, input.IntValue)
	suite.Equal(104, input.IntRenamed)
	suite.Equal(105, *input.IntPointer)

	suite.Equal(uint(111), input.Uint)
	suite.Equal(uint(112), input.Uint2)
	suite.Equal(uint(113), input.UintValue)
	suite.Equal(uint(114), input.UintRenamed)
	suite.Equal(uint(115), *input.UintPointer)

	suite.Equal(uint8(120), input.Uint8)
	suite.Equal(uint16(121), input.Uint16)
	suite.Equal(uint32(122), input.Uint32)
	suite.Equal(uint64(123), input.Uint64)

	suite.Equal(int8(124), input.Int8)
	suite.Equal(int16(125), input.Int16)
	suite.Equal(int32(126), input.Int32)
	suite.Equal(int64(127), input.Int64)

	suite.Equal(false, input.Bool)
	suite.Equal(false, *input.BoolPointer)

	suite.Equal(float32(12345.5), input.Float32)
	suite.Equal(float32(8080.90), *input.Float32Pointer)

	suite.Equal(99.123, input.Float64)
	suite.Equal(98765.4321, *input.Float64Pointer)

	suite.ElementsMatch([]string{"New-Foo1", "New-Bar1"}, input.StringSlice)
	suite.ElementsMatch([]string{"New-Foo2", "New-Bar2"}, input.StringSlice2)
	suite.ElementsMatch([]string{"New-Foo3", "New-Bar3"}, input.StringSliceRenamed)

	suite.Equal(1*time.Hour, input.Duration)
	suite.Equal(2*time.Hour, *input.DurationPointer)

	suite.Equal(time.Date(2000, 1, 1, 3, 4, 5, 6, time.UTC), input.Time)
	suite.Equal(time.Date(2100, 2, 3, 4, 5, 6, 7, time.UTC), *input.TimePointer)

	suite.Equal("New-AnonymousA", input.Anonymous.InnerString)
	suite.Equal("New-AnonymousB", input.Anonymous.InnerStringRenamed)
	suite.Equal("New-AnonymousC", *input.Anonymous.InnerStringPointer)

	suite.Equal("New-NestedInnerA", input.Nested.InnerString)
	suite.Equal("New-NestedInnerB", input.Nested.InnerStringRenamed)
	suite.Equal("New-NestedInnerC", *input.Nested.InnerStringPointer)
	suite.Equal(150, input.Nested.InnerInt)
	suite.Equal(uint(160), input.Nested.InnerUint)

	suite.Equal("New-NestedRenamedInnerA", input.NestedRenamed.InnerString)
	suite.Equal("New-NestedRenamedInnerB", input.NestedRenamed.InnerStringRenamed)
	suite.Equal("New-NestedRenamedInnerC", *input.NestedRenamed.InnerStringPointer)
	suite.Equal(151, input.NestedRenamed.InnerInt)
	suite.Equal(uint(161), input.NestedRenamed.InnerUint)

	suite.Nil(input.NestedPointerNil)

	suite.Require().NotNil(input.NestedPointer)
	suite.Equal("New-NestedPointerInnerA", input.NestedPointer.InnerString)
	suite.Equal("New-NestedPointerInnerB", input.NestedPointer.InnerStringRenamed)
	suite.Equal("New-NestedPointerInnerC", *input.NestedPointer.InnerStringPointer)
	suite.Equal(152, input.NestedPointer.InnerInt)
	suite.Equal(uint(162), input.NestedPointer.InnerUint)
}

func ExampleNewBinder() {
	// Source attribute names are by convention, so Host will use the string
	// value for "MYAPP_HOST" and Port will use the int value for "MYAPP_PORT".
	// You can customize what attribute we look for by using the 'conf' tag
	// on your field. In this case, Labels is populaed using the string slice
	// value for "MYAPP_TAGS" instead of "MYAPP_LABELS".
	config := struct {
		Host   string
		Port   int
		Labels []string `conf:"TAGS"`
	}{}
	source := configify.Map(configify.Values{
		"HOST": "localhost",
		"PORT": 1234,
		"TAGS": []string{"a", "b", "c"},
	})

	// Overlay the source's value onto our 'config' struct.
	binder := configify.NewBinder(source)
	binder.Bind(&config)

	fmt.Printf("Host: %s\n", config.Host)
	fmt.Printf("Port: %d\n", config.Port)
	fmt.Printf("Tags: %v\n", config.Labels)
	// Output: Host: localhost
	// Port: 1234
	// Tags: [a b c]
}
