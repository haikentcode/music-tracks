package handlers

import (
	"fmt"
	"musictracks/db"
	"musictracks/models"
	"musictracks/pkg/spotify"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTrackByISRC(c *gin.Context) {
	isrc := c.Param("isrc")

	tracks, err := spotify.GetTrackByISRC(isrc)
	if err != nil {
		// Handle the error appropriately, e.g., return an HTTP 500 response
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	var topTrack spotify.Item
	for _, track := range tracks.Tracks.Items {
		if topTrack.Popularity < track.Popularity {
			topTrack = track
		}
	}

	// Create Track and Artist instances from the extracted data
	track := models.Track{
		ISRC:            topTrack.ExternalIDs.ISRC,
		SpotifyImageURI: topTrack.Album.Images[0].URL, // Assuming the first image is the main one
		Title:           topTrack.Name,
		Popularity:      topTrack.Popularity,
	}

	var artists []models.Artist
	for _, artist := range topTrack.Artists {
		artists = append(artists, models.Artist{Name: artist.Name})
	}

	db := db.GetDB()
	tx := db.Begin()

	// Defer a function to handle the transaction result
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
	}()

	// Try to find an existing track with the same ISRC in the database
	var existingTrackInDB models.Track
	if tx.Where("isrc = ?", track.ISRC).First(&existingTrackInDB).Error == nil {
		// If the track exists, update its data
		existingTrackInDB.SpotifyImageURI = track.SpotifyImageURI
		existingTrackInDB.Title = track.Title
		existingTrackInDB.Popularity = track.Popularity
		tx.Save(&existingTrackInDB)

		// Update or create associated artists
		for _, artist := range artists {
			var existingArtist models.Artist
			if tx.Where("name = ?", artist.Name).First(&existingArtist).Error == nil {
				// If the artist exists, update its data
				existingArtist.Name = artist.Name
				tx.Save(&existingArtist)
			} else {
				// If the artist doesn't exist, create it
				tx.Create(&artist)
			}

			// Associate the updated or newly created artist with the track
			tx.Model(&existingTrackInDB).Association("Artists").Append(&existingArtist)
		}

		// Commit the transaction
		tx.Commit()

		// Return the updated track in the response
		c.JSON(http.StatusOK, gin.H{"top": existingTrackInDB, "all_traks": tracks})
		return
	}

	// Create the track if it doesn't exist
	tx.Create(&track)
	for _, artist := range artists {
		tx.FirstOrCreate(&artist, models.Artist{Name: artist.Name})
		tx.Model(&track).Association("Artists").Append(&artist)
	}

	// Commit the transaction
	tx.Commit()

	// Return the track in the response
	c.JSON(http.StatusOK, gin.H{"top": track, "all_traks": tracks})
}
