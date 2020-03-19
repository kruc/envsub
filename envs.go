package main

import (
	"os"
	"strings"
)

type envVars map[string]string

func prepareEnvs() envVars {
	envs := make(envVars)

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		envs[pair[0]] = pair[1]
	}
	return envs
}

func convert(envs map[string]string) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range envs {
		result[k] = v
	}

	return result
}
