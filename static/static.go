package static

import (
	"embed"
	"os"
)

//go:embed *
var content embed.FS

func Get(path string) ([]byte, error) {
	if fileExists(path) {
		return os.ReadFile(path)
	}

	return content.ReadFile(path)
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return os.IsExist(err)
}
