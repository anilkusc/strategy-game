package moves

import (
	"gorm.io/gorm"
)

type IMove interface {
	Create(*gorm.DB) error
	Read(*gorm.DB) error
	Update(*gorm.DB) error
	Delete(*gorm.DB) error
	HardDelete(*gorm.DB) error
	List(*gorm.DB) ([]Move, error)
}

type Move struct {
	gorm.Model
	GameID  uint
	BoardID uint
	PawnID  uint
	OldX    int16
	OldY    int16
	NewX    int16
	NewY    int16
}
