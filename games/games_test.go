package games

import (
	"testing"
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
	}
	db, _ = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	db.AutoMigrate(&Game{})
	return db, game
}
func Destruct() {
	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	db.Exec("DROP TABLE games")
}

func TestStart(t *testing.T) {
	db, game := Construct()

	tests := []struct {
		input Game
		err   error
	}{
		{
			input: game,
			err:   nil,
		},
	}
	for _, test := range tests {
		err := game.Start(db)
		if test.err != err {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
	}
	Destruct()
}
func TestEnd(t *testing.T) {
	db, game := Construct()
	game.Create(db)

	tests := []struct {
		input Game
		err   error
	}{
		{
			input: game,
			err:   nil,
		},
	}
	for _, test := range tests {
		err := game.End(db, 1)
		if test.err != err {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
	}
	Destruct()
}
