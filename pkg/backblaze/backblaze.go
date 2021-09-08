package backblaze

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
	"strings"
	"time"
)

// Backblaze represents the S3 session
type Backblaze struct {
	Session *s3.S3
}

// getFileName returns the constructed binary file name
func (backblaze *Backblaze) getFileName(file *os.File) string {
	name := strings.Split(file.Name(), "/")[2]
	return fmt.Sprintf("%s.bin", strings.Split(name, ".")[0])
}

// Upload uploads the file to the storage
func (backblaze *Backblaze) Upload(file *os.File) error {
	// open the file again, because it's closed. It's closed because we had to save in a different function
	file, err := os.Open(file.Name())

	if err != nil {
		return err
	}

	defer file.Close()

	input := &s3.PutObjectInput{
		Body:   file,
		Bucket: aws.String(os.Getenv("BACKBLAZE_BUCKET")),
		Key:    aws.String(fmt.Sprintf("%s/%s", time.Now().UTC().Format("2006-01-02"), backblaze.getFileName(file))),
	}

	_, err = backblaze.Session.PutObject(input)

	return err
}

// New starts a new S3 session with the given details
func New() (*Backblaze, error) {
	config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(os.Getenv("BACKBLAZE_KEY_ID"), os.Getenv("BACKBLAZE_APP_KEY"), ""),
		Endpoint:         aws.String(os.Getenv("BACKBLAZE_ENDPOINT")),
		Region:           aws.String(os.Getenv("BACKBLAZE_REGION")),
		DisableSSL:       aws.Bool(false),
		S3ForcePathStyle: aws.Bool(true),
	}

	s3session, err := session.NewSession(config)

	if err != nil {
		return nil, err
	}

	return &Backblaze{s3.New(s3session)}, nil
}
