package category

import (
	"github.com/sdvdxl/dbox/dbox/dao"
	"github.com/sdvdxl/dbox/dbox/dao/categoryDao"
	"github.com/sdvdxl/dbox/dbox/model"
)

func FindAll() []model.Category {
	return categoryDao.FindAll()
}

func CreateCategory(name string) *model.Category {
	return categoryDao.CreateCategory(dao.DB, name)

}
