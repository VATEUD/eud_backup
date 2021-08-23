package backup

import (
	"gopkg.in/yaml.v2"
	"os"
	"reflect"
	"testing"
)

func TestReadConfigFile(t *testing.T) {
	databases := []string{"test", "api_db"}
	if _, err := os.Stat("configs"); os.IsNotExist(err) {
		if err := os.Mkdir("configs", 0700); err != nil {
			t.Errorf("Couldn't create the file. Error: %s.", err.Error())
		}
	}

	defer os.RemoveAll("configs")

	file, err := os.Create("configs/config.yaml")

	if err != nil {
		t.Errorf("Couldn't create the file. Error: %s.", err.Error())
		return
	}

	c := Config{databases}
	bytes, err := yaml.Marshal(c)

	if err != nil {
		t.Errorf("Failed to marshal the config. Error: %s.", err.Error())
		return
	}

	if _, err = file.Write(bytes); err != nil {
		if err = file.Close(); err != nil {
			t.Errorf("Failed to close the file. Error: %s.", err.Error())
			return
		}
		t.Errorf("Failed to write the file. Error: %s.", err.Error())
		return
	}

	if err = file.Close(); err != nil {
		t.Errorf("Failed to close the file. Error: %s.", err.Error())
		return
	}

	config, err := readConfigFile()

	if err != nil {
		t.Errorf("Error occurred while reading the file. Error: %s.", err.Error())
		return
	}

	if typeName := reflect.TypeOf(config); typeName.String() != "backup.Config" {
		t.Errorf("File not config.Backup type.")
		return
	}

	for i, value := range config.Databases {
		if value != databases[i] {
			t.Errorf("Database doesn't equal the config database. %s != %s.", value, databases[i])
			return
		}
	}
}
