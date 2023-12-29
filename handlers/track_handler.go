package handlers

import (
	"musictracks/db"
	"musictracks/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetTodos retrieves all todos
func GetTracks(c *gin.Context) {
	var tracks []models.Track
	db.GetDB().Preload("Artists").Find(&tracks)
	c.JSON(http.StatusOK, tracks)
}

func GetTrackByISRC(c *gin.Context) {
	isrc := c.Param("isrc")

	db := db.GetDB()
	existingTrack := models.Track{}
	if db.Where("isrc = ?", isrc).Preload("Artists").First(&existingTrack).Error == nil {

		c.JSON(http.StatusOK, existingTrack)
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Track not found"})
}

func GetTracksByArtist(c *gin.Context) {
	artistName := c.Param("artist")

	// Perform a "like" search in the DB for tracks associated with the given artist
	var tracks []models.Track
	db := db.GetDB() // Adjust this based on how you manage your database connections
	db.Preload("Artists", "name LIKE ?", "%"+artistName+"%").Find(&tracks)

	// Return the list/array of tracks associated with the artist
	c.JSON(http.StatusOK, tracks)
}
