package api

import (
	"context"
	"log"
	"net"
	"strategy-game/api/protos"
	"strategy-game/boards"
	"strategy-game/boards/moves"
	"strategy-game/games"
	"strategy-game/games/users"
	"strategy-game/pawns"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func Construct() App {
	a := App{}
	a.DB, _ = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	a.DB.AutoMigrate(&games.Game{}, &users.User{}, &pawns.Pawn{}, &boards.Board{}, &moves.Move{})

	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	protos.RegisterStrategyGameServer(s, &a)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
	return a
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}
func Destruct() {
	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	db.Exec("DROP TABLE games")
	db.Exec("DROP TABLE users")
	db.Exec("DROP TABLE pawns")
	db.Exec("DROP TABLE boards")
	db.Exec("DROP TABLE moves")
}

func TestCreateGame(t *testing.T) {
	Construct()
	tests := []struct {
		input  uint32
		output uint64
		err    error
	}{
		{
			input:  1,
			output: 1,
			err:    nil,
		},
	}
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := protos.NewStrategyGameClient(conn)

	for _, test := range tests {
		req := &protos.CreateGameInputs{User1Id: test.input}
		resp, err := client.CreateGame(ctx, req)
		if test.err != err {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
		if resp.Gameid != test.output {
			t.Errorf("Result is: %v . Expected: %v", resp.Gameid, test.output)
		}

	}
	Destruct()
}
