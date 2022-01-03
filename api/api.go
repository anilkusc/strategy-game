package api

import (
	"context"
	"strategy-game/api/protos"
	"strategy-game/boards"
	"strategy-game/games"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (a *App) CreateGame(ctx context.Context, in *protos.CreateGameInputs) (*protos.CreateGameOutputs, error) {

	game := games.Game{
		User1ID: uint(in.Userid),
		User2ID: 0,
		BoardID: 0,
		Status:  -2,
	}
	err := game.CreateNewGame(a.DB)
	if err != nil {
		log.Error(err)
		return &protos.CreateGameOutputs{Gameid: 0}, err
	} else {
		return &protos.CreateGameOutputs{Gameid: uint64(game.ID)}, nil
	}

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

/*
func (a *App) MakeMove(ctx context.Context, in *protos.MoveInputs) (*protos.MoveOutputs, error) {
	//uint32 userid=1;
	//uint32 gameid=2;
	//uint32 pawnid=3;
	//uint32 x=4;
	//uint32 y=5;
	game := games.Game{
		Model: gorm.Model{
			ID: uint(in.Gameid),
		},
	}
	board := boards.Board{
		Model: gorm.Model{
			ID: game.BoardID,
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
*/
