package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Password string
}

type Emailregister struct {
	gorm.Model
	Username  string
	Emailcode string
}
