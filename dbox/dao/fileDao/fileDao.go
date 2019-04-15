package fileDao

import (
	"github.com/jinzhu/gorm"
	"github.com/sdvdxl/dbox/dbox/ex"
	"github.com/sdvdxl/dbox/dbox/model"
)

func Save(session *gorm.DB, file *model.File) {
	ex.Check(session.Create(file).Error)
}

func FindByCategoryID(session *gorm.DB, categoryID uint) []model.File {
	var files []model.File
	ex.Check(session.Where(map[string]interface{}{"category_id": categoryID}).Find(&files).Error)
	return files
}
