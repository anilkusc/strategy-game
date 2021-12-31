package pawns

import (
	"gorm.io/gorm"
)

func (p *Pawn) Create(db *gorm.DB) error {
	result := db.Create(p)
	return result.Error
}
func (p *Pawn) Read(db *gorm.DB) error {
	result := db.First(&p)
	return result.Error

}
func (p *Pawn) Update(db *gorm.DB) error {
	result := db.Save(p)
	return result.Error
}
func (p *Pawn) Delete(db *gorm.DB) error {

	result := db.Delete(&Pawn{}, p.ID)
	return result.Error
}
func (p *Pawn) HardDelete(db *gorm.DB) error {

	result := db.Unscoped().Delete(&Pawn{}, p.ID)
	return result.Error
}
func (p *Pawn) List(db *gorm.DB) ([]Pawn, error) {
	var pawns []Pawn
	result := db.Where("game_id = ? ", p.GameID).Find(&pawns)
	return pawns, result.Error
}
