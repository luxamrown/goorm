package main

import (
	"fmt"

	"mohamadelabror.me/gorm/repository"
	"mohamadelabror.me/gorm/repository/infra"
)

func main() {
	songInfra := infra.NewInfra()
	songRepo := repository.NewSongRepo(songInfra.Connect())

	// Retrieving a single object
	fmt.Println(songRepo.GetFirstNoOrder())
	fmt.Println(songRepo.GetFirstPrimaryKey())
	fmt.Println(songRepo.GetLastPrimaryKeyDesc())
}
