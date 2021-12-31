package pawns

import (
	"testing"
	"time"

	"gorm.io/gorm"
)

func TestCreate(t *testing.T) {
	db, pawn := Construct()

	tests := []struct {
		input Pawn
		//output error
		err error
	}{
		{input: pawn, err: nil},
	}
	for _, test := range tests {
		err := test.input.Create(db)
		if test.err != err {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
	}
	Destruct()
}
func TestRead(t *testing.T) {
	db, pawn := Construct()
	pawn.Create(db)

	tests := []struct {
		//input  uint
		output Pawn
		err    error
	}{
		{
			output: pawn, err: nil},
	}
	for _, test := range tests {
		test.output.ID = 1
		res := test.output
		err := res.Read(db)
		if test.err != err {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
		res.CreatedAt = time.Time{}
		res.UpdatedAt = time.Time{}
		res.DeletedAt = gorm.DeletedAt{Time: time.Time{}, Valid: false}
		test.output.CreatedAt = time.Time{}
		test.output.UpdatedAt = time.Time{}
		test.output.DeletedAt = gorm.DeletedAt{Time: time.Time{}, Valid: false}
		if res != test.output {
			t.Errorf("Result is: %v . Expected: %v", res, test.output)
		}
	}
	Destruct()
}
func TestList(t *testing.T) {
	db, pawn := Construct()
	pawn.Create(db)
	tests := []struct {
		input  Pawn
		output []Pawn
		err    error
	}{
		{
			input:  Pawn{GameID: 1},
			output: []Pawn{pawn}, err: nil},
	}
	for _, test := range tests {
		test.output[0].ID = 1
		res, err := test.input.List(db)
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

func TestUpdate(t *testing.T) {
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
		err := test.input.Update(db)
		if test.err != err {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
	}
	Destruct()
}
func TestDelete(t *testing.T) {
	db, pawn := Construct()
	pawn.Create(db)
	tests := []struct {
		input Pawn
		//output error
		err error
	}{
		{input: Pawn{
			Model: gorm.Model{
				ID: 1,
			},
		}, err: nil},
	}
	for _, test := range tests {
		err := test.input.Delete(db)
		if test.err != err {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
	}
	Destruct()
}
func TestHardDelete(t *testing.T) {
	db, pawn := Construct()
	pawn.Create(db)
	tests := []struct {
		input Pawn
		//output error
		err error
	}{
		{input: Pawn{
			Model: gorm.Model{
				ID: 1,
			},
		}, err: nil},
	}
	for _, test := range tests {
		err := test.input.HardDelete(db)
		if test.err != err {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
	}
	Destruct()
}
