package service

import (
	"github.com/sdvdxl/dbox/api/dao"
	"github.com/sdvdxl/dbox/api/model"
)

type Category struct {
}

func (*Category) FindAll() []model.Category {
	return categoryDao().FindAll()
}

func (*Category) CreateCategory(name string) *model.Category {
	return categoryDao().CreateCategory(name)
}

func categoryDao() *dao.CategoryDao {
	db := dao.NewDB()
	return &dao.CategoryDao{DB: db}
}
