package s3

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"io"
	"time"
)

func Singleton() Service {
	return Service{
		PutObject:          PutObject,
		GetObject:          GetObject,
		ToJsonFile:         ToJsonFile,
		GetGetPresignedUrl: GetGetPresignedUrl,
	}
}
func CreateConnection() (*s3.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}
	return s3.NewFromConfig(cfg), nil
}

func PutObject(bucket string, key string, final []byte) error {
	svc, err := CreateConnection()
	if err != nil {
		return err
	}

	input := &s3.PutObjectInput{
		Body:                 bytes.NewReader(final),
		Bucket:               aws.String(bucket),
		Key:                  aws.String(key),
		ServerSideEncryption: "AES256",
		StorageClass:         "STANDARD_IA",
	}

	_, err = svc.PutObject(context.TODO(), input)

	return err
}

func GetObject(bucket string, key string) ([]byte, error) {
	svc, err := CreateConnection()
	if err != nil {
		return []byte{}, err
	}

	input := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	output, err := svc.GetObject(context.TODO(), input)

	if err != nil {
		return nil, err
	}

	raw, err2 := io.ReadAll(output.Body)

	return raw, err2
}

func ToJsonFile(bucket string, key string, data interface{}) error {
	str, err := json.Marshal(data)

	if err != nil {
		return err
	}

	if 0 >= len(bucket) {
		return errors.New("no bucket specified")
	}
	err = PutObject(bucket, key, str)

	if err != nil {
		return err
	}

	return nil
}

func GetGetPresignedUrl(bucket string, key string, expiration time.Duration) (string, error) {
	svc, err := CreateConnection()
	if err != nil {
		return "", err
	}

	presignClient := s3.NewPresignClient(svc)
	req, err := presignClient.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}, func(opts *s3.PresignOptions) {
		opts.Expires = expiration
	})
	if err != nil {
		return "", err
	}
	if req == nil {
		return "", errors.New("unable to prepare \"get\" presigned url")
	}
	return req.URL, nil
}
