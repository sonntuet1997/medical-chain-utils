package cockroach

import (
	"strings"

	"github.com/sonntuet1997/medical-chain-utils/common"
	"gorm.io/gorm"
)

type ExtendGorm struct {
	gorm.DB
}

func (s *ExtendGorm) ApplyPagination(defaultSearchModel common.DefaultSearchModel) *ExtendGorm {
	db := &s.DB
	if defaultSearchModel.Limit > 0 {
		db = db.Limit(defaultSearchModel.Limit)
	}
	if defaultSearchModel.Skip > 0 {
		db = db.Offset(defaultSearchModel.Skip)
	}
	return &ExtendGorm{
		DB: *db,
	}
}

func (s *ExtendGorm) ApplySort(defaultSearchModel common.DefaultSearchModel) *ExtendGorm {
	db := &s.DB
	if defaultSearchModel.OrderBy != "" {
		orderByList := strings.Fields(defaultSearchModel.OrderBy)
		orderTypeList := strings.Fields(defaultSearchModel.OrderType)
		for i, orderBy := range orderByList {
			orderType := "asc"
			if orderTypeList[i] != "asc" {
				orderType = "desc"
			}
			db = db.Order(orderBy + " " + orderType)
		}
	}
	return &ExtendGorm{
		DB: *db,
	}
}

func (s *ExtendGorm) ApplyFiled(defaultSearchModel common.DefaultSearchModel) *ExtendGorm {
	db := &s.DB
	if defaultSearchModel.Fields != nil || len(defaultSearchModel.Fields) > 0 {
		db = db.Select(defaultSearchModel.Fields)
	}
	return &ExtendGorm{
		DB: *db,
	}
}

//
//func ApplyPaginationCDB(defaultSearchModel common.DefaultSearchModel, db *gorm.DB) *gorm.DB {
//	if defaultSearchModel.Limit > 0 {
//		db = db.Limit(defaultSearchModel.Limit)
//	}
//	if defaultSearchModel.Skip > 0 {
//		db = db.Offset(defaultSearchModel.Skip)
//	}
//	return db
//}
//
//func ApplySortCDB(defaultSearchModel common.DefaultSearchModel, db *gorm.DB) *gorm.DB {
//	if defaultSearchModel.OrderBy != "" {
//		orderByList := strings.Fields(defaultSearchModel.OrderBy)
//		orderTypeList := strings.Fields(defaultSearchModel.OrderType)
//		for i, orderBy := range orderByList {
//			orderType := "asc"
//			if orderTypeList[i] != "asc" {
//				orderType = "desc"
//			}
//			db = db.Order(orderBy + " " + orderType)
//		}
//	}
//	return db
//}
