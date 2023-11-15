package util

import (
	"context"
	"fmt"
	"io"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"github.com/tae2089/gin-boilerplate/common/config"
)

type CloudStorageUtil interface {
	UploadFileToCloud(file io.Reader, fileName string) (string, error)
}

type S3StorageUtil struct {
	s3Client *s3.Client
	bucket   string
}

func NewS3StorageUtil() CloudStorageUtil {
	cfg := config.NewAwsConfig()
	client := s3.NewFromConfig(cfg)
	return &S3StorageUtil{
		s3Client: client,
	}
}

func (s *S3StorageUtil) UploadFileToCloud(file io.Reader, fileName string) (string, error) {
	date, err := formatDate()
	if err != nil {
		panic(err)
	}
	uuid := uuid.New()
	ext := filepath.Ext(fileName)
	_, err = s.s3Client.PutObject(context.Background(), &s3.PutObjectInput{
		Key:    aws.String(fmt.Sprintf("%s/%s%s", date, uuid.String(), ext)),
		Bucket: &s.bucket,
		Body:   file,
	})
	if err != nil {
		return "", err
	}

	objectURL := fmt.Sprintf("https://%s.s3.ap-northeast-2.amazonaws.com/%s/%s%s", s.bucket, date, uuid.String(), ext)
	return objectURL, nil
}

func formatDate() (string, error) {
	location, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		return "", err
	}

	currentTime := time.Now().In(location)
	year := currentTime.Year()
	month := int(currentTime.Month())
	day := currentTime.Day()

	// 월과 일이 한 자리 수인 경우 앞에 0을 추가하여 두 자리로 만듦
	monthStr := fmt.Sprintf("%02d", month)
	dayStr := fmt.Sprintf("%02d", day)

	dateStr := fmt.Sprintf("%d%s%s", year, monthStr, dayStr)
	return dateStr, nil
}
