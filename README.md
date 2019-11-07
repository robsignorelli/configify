# What is Configify

Configify is a lean library that helps you load configuration values from
any number of key/value stores in a consistent manner. Whether you're 
reading values from the environment, Consul, or a map, Configify provides
a simple API for accessing individual values.

## Basic Usage

In order to fetch config values, you need to create a `Source`. Different
sources interact with different types of key/value stores. 

```
import (
	"github.com/robsignorelli/configify"
)

func main() {
	env := configify.Environment(configify.Options{})
	host := env.String("HTTP_HOST")
	port := env.Uint("HTTP_PORT")
	debugMode := env.Bool("DEBUG_MODE")
	timeout := env.Duration("HTTP_TIMEOUT")
	labels := env.StringSlice("LABELS")
	startTime := env.Time("START_TIME")
	...	
}
```

 ## Setting Default Values
 
 It's quite common to want to have your Source fall back to a known
 default when it does not contain an explicit value for your key. In
 Configify, you can actually define another Source to fall back to.