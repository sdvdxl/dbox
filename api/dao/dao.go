package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sdvdxl/dbox/api/config"
	"github.com/sdvdxl/dbox/api/ex"
	"github.com/sdvdxl/dbox/api/log"
	"github.com/sdvdxl/dbox/api/model"
	"strings"
	"time"
)

// global db or gorm db
var (
	DB *gorm.DB
)

func NewDB() *gorm.DB {
	return DB.Begin()
}

func Init() {
	var err error
	DB, err = gorm.Open("sqlite3", config.GetDBFile())
	if err != nil {
		panic("failed to connect database, error msg:" + err.Error())
	}

	DB.SetLogger(log.Logger(*log.Log))
	DB.Debug()
	DB.LogMode(true)
	ex.Check(DB.DB().Ping())
	DB.DB().SetMaxIdleConns(2)
	DB.DB().SetMaxOpenConns(20)
	DB.DB().SetConnMaxLifetime(time.Hour)

	initSchemas()

	log.Log.Info("db init success")

}

func initSchemas() {
	DB.AutoMigrate(model.Category{})
	DB.AutoMigrate(model.File{})
}

// Close 关闭
func Close() {
	if DB == nil {
		return
	}

	if err := DB.Close(); err != nil {
		panic(err)
	}

}

func RollBackIfPanic(db *gorm.DB) {
	ex.Check(db.Rollback().Error)
}

func EscapeLike(colVal string)  string{
	return strings.Replace(strings.Replace(colVal,"%","\\%",-1),"_","\\_",-1)
}