package moves

import (
	"errors"

	"gorm.io/gorm"
)

type Mover interface {
	Create(*gorm.DB) error
	Read(*gorm.DB) error
	Update(*gorm.DB) error
	Delete(*gorm.DB) error
	HardDelete(*gorm.DB) error
	List(*gorm.DB) ([]Move, error)
	AppendMove(*gorm.DB, uint, uint, uint, int16, int16, uint8) (error, int8)
}

type Move struct {
	gorm.Model
	GameID    uint
	BoardID   uint
	PawnID    uint
	UserID    uint
	X         int16
	Y         int16
	Direction uint8
	Round     uint16
}

func (m *Move) AppendMove(db *gorm.DB, gameid uint, userid uint, pawnid uint, x int16, y int16, direction uint8, round uint16, boardid uint, gamestatus int8, user1id uint, user2id uint) (int8, error) {

	move := Move{
		GameID:    gameid,
		BoardID:   boardid,
		PawnID:    pawnid,
		UserID:    userid,
		X:         x,
		Y:         y,
		Direction: direction,
		Round:     round,
	}

	if gamestatus == -3 && userid == user1id || gamestatus == -2 && userid == user2id {
		return -100, errors.New("user has already sent the moves")
	}
	err := move.Create(db)
	if err != nil {
		return -100, err
	}

	switch gamestatus {
	case -3, -4:
		gamestatus = -5
	case -1:
		if userid == user2id {
			gamestatus = -4
		} else if userid == user1id {
			gamestatus = -3
		}
	}

	return gamestatus, nil
}
