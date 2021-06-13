package upload

type Uploader interface {
	MaxFileSize() uint64
	StoreFile(fileName string, file []byte) (*string, error)
	FileExists(fileName string) bool
	GetFile(fileName string) ([]byte, error)
}
