package database

import (
	"io/ioutil"
	"os/exec"
	"time"
)

type Database struct {
	Name string
	DumpedBytes []byte
	AddedAt time.Time
}

func (database *Database) Dump() error {
	cmd := exec.Command("mysqldump", "-uroot", "-p1234", database.Name)

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		return err
	}

	if err = cmd.Start(); err != nil {
		return err
	}

	bytes, err := ioutil.ReadAll(stdout)

	if err != nil {
		return err
	}

	database.DumpedBytes = bytes

	return nil
}

