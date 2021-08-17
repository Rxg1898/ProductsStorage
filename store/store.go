package store

type Uploader interface {
	UploadFile(bucketName, objectKey, filePath string) error
}
