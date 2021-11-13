package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var (
	DB *gorm.DB
)
func InitMySQL()  {
	dsn:= "t:1766468434@(47.101.194.177:3306)/webgal?charset=utf8&parseTime=True&loc=Local"
	Db, err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err!= nil{
		panic(err)
	}
	fmt.Println("初始化数据库成功......")
	DB=Db
}
