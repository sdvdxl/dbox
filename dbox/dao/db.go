package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sdvdxl/dbox/dbox/ex"
	"github.com/sdvdxl/dbox/dbox/log"
	"github.com/sdvdxl/dbox/dbox/model"
)

// global db or gorm db
var (
	DB        *gorm.DB
	dbDialect dbType
	dbArgs    []interface{}
)

type dbType string

const (
	// DialectSqlite3 sqlite3
	DialectSqlite3 = dbType("sqlite3")
	// DialectMysql mysql
	DialectMysql = dbType("mysql")
)

func Use(dialect dbType, args ...interface{}) {
	dbDialect = dialect
	dbArgs = args
}

func Init() {
	var err error
	DB, err = gorm.Open(string(dbDialect), dbArgs...)
	if err != nil {
		panic("failed to connect database, error msg:" + err.Error())
	}

	DB.SetLogger(log.Logger(*log.Log))
	DB.Debug()
	DB.LogMode(true)

	initSchemas()

	log.Log.Info(dbDialect, "db init success")

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

func Transaction(session *gorm.DB) {
	if err := recover(); err != nil {
		ex.Check(session.Rollback().Error)
		return
	}

	ex.Check(session.Commit().Error)
	log.Log.Debug("commit success")
}
