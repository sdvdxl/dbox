package categoryDao

import (
	"github.com/jinzhu/gorm"
	. "github.com/sdvdxl/dbox/api/dao"
	"github.com/sdvdxl/dbox/api/ex"
	"github.com/sdvdxl/dbox/api/model"
)

func Save(session *gorm.DB, name string) *model.Category {
	var c model.Category
	ex.Check(session.Where(model.Category{Name: name}).FirstOrCreate(&c).Error)
	return &c
}

func FindAll() []model.Category {
	var models []model.Category
	ex.Check(DB.Find(&models).Error)
	return models

}

func CreateCategory(session *gorm.DB, name string) *model.Category {
	c := &model.Category{Name: name}
	ex.Check(session.Where(map[string]interface{}{"name": name}).FirstOrCreate(c).Error)
	return c
}
