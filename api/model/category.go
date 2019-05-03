package model

import "github.com/jinzhu/gorm"

const CatetoryRoot = "default"

// Category 目录，暂时只支持一级
type Category struct {
	gorm.Model
	Name string
}
