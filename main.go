package main

import (
	"fmt"

	"educabot.com/bookshop/handlers"
	"educabot.com/bookshop/providers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.SetTrustedProxies(nil)

	booksProvider := providers.NewApiBooksProvider("https://6781684b85151f714b0aa5db.mockapi.io/api/v1/books")
	metricsHandler := handlers.NewGetMetrics(booksProvider)
	router.GET("/", metricsHandler.Handle())
	router.Run(":3000")
	fmt.Println("Starting server on :3000")
}
