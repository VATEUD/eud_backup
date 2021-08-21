package config

// Config represents the config that holds database information
type Config struct {
	Database Database
}

// Database represents database authentication details
type Database struct {
	Username, Password, Host, Port string
}
