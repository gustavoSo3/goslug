package utils

import (
	"goslug/internal/logger"
	"os"
)

func SetEnv(variable string, def string) string {
	env, exist := os.LookupEnv(variable)
	if exist {
		logger.Log("ENV_FOUND " + variable + ":" + env)
		return env
	}

	logger.Log("ENV_NOT_FOUND::USING_DEFAULT " + variable + ":" + def)
	return def
}
