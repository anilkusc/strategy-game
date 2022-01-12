package logic

import (
	"strategy-game/api/protos"
	"strategy-game/boards"
	"strategy-game/boards/moves"
	"strategy-game/games"
	"strategy-game/pawns"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

//TODO : add logs here for != nil expressions.
func CreateNewGame(db *gorm.DB, userid uint) (uint, error) {
	var err error
	game := games.Game{
		User1ID: userid,
		Round:   0,
		Status:  -2,
	}

	err = game.Create(db)
	if err != nil {
		return 0, err
	}

	err = game.Read(db)
	if err != nil {
		return 0, err
	}

	board := boards.Board{
		Type:   "flat",
		GameID: game.ID,
	}
	err = board.CreateBoard(db)
	if err != nil {
		return 0, err
	}

	err = board.Read(db)
	if err != nil {
		return 0, err
	}

	pawn := pawns.Pawn{
		UserID:  uint(userid),
		GameID:  game.ID,
		BoardID: board.ID,
		Type:    "cavalry",
		X:       5,
		Y:       10,
	}
	err = pawn.InitiatePawn()
	if err != nil {
		return 0, err
	}
	err = pawn.Create(db)
	if err != nil {
		return 0, err
	}

	game.BoardID = board.ID
	err = game.Update(db)
	if err != nil {
		return 0, err
	}

	err = board.DeployPawn(db, pawn.ID, pawn.X, pawn.Y)
	if err != nil {
		return 0, err
	}
	return game.ID, nil
}
func JoinAGame(db *gorm.DB, user2id uint, gameid uint) (uint, string, string, error) {
	game := games.Game{
		Model: gorm.Model{
			ID: gameid,
		},
	}
	err := game.Read(db)
	if err != nil {
		return 0, "", "", err
	}
	game.Status = -1
	game.User2ID = user2id

	pawn := pawns.Pawn{
		UserID:  user2id,
		GameID:  gameid,
		BoardID: game.BoardID,
		Type:    "cavalry",
		X:       15,
		Y:       10,
	}
	err = pawn.InitiatePawn()
	if err != nil {
		return 0, "", "", err
	}
	err = pawn.Create(db)
	if err != nil {
		return 0, "", "", err
	}
	board := boards.Board{Model: gorm.Model{ID: game.BoardID}}

	err = board.Read(db)
	if err != nil {
		return 0, "", "", err
	}
	err = board.DeployPawn(db, pawn.ID, pawn.X, pawn.Y)
	if err != nil {
		return 0, "", "", err
	}

	err = game.Update(db)
	if err != nil {
		return 0, "", "", err
	}

	return game.User1ID, board.TerrainJson, board.FeaturedMapJson, nil

}
func MakeMoves(db *gorm.DB, in *protos.MoveInputs) error {
	var gamestatus int8

	game := games.Game{}
	game.ID = uint(in.Gameid)
	err := game.Read(db)
	if err != nil {
		log.Error(err)
		return err
	}
	board := boards.Board{}
	board.ID = game.BoardID
	err = board.Read(db)
	if err != nil {
		log.Error(err)
		return err
	}
	for _, input := range in.Moveinput {
		for _, move := range input.Move {
			m := moves.Move{}
			gamestatus, err = m.AppendMove(db, uint(in.Gameid), uint(in.Userid), uint(input.Pawnid), int16(move.X), int16(move.Y), uint8(move.Direction), game.Round, game.BoardID, game.Status, game.User1ID, game.User2ID)
			if err != nil {
				log.Error(err)
				return err
			}
		}
	}
	if gamestatus == -5 {
		var movesList []moves.Move
		var user1Moves []moves.Move
		var user2Moves []moves.Move
		m := moves.Move{
			GameID: game.ID,
			Round:  game.Round,
		}
		movesList, err = m.List(db)
		if err != nil {
			log.Error(err)
			return err
		}

		for _, mv := range movesList {
			if mv.UserID == game.User1ID {
				user1Moves = append(user1Moves, mv)
			} else {
				user2Moves = append(user2Moves, mv)
			}
		}
		var maxlength int
		if len(user1Moves) > len(user2Moves) {
			maxlength = len(user1Moves)
		} else {
			maxlength = len(user2Moves)
		}

		for i := 0; i < maxlength; i++ {
			pawn1 := pawns.Pawn{}
			pawn2 := pawns.Pawn{}
			pawn1.ID = user1Moves[i].PawnID
			pawn2.ID = user2Moves[i].PawnID
			err = pawn1.Read(db)
			if err != nil {
				log.Error(err)
				return err
			}
			err = pawn2.Read(db)
			if err != nil {
				log.Error(err)
				return err
			}

			board.Terrain[pawn1.Y][pawn1.X] = 0
			pawn1.X = pawn1.X + user1Moves[i].X
			pawn1.Y = pawn1.Y + user1Moves[i].Y
			board.Terrain[pawn1.Y][pawn1.X] = int16(pawn1.ID)
			err = pawn1.Update(db)
			if err != nil {
				log.Error(err)
				return err
			}
			board.Terrain[pawn2.Y][pawn2.X] = 0
			pawn2.X = pawn2.X + user2Moves[i].X
			pawn2.Y = pawn2.Y + user2Moves[i].Y
			board.Terrain[pawn2.Y][pawn2.X] = int16(pawn2.ID)

			err = pawn2.Update(db)
			if err != nil {
				log.Error(err)
				return err
			}

		}
		game.Status = gamestatus
		game.Round++
	} else {
		game.Status = gamestatus
	}
	err = board.Update(db)
	if err != nil {
		log.Error(err)
		return err
	}
	err = game.Update(db)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
