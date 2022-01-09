package moves

import (
	"strategy-game/games"

	"gorm.io/gorm"
)

type IMove interface {
	Create(*gorm.DB) error
	Read(*gorm.DB) error
	Update(*gorm.DB) error
	Delete(*gorm.DB) error
	HardDelete(*gorm.DB) error
	List(*gorm.DB) ([]Move, error)
	AppendMove(*gorm.DB, uint, uint, uint, int16, int16, uint8) error
}

//  message Move{
//	uint32 pawnid=1;
//	uint32 x=2;
//	uint32 y=3;
//	uint32 direction=4;
//  }
//
//  message MoveInputs {
//	uint32 userid=1;
//	uint32 gameid=2;
//	repeated Move move=3;
//  }
//
//  message MoveOutputs {
//	bool OK=1;
//  }

type Move struct {
	gorm.Model
	GameID    uint
	BoardID   uint
	PawnID    uint
	X         int16
	Y         int16
	Direction uint8
	Round     uint16
}

func (m *Move) AppendMove(db *gorm.DB, gameid uint, boardid uint, pawnid uint, x int16, y int16, direction uint8) error {
	game := games.Game{}
	game.ID = gameid
	err := game.Read(db)
	if err != nil {
		return err
	}
	move := Move{
		GameID:    gameid,
		BoardID:   boardid,
		PawnID:    pawnid,
		X:         x,
		Y:         y,
		Direction: direction,
		Round:     game.Round,
	}
	err = move.Create(db)
	if err != nil {
		return err
	}
	return nil
}
