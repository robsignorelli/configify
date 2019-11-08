# What is Configify

Configify is a lean library that helps you load configuration values from
any number of key/value stores in a consistent manner. Whether you're 
reading values from the environment, Consul, or a map, Configify provides
a simple API for accessing individual values.

Configify is NOT a way to suck in values JSON, YAML, or some other specialized
type of configuration. If you need that then you might be better off using Viper
or something similar. Configify is for those of us that try to keep our complexity
to a minimum so we let our orchestrator (e.g. Kubernetes) feed values to the environment
or some central store like Consul. This library makes it so you don't have to care
where the values come from - just fetch them and populate your data structures.

## Getting Started

```
go get github.com/robsignorelli/configify
```

## Basic Usage

In order to fetch config values, you need to create a `Source`. Different
sources interact with different types of key/value stores. 

```
import (
	"github.com/robsignorelli/configify"
)

func main() {
	// The 'ok' return value indicates whether we found that value in the
	// source or not. For instance, is "DEBUG_MODE" false because it wasn't
	// in the environment or did you explicitly set it to "false"?
	env, ok := configify.Environment(configify.Options{})
	host, ok := env.String("HTTP_HOST")
	port, ok := env.Uint("HTTP_PORT")
	debugMode, ok := env.Bool("DEBUG_MODE")
	timeout, ok := env.Duration("HTTP_TIMEOUT")
	labels, ok := env.StringSlice("LABELS")
	startTime, ok := env.Time("START_TIME")
	...	
}
```

## Struct Binding

When you pull values from the environment, you typically don't store them in a mess
of free-floating variables. Configify can auto-populate struct values for you.

```
type ServiceConfig struct {
	Host   string `conf:"HTTP_HOST"`,
	Port   uint16 `conf:"HTTP_PORT"`,
	Debug  bool   `conf:"DEBUG_MODE"`

	// Automatically uses the key "LABELS", so no need for 'conf' 
	Labels []string,
}

func main() {
	// Sample environment:
	// HTTP_PORT=1234
	// DEBUG_MODE=true
	// LABELS=foo,bar,baz

	env := configify.Environment(configify.Options{})
	binder := configify.NewBinder(env)

	// You can provide starting values and the binder will only replace
	// what is explicitly defined in the source.
	serviceConfig := ServiceConfig{
		Host: "localhost",
		Port: uint16(9999),
	}
	binder.Bind(&serviceConfig)
	
	// Will start the service w/ these values
	// serviceConfig.Host == "localhost"
	// serviceConfig.Port == uint16(1234)
	// serviceConfig.Debug == true
	// serviceConfig.Labels == [foo bar baz]
	service.Start(serviceConfig)
}
```

## Setting Default Values
 
It's quite common to want to have your Source fall back to a known
default when it does not contain an explicit value for your key. The `Map`
source can be used to provide hard-coded fallback values.


```
func main() {
	// Sample environment:
	// HTTP_PORT=1234
	// DEBUG_MODE=true
	// LABELS=foo,bar,baz

	defaults := configify.Map(configify.Values{
		"HTTP_HOST": "google.com"
		"HTTP_PORT": 9999,
		"LABELS":    []string{"a", "b", "c"}
	})
	env := configify.Environment(configify.Options{
		Defaults: defaults
	})

	// "google.com" 
	host, ok := env.String("HTTP_HOST")
	// 1234 (not 9999 b/c it was defined in the environment)
	port, ok := env.Uint("HTTP_PORT")
	// true
	debugMode, ok := env.Bool("DEBUG_MODE")
	// [a b c]
	labels, ok := env.StringSlice("LABELS")
	...	
}
```

## Namespaces

It's fairly common to provide standard prefixes to all of your keys to avoid conflicts
with common keys like "NAME". So if your program is called "HELLO", you might define
keys like "HELLO_NAME" and "HELLO_PORT". Configify supports defining a namespace in the
Options so you can ask for "NAME" and have "HELLO_" automatically prepended.

```
func main() {
	// Sample environment:
	// HELLO_HTTP_HOST=hello.example.com
	// HELLO_HTTP_PORT=1234
	// HELLO_DEBUG_MODE=true
	// GOODBYE_HTTP_HOST=goodbye.example.com

	// I'm being explicit, but you can leave off `Delimiter: "_"` because that's the
	// default delimiter already when defining a namespace. 
	env := configify.Environment(configify.Options{
		Namespace: configify.Namespace{Name: "HELLO", Delimiter: "_"})
	})

	// When you request values, simply provide the unqualified key name

	// "hello.example.com"
	host, ok := env.String("HTTP_HOST")
	// 1234
	port := env.Uint("HTTP_PORT")
	// true
	debugMode := env.Bool("DEBUG_MODE")
	// We looked up "HELLO_GOODBYE_HTTP_HOST", so this one is empty
	empty, ok := env.String("GOODBYE_HTTP_HOST")
	...	
}
``` 