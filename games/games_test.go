package games

import (
	"strategy-game/boards"
	"strategy-game/pawns"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Construct() (*gorm.DB, Game) {
	var db *gorm.DB
	var game = Game{
		Model: gorm.Model{
			//ID:        1,
			UpdatedAt: time.Time{},
			CreatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{Time: time.Time{}, Valid: false},
		},
		User1ID: 1,
		User2ID: 2,
		BoardID: 1,
		Round:   0,
	}
	db, _ = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	db.AutoMigrate(&Game{}, &boards.Board{}, &pawns.Pawn{})
	return db, game
}
func Destruct() {
	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	db.Exec("DROP TABLE games")
	db.Exec("DROP TABLE boards")
	db.Exec("DROP TABLE pawns")
}
