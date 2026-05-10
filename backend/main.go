package main

import (
	_ "embed"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"alphabetbengali/handlers"
)

//go:embed web/index.html
var indexHTML []byte

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", indexHTML)
	})

	api := r.Group("/api")
	{
		api.GET("/characters", handlers.GetCharacters)
		api.GET("/characters/:id", handlers.GetCharacterByID)
		api.GET("/conjuncts", handlers.GetConjuncts)
		api.GET("/conjuncts/:id", handlers.GetConjunctByID)
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
