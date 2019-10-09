package configify

// defaults lets us "dog food" our own Source representation so that when some "real" source
// fails to resolve a value, you can always fall back to this source which is nothing but hard-coded
// values such as "all ints default to 0" and "all strings default to empty".
type defaults struct{}

func (defaults) GetString(key string) string {
	return ""
}

func (defaults) GetStringSlice(key string) []string {
	return emptyStringSlice
}

func (defaults) GetInt(key string) int {
	return 0
}

func (defaults) GetUint(key string) uint {
	return uint(0)
}
