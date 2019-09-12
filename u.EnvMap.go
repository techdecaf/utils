package utils

import (
	"os"
	"strings"
)

// EnvMap is a converts the env to a map fo string[interface]
func EnvMap() map[string]interface{} {
	env := make(map[string]interface{})

	for _, v := range os.Environ() {
		s := strings.Split(v, "=")
		env[s[0]] = strings.Join(s[1:], "=")
	}

	return env
}
