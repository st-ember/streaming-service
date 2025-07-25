package diskstorage

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

type DiskStorage struct {
	BasePath string
}

func NewDiskStorage(basePath string) *DiskStorage {
	return &DiskStorage{
		BasePath: basePath,
	}
}

func (d *DiskStorage) Store(file io.Reader, extension string) (string, error) {
	storedName := uuid.NewString()
	fullPath := filepath.Join(d.BasePath, storedName, extension)
	out, err := os.Create(fullPath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %w", err)
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		return "", fmt.Errorf("failed to write file: %w", err)
	}

	return storedName, nil
}
