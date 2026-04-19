package s3

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	log "github.com/sirupsen/logrus"
)

type S3Storage struct {
	Client *minio.Client
}

func NewS3Storage(endpoint, accessKeyID, secretAccessKey string) (*S3Storage, error) {
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: true,
	})
	if err != nil {
		log.Debug(err)
		return nil, err
	}
	return &S3Storage{Client: minioClient}, nil
}

func (s *S3Storage) MakeBucket(bucketName string, location string) error {
	ctx := context.Background()
	err := s.Client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		exists, err := s.Client.BucketExists(ctx, bucketName)
		if err != nil {
			return err
		}
		if exists {
			log.Info(fmt.Sprintf("Bucket %s already exists", bucketName))
			return nil
		}
		return err
	}
	log.Debug("Bucket created")
	return nil
}

func (s *S3Storage) UploadFile(location string, filename string) (string, error) {
	bucketName := "images"
	ctx := context.Background()
	_, err := s.Client.FPutObject(ctx, bucketName, location, filename, minio.PutObjectOptions{
		ContentType: "application/octet-stream",
	})
	if err != nil {
		log.Debug(err)
	}
	//ЗАГЛУШКА
	return bucketName, err
}
