package api

import (
	"context"
	"strategy-game/api/protos"
	"strategy-game/boards/moves"
	"strategy-game/games"
	logic "strategy-game/logic"

	log "github.com/sirupsen/logrus"
)

func (a *App) CreateGame(ctx context.Context, in *protos.CreateGameInputs) (*protos.CreateGameOutputs, error) {

	err, gameid := logic.CreateNewGame(a.DB, uint(in.Userid))
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

	for _, input := range in.Moveinput {
		for _, move := range input.Move {
			m := moves.Move{}
			err := m.AppendMove(a.DB, uint(in.Gameid), uint(in.Userid), uint(input.Pawnid), int16(move.X), int16(move.Y), uint8(move.Direction))
			if err != nil {
				log.Error(err)
				return &protos.MoveOutputs{OK: false}, nil
			}
		}
	}
	game := games.Game{}
	game.ID = uint(in.Gameid)
	err := game.Read(a.DB)
	if err != nil {
		log.Error(err)
		return &protos.MoveOutputs{OK: false}, nil
	}
	if game.Status == -5 {
		//
	}
	return &protos.MoveOutputs{OK: true}, nil

}
