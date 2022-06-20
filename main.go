package main

import (
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

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

func main() {
	basketCase := NewSong(uuid.New().String(), "Green Day", "Dookie", "Basket Case")
	dsn := "host=localhost user=postgres password=stauffenberg dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Song{})
	db.Create(basketCase)
}
