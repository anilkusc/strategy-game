package games

import (
	"gorm.io/gorm"
)

type Gamer interface {
	Create(*gorm.DB) error
	Read(*gorm.DB) error
	Update(*gorm.DB) error
	Delete(*gorm.DB) error
	HardDelete(*gorm.DB) error
	List(*gorm.DB) ([]Game, error)
}

type Game struct {
	gorm.Model
	User1ID uint
	User2ID uint
	BoardID uint
	Round   uint16
	// -5: ready for play
	// -4: only user2 sent moves
	// -3: only user1 sent moves
	// -2: 1 User Ready
	// -1: Started
	//  0: Draw
	//  1: User1 is winner
	//  2: User2 is winner
	Status int8
}
