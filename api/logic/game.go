package logic

import (
	"strategy-game/api/protos"
	"strategy-game/boards"
	"strategy-game/boards/moves"
	"strategy-game/games"
	"strategy-game/pawns"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

//TODO : add logs here for != nil expressions.
func CreateNewGame(db *gorm.DB, userid uint) (uint, error) {
	var err error
	game := games.Game{
		User1ID: userid,
		Round:   0,
		Status:  -2,
	}

	err = game.Create(db)
	if err != nil {
		return 0, err
	}

	err = game.Read(db)
	if err != nil {
		return 0, err
	}

	board := boards.Board{
		Type:   "flat",
		GameID: game.ID,
	}
	err = board.CreateBoard(db)
	if err != nil {
		return 0, err
	}

	err = board.Read(db)
	if err != nil {
		return 0, err
	}

	pawn := pawns.Pawn{
		UserID:  uint(userid),
		GameID:  game.ID,
		BoardID: board.ID,
		Type:    "cavalry",
	}
	err = pawn.InitiatePawn()
	if err != nil {
		return 0, err
	}
	err = pawn.Create(db)
	if err != nil {
		return 0, err
	}

	game.BoardID = board.ID
	err = game.Update(db)
	if err != nil {
		return 0, err
	}

	err = board.DeployPawn(db, pawn.ID, 5, 10)
	if err != nil {
		return 0, err
	}
	return game.ID, nil
}
func JoinAGame(db *gorm.DB, user2id uint, gameid uint) (uint, string, string, error) {
	game := games.Game{
		Model: gorm.Model{
			ID: gameid,
		},
	}
	err := game.Read(db)
	if err != nil {
		return 0, "", "", err
	}
	game.Status = -1
	game.User2ID = user2id

	pawn := pawns.Pawn{
		UserID:  user2id,
		GameID:  gameid,
		BoardID: game.BoardID,
		Type:    "cavalry",
	}
	err = pawn.InitiatePawn()
	if err != nil {
		return 0, "", "", err
	}
	err = pawn.Create(db)
	if err != nil {
		return 0, "", "", err
	}
	board := boards.Board{Model: gorm.Model{ID: game.BoardID}}

	err = board.Read(db)
	if err != nil {
		return 0, "", "", err
	}
	err = board.DeployPawn(db, pawn.ID, 15, 10)
	if err != nil {
		return 0, "", "", err
	}
	return game.User1ID, board.TerrainJson, board.FeaturedMapJson, nil

}
func MakeMoves(db *gorm.DB, in *protos.MoveInputs) error {
	var gamestatus int8
	game := games.Game{}
	game.ID = uint(in.Gameid)
	err := game.Read(db)
	if err != nil {
		log.Error(err)
		return err
	}

	for _, input := range in.Moveinput {
		for _, move := range input.Move {
			m := moves.Move{}
			gamestatus, err = m.AppendMove(db, uint(in.Gameid), uint(in.Userid), uint(input.Pawnid), int16(move.X), int16(move.Y), uint8(move.Direction), game.Round, game.BoardID, game.Status, game.User1ID, game.User2ID)
			if err != nil {
				log.Error(err)
				return err
			}
		}
	}
	if gamestatus == -5 {
		var movesList []moves.Move
		var user1Moves []moves.Move
		var user2Moves []moves.Move
		m := moves.Move{
			GameID: game.ID,
			Round:  game.Round,
		}
		movesList, err = m.List(db)
		if err != nil {
			log.Error(err)
			return err
		}
		//TODO play the game
		game.Round++
	} else {
		game.Status = gamestatus
	}

	err = game.Update(db)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
