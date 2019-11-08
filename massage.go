package configify

import (
	"strconv"
	"strings"
	"time"
)

// Massage standardizes how we try to convert values from a source into raw values to
// feed to your program.
type Massage struct {
}

// StringToSlice splits the value by commas, stripping any spaces in the tokens.
func (Massage) StringToSlice(value string) ([]string, bool) {
	if value == "" {
		return []string{}, true
	}
	slice := strings.Split(value, ",")
	for i := range slice {
		slice[i] = strings.TrimSpace(slice[i])
	}
	return slice, true
}

// StringToInt64 parses the value as an integer. This will strip out any commas and
// decimal info before performing the actual parse.
func (m Massage) StringToInt64(value string) (int64, bool) {
	number, err := strconv.ParseInt(m.normalizeInteger(value), 10, 64)
	if err != nil {
		return 0, false
	}
	return number, true
}

// StringToUint64 parses the value as an integer. This will strip out any commas and
// decimal info before performing the actual parse.
func (m Massage) StringToUint64(value string) (uint64, bool) {
	number, err := strconv.ParseUint(m.normalizeInteger(value), 10, 64)
	if err != nil {
		return 0, false
	}
	return number, true
}

// StringToFloat64 parses the value as a floating point number.
func (m Massage) StringToFloat64(value string) (float64, bool) {
	number, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, false
	}
	return number, true
}

// StringToBool converts the strings "true" and "false" (case insensitive) into the
// raw boolean values they represent.
func (m Massage) StringToBool(value string) (bool, bool) {
	switch strings.ToLower(value) {
	case "true":
		return true, true
	case "false":
		return false, true
	default:
		return false, false
	}
}

// StringToDuration parses Go duration strings such as "5m30s" into raw durations.
func (m Massage) StringToDuration(value string) (time.Duration, bool) {
	duration, err := time.ParseDuration(value)
	if err != nil {
		return time.Duration(0), false
	}
	return duration, true
}

// StringToTime parses date/times to the raw time instance it represents. We support
// YYYY-MM-DD strings as well as RFC3339 strings.
func (m Massage) StringToTime(value string) (time.Time, bool) {
	var t time.Time
	var err error

	switch len(value) {
	case 0:
		return time.Time{}, false
	case 10:
		t, err = time.Parse("2006-01-02", value)
	default:
		t, err = time.Parse(time.RFC3339, value)
	}

	if err != nil {
		return time.Time{}, false
	}
	return t, true
}

// normalizeInteger strips all commas and decimal points so you get the raw integer
// encoded in this string.
func (m Massage) normalizeInteger(value string) string {
	decimalPos := strings.IndexRune(value, '.')
	if decimalPos == 0 {
		return ""
	}
	if decimalPos > 0 {
		value = value[:decimalPos]
	}
	return strings.ReplaceAll(value, string(','), "")
}
