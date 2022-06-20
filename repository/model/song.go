package model

type Song struct {
	Id     string `gorm:"primaryKey"`
	Artist string
	Album  string
	Title  string
}

func NewSong(id, artist, album, title string) Song {
	return Song{
		Id:     id,
		Artist: artist,
		Album:  album,
		Title:  title,
	}
}
