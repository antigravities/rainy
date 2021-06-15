package static

import (
	"embed"
	"os"
)

//go:embed *
var content embed.FS

//go:embed html/modern/static
var ModernContent embed.FS

func Get(path string) ([]byte, error) {
	if fileExists(path) {
		return os.ReadFile(path)
	}

	return content.ReadFile(path)
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil // do not !os.IsNotExist(): https://stackoverflow.com/a/12518877
}
