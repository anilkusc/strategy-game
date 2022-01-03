package games

import (
	"strategy-game/boards"

	"gorm.io/gorm"
)

type Gamer interface {
	Create(*gorm.DB) error
	Read(*gorm.DB) error
	Update(*gorm.DB) error
	Delete(*gorm.DB) error
	HardDelete(*gorm.DB) error
	List(*gorm.DB) ([]Game, error)
	CreateNewGame(*gorm.DB) error
	JoinGame(*gorm.DB, uint) error
	End(*gorm.DB) error
}

type Game struct {
	gorm.Model
	User1ID uint
	User2ID uint
	BoardID uint
	Status  int8 // -2: 1 User Ready,-1: Started,0: Draw , 1: User1 is winner , 2: User2 is winner
}

func (g *Game) CreateNewGame(db *gorm.DB) error {
	var err error
	board := boards.Board{
		Type:   "flat",
		GameID: g.ID,
	}
	err = board.CreateBoard(db)
	if err != nil {
		return err
	}

	g.BoardID = board.ID
	err = g.Create(db)
	if err != nil {
		return err
	}
	err = g.Read(db)
	if err != nil {
		return err
	}
	return nil
}

func (g *Game) JoinGame(db *gorm.DB, user2id uint) error {
	g.Read(db)
	g.Status = -1
	g.User2ID = user2id
	err := g.Update(db)
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
