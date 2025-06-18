package main

import (
	"github.com/gin-gonic/gin"
	"st-ember.github.com/streamingservice/internal/application"
	"st-ember.github.com/streamingservice/internal/delivery"
)

func main() {
	usecase := application.NewVideoUseCase()
	handler := delivery.NewVideoHandler(usecase)

	r := gin.Default()
	handler.RegisterRoutes(r)

	r.Run(":8080")
}
