package logic

import (
	"strategy-game/boards"
	"strategy-game/games"
	"strategy-game/pawns"

	"gorm.io/gorm"
)

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
