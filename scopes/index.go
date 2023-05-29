package scopes

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Paginate(page, size int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		offset := (page - 1) * size

		return db.Offset(offset).Limit(size)
	}
}

func Include(model string, omit ...string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Preload(model, func(db *gorm.DB) *gorm.DB {
			return db.Omit(omit...)
		})
	}
}

func Exclude(omit ...string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Omit(omit...)
	}
}

func Search(search string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("name LIKE ?", search)
	}
}

func Sort(sort, order string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(clause.OrderByColumn{
			Column: clause.Column{Name: sort},
			Desc:   order == "desc",
		})
	}
}

func Count(count *int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(-1).Limit(-1).Count(count)
	}
}
