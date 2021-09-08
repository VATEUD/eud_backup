package zipper

import (
	"archive/zip"
	backblaze2 "eud_backup/pkg/backblaze"
	"eud_backup/pkg/database"
	"eud_backup/pkg/encryption"
	"eud_backup/pkg/minio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

// Archive represents the zip file
type Archive struct {
	ZipFile *os.File
}

// Upload uploads the zip file to the storage
func (archive *Archive) Upload() error {
	// start the session
	session, err := minio.New()

	if err != nil {
		return err
	}

	backblaze, err := backblaze2.New()

	if err != nil {
		return err
	}

	// read the zip file
	bytes, err := ioutil.ReadFile(archive.ZipFile.Name())

	if err != nil {
		return err
	}

	// defer the file removal (until end of the function)
	defer os.Remove(archive.ZipFile.Name())

	// construct a new cipher
	key, err := encryption.New()

	if err != nil {
		return err
	}

	//encrypt the data
	encrypted, err := key.Encrypt(bytes)

	if err != nil {
		return err
	}

	// create a new temporary binary file
	file, err := os.CreateTemp("", archive.toBin())

	if err != nil {
		return err
	}

	// defer the binary file removal
	defer os.Remove(file.Name())

	// write the binary file
	if _, err = file.Write(encrypted); err != nil {
		return err
	}

	// close (save) the binary file
	if err = file.Close(); err != nil {
		return err
	}

	// upload the binary file
	if err = session.Upload(file); err != nil {
		return err
	}

	if err = backblaze.Upload(file); err != nil {
		return err
	}

	return nil
}

// toBin returns the constructed binary file name
func (archive *Archive) toBin() string {
	name := strings.Split(archive.ZipFile.Name(), "/")[2]
	return fmt.Sprintf("%s.bin", strings.Split(name, ".")[0])
}

// Zip adds database files to the zip file
func (archive *Archive) Zip(databases []*database.Database) []error {
	// store all errors
	var errors []error
	writer := zip.NewWriter(archive.ZipFile)
	defer writer.Close()

	// go through databases
	for _, db := range databases {
		// creates a new file inside of the zip
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

// New creates and returns the temporary zip file that'll be used to store the backup
func New() (*Archive, error) {
	file, err := os.CreateTemp("", fmt.Sprintf("database_backup_%s.zip", time.Now().UTC().Format("2006_01_02")))

	if err != nil {
		return nil, err
	}

	return &Archive{file}, nil
}
