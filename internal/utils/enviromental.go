package utils

import (
	"goslug/internal/logger"
	"os"
)

// Looks for the variable if found returns its value otherwise returns def
func SetEnv(variable string, def string) string {
	env, exist := os.LookupEnv(variable)
	if exist {
		logger.Log("ENV_FOUND " + variable + ": " + env)
		return env
	}

	logger.Log("ENV_NOT_FOUND::USING_DEFAULT " + variable + ": " + def)
	return def
}
