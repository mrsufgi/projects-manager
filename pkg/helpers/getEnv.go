package helpers

import (
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func GetEnvInt(key, fallback string) int {
	val := GetEnv(key, fallback)
	res, err := strconv.Atoi(val)
	if err != nil {
		log.Fatalf("value with key %v cannot be parsed: %v", key, err)
	}
	return res
}
