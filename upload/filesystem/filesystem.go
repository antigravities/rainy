package filesystem

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"get.cutie.cafe/rainy/upload"
)

type FilesystemUploader struct {
	path string
}

func New(path string) (upload.Uploader, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, 0666)

		if err != nil {
			return nil, err
		}
	}

	fx := &FilesystemUploader{
		path: path,
	}

	return fx, nil
}

func (f *FilesystemUploader) MaxFileSize() uint64 {
	num, err := strconv.ParseUint(os.Getenv("RAINY_MAX_FILE_SIZE"), 10, 64)

	if err != nil {
		return 0
	}

	return num
}

func (f *FilesystemUploader) StoreFile(fileName string, file []byte) (*string, error) {
	os.WriteFile(fmt.Sprintf("%s/%s", f.path, fileName), file, 0666)

	fp := fmt.Sprintf("%s/%s", os.Getenv("RAINY_FILE_PATH"), fileName)

	return &fp, nil
}

func (f *FilesystemUploader) FileExists(fileName string) bool {
	_, err := os.Stat(fmt.Sprintf("%s/%s", f.path, fileName))
	return os.IsExist(err)
}

func (f *FilesystemUploader) GetFile(fileName string) ([]byte, error) {
	if !f.FileExists(fileName) {
		return []byte{}, errors.New("file does not exist")
	}

	return os.ReadFile(fileName)
}
