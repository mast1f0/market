package s3

import (
	"bytes"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Client struct {
	client *s3.Client
	Bucket string
}

func NewS3Client(bucket string) *S3Client {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}

	client := s3.NewFromConfig(cfg)
	return &S3Client{
		Bucket: bucket,
		client: client,
	}
}

func (s *S3Client) UploadFile(filename string, data []byte) (string, error) {
	_, err := s.client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &s.Bucket,
		Key:    &filename,
		Body:   bytes.NewReader(data),
	})

	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", s.Bucket, filename)
	return url, nil
}
