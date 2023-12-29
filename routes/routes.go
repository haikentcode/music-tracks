package routes

import (
	"musictracks/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/tracks", handlers.GetTracks)
		v1.PUT("/create-track/:isrc", handlers.CreateTrackByISRC)
		v1.GET("/tracks/:isrc", handlers.GetTrackByISRC)
		v1.GET("/tracks/by-artist/:artist", handlers.GetTracksByArtist)
	}
}
