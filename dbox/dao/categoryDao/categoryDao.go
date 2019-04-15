package categoryDao

import (
	"github.com/jinzhu/gorm"
	. "github.com/sdvdxl/dbox/dbox/dao"
	"github.com/sdvdxl/dbox/dbox/ex"
	"github.com/sdvdxl/dbox/dbox/model"
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
