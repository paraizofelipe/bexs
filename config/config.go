package config

import (
	"os"
	"strconv"
)

var (
	Host    string = os.Getenv("HOST")
	Port    int    = EnvToInt(os.Getenv("PORT"))
	Storage string = os.Getenv("FILE")
	Debug   bool   = EnvToBool(os.Getenv("DEBUG"))
)

func EnvToBool(env string) bool {
	b, err := strconv.ParseBool(env)
	if err != nil {
		return false
	}
	return b
}

func EnvToInt(env string) int {
	i, err := strconv.Atoi(env)
	if err != nil {
		return 0
	}
	return i
}
