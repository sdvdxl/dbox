package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sdvdxl/dbox/api/ex"
	"github.com/sdvdxl/dbox/api/log"
	"github.com/sdvdxl/dbox/api/model"
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
	DB, err = gorm.Open("sqlite3")
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

func RollBackIfPanic(tx *gorm.DB) {

}
