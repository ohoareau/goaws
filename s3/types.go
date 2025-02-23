package s3

import "time"

type Service struct {
	PutObject          func(bucket string, key string, final []byte) error
	GetObject          func(bucket string, key string) ([]byte, error)
	ToJsonFile         func(bucket string, key string, data interface{}) error
	GetGetPresignedUrl func(bucket string, key string, expiration time.Duration) (string, error)
}
