package boards

import (
	"errors"

	"gorm.io/gorm"
)

type Boarder interface {
	Create(*gorm.DB) error
	Read(*gorm.DB) error
	Update(*gorm.DB) error
	Delete(*gorm.DB) error
	HardDelete(*gorm.DB) error
	List(*gorm.DB) ([]Board, error)
	ArrayToJson([][]int16) (string, error)
	JsonToArray(string) ([][]int16, error)
	DeployPawn(*gorm.DB, uint, int16, int16) error
	CreateBoard(db *gorm.DB) error
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
func (b *Board) CreateBoard(db *gorm.DB) error {
	switch b.Type {
	case "flat":
		var err error
		b.LengthX = 20
		b.LengthY = 20
		for i := 0; i < int(b.LengthX); i++ {
			m := []int16{}
			for j := 0; j < int(b.LengthY); j++ {
				m = append(m, 0)
			}
			b.FeaturedMap = append(b.FeaturedMap, m)
			b.Terrain = append(b.Terrain, m)
		}
		b.FeaturedMapJson, err = b.ArrayToJson(b.FeaturedMap)
		if err != nil {
			return err
		}
		b.TerrainJson, err = b.ArrayToJson(b.Terrain)
		if err != nil {
			return err
		}
		err = b.Create(db)
		if err != nil {
			return err
		}
		return nil
	default:
		return errors.New("board type is not recognized")
	}
}
