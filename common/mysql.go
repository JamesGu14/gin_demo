package common

import (
	"gin_demo/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectMySQL() *gorm.DB {
	database, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:root@tcp(127.0.0.1:3306)/gin_demo?charset=utf8mb4&parseTime=True&loc=Local", // data source name
		DefaultStringSize:         256,                                                                               // default size for string fields
		DisableDatetimePrecision:  true,                                                                              // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,                                                                              // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,                                                                              // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,                                                                             // auto configure based on currently MySQL version
	}), &gorm.Config{})

	if err != nil {
		panic("MySQL connection failed")
	}
	DB = database
	return DB
}

// Using this function to get a connection, you can create your connection pool here.
func GetDB() *gorm.DB {
	return DB
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&entities.Teacher{})
	db.AutoMigrate(&entities.Student{})
	db.AutoMigrate(&entities.Class{})
	db.AutoMigrate(&entities.Course{})
}
