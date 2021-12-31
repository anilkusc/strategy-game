package users

import (
	"reflect"
	"testing"
	"time"

	"gorm.io/gorm"
)

func TestCreate(t *testing.T) {
	db, user := Construct()

	tests := []struct {
		input User
		//output error
		err error
	}{
		{input: user, err: nil},
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
	db, user := Construct()
	user.Create(db)

	tests := []struct {
		output User
		err    error
	}{
		{output: user, err: nil},
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

		if reflect.DeepEqual(res, test.output) == false {
			t.Errorf("Result is: %v . Expected: %v", res, test.output)
		}
	}
	Destruct()
}

func TestList(t *testing.T) {
	db, user := Construct()
	user.Create(db)

	tests := []struct {
		//input  Game
		output []User
		err    error
	}{
		{
			output: []User{user}, err: nil},
	}
	for _, test := range tests {
		res, err := user.List(db)
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
	db, user := Construct()
	user.Create(db)

	tests := []struct {
		input User
		//output error
		err error
	}{
		{input: user, err: nil},
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
	db, user := Construct()
	user.Create(db)
	tests := []struct {
		input User
		//output error
		err error
	}{
		{input: User{
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
	db, user := Construct()
	user.Create(db)
	tests := []struct {
		input User
		//output error
		err error
	}{
		{input: User{
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
