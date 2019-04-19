package fileDao

import (
	"github.com/jinzhu/gorm"
	"github.com/sdvdxl/dbox/dbox/ex"
	"github.com/sdvdxl/dbox/dbox/log"
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

func FindByMD5(session *gorm.DB, md5 string) *model.File {
	var file model.File
	err := session.Where(map[string]interface{}{"md5": md5}).Find(&file).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Log.Debug("md5:", md5, ", not found")
			return nil
		}

		ex.Check(err)
	}

	return &file
}
