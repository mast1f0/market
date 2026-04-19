package ports

type S3Repository interface {
	UploadFile(location string, filename string) (string, error)
	MakeBucket(bucketName string, location string) error
}
