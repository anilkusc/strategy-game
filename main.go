package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	game := App{}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	game.DB = db
	game.Init()
	//game.Start()
}
