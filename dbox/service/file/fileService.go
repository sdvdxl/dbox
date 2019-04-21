package fileService

import (
	"github.com/gohugoio/hugo/helpers"
	"github.com/sdvdxl/dbox/dbox/dao"
	"github.com/sdvdxl/dbox/dbox/dao/categoryDao"
	"github.com/sdvdxl/dbox/dbox/dao/fileDao"
	"github.com/sdvdxl/dbox/dbox/ex"
	. "github.com/sdvdxl/dbox/dbox/log"
	"github.com/sdvdxl/dbox/dbox/model"
	"github.com/sdvdxl/dbox/dbox/service/cloud"
	"os"
	"strings"
)

type (
	FileService interface {
		// UploadLocalFile 上传文件
		UploadLocalFile(file, category string) error
		// FindByCategory  根据目录查找文件
		FindByCategory(categoryID uint) []model.File

		// FindByName 根据文件名字模糊搜索
		FindByNameFuzz(name string) []model.File
	}
)

// UploadLocalFile 上传本地文件
func UploadLocalFile(file, category string) error {
	tx := dao.DB.Begin()
	ex.Check(tx.Error)
	defer dao.RollBackIfPanic(tx)

	Log.Infow("upload file info", "file", file, "category", category)
	stat, err := os.Stat(file)
	if err != nil {
		return ex.FileNotExistErr.Arg(", file:", file)
	}

	if stat.IsDir() {
		return ex.FileErr.Arg("not support dir, ", file)
	}

	f, err := os.Open(file)
	if err != nil {
		return ex.FileErr.Arg(err)
	}

	md5Sum, err := helpers.MD5FromFile(f)
	if err != nil {
		return ex.FileErr.Arg(err)
	}
	Log.Info("file:", file, ", md5:", md5Sum)

	category = strings.TrimSpace(category)
	if category == "" {
		Log.Warn("category is blank, will use default")
		category = model.CatetoryRoot
	}

	c := categoryDao.Save(tx, category)
	existFile := fileDao.FindByMD5(tx, md5Sum)
	if existFile != nil {
		tx.Rollback()
		return ex.FileExistErr.Arg(", file:", file)
	}
	fileDao.Save(tx, &model.File{Name: file, CategoryID: c.ID, MD5: md5Sum, Path: category})
	Log.Info("##################", tx)

	if err := cloudService.Upload(file, category); err != nil {
		tx.Rollback()
		return err
	}

	Log.Error("##################", tx)
	return nil
}

func FindByCategoryID(categoryID uint) []model.File {
	return fileDao.FindByCategoryID(dao.DB, categoryID)
}

func FindByFuzz(f model.FileDTO) []model.File {
	Log.Info(f)
	var files []model.File
	if f.Category != "" {
		ex.Check(dao.DB.Table("files").Select("files.*").
			Joins("join categories on files.category_id=categories.id").
			Where("files.name like ? and categories.name = ?",
				"%"+f.Name+"%", f.Category).
			Find(&files).Error)

	} else {
		ex.Check(dao.DB.Where("name like ?", "%"+f.Name+"%").Find(&files).Error)
	}

	return files

}
