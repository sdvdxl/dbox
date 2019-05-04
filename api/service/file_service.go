package service

import (
	"github.com/getlantern/errors"
	"github.com/gohugoio/hugo/helpers"
	"github.com/sdvdxl/dbox/api/config"
	"github.com/sdvdxl/dbox/api/dao"
	"github.com/sdvdxl/dbox/api/ex"
	. "github.com/sdvdxl/dbox/api/log"
	"github.com/sdvdxl/dbox/api/model"
	"os"
	"path/filepath"
	"strings"
)

type FileService struct {
}

// UploadLocalFile 上传本地文件
func (s *FileService) UploadLocalFile(file, fileName, category string) (*model.File, error) {
	db := dao.NewDB()
	fileDao := &dao.FileDao{DB: db}
	categoryDao := &dao.CategoryDao{DB: db}

	defer func() {
		if err := recover(); err != nil {
			Log.Error(err)
			db.Rollback()
		}
	}()

	Log.Infow("upload file info", "file", file, "category", category)
	stat, err := os.Stat(file)
	if err != nil {
		return nil, ex.FileNotExistErr.Arg(", file:", file)
	}

	if stat.IsDir() {
		return nil, ex.FileErr.Arg("not support dir, ", file)
	}

	f, err := os.Open(file)
	if err != nil {
		return nil, ex.FileErr.Arg(err)
	}

	md5Sum, err := helpers.MD5FromFile(f)
	if err != nil {
		return nil, ex.FileErr.Arg(err)
	}
	Log.Info("file:", file, ", md5:", md5Sum)

	category = strings.TrimSpace(category)
	if category == "" {
		Log.Warn("category is blank, will use default")
		category = model.CatetoryRoot
	}

	c := categoryDao.Save(category)
	existFile := fileDao.FindByMD5(md5Sum)
	if existFile != nil {
		db.Rollback()
		return existFile, ex.FileExistErr.Arg(", file:", file)
	}

	if fileName == "" {
		fileName = filepath.Base(file)
	}

	existFile = fileDao.FindByName(fileName)
	if existFile != nil {
		db.Rollback()
		return existFile, ex.FileExistErr.Arg(", file:", file)
	}

	fileDao.Save(&model.File{Name: fileName, CategoryID: c.ID, MD5: md5Sum, Path: category})

	if err := cfm.Upload(file, fileName, category); err != nil {
		db.Rollback()
		return nil, err
	}

	existFile = fileDao.FindByName(fileName)
	ex.Check(db.Commit().Error)
	return existFile, nil
}

func (s *FileService) FindByCategoryID(categoryID uint) []model.File {
	return fileDao().FindByCategoryID(dao.DB, categoryID)
}

func (s *FileService) FindByFuzz(f model.FileDTO) []model.FileDTO {
	return fileDao().FindByFuzz(f)
}

func (s *FileService) Download(id int, folder, filename string) (string, error) {
	file := fileDao().FindByID(id)
	if filename == "" {
		filename = file.Name
	}

	if folder == "" {
		folder = config.GetHomeDir()
	}

	fpath := filepath.Clean(folder) + string(filepath.Separator) + filename
	if file == nil {
		return fpath, ex.FileNotExistErr
	}

	_, err := os.Stat(fpath)

	if err == nil {
		return fpath, ex.FileExistErr
	}

	return fpath, cfm.Download((*file).Path, fpath)
}

func (s *FileService) SyncDBFile(command string) error {
	switch command {
	case "upload":
		_, err := os.Stat(config.GetDBFile())
		if os.IsNotExist(err) {
			return ex.FileNotExistErr
		}

		return cfm.Upload(config.GetDBFile(), "meta.db", "meta")
	case "download":
		return cfm.Download("meta/meta.db", config.GetDBFile())
	case "merge":
		return nil
	default:
		return errors.New("command not support, available is : upload, download or merge")
	}

}

func fileDao() *dao.FileDao {
	return &dao.FileDao{DB: dao.NewDB()}
}
