package model

type User struct {
	// primary key
	ID int64 `gorm:"primary_key;not_null;auto_increment"`
	// user name
	UserName string `gorm:"unique_index;not_null"`
	// required text
	FirstName    string
	HashPassword string
}
