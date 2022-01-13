package logic

import (
	"strategy-game/boards"
	"strategy-game/games"

	"gorm.io/gorm"
)

func GetGameAndBoard(db *gorm.DB, gameid uint) (games.Game, boards.Board, error) {
	game := games.Game{}
	board := boards.Board{}

	game.ID = uint(gameid)
	err := game.Read(db)
	if err != nil {
		return game, board, err
	}

	board.ID = game.BoardID
	err = board.Read(db)
	if err != nil {
		return game, board, err
	}
	return game, board, nil
}
func UpdateGameAndBoard(db *gorm.DB, board *boards.Board, game *games.Game) error {
	err := board.Update(db)
	if err != nil {
		return err
	}
	err = game.Update(db)
	if err != nil {
		return err
	}
	return nil
}
