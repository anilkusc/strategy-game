package api

import (
	"context"
	"strategy-game/api/logic"
	"strategy-game/api/protos"

	log "github.com/sirupsen/logrus"
)

func (a *App) CreateGame(ctx context.Context, in *protos.CreateGameInputs) (*protos.CreateGameOutputs, error) {

	gameid, featuredmap, err := logic.CreateNewGame(a.DB, uint(in.Userid))
	if err != nil {
		log.Error(err)
		return &protos.CreateGameOutputs{Gameid: 0}, err
	}
	return &protos.CreateGameOutputs{Gameid: uint32(gameid), Featuredmap: featuredmap}, nil

}

func (a *App) JoinGame(ctx context.Context, in *protos.JoinGameInputs) (*protos.JoinGameOutputs, error) {

	user1ID, featured, err := logic.JoinAGame(a.DB, uint(in.Userid), uint(in.Gameid))
	if err != nil {
		log.Error(err)
		return &protos.JoinGameOutputs{Otherusersid: 0, Featuredmap: "0"}, err
	}

	return &protos.JoinGameOutputs{Otherusersid: uint32(user1ID), Featuredmap: featured}, nil

}

func (a *App) MakeMove(ctx context.Context, in *protos.MoveInputs) (*protos.MoveOutputs, error) {

	err := logic.MakeMoves(a.DB, in)
	if err != nil {
		log.Error(err)
		return &protos.MoveOutputs{OK: false}, err
	}

	return &protos.MoveOutputs{OK: true}, nil

}
func (a *App) GetLastMoves(ctx context.Context, in *protos.LastMovesInputs) (*protos.LastMovesOutputs, error) {

	movesList, err := logic.GetLastMoves(a.DB, in)
	if err != nil {
		log.Error(err)
		return &protos.LastMovesOutputs{}, err
	}

	var mvs []*protos.LastMovesOutput
	for _, moveList := range movesList {
		mvs = append(mvs, &protos.LastMovesOutput{Userid: uint32(moveList.UserID), Gameid: uint32(moveList.GameID), Pawnid: uint32(moveList.PawnID), X: int32(moveList.X), Y: int32(moveList.Y), Direction: int32(moveList.Direction), Round: int32(moveList.Round)})
	}

	return &protos.LastMovesOutputs{Lastmovesoutput: mvs}, nil

}
