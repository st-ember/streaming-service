package application

import (
	"io"

	"github.com/google/uuid"
)

type UploadInput struct {
	UserId      uuid.UUID
	Title       string
	Description string
	FileName    string
	File        io.Reader
	Extension   string
}
