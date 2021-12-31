package boards

import (
	"reflect"
	"testing"
	"time"

	"gorm.io/gorm"
)

func TestCreate(t *testing.T) {
	db, board := Construct()

	tests := []struct {
		input Board
		//output error
		err error
	}{
		{input: board, err: nil},
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
	db, board := Construct()
	board.Create(db)

	tests := []struct {
		output Board
		err    error
	}{
		{output: board, err: nil},
	}
	for _, test := range tests {
		test.output.ID = 1
		res := test.output
		test.output.FeaturedMap, _ = test.output.JsonToArray(test.output.FeaturedMapJson)
		test.output.Terrain, _ = test.output.JsonToArray(test.output.TerrainJson)

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

		if reflect.DeepEqual(res, test.output) == false {
			t.Errorf("Result is: %v . Expected: %v", res, test.output)
		}
	}
	Destruct()
}

func TestList(t *testing.T) {
	db, board := Construct()
	board.Create(db)

	tests := []struct {
		input  Board
		output []Board
		err    error
	}{
		{
			input:  Board{GameID: 1},
			output: []Board{board}, err: nil},
	}
	for _, test := range tests {
		test.output[0].ID = 1
		test.output[0].FeaturedMap, _ = test.output[0].JsonToArray(test.output[0].FeaturedMapJson)
		test.output[0].Terrain, _ = test.output[0].JsonToArray(test.output[0].TerrainJson)
		res, err := board.List(db)
		if test.err != err {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
		for i := range res {
			res[i].CreatedAt = time.Time{}
			res[i].UpdatedAt = time.Time{}
			res[i].DeletedAt = gorm.DeletedAt{Time: time.Time{}, Valid: false}
			test.output[i].CreatedAt = time.Time{}
			test.output[i].UpdatedAt = time.Time{}
			test.output[i].DeletedAt = gorm.DeletedAt{Time: time.Time{}, Valid: false}

			if reflect.DeepEqual(res[i], test.output[i]) == false {
				t.Errorf("Result is: %v . Expected: %v", res[i], test.output[i])
				t.Errorf("Result list is: %v . Expected list: %v", res, test.output)
			}
		}

	}
	Destruct()
}

func TestUpdate(t *testing.T) {
	db, board := Construct()
	board.Create(db)

	tests := []struct {
		input Board
		//output error
		err error
	}{
		{input: board, err: nil},
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
	db, board := Construct()
	board.Create(db)
	tests := []struct {
		input Board
		//output error
		err error
	}{
		{input: Board{
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
	db, board := Construct()
	board.Create(db)
	tests := []struct {
		input Board
		//output error
		err error
	}{
		{input: Board{
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
