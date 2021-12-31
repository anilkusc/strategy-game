package boards

import (
	"errors"

	"gorm.io/gorm"
)

type IBoard interface {
	Create(*gorm.DB) error
	Read(*gorm.DB) error
	Update(*gorm.DB) error
	Delete(*gorm.DB) error
	HardDelete(*gorm.DB) error
	List(*gorm.DB) ([]Board, error)
	ArrayToJson([][]int16) (string, error)
	JsonToArray(string) ([][]int16, error)
	DeployPawn(*gorm.DB, uint, int16, int16) error
}

type Board struct {
	gorm.Model
	GameID          uint
	LengthX         int16
	LengthY         int16
	Type            string
	TerrainJson     string
	FeaturedMapJson string
	Terrain         [][]int16 `gorm:"-"`
	FeaturedMap     [][]int16 `gorm:"-"`
}

func (b *Board) DeployPawn(db *gorm.DB, pawnID uint, X int16, Y int16) error {
	err := b.Read(db)
	if err != nil {
		return err
	}
	if b.Terrain[Y][X] == 0 {
		b.Terrain[Y][X] = int16(pawnID)
		err = b.Update(db)
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("cannot deploy pawn since there is another pawn in the location")
	}

}
