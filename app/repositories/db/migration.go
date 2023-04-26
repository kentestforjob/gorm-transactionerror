package db

import (
	"test/gormtransactionerr/app/domains"

	"gorm.io/gorm"
)

func MysqlMigration(db *gorm.DB) {

	db.AutoMigrate(
		&domains.Dummy{},
	)
}
