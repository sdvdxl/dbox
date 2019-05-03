package category

import (
	"github.com/sdvdxl/dbox/api/dao"
	"github.com/sdvdxl/dbox/api/dao/categoryDao"
	"github.com/sdvdxl/dbox/api/model"
)

func FindAll() []model.Category {
	return categoryDao.FindAll()
}

func CreateCategory(name string) *model.Category {
	return categoryDao.CreateCategory(dao.DB, name)

}
