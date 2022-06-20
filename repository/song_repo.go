package repository

import (
	"gorm.io/gorm"
	"mohamadelabror.me/gorm/repository/model"
)

type SongRepo interface {
	CreateSchema() error
	CreateSong(song model.Song) error
	GetFirstPrimaryKey() model.Song
	GetFirstNoOrder() model.Song
	GetLastPrimaryKeyDesc() model.Song
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

func NewSongRepo(songDb *gorm.DB) SongRepo {
	return &songRepoImpl{
		songDb: songDb,
	}
}
