package service

import (
	"github.com/sdvdxl/dbox/api/dao"
	"github.com/sdvdxl/dbox/api/model"
)

func FindAll() []model.Category {
	return categoryDao().FindAll()
}

func CreateCategory(name string) *model.Category {
	return categoryDao().CreateCategory(name)
}

func categoryDao() *dao.CategoryDao {
	db := dao.NewDB()
	return &dao.CategoryDao{DB: db}
}
