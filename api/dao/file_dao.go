package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/sdvdxl/dbox/api/ex"
	"github.com/sdvdxl/dbox/api/log"
	"github.com/sdvdxl/dbox/api/model"
	"strings"
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

func (dao *FileDao) FindByFuzz(f model.FileDTO) []model.FileDTO {
	sess := dao.DB.Table("files").Select("files.*,categories.name as category").
		Joins("join categories on files.category_id=categories.id")
	log.Log.Info(f)
	var files []model.FileDTO
	name := strings.ToUpper(EscapeLike(f.Name))
	if name == "" {
		name = "%"
	}

	if f.Category != "" {
		sess = sess.Where("UPPER(files.name) like ? and UPPER(categories.name) = ?",
			"%"+name+"%", strings.ToUpper(f.Category))
	} else {
		sess = sess.Where("UPPER(files.name) like ?  ESCAPE '\\'",
			"%"+name+"%")
	}

	ex.Check(sess.Find(&files).Error)
	return files
}

func (dao *FileDao) FindByName(name string) *model.File {
	var file model.File
	err := dao.DB.Table("files").Where("name = ?", name).Find(&file).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	ex.Check(err)
	return &file
}

func (dao *FileDao) FindByID(id int) *model.File {
	var file model.File
	ex.Check(dao.DB.Table("files").Where("id=?", id).Find(&file).Error)
	return &file
}
