package repository

import (
	"fmt"

	"gorm.io/gorm"
	"mohamadelabror.me/gorm/repository/model"
)

type SongRepo interface {
	CreateSchema() error
	CreateSong(song model.Song) error
	GetFirstPrimaryKey() model.Song
	GetFirstNoOrder() model.Song
	GetLastPrimaryKeyDesc() model.Song
	GetAll() []model.Song
	GetSongByTitle(title string) (model.Song, error)
	GetSongByAlbum(album string) ([]model.Song, error)
	GetSongByMultipleAlbum(albums []string) ([]model.Song, error)
	SearchSongByTitle(title string) ([]model.Song, error)
}

type songRepoImpl struct {
	songDb *gorm.DB
}

func (s *songRepoImpl) CreateSchema() error {
	err := s.songDb.AutoMigrate(&model.Song{})
	if err != nil {
		panic(err)
	}
	return nil
}

func (s *songRepoImpl) CreateSong(song model.Song) error {
	res := s.songDb.Create(song)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// Get the first record ordered by primary key
func (s *songRepoImpl) GetFirstPrimaryKey() model.Song {
	var song model.Song
	result := s.songDb.First(&song)
	if result.Error != nil {
		panic(result.Error)
	}
	return song
}

// Get one Record, no specifer order
func (s *songRepoImpl) GetFirstNoOrder() model.Song {
	var song model.Song
	result := s.songDb.Take(&song)
	if result.Error != nil {
		panic(result.Error)
	}
	return song
}

// Get last record, ordered by primary key desc
func (s *songRepoImpl) GetLastPrimaryKeyDesc() model.Song {
	var song model.Song
	result := s.songDb.Last(&song)
	if result.Error != nil {
		panic(result.Error)
	}
	return song
}

// Get all records
func (s *songRepoImpl) GetAll() []model.Song {
	var songs []model.Song
	result := s.songDb.Find(&songs)
	if result.Error != nil {
		panic(result.Error)
	}
	return songs
}

// Get first matcher record
func (s *songRepoImpl) GetSongByTitle(title string) (model.Song, error) {
	var song model.Song
	result := s.songDb.Where("title = ?", title).First(&song)
	if result.Error != nil {
		return model.Song{}, result.Error
	}
	return song, nil
}

// Get all matched record
func (s *songRepoImpl) GetSongByAlbum(album string) ([]model.Song, error) {
	var songs []model.Song
	result := s.songDb.Where("album = ?", album).Find(&songs)
	if result.Error != nil {
		return []model.Song{}, result.Error
	}
	return songs, nil
}

// IN
func (s *songRepoImpl) GetSongByMultipleAlbum(albums []string) ([]model.Song, error) {
	var songs []model.Song
	result := s.songDb.Where("album IN ?", albums).Find(&songs)
	if result.Error != nil {
		return []model.Song{}, result.Error
	}

	return songs, nil
}

// Like
func (s *songRepoImpl) SearchSongByTitle(title string) ([]model.Song, error) {
	var songs []model.Song
	selectedTitle := fmt.Sprintf("%%%s%%", title)
	result := s.songDb.Where("title LIKE ?", selectedTitle).Find(&songs)
	if result.Error != nil {
		return []model.Song{}, result.Error
	}
	return songs, nil
}

func NewSongRepo(songDb *gorm.DB) SongRepo {
	return &songRepoImpl{
		songDb: songDb,
	}
}
