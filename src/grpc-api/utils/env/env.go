package env

import (
	"os"
	"strconv"
)

func GetEnvString(key string) string {
	return os.Getenv(key)
}

func GetEnvInt(key string) int {
	i, _ := strconv.Atoi(key)
	return i
}
