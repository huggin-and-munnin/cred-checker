package config

import (
	"os"
)

type ConfigKey string

var (
	Port ConfigKey = "PORT"

	HealthPath ConfigKey = "HEALTH_PATH"
	HealthPort ConfigKey = "HEALTH_PORT"
)

type ConfigValue struct {
	value any
}

var defaultValues map[ConfigKey]string = map[ConfigKey]string{
	Port: "8080",

	HealthPath: "healthy",
	HealthPort: "8081",
}

func GetValue(key ConfigKey) ConfigValue {
	value := os.Getenv(string(key))
	if defaultValue, hasDefault := defaultValues[key]; len(value) == 0 && hasDefault {
		value = defaultValue
	}
	return ConfigValue{
		value: value,
	}
}

func (v ConfigValue) String() string {
	if val, ok := v.value.(string); ok {
		return val
	}
	panic("ConfigValue is not string!")
}
