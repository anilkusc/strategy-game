package users

import (
	"gorm.io/gorm"
)

func (u *User) Create(db *gorm.DB) error {
	result := db.Create(u)
	return result.Error
}
func (u *User) Read(db *gorm.DB) error {
	result := db.First(&u)
	return result.Error

}
func (u *User) Update(db *gorm.DB) error {
	result := db.Save(u)
	return result.Error
}
func (u *User) Delete(db *gorm.DB) error {

	result := db.Delete(&User{}, u.ID)
	return result.Error
}
func (u *User) HardDelete(db *gorm.DB) error {

	result := db.Unscoped().Delete(&User{}, u.ID)
	return result.Error
}
func (u *User) List(db *gorm.DB) ([]User, error) {
	var users []User
	result := db.Find(&users)
	return users, result.Error
}
