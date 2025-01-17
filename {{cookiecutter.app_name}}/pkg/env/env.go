package env

import (
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type commonValue interface {
	int | bool | string
}

func Get[T commonValue](key string, def T) T {
	var result any = def
	valueStr := strings.TrimSpace(os.Getenv(key))

	switch reflect.TypeOf(def).Kind() {
	case reflect.String:
		if valueStr != "" {
			result = valueStr
		}

	case reflect.Bool:
		if valueStr != "" {
			val, err := strconv.ParseBool(strings.ToLower(valueStr))
			if err != nil {
				return def
			}
			result = val
		}

	case reflect.Int:
		if valueStr != "" {
			val, err := strconv.Atoi(valueStr)
			if err != nil {
				return def
			}
			result = val
		}
	}

	return result.(T)
}

func GetDuration(key string, def time.Duration) time.Duration {
	valueStr := Get(key, "")
	if value, err := time.ParseDuration(valueStr); err == nil {
		return value
	}

	return def
}
