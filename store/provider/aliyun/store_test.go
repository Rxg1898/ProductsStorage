package aliyun_test

import (
	"testing"

	"github.com/Rxg1898/ProductsStorage/store/provider/aliyun"
	"github.com/stretchr/testify/assert"
)

var (
	endpoint   = ""
	ak         = ""
	sk         = ""
	bucketName = ""
	objectKey  = ""
	filePath   = ""
)

func TestUploadFile(t *testing.T) {
	should := assert.New(t)

	uploader, err := aliyun.NewUploader(endpoint, ak, sk)
	if should.NoError(err) {
		err := uploader.UploadFile(bucketName, objectKey, filePath)
		should.NoError(err)
	}
}
