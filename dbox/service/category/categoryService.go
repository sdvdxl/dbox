package category

import (
	"github.com/sdvdxl/dbox/dbox/dao/categoryDao"
	"github.com/sdvdxl/dbox/dbox/model"
)

func FindAll() []model.Category {
	return categoryDao.FindAll()
}
