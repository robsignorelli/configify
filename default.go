package configify

var emptyStringSlice = make([]string, 0)

// Defaults lets us "dog food" our own Source representation so that when some "real" source
// fails to resolve a value, you can always fall back to this source which is nothing but hard-coded
// values such as "all ints default to 0" and "all strings default to empty".
type Defaults struct{}

func (Defaults) Options() Options {
	return Options{}
}

func (Defaults) String(string) (string, bool) {
	return "", false
}

func (Defaults) StringSlice(string) ([]string, bool) {
	return emptyStringSlice, false
}

func (Defaults) Int(string) (int, bool) {
	return 0, false
}

func (Defaults) Uint(string) (uint, bool) {
	return uint(0), false
}

func (Defaults) GetString(string) string {
	return ""
}

func (Defaults) GetStringSlice(string) []string {
	return emptyStringSlice
}

func (Defaults) GetInt(string) int {
	return 0
}

func (Defaults) GetUint(string) uint {
	return uint(0)
}
