package delivery

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"st-ember.github.com/streamingservice/internal/application"
)

type VideoHandler struct {
	usecase *application.VideoUseCase
}

func NewVideoHandler(u *application.VideoUseCase) *VideoHandler {
	return &VideoHandler{usecase: u}
}

func (h *VideoHandler) RegisterRoutes(r *gin.Engine) {
	r.GET("/videos/:id/manifest.mpd", h.ServeManifest)
	r.GET("/videos/:id/segment_:num.m4s", h.ServeSegment)
	r.POST("/videos/create", h.HandleUpload)
}

func (h *VideoHandler) ServeManifest(c *gin.Context) {
	id := c.Param("id")
	path, err := h.usecase.GetManifestPath(id)
	if err != nil {
		c.Status(404)
		return
	}
	c.File(path)
}

func (h *VideoHandler) ServeSegment(c *gin.Context) {
	id := c.Param("id")
	num := c.Param("num")
	path, err := h.usecase.GetSegmentPath(id, num)
	if err != nil {
		c.Status(404)
		return
	}
	c.File(path)
}

func (h *VideoHandler) HandleUpload(c *gin.Context) {
	if _, err := c.Request.MultipartReader(); err != nil {
		http.Error(c.Writer, "request must be multipart/form-data", http.StatusUnsupportedMediaType)
		return
	}

	userId, err := uuid.Parse(c.Request.FormValue("user_id"))
	if err != nil {
		http.Error(c.Writer, "cannot parse user id", http.StatusInternalServerError)
		return
	}

	title := c.Request.FormValue("title")
	description := c.Request.FormValue("description")
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		http.Error(c.Writer, "cannot read file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	input := application.UploadInput{
		UserId:      userId,
		Title:       title,
		Description: description,
		File:        file,
		FileName:    header.Filename,
		Extension:   filepath.Ext(header.Filename),
	}

	h.usecase.CreateResource(input)

	c.Writer.WriteHeader(http.StatusOK)
}
