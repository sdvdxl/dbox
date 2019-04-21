package model

import "github.com/jinzhu/gorm"

// File 文件信息
type File struct {
	gorm.Model
	// 文件名
	Name string
	// 目录id 和 文件是一对多关系
	CategoryID uint

	// tag
	MD5 string

	// 相对地址
	Path string
}

type FileDTO struct {
	Name     string
	Category string
}
