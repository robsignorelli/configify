package functional

// Option provides access to generic, standardized way to include functional arguments/options
// when setting up your application components.
type Option[T any] func(component *T)

// Apply runs the given component through the gamut of args/options, populating your component
// config with the desired values. Typically, 'T' is a pointer to some data structure, so you
// should expect that this will mutate the component you supply (which is standard for this pattern).
func Apply[T any, F ~func(*T)](component *T, options ...F) {
	for _, option := range options {
		option(component)
	}
}
