package models

type Track struct {
	ID              uint     `gorm:"primaryKey"`
	ISRC            string   `gorm:"unique"`
	SpotifyImageURI string   `gorm:"not null"`
	Title           string   `gorm:"not null"`
	Popularity      int      `gorm:"not null"`
	Artists         []Artist `gorm:"many2many:track_artists"`
}

type Artist struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique;not null"`
	// Tracks []Track `gorm:"many2many:track_artists"`
}
