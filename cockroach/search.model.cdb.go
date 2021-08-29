package cockroach

import (
	"github.com/sonntuet1997/medical-chain-utils/common"
	"gorm.io/gorm"
	"strings"
)

type SearchModelCDB struct {
	*gorm.DB
	common.DefaultSearchModel
}

func (s *SearchModelCDB) ApplyPagination() *SearchModelCDB {
	db := s.DB
	if s.DefaultSearchModel.Limit > 0 {
		db = db.Limit(s.DefaultSearchModel.Limit)
	}
	if s.DefaultSearchModel.Skip > 0 {
		db = db.Offset(s.DefaultSearchModel.Skip)
	}
	return &SearchModelCDB{
		DB:                 db,
		DefaultSearchModel: s.DefaultSearchModel,
	}
}

func (s *SearchModelCDB) ApplySort() *SearchModelCDB {
	db := s.DB
	if s.DefaultSearchModel.OrderBy != "" {
		orderByList := strings.Fields(s.DefaultSearchModel.OrderBy)
		orderTypeList := strings.Fields(s.DefaultSearchModel.OrderType)
		for i, orderBy := range orderByList {
			orderType := "asc"
			if orderTypeList[i] != "asc" {
				orderType = "desc"
			}
			db = db.Order(orderBy + " " + orderType)
		}
	}
	return &SearchModelCDB{
		DB:                 db,
		DefaultSearchModel: s.DefaultSearchModel,
	}
}

func ApplyPaginationCDB(defaultSearchModel common.DefaultSearchModel, db *gorm.DB) *gorm.DB {
	if defaultSearchModel.Limit > 0 {
		db = db.Limit(defaultSearchModel.Limit)
	}
	if defaultSearchModel.Skip > 0 {
		db = db.Offset(defaultSearchModel.Skip)
	}
	return db
}

func ApplySortCDB(defaultSearchModel common.DefaultSearchModel, db *gorm.DB) *gorm.DB {
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
	return db
}
