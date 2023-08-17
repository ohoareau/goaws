package s3

type Service struct {
	PutObject  func(bucket string, key string, final []byte) error
	GetObject  func(bucket string, key string) ([]byte, error)
	ToJsonFile func(bucket string, key string, data interface{}) error
}
