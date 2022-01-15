package pawns

import (
	"strategy-game/boards"
	"testing"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Construct() (*gorm.DB, Pawn) {
	var pawn = Pawn{
		Model: gorm.Model{
			//ID:        1,
			UpdatedAt: time.Time{},
			CreatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{Time: time.Time{}, Valid: false},
		},
		UserID:    1,
		BoardID:   1,
		GameID:    1,
		X:         1,
		Y:         1,
		Direction: 1,
		Health:    100,
		Defense:   50,
		Attack:    60,
		Speed:     10,
		Affect:    1,
		Agility:   1,
		Range:     1,
		Type:      "cavalry",
	}
	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	db.AutoMigrate(&Pawn{}, &boards.Board{})
	return db, pawn
}
func Destruct() {
	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	db.Exec("DROP TABLE pawns")
	db.Exec("DROP TABLE boards")
}
func TestAttackTo(t *testing.T) {
	db, pawn := Construct()
	pawn.Create(db)

	tests := []struct {
		input Pawn
		//output error
		err error
	}{
		{input: pawn, err: nil},
	}

	for _, test := range tests {
		defenderPawn := test.input
		err := test.input.AttackTo(db, defenderPawn.ID)
		if test.err != err {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
	}
	Destruct()
}
func TestInitiatePawn(t *testing.T) {
	_, pawn := Construct()

	tests := []struct {
		input  Pawn
		output Pawn
		err    error
	}{
		{
			input: pawn,
			output: Pawn{
				Model: gorm.Model{
					UpdatedAt: time.Time{},
					CreatedAt: time.Time{},
					DeletedAt: gorm.DeletedAt{Time: time.Time{}, Valid: false},
				},
				UserID:    1,
				BoardID:   1,
				GameID:    1,
				X:         1,
				Y:         1,
				Direction: 1,
				Health:    100,
				Defense:   50,
				Attack:    60,
				Speed:     6,
				Affect:    1,
				Range:     1,
				Agility:   1,
				Type:      "cavalry",
			}, err: nil},
	}

	for _, test := range tests {
		err := test.input.InitiatePawn(test.input.Direction)
		if test.err != err {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
		if test.output != test.input {
			t.Errorf("Result is: %v . Expected: %v", test.input, test.output)
		}
	}
	Destruct()
}
func TestShufflePawns(t *testing.T) {
	db, pawn := Construct()
	pawn.Create(db)
	tests := []struct {
		input  []uint
		output []Pawn
		err    error
	}{
		{
			input:  []uint{1},
			output: []Pawn{pawn},
			err:    nil},
	}

	for _, test := range tests {
		res, err := pawn.ShufflePawns(db, test.input)
		if test.err != err {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
		for i, element := range res {
			element.CreatedAt = time.Time{}
			element.UpdatedAt = time.Time{}
			element.DeletedAt = gorm.DeletedAt{Time: time.Time{}, Valid: false}
			test.output[i].CreatedAt = time.Time{}
			test.output[i].UpdatedAt = time.Time{}
			test.output[i].DeletedAt = gorm.DeletedAt{Time: time.Time{}, Valid: false}
			if element != test.output[i] {
				t.Errorf("Result is: %v . Expected: %v", element, test.output[i])
				t.Errorf("Result list is: %v . Expected list: %v", res, test.output)
			}
		}
	}
	Destruct()
}

func TestIsPawnMoveValid(t *testing.T) {
	_, pawn := Construct()
	tests := []struct {
		input     Pawn
		Direction uint8
		X         int16
		Y         int16
		output    bool
	}{
		{
			input:     pawn,
			Direction: 1,
			X:         1,
			Y:         0,
			output:    true},
	}

	for _, test := range tests {
		res := pawn.IsPawnMoveValid(test.Direction, test.X, test.Y)
		if res != test.output {
			t.Errorf("Result is: %v . Expected: %v", res, test.output)
		}

	}
	Destruct()
}
func TestCanPawnAttackTo(t *testing.T) {
	_, pawn := Construct()
	pawn2 := pawn
	pawn2.X = 2
	pawn2.Y = 1

	tests := []struct {
		input  Pawn
		output bool
	}{
		{
			input:  pawn2,
			output: true},
	}

	for _, test := range tests {
		res := pawn.CanPawnAttackTo(&test.input)
		if res != test.output {
			t.Errorf("Result is: %v . Expected: %v", res, test.output)
		}

	}
	Destruct()
}
func TestChangeDirection(t *testing.T) {
	_, pawn := Construct()

	tests := []struct {
		input uint8
		err   error
	}{
		{
			input: 1,
			err:   nil},
	}

	for _, test := range tests {
		err := pawn.ChangeDirection(test.input)
		if err != test.err {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}

	}
	Destruct()
}
