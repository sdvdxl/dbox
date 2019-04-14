package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sdvdxl/dbox/dbox/model"
)

// global db or gorm db
var (
	gdb       *gorm.DB
	dbDialect dbType
	dbArgs    []interface{}
	File      = &fileRepo{}
	Category  = &categoryRepo{}
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
	gdb, err = gorm.Open(string(dbDialect), dbArgs...)
	if err != nil {
		panic("failed to connect database, error msg:" + err.Error())
	}

	initSchemas()

	fmt.Println(dbDialect, "db init success")

}

func initSchemas() {
	gdb.AutoMigrate(model.Category{})
	gdb.AutoMigrate(model.File{})
}

// Close 关闭
func Close() {
	if gdb == nil {
		return
	}

	if err := gdb.Close(); err != nil {
		panic(err)
	}

}
