package moves

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Construct() (*gorm.DB, Move) {
	var db *gorm.DB
	var move = Move{
		Model: gorm.Model{
			//ID:        1,
			UpdatedAt: time.Time{},
			CreatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{Time: time.Time{}, Valid: false},
		},
		GameID:  1,
		BoardID: 1,
		PawnID:  1,
		OldX:    1,
		OldY:    1,
		NewX:    2,
		NewY:    2,
	}
	db, _ = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	db.AutoMigrate(&Move{})
	return db, move
}
func Destruct() {
	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	db.Exec("DROP TABLE moves")
}
