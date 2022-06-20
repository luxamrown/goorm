package main

import (
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID    string `gorm:"primaryKey"`
	Name  string
	Email string
	Age   int
}

func NewUser(id, name, email string, age int) User {
	return User{
		ID:    id,
		Name:  name,
		Email: email,
		Age:   age,
	}
}

func main() {
	luxam := NewUser(uuid.New().String(), "luxam", "luxam@gmail.com", 18)
	dsn := "host=localhost user=postgres password=stauffenberg dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{})
	db.Create(luxam)
}
