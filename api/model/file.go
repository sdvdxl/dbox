package model

import "github.com/jinzhu/gorm"

// File 文件信息
type File struct {
	gorm.Model
	// 文件名
	Name string `gorm:"unique;not null"`
	// 目录id 和 文件是一对多关系
	CategoryID uint

	Category Category `gorm:"foreignkey:CategoryID"`

	// tag
	MD5 string `gorm:"unique;not null"`

	// 相对地址
	Path string
}

type FileDTO struct {
	File
	Category string
}
