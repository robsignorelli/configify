package functional_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/robsignorelli/configify/functional"
	"github.com/stretchr/testify/suite"
)

func TestOptionsSuite(t *testing.T) {
	suite.Run(t, new(OptionsSuite))
}

type OptionsSuite struct {
	suite.Suite
}

func (s *OptionsSuite) TestApply_keepDefaults() {
	nop := func(config *Config) {}
	nopOption := ConfigOption(nop)

	// Don't supply any options - since they're generic you need to specify type
	// parameters manually.
	config := Config{}
	functional.Apply[Config, ConfigOption](&config)
	s.Equal(0, config.A)
	s.Equal("", config.B)

	// Standard usage where you'd use varargs to pass 0-N option functions.
	config = Config{A: 42, B: "Hello"}
	functional.Apply(&config, noOptions...)
	s.Equal(42, config.A)
	s.Equal("Hello", config.B)

	// Specify an alias to the function type
	config = Config{A: 42, B: "Hello"}
	functional.Apply(&config, nopOption)
	s.Equal(42, config.A)
	s.Equal("Hello", config.B)

	// Specify a raw function as the options.
	config = Config{A: 42, B: "Hello"}
	functional.Apply(&config, nop)
	s.Equal(42, config.A)
	s.Equal("Hello", config.B)

	// Specify a mix of raw functions and reasonable aliases.
	config = Config{A: 42, B: "Hello"}
	functional.Apply(&config, nop, nopOption)
	s.Equal(42, config.A)
	s.Equal("Hello", config.B)
}

func (s *OptionsSuite) TestApply_overrideDefaults() {
	// Override A, but leave B
	config := Config{A: 42, B: "Hello"}
	functional.Apply(&config,
		func(c *Config) { c.A = 1024 },
	)
	s.Equal(1024, config.A)
	s.Equal("Hello", config.B)

	// Override B, but leave A
	config = Config{A: 42, B: "Hello"}
	functional.Apply(&config,
		func(c *Config) { c.B = "Goodbye" },
	)
	s.Equal(42, config.A)
	s.Equal("Goodbye", config.B)

	// Override everything
	config = Config{A: 42, B: "Hello"}
	functional.Apply(&config,
		func(c *Config) { c.B = "Goodbye" },
		func(c *Config) { c.A = 1234 },
	)
	s.Equal(1234, config.A)
	s.Equal("Goodbye", config.B)

	// Override everything multiple times (last in wins)
	config = Config{A: 42, B: "Hello"}
	functional.Apply(&config,
		func(c *Config) { c.B = "Goodbye" },
		func(c *Config) { c.A = 1234 },
		func(c *Config) { c.B = "The Empire Strikes First" },
		func(c *Config) { c.A = 1984 },
	)
	s.Equal(1984, config.A)
	s.Equal("The Empire Strikes First", config.B)
}

func (s *OptionsSuite) TestApply_alias() {
	type Config struct {
		A int
		B string
	}

	// Make sure that we allow users to define options either as direct functions or an
	// alias to our convenience Option[T] type.
	type ConfigOption functional.Option[Config]
	type AnotherAlias func(*Config)

	optionA := ConfigOption(func(c *Config) { c.A = 8888 })
	optionB := ConfigOption(func(c *Config) { c.B = "Chicken Butt" })

	config := Config{A: 42, B: "Hello"}
	functional.Apply(&config, optionA, optionB)
	s.Equal(8888, config.A)
	s.Equal("Chicken Butt", config.B)

	anotherA := AnotherAlias(func(c *Config) { c.A = 8888 })
	anotherB := AnotherAlias(func(c *Config) { c.B = "Chicken Butt" })

	config = Config{A: 42, B: "Hello"}
	functional.Apply(&config, anotherA, anotherB)
	s.Equal(8888, config.A)
	s.Equal("Chicken Butt", config.B)
}

func (s *OptionsSuite) TestApply_rawFuncs() {
	type Config struct {
		A int
		B string
	}

	optionA := func(c *Config) { c.A = 8888 }
	optionB := func(c *Config) { c.B = "Chicken Butt" }

	config := Config{A: 42, B: "Hello"}
	functional.Apply(&config, optionA, optionB)
	s.Equal(8888, config.A)
	s.Equal("Chicken Butt", config.B)
}

// Here's a simple example showing how you can use functional option
// style (https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis)
// to configure components without needing custom types and without the copy/pasted
// for loop to invoke all of your option functions.
func ExampleApply() {
	type Client struct {
		Address string
		Port    uint16
		Timeout time.Duration
	}

	var WithPort = func(port uint16) functional.Option[Client] {
		return func(client *Client) { client.Port = port }
	}

	var WithTimeout = func(timeout time.Duration) functional.Option[Client] {
		return func(client *Client) { client.Timeout = timeout }
	}

	var NewClient = func(address string, options ...functional.Option[Client]) Client {
		client := Client{
			Address: address,
			Port:    443,
			Timeout: 1 * time.Minute,
		}
		functional.Apply(&client, options...)
		return client
	}

	clientA := NewClient("https://go.dev",
		WithTimeout(5*time.Second),
	)
	clientB := NewClient("https://some-random-api.io",
		WithPort(9000),
		WithTimeout(10*time.Second),
	)

	fmt.Printf("Client A: %s [%d] [%v]\n", clientA.Address, clientA.Port, clientA.Timeout)
	fmt.Printf("Client B: %s [%d] [%v]\n", clientB.Address, clientB.Port, clientB.Timeout)

	// Output:
	// Client A: https://go.dev [443] [5s]
	// Client B: https://some-random-api.io [9000] [10s]
}

// You can also define custom types for your options if you feel that provides more
// clarity to your final code.
func ExampleApply_customType() {
	type Client struct {
		Address string
		Port    uint16
		Timeout time.Duration
	}
	type ClientOption func(*Client)

	var WithPort = func(port uint16) ClientOption {
		return func(client *Client) { client.Port = port }
	}

	var WithTimeout = func(timeout time.Duration) ClientOption {
		return func(client *Client) { client.Timeout = timeout }
	}

	var NewClient = func(address string, options ...ClientOption) Client {
		client := Client{
			Address: address,
			Port:    443,
			Timeout: 1 * time.Minute,
		}
		functional.Apply(&client, options...)
		return client
	}

	clientA := NewClient("https://go.dev",
		WithTimeout(5*time.Second),
	)
	clientB := NewClient("https://some-random-api.io",
		WithPort(9000),
		WithTimeout(10*time.Second),
	)

	fmt.Printf("Client A: %s [%d] [%v]\n", clientA.Address, clientA.Port, clientA.Timeout)
	fmt.Printf("Client B: %s [%d] [%v]\n", clientB.Address, clientB.Port, clientB.Timeout)

	// Output:
	// Client A: https://go.dev [443] [5s]
	// Client B: https://some-random-api.io [9000] [10s]
}

type Config struct {
	A int
	B string
}
type ConfigOption func(*Config)

var noOptions []ConfigOption
