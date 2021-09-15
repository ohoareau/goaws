package s3

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io/ioutil"
)

func Singleton() Service {
	return Service{
		PutObject: PutObject,
		GetObject: GetObject,
	}
}
func CreateConnection() *s3.S3 {
	sess := session.Must(session.NewSession())
	return s3.New(sess)
}

func PutObject(bucket string, key string, final []byte) error {
	svc := CreateConnection()

	input := &s3.PutObjectInput{
		Body:                 bytes.NewReader(final),
		Bucket:               aws.String(bucket),
		Key:                  aws.String(key),
		ServerSideEncryption: aws.String("AES256"),
		StorageClass:         aws.String("STANDARD_IA"),
	}

	_, err := svc.PutObject(input)

	return err
}

func GetObject(bucket string, key string) ([]byte, error) {
	svc := CreateConnection()

	input := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	output, err := svc.GetObject(input)

	if err != nil {
		return nil, err
	}

	raw, err2 := ioutil.ReadAll(output.Body)

	return raw, err2
}