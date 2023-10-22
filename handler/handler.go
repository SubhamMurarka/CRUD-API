package handler

import (
	"net/http"

	"github.com/SubhamMurarka/crud/database"
	"github.com/SubhamMurarka/crud/models"
	"github.com/gin-gonic/gin"
)

// GetMovies handles the GET request to fetch all movies
func GetMovies(c *gin.Context) {
	// Use function from the database package to get a list of movies from the database
	movies, err := database.GetMoviesFromDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch movies"})
		return
	}
	c.JSON(http.StatusOK, movies)
}

// GetMovie handles the GET request to fetch a single movie by ID
func GetMovie(c *gin.Context) {
	movieID := c.Param("id")
	// Use function from the database package to get a specific movie by ID
	movie, err := models.GetMovieByIDFromDB(movieID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}
	c.JSON(http.StatusOK, movie)
}

// CreateMovie handles the POST request to create a new movie
func CreateMovie(c *gin.Context) {
	var newMovie models.Movie
	if err := c.ShouldBindJSON(&newMovie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}
	// Use functions from the database package to create a new movie in the database
	if err := database.CreateMovieInDB(newMovie); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create movie"})
		return
	}
	c.JSON(http.StatusCreated, newMovie)
}

// UpdateMovie handles the PUT request to update an existing movie
func UpdateMovie(c *gin.Context) {
	movieID := c.Param("id")
	var updatedMovie models.Movie
	if err := c.ShouldBindJSON(&updatedMovie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}
	// Use function from the database package to update the movie in the database
	if err := database.UpdateMovieInDB(movieID, updatedMovie); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update movie"})
		return
	}
	c.JSON(http.StatusOK, updatedMovie)
}

// DeleteMovie handles the DELETE request to delete a movie by ID
func DeleteMovie(c *gin.Context) {
	movieID := c.Param("id")
	// Use function from the database package to delete the movie from the database
	if err := database.DeleteMovieFromDB(movieID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete movie"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Movie deleted successfully"})
}
