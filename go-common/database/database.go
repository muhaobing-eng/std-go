package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func New(opt *Option) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(opt.GetURL()))
}
