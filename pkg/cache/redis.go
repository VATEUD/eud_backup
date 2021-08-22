package cache

import (
	"github.com/go-redis/redis/v8"
	"os"
	"strconv"
)

const (
	defaultAddressFieldName  = "REDIS_ADDRESS"
	defaultPasswordFieldName = "REDIS_PASSWORD"
	defaultDatabaseFieldName = "REDIS_DB"
	defaultAddress           = "localhost:6379"
	defaultDatabase          = 0
)

// New constructs a new Redis client
func New() *redis.Client {
	return redis.NewClient(retrieveConfig())
}

// retrieveConfig returns Redis options
func retrieveConfig() *redis.Options {
	return &redis.Options{
		Addr:     getAddress(),
		Password: os.Getenv(defaultPasswordFieldName),
		DB:       getDatabase(),
	}
}

// getAddress returns the address from the environment file. If it's not available, it'll return the default one
func getAddress() string {
	if address := os.Getenv(defaultAddressFieldName); address != "" {
		return address
	}

	return defaultAddress
}

// getDatabase returns the database from the environment file. If it's not available, it'll return the default one
func getDatabase() int {
	if database := os.Getenv(defaultDatabaseFieldName); database != "" {
		db, err := strconv.Atoi(database)
		if err != nil {
			return defaultDatabase
		}

		return db
	}

	return defaultDatabase
}
