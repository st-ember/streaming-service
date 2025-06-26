package main

import (
	"github.com/gin-gonic/gin"
	"st-ember.github.com/streamingservice/internal/application"
	"st-ember.github.com/streamingservice/internal/delivery"
	"st-ember.github.com/streamingservice/internal/infra/ffmpeg"
)

func main() {
	transcoder := ffmpeg.NewTranscoder()
	usecase := application.NewVideoUseCase(transcoder)
	handler := delivery.NewVideoHandler(usecase)

	r := gin.Default()
	handler.RegisterRoutes(r)

	r.Run(":8080")
}
