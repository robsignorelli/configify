package configify_test

import (
	"fmt"
	"testing"

	"github.com/robsignorelli/configify"
)

type Foo struct {
	Name      string `conf:"NAMEAGOO"`
	FirstName string
	LastName  string
	Age       int
	Labels    []string
	Labels2   []string
	Numbers   []int
	Http      HttpOptions
	Http2     *HttpOptions
}

type HttpOptions struct {
	Host string
	Port int `conf:"PORK"`
}

func TestModelBinder_Bind(t *testing.T) {
	foo := Foo{
		FirstName: "Jeff",
		LastName:  "Boomstick",
		Labels2:   []string{"Hello", "World"},
		Http2:     &HttpOptions{Host: "farts"},
	}
	fmt.Printf("========\n%+v\n========\n", foo)

	source := configify.Fixed(configify.Options{Namespace: "FOO"}, map[string]interface{}{
		"NAMEAGOO":   "Billyclub",
		"FIRST_NAME": "Bobby",
		"LABELS":     []string{"cow", "sheep", "fox"},
		"HTTP_HOST":  "localhost",
		"HTTP_PORK":  8877,
		"AGE":        22,
	})
	binder := configify.NewBinder(source)
	binder.Bind(&foo)

	fmt.Printf("========\n%+v\n%+v\n========\n", foo, *(foo.Http2))
}
