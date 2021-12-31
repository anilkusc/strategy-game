package games

import (
	"gorm.io/gorm"
)

func (g *Game) Create(db *gorm.DB) error {
	result := db.Create(g)
	return result.Error
}
func (g *Game) Read(db *gorm.DB) error {
	result := db.First(&g)
	return result.Error

}
func (g *Game) Update(db *gorm.DB) error {
	result := db.Save(g)
	return result.Error
}
func (g *Game) Delete(db *gorm.DB) error {

	result := db.Delete(&Game{}, g.ID)
	return result.Error
}
func (g *Game) HardDelete(db *gorm.DB) error {

	result := db.Unscoped().Delete(&Game{}, g.ID)
	return result.Error
}
func (g *Game) List(db *gorm.DB) ([]Game, error) {
	var games []Game
	result := db.Find(&games)
	return games, result.Error
}
