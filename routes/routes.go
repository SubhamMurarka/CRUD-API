package routes

import (
	"github.com/SubhamMurarka/crud/handler"
	"github.com/gin-gonic/gin"
)

// SetupRoutes sets up the API routes and attaches handler functions
func SetupRoutes(r *gin.Engine) {
	r.GET("/movies", handler.GetMovies)
	r.GET("/movies/:id", handler.GetMovie)
	r.POST("/movies", handler.CreateMovie)
	r.PUT("/movies/:id", handler.UpdateMovie)
	r.DELETE("/movies/:id", handler.DeleteMovie)
}
