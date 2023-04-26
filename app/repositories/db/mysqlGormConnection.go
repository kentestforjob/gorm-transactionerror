package db

import (
	"fmt"

	"gorm.io/driver/mysql"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

// var db *gorm.DB

func ConnectMysqlGormDatabase() (db *gorm.DB) {

	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetInt(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&collation=utf8mb4_unicode_520_ci&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)

	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		fmt.Println("connection to mysql failed:", err)
		panic("exit")

	}

	MysqlMigration(gormDB)

	return gormDB
}
