package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/sdvdxl/dbox/api/ex"
	"github.com/sdvdxl/dbox/api/log"
	"github.com/sdvdxl/dbox/api/model"
)

type FileDao struct {
	DB *gorm.DB
}

func (dao *FileDao) Save(file *model.File) {
	ex.Check(dao.DB.Create(file).Error)
}

func (dao *FileDao) FindByCategoryID(session *gorm.DB, categoryID uint) []model.File {
	var files []model.File
	ex.Check(session.Where(map[string]interface{}{"category_id": categoryID}).Find(&files).Error)
	return files
}

func (dao *FileDao) FindByMD5(md5 string) *model.File {
	var file model.File
	err := dao.DB.Where(map[string]interface{}{"md5": md5}).Find(&file).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Log.Debug("md5:", md5, ", not found")
			return nil
		}

		ex.Check(err)
	}

	return &file
}