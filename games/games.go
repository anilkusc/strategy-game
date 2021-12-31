package games

import (
	"gorm.io/gorm"
)

type IGame interface {
	Create(*gorm.DB) error
	Read(*gorm.DB) error
	Update(*gorm.DB) error
	Delete(*gorm.DB) error
	HardDelete(*gorm.DB) error
	List(*gorm.DB) ([]Game, error)
	Start(*gorm.DB) error
	End(*gorm.DB) error
}

type Game struct {
	gorm.Model
	User1ID uint
	User2ID uint
	BoardID uint
	Status  int8 // -1: Not started,0: Draw , 1: User1 is winner , 2: User2 is winner
}

func (g *Game) Start(db *gorm.DB) error {
	err := g.Create(db)
	if err != nil {
		return err
	}
	return nil

}
func (g *Game) End(db *gorm.DB, winner int8) error {
	g.Status = winner
	err := g.Update(db)
	if err != nil {
		return err
	}
	return nil

}
