package api

import (
	"net"
	"strategy-game/api/protos"
	"strategy-game/boards"
	"strategy-game/boards/moves"
	"strategy-game/games"
	"strategy-game/games/users"
	"strategy-game/pawns"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type App struct {
	DB *gorm.DB
	protos.UnimplementedStrategyGameServer
}

func (a *App) Init() {
	log.Info("application is starting...")
	var err error
	log.Info("connecting to database..")
	a.DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Error("error while connecting database.")
	}
	log.Info("connected to database..")
	log.Info("migrating database tables...")
	a.DB.AutoMigrate(&games.Game{}, &users.User{}, &pawns.Pawn{}, &boards.Board{}, &moves.Move{})
	log.Info("migrated...")
	log.Info("creating grpc api...")
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	protos.RegisterStrategyGameServer(s, a)
	log.Info("server listening at ", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (a *App) Start() {
	a.Init()

}
