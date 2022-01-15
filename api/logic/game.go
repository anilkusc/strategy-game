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

func CreateNewGame(db *gorm.DB, userid uint) (uint, string, error) {
	var err error
	game := games.Game{
		User1ID: userid,
		Round:   0,
		Status:  -2,
	}

	err = game.Create(db)
	if err != nil {
		log.Error(err)
		return 0, "", err
	}

	err = game.Read(db)
	if err != nil {
		log.Error(err)
		return 0, "", err
	}

	board := boards.Board{
		Type:   "flat",
		GameID: game.ID,
	}
	err = board.CreateBoard(db)
	if err != nil {
		log.Error(err)
		return 0, "", err
	}

	err = board.Read(db)
	if err != nil {
		log.Error(err)
		return 0, "", err
	}

	pawn := pawns.Pawn{
		UserID:  uint(userid),
		GameID:  game.ID,
		BoardID: board.ID,
		Type:    "cavalry",
		X:       5,
		Y:       10,
	}
	for i := 3; i < 13; i = i + 3 {
		p := pawn
		p.Y = int16(i)
		err = p.InitiatePawn(1)
		if err != nil {
			log.Error(err)
			return 0, "", err
		}
		err = p.Create(db)
		if err != nil {
			log.Error(err)
			return 0, "", err
		}
		err = p.Read(db)
		if err != nil {
			log.Error(err)
			return 0, "", err
		}

		err = board.DeployPawn(db, p.ID, p.X, p.Y)
		if err != nil {
			log.Error(err)
			return 0, "", err
		}
	}

	game.BoardID = board.ID
	err = game.Update(db)
	if err != nil {
		log.Error(err)
		return 0, "", err
	}

	return game.ID, board.FeaturedMapJson, nil
}
func JoinAGame(db *gorm.DB, user2id uint, gameid uint) (uint, string, error) {
	game := games.Game{
		Model: gorm.Model{
			ID: gameid,
		},
	}
	err := game.Read(db)
	if err != nil {
		log.Error(err)
		return 0, "", err
	}
	game.Status = -1
	game.User2ID = user2id

	pawn := pawns.Pawn{
		UserID:  user2id,
		GameID:  gameid,
		BoardID: game.BoardID,
		Type:    "cavalry",
		X:       15,
		Y:       10,
	}
	board := boards.Board{Model: gorm.Model{ID: game.BoardID}}

	err = board.Read(db)
	if err != nil {
		log.Error(err)
		return 0, "", err
	}
	for i := 3; i < 13; i = i + 3 {
		p := pawn
		p.Y = int16(i)
		err = p.InitiatePawn(3)
		if err != nil {
			log.Error(err)
			return 0, "", err
		}
		err = p.Create(db)
		if err != nil {
			log.Error(err)
			return 0, "", err
		}
		err = board.DeployPawn(db, p.ID, p.X, p.Y)
		if err != nil {
			log.Error(err)
			return 0, "", err
		}
	}

	game.Round = 1
	err = game.Update(db)
	if err != nil {
		log.Error(err)
		return 0, "", err
	}

	return game.User1ID, board.FeaturedMapJson, nil

}
func MakeMoves(db *gorm.DB, in *protos.MoveInputs) error {
	var gamestatus int8
	game, board, err := GetGameAndBoard(db, uint(in.Gameid))
	if err != nil {
		log.Error(err)
		return err
	}
	for _, input := range in.Moveinput {
		for _, move := range input.Move {
			m := moves.Move{}
			p := pawns.Pawn{}

			if p.IsPawnMoveValid(uint8(move.Direction), int16(move.X), int16(move.Y)) {
				gamestatus, err = m.AppendMove(db, uint(in.Gameid), uint(in.Userid), uint(input.Pawnid), int16(move.X), int16(move.Y), uint8(move.Direction), game.Round, game.BoardID, game.Status, game.User1ID, game.User2ID)
				if err != nil {
					log.Error(err)
					return err
				}
			} else {
				log.Error("invalid move")
			}
		}
	}
	if gamestatus == -5 {
		err := GameSimulation(db, &board, &game)
		if err != nil {
			log.Error(err)
			return err
		}

	} else {
		game.Status = gamestatus
	}
	err = UpdateGameAndBoard(db, &board, &game)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
func GetLastMoves(db *gorm.DB, in *protos.LastMovesInputs) ([]moves.Move, error) {

	move := moves.Move{
		GameID: uint(in.Gameid),
		Round:  uint16(in.Round),
	}
	moveList, err := move.List(db)
	if err != nil {
		return moveList, err
	}
	return moveList, nil
}
func GetMaps(db *gorm.DB, gameid uint32) (string, string, error) {

	_, board, err := GetGameAndBoard(db, uint(gameid))
	if err != nil {
		return "", "", err
	}
	return board.TerrainJson, board.FeaturedMapJson, nil
}
