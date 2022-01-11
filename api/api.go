package api

import (
	"context"
	"strategy-game/api/logic"
	"strategy-game/api/protos"

	log "github.com/sirupsen/logrus"
)

func (a *App) CreateGame(ctx context.Context, in *protos.CreateGameInputs) (*protos.CreateGameOutputs, error) {

	gameid, err := logic.CreateNewGame(a.DB, uint(in.Userid))
	if err != nil {
		log.Error(err)
		return &protos.CreateGameOutputs{Gameid: 0}, err
	}
	return &protos.CreateGameOutputs{Gameid: uint32(gameid)}, nil

}

func (a *App) JoinGame(ctx context.Context, in *protos.JoinGameInputs) (*protos.JoinGameOutputs, error) {

	user1ID, terrain, featured, err := logic.JoinAGame(a.DB, uint(in.Userid), uint(in.Gameid))
	if err != nil {
		log.Error(err)
		return &protos.JoinGameOutputs{Otherusersid: 0, Terrainmap: "0", Featuredmap: "0"}, err
	}

	return &protos.JoinGameOutputs{Otherusersid: uint32(user1ID), Terrainmap: terrain, Featuredmap: featured}, nil

}

func (a *App) MakeMove(ctx context.Context, in *protos.MoveInputs) (*protos.MoveOutputs, error) {

	return &protos.MoveOutputs{OK: true}, nil

}
