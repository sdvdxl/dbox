package main

import (
	"fmt"
	"github.com/sdvdxl/dbox/dbox/config"
	"github.com/sdvdxl/dbox/dbox/dao"
	"github.com/sdvdxl/dbox/dbox/ex"
	. "github.com/sdvdxl/dbox/dbox/log"
	"github.com/sdvdxl/dbox/dbox/model"
	categoryService "github.com/sdvdxl/dbox/dbox/service/category"
	"github.com/sdvdxl/dbox/dbox/service/cloud"
	fileService "github.com/sdvdxl/dbox/dbox/service/file"
	"os"
)

const (
//metaFile = "meta.db"
)

func main() {
	err := config.Parse("cfg.yml")
	if err != nil {
		Log.Error("parse config error,", err)
		os.Exit(-1)
	}

	dao.Use(dao.DialectSqlite3, "meta.db")
	dao.Init()
	defer dao.Close()
	err = cloudService.UseCloudFileManager(
		&cloudService.AliOssFileManager{})
	if err != nil {
		Log.Error("oss config error", err)
		os.Exit(-1)
	}
	upLoad("测试 as")
	Log.Info("FindByCategory 1", FindByCategory(1))

	FindAllCategory()
}

func FindByCategory(categoryID uint) []model.File {
	return fileService.FindByCategoryID(categoryID)
}

func FindAllCategory() {
	Log.Info("FindAllCategory")
	for _, v := range categoryService.FindAll() {
		fmt.Println(v)
	}
}

func upLoad(category string) {
	err := fileService.UploadLocalFile("/tmp/a", category)

	if err != nil {
		if e, ok := err.(ex.Error); ok {
			Log.Info(e)
			return
		}

		Log.Error(err)
	}
}
