package ports

type S3Repository interface {
	UploadFile(filename string, data []byte) (string, error)
}
