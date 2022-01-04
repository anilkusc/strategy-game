package games

import (
	"strategy-game/boards"
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
	db.AutoMigrate(&Game{}, &boards.Board{})
	return db, game
}
func Destruct() {
	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	db.Exec("DROP TABLE games")
	db.Exec("DROP TABLE boards")
}

func TestCreateNewGame(t *testing.T) {
	db, game := Construct()
	game.ID = 1
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
		err := game.CreateNewGame(db, game.ID)
		if test.err != err {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
	}
	Destruct()
}
func TestJoinGame(t *testing.T) {
	db, game := Construct()
	game.ID = 1
	game.CreateNewGame(db, game.ID)

	tests := []struct {
		input uint
		err   error
	}{
		{
			input: 2,
			err:   nil,
		},
	}
	for _, test := range tests {
		err := game.JoinGame(db, test.input)
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
