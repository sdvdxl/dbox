package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/sdvdxl/dbox/api/ex"
	"github.com/sdvdxl/dbox/api/model"
)

type CategoryDao struct {
	DB *gorm.DB
}

func (dao *CategoryDao) Save(name string) *model.Category {
	var c model.Category
	ex.Check(dao.DB.Where(model.Category{Name: name}).FirstOrCreate(&c).Error)
	return &c
}

func (dao *CategoryDao) FindAll() []model.Category {
	var models []model.Category
	ex.Check(DB.Find(&models).Error)
	return models

}

func (dao *CategoryDao) CreateCategory(name string) *model.Category {
	c := &model.Category{Name: name}
	ex.Check(dao.DB.Where(map[string]interface{}{"name": name}).FirstOrCreate(c).Error)
	return c
}
