package main

import (
	"fmt"

	"mohamadelabror.me/gorm/repository"
	"mohamadelabror.me/gorm/repository/infra"
)

func main() {
	songInfra := infra.NewInfra()
	songRepo := repository.NewSongRepo(songInfra.Connect())

	fmt.Println("-------------------------------")
	fmt.Println("Retrieving a single object")
	// // Retrieving a single object
	fmt.Println(songRepo.GetFirstNoOrder())
	fmt.Println(songRepo.GetFirstPrimaryKey())
	fmt.Println(songRepo.GetLastPrimaryKeyDesc())

	fmt.Println("-------------------------------")
	fmt.Println("Retrieving multiple object")
	// // Retrieving multiple object
	fmt.Println(songRepo.GetAll())

	fmt.Println("-------------------------------")
	fmt.Println("Retrieving conditional object")
	// //  First match record
	fmt.Println(songRepo.GetSongByTitle("Paranoid Android"))
	// // All matched record
	fmt.Println(songRepo.GetSongByAlbum("American Idiot"))

	// IN
	albums := []string{"American Idiot", "OK Computer"}
	fmt.Println(songRepo.GetSongByMultipleAlbum(albums))

	//Like
	fmt.Println(songRepo.SearchSongByTitle("Andro"))
}
