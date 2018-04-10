package models

// Video struct
type Video struct {
	Model
	Name        string `gorm:"type:varchar(100)"`
	Description string `gorm:"type:text"`
	ChannelID   int
	Channel     Channel
	Playlists   []Playlist `gorm:"many2many:playlist_video"`
	Supplies    []Supply   `gorm:"many2many:supply_video"`
}
