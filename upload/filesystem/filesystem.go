package filesystem

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"get.cutie.cafe/rainy/conf"
	"get.cutie.cafe/rainy/upload"
)

type FilesystemUploader struct {
	path string
}

func New(path string) (upload.Uploader, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, 0755)

		if err != nil {
			return nil, err
		}
	}

	fx := &FilesystemUploader{
		path: path,
	}

	if !fx.FileExists("index.html") {
		fx.StoreFile("index.html", []byte{})
	}

	return fx, nil
}

func (f *FilesystemUploader) MaxFileSize() uint64 {
	return conf.GetUInt64("MAX_FILE_SIZE")
}

func (f *FilesystemUploader) StoreFile(fileName string, file []byte) (*string, error) {
	// TODO: make perms configurable
	os.WriteFile(fmt.Sprintf("%s/%s", f.path, fileName), file, 0666)

	fp := fmt.Sprintf("%s/%s", conf.GetString("PUBLIC_FILE_PATH"), fileName)

	return &fp, nil
}

func (f *FilesystemUploader) FileExists(fileName string) bool {
	_, err := os.Stat(fmt.Sprintf("%s/%s", f.path, fileName))
	return err == nil // do not !os.IsNotExist(): https://stackoverflow.com/a/12518877
}

func (f *FilesystemUploader) GetFile(fileName string) ([]byte, error) {
	if !f.FileExists(fileName) {
		return []byte{}, errors.New("file does not exist")
	}

	// prevent directory traversal attacks
	fileName = strings.ReplaceAll(fileName, "..", ".")
	fileName = strings.ReplaceAll(fileName, "/", "")

	return os.ReadFile(fmt.Sprintf("%s/%s", f.path, fileName))
}

func (f *FilesystemUploader) StoreFileStream(fileName string, stream io.ReadCloser) (*string, error) {
	file, err := os.OpenFile(fmt.Sprintf("%s/%s", f.path, fileName), os.O_WRONLY|os.O_CREATE, 0660)
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(file, stream)
	if err != nil {
		return nil, err
	}

	fp := fmt.Sprintf("%s/%s", conf.GetString("PUBLIC_UPLOAD_URL"), fileName)

	return &fp, nil
}
