package util

import (
	"fmt"
	"os"
	"reflect"
)

type ServerConfig struct {
	Environment string `name:"ENVIRONMENT" default:"production"`
	Port        string `name:"PORT" default:"8080"`
	DB          string `name:"DB"`
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

		if value == "" {
			panic(fmt.Sprintf("Configuration for %s is missing", name))
		}

		v.Elem().
			FieldByName(field.Name).
			SetString(value)
	}
}
