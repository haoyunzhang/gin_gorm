package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func init() {
	db, err := gorm.Open("mysql", "root:mahui@/gorm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	DB = db
	//DB.Model(&Role{}).Related(&Perm{})
	DB.AutoMigrate(&Perm{}, &Role{}, &User{})
}

