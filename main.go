package main

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sdvdxl/dbox/dbox/db"
	"github.com/sdvdxl/dbox/dbox/ex"
	. "github.com/sdvdxl/dbox/dbox/log"
	"github.com/sdvdxl/dbox/dbox/service"
	"os"
)

const (
//metaFile = "meta.db"
)

func main() {
	db.Use(db.DialectSqlite3, "meta.db")
	db.Init()
	defer db.Close()
	err := service.UseCloudFileManager(
		&service.AliOssFileManager{
			Endpoint:        os.Args[1],
			AccessKeyID:     os.Args[2],
			AccessKeySecret: os.Args[3],
			Bucket:          os.Args[4]})
	if err != nil {
		Log.Error("oss config error", err)
		os.Exit(-1)
	}

	err = service.UploadLocalFile("/tmp/a", "aa", "aa")

	if err != nil {
		if e, ok := err.(ex.Error); ok {
			Log.Info(e)
			return
		}

		Log.Error(err)
	}
}
