package moves

import (
	"gorm.io/gorm"
)

func (m *Move) Create(db *gorm.DB) error {
	result := db.Create(m)
	return result.Error
}
func (m *Move) Read(db *gorm.DB) error {
	result := db.First(&m)
	return result.Error

}
func (m *Move) Update(db *gorm.DB) error {
	result := db.Save(m)
	return result.Error
}
func (m *Move) Delete(db *gorm.DB) error {

	result := db.Delete(&Move{}, m.ID)
	return result.Error
}
func (m *Move) HardDelete(db *gorm.DB) error {

	result := db.Unscoped().Delete(&Move{}, m.ID)
	return result.Error
}
func (m *Move) List(db *gorm.DB) ([]Move, error) {
	var moves []Move
	result := db.Where("game_id = ? ", m.GameID).Find(&moves)
	return moves, result.Error
}
