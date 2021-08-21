package zipper

import (
	"archive/zip"
	"eud_backup/pkg/database"
	"eud_backup/pkg/encryption"
	"eud_backup/pkg/minio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type Archive struct {
	ZipFile *os.File
}

func (archive *Archive) Upload() error {
	session, err := minio.New()

	if err != nil {
		return err
	}

	bytes, err := ioutil.ReadFile(archive.ZipFile.Name())

	defer os.Remove(archive.ZipFile.Name())

	key, err := encryption.New()

	if err != nil {
		return err
	}

	encrypted, err := key.Encrypt(bytes)

	if err != nil {
		return err
	}

	file, err := os.CreateTemp("", archive.ToBin())

	if err != nil {
		return err
	}

	defer os.Remove(file.Name())

	if _, err = file.Write(encrypted); err != nil {
		return err
	}

	if err = session.Upload(file); err != nil {
		return err
	}

	return nil
}

func (archive *Archive) ToBin() string {
	name := strings.Split(archive.ZipFile.Name(), "/")[2]
	return fmt.Sprintf("%s.bin", strings.Split(name, ".")[0])
}

func (archive *Archive) Zip(databases []*database.Database) []error {
	var errors []error
	writer := zip.NewWriter(archive.ZipFile)
	defer writer.Close()

	for _, db := range databases {
		file, err := writer.Create(db.FileName())

		if err != nil {
			errors = append(errors, err)
			continue
		}

		if _, err := file.Write(db.DumpedBytes); err != nil {
			errors = append(errors, err)
		}
	}

	return errors
}

func New() (*Archive, error) {
	file, err := os.CreateTemp("", fmt.Sprintf("database_backup_%s.zip", time.Now().UTC().Format("2006_01_02")))

	if err != nil {
		return nil, err
	}

	return &Archive{file}, nil
}
