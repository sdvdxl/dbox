package db

import (
	"github.com/sdvdxl/dbox/dbox/ex"
	"github.com/sdvdxl/dbox/dbox/model"
)

type fileRepo struct {
}

func (*fileRepo) Save(file *model.File) {
	ex.Check(gdb.Create(file).Error)
}
