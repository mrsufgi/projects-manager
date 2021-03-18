package helpers

import "fmt"

// GetConnectionString returns string needed to connect to pg, for either test (store\integration) or main app
func GetConnectionString() string {
	dbname := GetEnv("POSTGRES_DB", "postgres")
	user := GetEnv("POSTGRES_USER", "postgres")
	password := GetEnv("POSTGRES_PASSWORD", "")
	host := GetEnv("POSTGRES_HOST", "localhost")
	port := GetEnv("POSTGRES_PORT", "5432")
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
}
