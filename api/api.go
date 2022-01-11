package api

import (
	"context"
	"strategy-game/api/protos"
	"strategy-game/boards"
	"strategy-game/boards/moves"
	"strategy-game/games"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (a *App) CreateGame(ctx context.Context, in *protos.CreateGameInputs) (*protos.CreateGameOutputs, error) {

	game := games.Game{
		User1ID: uint(in.Userid),
		User2ID: 0,
		Status:  -2,
	}

	err := game.CreateNewGame(a.DB, game.User1ID)
	if err != nil {
		log.Error(err)
		return &protos.CreateGameOutputs{Gameid: 0}, err
	}
	return &protos.CreateGameOutputs{Gameid: uint32(game.ID)}, nil

}

func (a *App) JoinGame(ctx context.Context, in *protos.JoinGameInputs) (*protos.JoinGameOutputs, error) {
	game := games.Game{
		Model: gorm.Model{
			ID: uint(in.Gameid),
		},
	}

	err := game.JoinGame(a.DB, uint(in.Userid))
	if err != nil {
		log.Error(err)
		return &protos.JoinGameOutputs{Otherusersid: 0, Terrainmap: "0", Featuredmap: "0"}, err
	}
	board := boards.Board{
		Model: gorm.Model{
			ID: uint(in.Gameid),
		},
	}
	err = board.Read(a.DB)
	if err != nil {
		log.Error(err)
		return &protos.JoinGameOutputs{Otherusersid: 0, Terrainmap: "0", Featuredmap: "0"}, err
	}

	return &protos.JoinGameOutputs{Otherusersid: uint32(game.User1ID), Terrainmap: board.TerrainJson, Featuredmap: board.FeaturedMapJson}, err

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
