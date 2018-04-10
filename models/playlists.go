package models

// Playlist struct
type Playlist struct {
	Model
	Name        string `gorm:"type:varchar(50)"`
	Description string `gorm:"type:text"`
	ChannelID   int
	Channel     Channel
	Videos      []Video `gorm:"many2many:playlist_video"`
}
