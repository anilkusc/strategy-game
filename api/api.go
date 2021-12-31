package api

import (
	"context"
	"log"
	"strategy-game/api/protos"
)

func (a *App) Hello(ctx context.Context, in *protos.Values) (*protos.Result, error) {
	log.Println("Received Message")
	log.Println(in)
	return &protos.Result{Result: "hello"}, nil
}
