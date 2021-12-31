package users

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Construct() (*gorm.DB, User) {
	var db *gorm.DB
	var user = User{
		Model: gorm.Model{
			//ID:        1,
			UpdatedAt: time.Time{},
			CreatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{Time: time.Time{}, Valid: false},
		},
		Username: "testuser",
		Password: "testpassword",
		Role:     "testrole",
	}
	db, _ = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	db.AutoMigrate(&User{})
	return db, user
}
func Destruct() {
	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	db.Exec("DROP TABLE users")
}
