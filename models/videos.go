package models

import "github.com/jinzhu/gorm"

// Video struct
type Video struct {
	gorm.Model
	Name        string `gorm:"type:varchar(100)"`
	Description string `gorm:"type:text"`
	ChannelID   int
	Channel     Channel
	Playlists   []Playlist `gorm:"many2many:playlist_video"`
	Supplies    []Supply   `gorm:"many2many:supply_video"`
}
