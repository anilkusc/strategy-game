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
	DetectPawns(*gorm.DB) ([]uint, error)
	MovePawnTo(int16, int16, int16, int16) error
	CreateBoard(*gorm.DB) error
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
	if b.Terrain[Y][X] == b.FeaturedMap[Y][X] {
		b.Terrain[Y][X] = int16(pawnID)
		b.TerrainJson, err = b.ArrayToJson(b.Terrain)
		if err != nil {
			return err
		}
	} else {
		return errors.New("cannot deploy pawn since there is another pawn in the location")
	}

	err = b.Update(db)
	if err != nil {
		return err
	}
	return nil
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
	default:
		return errors.New("board type is not recognized")
	}

	err := b.Create(db)
	if err != nil {
		return err
	}
	return nil
}

func (b *Board) DetectPawns(db *gorm.DB) ([]uint, error) {
	err := b.Read(db)
	if err != nil {
		return []uint{0}, err
	}
	var pawns []uint
	for i, terrain := range b.Terrain {
		for j, point := range terrain {
			if point != b.FeaturedMap[i][j] {
				pawns = append(pawns, uint(point))
			}
		}
	}
	return pawns, nil
}

func (b *Board) MovePawnTo(fromX int16, fromY int16, toX int16, toY int16) error {
	var err error
	if b.Terrain[toY][toX] == b.FeaturedMap[toY][toX] && b.Terrain[fromY][fromX] != b.FeaturedMap[fromY][fromX] {
		pawn := b.Terrain[fromY][fromX]
		b.Terrain[fromY][fromX] = b.FeaturedMap[fromY][fromX]
		b.Terrain[toY][toX] = pawn
		b.TerrainJson, err = b.ArrayToJson(b.Terrain)
		if err != nil {
			return err
		}
	} else {
		return errors.New("cannot move pawn")
	}

	return nil
}
