package util

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"sirclo/project/capstone/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func UploadToS3(file *multipart.File, filename string) (string, error) {
	config := config.GetConfig()

	// Destination
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, *file); err != nil {
		return "", err
	}

	s3Config := &aws.Config{
		Region:      aws.String(config.S3Config.Region),
		Credentials: credentials.NewStaticCredentials(config.S3Config.KeyID, config.S3Config.AccessKey, ""),
	}

	s3Session := session.New(s3Config)

	uploader := s3manager.NewUploader(s3Session)

	input := &s3manager.UploadInput{
		Bucket:      aws.String(config.S3Config.BucketName), // bucket's name
		Key:         aws.String(filename),                   // files destination location
		Body:        bytes.NewReader(buf.Bytes()),           // content of the file
		ContentType: aws.String("image/*"),                  // content type
	}

	output, err := uploader.UploadWithContext(context.Background(), input)
	if err != nil {
		return "", err
	}

	return output.Location, nil
}