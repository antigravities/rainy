package upload

import "io"

type Uploader interface {
	MaxFileSize() uint64
	StoreFile(fileName string, file []byte) (*string, error)
	StoreFileStream(fileName string, stream io.ReadCloser) (*string, error)
	FileExists(fileName string) bool
	GetFile(fileName string) ([]byte, error)
}
