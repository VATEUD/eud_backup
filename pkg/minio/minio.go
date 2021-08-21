package minio

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

// Minio represents the S3 session
type Minio struct {
	Session *s3.S3
}

// getFileName returns the constructed binary file name
func (minio *Minio) getFileName(file *os.File) string {
	name := strings.Split(file.Name(), "/")[2]
	return fmt.Sprintf("%s.bin", strings.Split(name, ".")[0])
}

// Upload uploads the file to the storage
func (minio *Minio) Upload(file *os.File) error {
	// open the file again, because it's closed. It's closed because we had to save in a different function
	file, err := os.Open(file.Name())

	if err != nil {
		return err
	}

	defer file.Close()

	input := &s3.PutObjectInput{
		Body:   file,
		Bucket: aws.String(os.Getenv("MINIO_BUCKET")),
		Key:    aws.String(fmt.Sprintf("%s/%s", time.Now().UTC().Format("2006-01-02"), minio.getFileName(file))),
	}

	_, err = minio.Session.PutObject(input)

	return err
}

// New starts a new S3 session with the given details
func New() (*Minio, error) {
	config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(os.Getenv("MINIO_ACCESS_KEY"), os.Getenv("MINIO_SECRET_KEY"), ""),
		Endpoint:         aws.String(os.Getenv("MINIO_ENDPOINT")),
		Region:           aws.String("us-east-1"),
		DisableSSL:       aws.Bool(false),
		S3ForcePathStyle: aws.Bool(true),
	}

	s3session, err := session.NewSession(config)

	if err != nil {
		return nil, err
	}

	return &Minio{s3.New(s3session)}, nil
}
