package goaws

import (
	"github.com/ohoareau/goaws/s3"
)

//goland:noinspection GoUnusedExportedFunction
func S3() S3Service {
	return s3.Singleton()
}
