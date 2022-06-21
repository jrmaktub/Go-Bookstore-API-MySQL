package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//file returns a variable db to help other files interact with DB
var (
	db *gorm.DB
)

func Connect() {
	//might need to use other DB
	//new db
	d, err := gorm.Open("mysql", "akhil:Axlesharma@12@/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

//help other files talk to DB
func GetDB() *gorm.DB {
	return db
}
