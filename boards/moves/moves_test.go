package moves

import (
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
	db.AutoMigrate(&Move{})
	return db, move
}
func Destruct() {
	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	db.Exec("DROP TABLE moves")
}
func TestAppendMove(t *testing.T) {
	db, move := Construct()

	type ExtraParams struct {
		round      uint16
		boardid    uint
		gamestatus int8
		user1id    uint
		user2id    uint
	}
	tests := []struct {
		input  Move
		input2 ExtraParams
		result int8
		err    error
	}{
		{
			input: move,
			input2: ExtraParams{
				round:      1,
				boardid:    1,
				gamestatus: -1,
				user1id:    1,
				user2id:    2,
			},
			result: -3,
			err:    nil,
		},
	}
	for _, test := range tests {
		in := test.input
		in2 := test.input2
		res, err := move.AppendMove(db, in.GameID, in.BoardID, in.PawnID, in.X, in.Y, in.Direction, in2.round, in2.boardid, in2.gamestatus, in2.user1id, in2.user2id)
		if test.err != err {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
		if test.result != res {
			t.Errorf("Result is: %v . Expected: %v", res, test.result)
		}
	}

	Destruct()
}
