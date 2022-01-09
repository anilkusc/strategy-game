package moves

import (
	"strategy-game/games"
	"testing"
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
		GameID:    1,
		BoardID:   1,
		PawnID:    1,
		X:         1,
		Y:         0,
		Direction: 1,
	}
	db, _ = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	db.AutoMigrate(&Move{}, &games.Game{})
	return db, move
}
func Destruct() {
	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	db.Exec("DROP TABLE moves")
}
func TestAppendMove(t *testing.T) {
	db, move := Construct()
	game := games.Game{}
	game.Create(db)
	tests := []struct {
		input Move
		err   error
	}{
		{
			input: move,
			err:   nil,
		},
	}
	for _, test := range tests {
		err := move.AppendMove(db, move.GameID, move.BoardID, move.PawnID, move.X, move.Y, move.Direction)
		if test.err != err {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}

	}

	Destruct()
}
