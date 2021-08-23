package cache

import (
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"os"
	"reflect"
	"strconv"
	"testing"
)

const (
	testingPassword = "test"
	testingAddress  = "vateud.net:3102"
	testingDatabase = 10
)

func TestNewDefault(t *testing.T) {
	client := New()

	if typeDetails := reflect.TypeOf(client); typeDetails.String() != "*redis.Client" {
		t.Errorf("Type doesn't match. Expected *redis.Client, got %s.", typeDetails)
		return
	}

	rClient := redis.NewClient(&redis.Options{
		Addr:     defaultAddress,
		Password: "",
		DB:       defaultDatabase,
	})

	if ok := assert.Equal(t, rClient.Options().Addr, client.Options().Addr); !ok {
		t.Errorf("Addresses don't match. Expected %s, got %s.", rClient.Options().Addr, client.Options().Addr)
		return
	}

	if ok := assert.Equal(t, rClient.Options().Password, client.Options().Password); !ok {
		t.Errorf("Passwords don't match. Expected %s, got %s.", rClient.Options().Password, client.Options().Password)
		return
	}

	if ok := assert.Equal(t, rClient.Options().DB, client.Options().DB); !ok {
		t.Errorf("Databases don't match. Expected %d, got %d.", rClient.Options().DB, client.Options().DB)
		return
	}
}

func TestNewEnvironment(t *testing.T) {
	if err := os.Setenv("REDIS_ADDRESS", testingAddress); err != nil {
		t.Errorf("Failed to set the environment variable. Error: %s.", err.Error())
		return
	}
	if err := os.Setenv("REDIS_PASSWORD", testingPassword); err != nil {
		t.Errorf("Failed to set the environment variable. Error: %s.", err.Error())
		return
	}
	if err := os.Setenv("REDIS_DB", strconv.Itoa(testingDatabase)); err != nil {
		t.Errorf("Failed to set the environment variable. Error: %s.", err.Error())
		return
	}
	client := New()

	if typeDetails := reflect.TypeOf(client); typeDetails.String() != "*redis.Client" {
		t.Errorf("Type doesn't match. Expected *redis.Client, got %s.", typeDetails)
		return
	}

	rClient := redis.NewClient(&redis.Options{
		Addr:     testingAddress,
		Password: testingPassword,
		DB:       testingDatabase,
	})

	if ok := assert.Equal(t, rClient.Options().Addr, client.Options().Addr); !ok {
		t.Errorf("Addresses don't match. Expected %s, got %s.", rClient.Options().Addr, client.Options().Addr)
		return
	}

	if ok := assert.Equal(t, rClient.Options().Password, client.Options().Password); !ok {
		t.Errorf("Passwords don't match. Expected %s, got %s.", rClient.Options().Password, client.Options().Password)
		return
	}

	if ok := assert.Equal(t, rClient.Options().DB, client.Options().DB); !ok {
		t.Errorf("Databases don't match. Expected %d, got %d.", rClient.Options().DB, client.Options().DB)
		return
	}
}

func TestConfigDefault(t *testing.T) {
	if err := os.Setenv("REDIS_ADDRESS", defaultAddress); err != nil {
		t.Errorf("Failed to set the environment variable. Error: %s.", err.Error())
		return
	}
	if err := os.Setenv("REDIS_PASSWORD", ""); err != nil {
		t.Errorf("Failed to set the environment variable. Error: %s.", err.Error())
		return
	}
	if err := os.Setenv("REDIS_DB", strconv.Itoa(defaultDatabase)); err != nil {
		t.Errorf("Failed to set the environment variable. Error: %s.", err.Error())
		return
	}
	config := retrieveConfig()

	if typeDetails := reflect.TypeOf(config); typeDetails.String() != "*redis.Options" {
		t.Errorf("Type doesn't match. Expected *redis.Options, got %s.", typeDetails)
		return
	}

	options := redis.Options{
		Addr:     defaultAddress,
		Password: "",
		DB:       defaultDatabase,
	}

	if ok := assert.Equal(t, config.Addr, options.Addr); !ok {
		t.Errorf("Addresses don't match. Expected %s, got %s.", options.Addr, config.Addr)
		return
	}

	if ok := assert.Equal(t, config.Password, options.Password); !ok {
		t.Errorf("Passwords don't match. Expected %s, got %s.", options.Password, config.Password)
		return
	}

	if ok := assert.Equal(t, config.DB, options.DB); !ok {
		t.Errorf("Databases don't match. Expected %d, got %d.", options.DB, config.DB)
		return
	}
}

func TestConfigEnvironment(t *testing.T) {
	if err := os.Setenv("REDIS_ADDRESS", testingAddress); err != nil {
		t.Errorf("Failed to set the environment variable. Error: %s.", err.Error())
		return
	}
	if err := os.Setenv("REDIS_PASSWORD", testingPassword); err != nil {
		t.Errorf("Failed to set the environment variable. Error: %s.", err.Error())
		return
	}
	if err := os.Setenv("REDIS_DB", strconv.Itoa(testingDatabase)); err != nil {
		t.Errorf("Failed to set the environment variable. Error: %s.", err.Error())
		return
	}
	config := retrieveConfig()

	if typeDetails := reflect.TypeOf(config); typeDetails.String() != "*redis.Options" {
		t.Errorf("Type doesn't match. Expected *redis.Options, got %s.", typeDetails)
		return
	}

	options := redis.Options{
		Addr:     testingAddress,
		Password: testingPassword,
		DB:       testingDatabase,
	}

	if ok := assert.Equal(t, config.Addr, options.Addr); !ok {
		t.Errorf("Addresses don't match. Expected %s, got %s.", options.Addr, config.Addr)
		return
	}

	if ok := assert.Equal(t, config.Password, options.Password); !ok {
		t.Errorf("Passwords don't match. Expected %s, got %s.", options.Password, config.Password)
		return
	}

	if ok := assert.Equal(t, config.DB, options.DB); !ok {
		t.Errorf("Databases don't match. Expected %d, got %d.", options.DB, config.DB)
		return
	}
}

func TestGetAddressDefault(t *testing.T) {
	if err := os.Setenv("REDIS_ADDRESS", defaultAddress); err != nil {
		t.Errorf("Failed to set the environment variable. Error: %s.", err.Error())
		return
	}
	address := getAddress()

	if typeDetails := reflect.TypeOf(address); typeDetails.String() != "string" {
		t.Errorf("Expected type string, got %s.", typeDetails)
		return
	}

	if address != defaultAddress {
		t.Errorf("Expected %s, got %s.", defaultAddress, address)
	}
}

func TestGetAddressEnvironment(t *testing.T) {
	if err := os.Setenv("REDIS_ADDRESS", testingAddress); err != nil {
		t.Errorf("Failed to set the environment variable. Error: %s.", err.Error())
		return
	}
	address := getAddress()

	if typeDetails := reflect.TypeOf(address); typeDetails.String() != "string" {
		t.Errorf("Expected type string, got %s.", typeDetails)
		return
	}

	if address != testingAddress {
		t.Errorf("Expected %s, got %s.", defaultAddress, address)
	}
}

func TestGetDatabaseDefault(t *testing.T) {
	if err := os.Setenv("REDIS_DB", strconv.Itoa(defaultDatabase)); err != nil {
		t.Errorf("Failed to set the environment variable. Error: %s.", err.Error())
		return
	}
	database := getDatabase()

	if typeDetails := reflect.TypeOf(database); typeDetails.String() != "int" {
		t.Errorf("Expected type string, got %s.", typeDetails)
		return
	}

	if database != defaultDatabase {
		t.Errorf("Expected %d, got %d.", defaultDatabase, database)
	}
}

func TestGetDatabaseEnvironment(t *testing.T) {
	if err := os.Setenv("REDIS_DB", strconv.Itoa(testingDatabase)); err != nil {
		t.Errorf("Failed to set the environment variable. Error: %s.", err.Error())
		return
	}
	database := getDatabase()

	if typeDetails := reflect.TypeOf(database); typeDetails.String() != "int" {
		t.Errorf("Expected type string, got %s.", typeDetails)
		return
	}

	if database != testingDatabase {
		t.Errorf("Expected %d, got %d.", testingDatabase, database)
	}
}
