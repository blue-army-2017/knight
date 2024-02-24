package util

import (
	"os"
	"reflect"
)

type ServerConfig struct {
	Port string `name:"PORT" default:"8080"`
}

var Config *ServerConfig

func init() {
	Config = &ServerConfig{}

	t := reflect.TypeOf(*Config)
	v := reflect.ValueOf(Config)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		name := field.Tag.Get("name")
		defaultValue := field.Tag.Get("default")

		value, exists := os.LookupEnv(name)
		if !exists && len(defaultValue) > 0 {
			value = defaultValue
		}

		v.Elem().
			FieldByName(field.Name).
			SetString(value)
	}
}
