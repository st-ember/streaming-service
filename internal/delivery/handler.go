package delivery

import (
	"github.com/gin-gonic/gin"
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
