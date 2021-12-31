package boards

import (
	"gorm.io/gorm"
)

func (b *Board) Create(db *gorm.DB) error {
	result := db.Create(b)
	return result.Error
}
func (b *Board) Read(db *gorm.DB) error {
	var err error
	result := db.First(&b)
	b.FeaturedMap, err = b.JsonToArray(b.FeaturedMapJson)
	if err != nil {
		return err
	}
	b.Terrain, err = b.JsonToArray(b.TerrainJson)
	if err != nil {
		return err
	}
	return result.Error

}
func (b *Board) Update(db *gorm.DB) error {
	result := db.Save(b)
	return result.Error
}
func (b *Board) Delete(db *gorm.DB) error {

	result := db.Delete(&Board{}, b.ID)
	return result.Error
}
func (b *Board) HardDelete(db *gorm.DB) error {

	result := db.Unscoped().Delete(&Board{}, b.ID)
	return result.Error
}
func (b *Board) List(db *gorm.DB) ([]Board, error) {
	var boards []Board
	var err error
	result := db.Where("game_id = ? ", b.GameID).Find(&boards)
	for i := range boards {
		boards[i].FeaturedMap, err = boards[i].JsonToArray(boards[i].FeaturedMapJson)
		if err != nil {
			return boards, err
		}
		boards[i].Terrain, err = boards[i].JsonToArray(boards[i].TerrainJson)
		if err != nil {
			return boards, err
		}
	}
	return boards, result.Error
}
