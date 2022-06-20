package infra

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Infra interface {
	Connect() *gorm.DB
}

type infra struct {
	db *gorm.DB
}

func (i *infra) Connect() *gorm.DB {
	return i.db
}

func NewInfra() Infra {
	dsn := "host=localhost user=postgres password=stauffenberg dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &infra{
		db: conn,
	}
}
