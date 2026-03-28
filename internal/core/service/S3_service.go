package service

import "market/internal/core/ports"

type S3Service struct {
	repo ports.S3Repository
}

func NewS3Service(repo ports.S3Repository) *S3Service {
	return &S3Service{repo: repo}
}

func (s *S3Service) UploadFile(filename string, data []byte) (string, error) {
	return s.repo.UploadFile(filename, data)
}
