package games

import (
	"errors"
	"fmt"
	"strategy-game/boards"
	"strategy-game/boards/moves"
	"strategy-game/pawns"

	"gorm.io/gorm"
)

type Gamer interface {
	Create(*gorm.DB) error
	Read(*gorm.DB) error
	Update(*gorm.DB) error
	Delete(*gorm.DB) error
	HardDelete(*gorm.DB) error
	List(*gorm.DB) ([]Game, error)
	CreateNewGame(*gorm.DB, uint) error
	JoinGame(*gorm.DB, uint) error
	SimulateGame(*gorm.DB) error
	End(*gorm.DB) error
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

func (g *Game) CreateNewGame(db *gorm.DB, userid uint) error {
	var err error

	err = g.Create(db)
	if err != nil {
		return err
	}
	err = g.Read(db)
	if err != nil {
		return err
	}
	board := boards.Board{
		Type:   "flat",
		GameID: g.ID,
	}
	err = board.CreateBoard(db, g.ID)
	if err != nil {
		return err
	}
	err = board.Read(db)
	if err != nil {
		return err
	}
	g.BoardID = board.ID
	err = g.Update(db)
	if err != nil {
		return err
	}

	pawn := pawns.Pawn{
		UserID:  uint(userid),
		GameID:  g.ID,
		BoardID: board.ID,
		Type:    "cavalry",
	}
	err = pawn.InitiatePawn()
	if err != nil {
		return err
	}
	err = pawn.Create(db)
	if err != nil {
		return err
	}
	err = board.DeployPawn(db, pawn.ID, 5, 10)
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
	pawn := pawns.Pawn{
		UserID:  g.User2ID,
		GameID:  g.ID,
		BoardID: g.BoardID,
		Type:    "cavalry",
	}
	err = pawn.InitiatePawn()
	if err != nil {
		return err
	}
	err = pawn.Create(db)
	if err != nil {
		return err
	}
	board := boards.Board{Model: gorm.Model{ID: g.BoardID}}

	err = board.Read(db)
	if err != nil {
		return err
	}
	err = board.DeployPawn(db, pawn.ID, 15, 10)
	if err != nil {
		return err
	}
	return nil

}
func (g *Game) SimulateGame(db *gorm.DB, gameid uint, round uint16) error {
	g.ID = gameid
	err := g.Read(db)
	if err != nil {
		return err
	}

	move := moves.Move{
		Model: gorm.Model{
			ID: gameid,
		},
		Round: round,
	}

	moveList, err := move.List(db)
	if err != nil {
		return err
	}
	var user1Moves []moves.Move
	var user2Moves []moves.Move

	for _, move := range moveList {
		if move.UserID == g.User1ID {
			user1Moves = append(user1Moves, move)
		} else if move.UserID == g.User2ID {
			user2Moves = append(user2Moves, move)
		} else {
			return errors.New("error grouping users")
		}
	}
	var simulCount int
	if len(user1Moves) > len(user2Moves) {
		simulCount = len(user1Moves)
	} else {
		simulCount = len(user2Moves)
	}
	//	GameID    uint
	//	BoardID   uint
	//	PawnID    uint
	//	UserID    uint
	//	X         int16
	//	Y         int16
	//	Direction uint8
	//	Round     uint16
	board := boards.Board{
		Model: gorm.Model{
			ID: g.BoardID,
		},
	}
	err = board.Read(db)
	if err != nil {
		return err
	}

	for i := 0; i < simulCount; i++ {

		//pawnid1 := board.Terrain[Y][X]
	}
	fmt.Println(board)
	return nil
}
