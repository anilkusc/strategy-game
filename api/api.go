package api

import (
	"context"
	"strategy-game/api/protos"
	"strategy-game/boards"
	"strategy-game/games"

	log "github.com/sirupsen/logrus"
)

func (a *App) CreateGame(ctx context.Context, in *protos.CreateGameInputs) (*protos.CreateGameOutputs, error) {

	game := games.Game{
		User1ID: uint(in.User1Id),
		User2ID: 0,
		BoardID: 0,
		Status:  -2,
	}
	err := game.Create(a.DB)
	if err != nil {
		log.Error(err)
		return &protos.CreateGameOutputs{Gameid: 0}, err
	}
	err = game.Read(a.DB)
	if err != nil {
		log.Error(err)
		return &protos.CreateGameOutputs{Gameid: 0}, err
	}

	board := boards.Board{
		Type:   "flat",
		GameID: game.ID,
	}
	err = board.Create(a.DB)
	if err != nil {
		log.Error(err)
		return &protos.CreateGameOutputs{Gameid: 0}, err
	}

	return &protos.CreateGameOutputs{Gameid: uint64(game.ID)}, nil
}
