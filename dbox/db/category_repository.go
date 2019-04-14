package db

import (
	"github.com/sdvdxl/dbox/dbox/ex"
	"github.com/sdvdxl/dbox/dbox/model"
)

type categoryRepo struct {
}

func (*categoryRepo) Save(name string) *model.Category {
	var c model.Category
	ex.Check(gdb.Where(model.Category{Name: name}).FirstOrCreate(&c).Error)
	return &c
}
