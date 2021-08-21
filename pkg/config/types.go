package config

type Config struct {
	Database Database
}

type Database struct {
	Username, Password, Host, Port string
}
