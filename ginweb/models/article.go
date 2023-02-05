package models

import (
	"gorm.io/gorm"
)

// 创建博客模型
type Post struct {
	gorm.Model
	Title       string
	Content     string `gorm:"type:text"`
	Tag         string
	Description string
	Ctime       string
}
