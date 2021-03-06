package sys

import (
	"os"
	"path/filepath"
)

func AutoCreateSubDirs(path string) error {
	filePath := filepath.FromSlash(path)
	filePath = filepath.Clean(filePath)

	if stat, err := os.Stat(filePath); os.IsExist(err) && !stat.IsDir() {
		return nil
	}

	err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm)

	return err
}
