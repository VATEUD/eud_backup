package backup

// Config represents the config which holds databases list
type Config struct {
	Databases []string `yaml:"databases"`
}
