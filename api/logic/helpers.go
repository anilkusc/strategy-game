package logic

import (
	"strategy-game/boards"
	"strategy-game/boards/moves"
	"strategy-game/games"
	"strategy-game/pawns"

	"gorm.io/gorm"
)

func GetGameAndBoard(db *gorm.DB, gameid uint) (games.Game, boards.Board, error) {
	game := games.Game{}
	board := boards.Board{}

	game.ID = uint(gameid)
	err := game.Read(db)
	if err != nil {
		return game, board, err
	}

	board.ID = game.BoardID
	err = board.Read(db)
	if err != nil {
		return game, board, err
	}
	return game, board, nil
}
func UpdateGameAndBoard(db *gorm.DB, board *boards.Board, game *games.Game) error {
	err := board.Update(db)
	if err != nil {
		return err
	}
	err = game.Update(db)
	if err != nil {
		return err
	}
	return nil
}
func GameSimulation(db *gorm.DB, board *boards.Board, game *games.Game) error {
	pawnList, err := board.DetectPawns(db)
	if err != nil {
		return err
	}
	pwn := pawns.Pawn{}
	pawns, err := pwn.ShufflePawns(db, pawnList)
	if err != nil {
		return err
	}
	mv := moves.Move{
		GameID: game.ID,
		Round:  game.Round,
	}
	moveList, err := mv.List(db)
	if err != nil {
		return err
	}

	for _, pawn := range pawns {

		for _, mve := range moveList {
			if mve.PawnID == pawn.ID {
				newX := pawn.X + mve.X
				newY := pawn.Y + mve.Y
				collisions := board.CollisionControl(pawn.X, pawn.Y, pawn.Range)
				if len(collisions) < 1 {
					err = board.MovePawnTo(pawn.X, pawn.Y, newX, newY)
					if err != nil {
						return err
					}
					pawn.X = newX
					pawn.Y = newY
					newcollisions := board.CollisionControl(pawn.X, pawn.Y, pawn.Range)
					if len(newcollisions) > 0 {
						for _, coll := range newcollisions {
							err = pawn.AttackTo(db, uint(coll))
							if err != nil {
								return err
							}
						}
					}
				} else {
					for _, collission := range collisions {
						err = pawn.AttackTo(db, uint(collission))
						if err != nil {
							return err
						}
					}
					if !board.IsPawnPathBlocked(pawn.ID, pawn.X, pawn.Y, newX, newY) {
						err = board.MovePawnTo(pawn.X, pawn.Y, newX, newY)
						if err != nil {
							return err
						}
						pawn.X = newX
						pawn.Y = newY
					}
				}
				err = pawn.Update(db)
				if err != nil {
					return err
				}
				continue
			}
		}
	}
	game.Status = -1
	game.Round++
	return nil
}
