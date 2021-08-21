package config

import (
	"fmt"
	"os"
	"strings"
)

const (
	defaultDatabaseUsername   = "DB_USERNAME"
	defaultDatabasePassword   = "DB_PASSWORD"
	defaultDatabaseHost       = "DB_HOST"
	defaultDatabasePort       = "DB_PORT"
	singleDatabaseCredentials = "SINGLE_DATABASE_CREDENTIALS"
)

func isUsingSingleDatabaseCredentials() bool {
	return os.Getenv(singleDatabaseCredentials) == "true"
}

func RetrieveDatabaseCredentials(database string) Config {
	if isUsingSingleDatabaseCredentials() {
		return Config{
			Database: Database{
				Username: os.Getenv(defaultDatabaseUsername),
				Password: os.Getenv(defaultDatabasePassword),
				Host:     os.Getenv(defaultDatabaseHost),
				Port:     os.Getenv(defaultDatabasePort),
			},
		}
	}

	database = strings.ToUpper(database)

	return Config{
		Database: Database{
			Username: os.Getenv(fmt.Sprintf("%s_%s", database, defaultDatabaseUsername)),
			Password: os.Getenv(fmt.Sprintf("%s_%s", database, defaultDatabasePassword)),
			Host:     os.Getenv(fmt.Sprintf("%s_%s", database, defaultDatabaseHost)),
			Port:     os.Getenv(fmt.Sprintf("%s_%s", database, defaultDatabasePort)),
		},
	}
}
