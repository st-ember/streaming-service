package delivery

import "github.com/google/uuid"

type UploadReq struct {
	UserId      uuid.UUID
	Title       string
	Description string
	FileName    string
}
