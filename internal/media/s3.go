package media

import (
	"bytes"
	"context"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/PNamGP1120/ougreenplus-go/internal/config"
)

func UploadToS3(cfg *config.Config, file *multipart.FileHeader) (string, error) {
	f, err := file.Open()
	if err != nil {
		return "", err
	}
	defer f.Close()

	buffer := bytes.NewBuffer(nil)
	_, err = buffer.ReadFrom(f)
	if err != nil {
		return "", err
	}

	client := s3.New(s3.Options{Region: cfg.AWSRegion})
	uploader := manager.NewUploader(client)

	key := fmt.Sprintf("media/%d_%s", time.Now().Unix(), file.Filename)

	_, err = uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(cfg.S3Bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(buffer.Bytes()),
	})
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("https://%s.s3.amazonaws.com/%s", cfg.S3Bucket, key), nil
}
