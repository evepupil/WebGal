package main

import (
	"WebGal/dao"
	"WebGal/models"
	"WebGal/routers"
	"fmt"
)

func main() {
	//1. 连接MySql数据库
	fmt.Println("连接MySql数据库......")
	dao.InitMySQL()
	fmt.Println("连接MySql数据库成功")
	//2. 自动迁移  把结构体和数据表进行对应
	dao.DB.AutoMigrate(&models.Todo{})
	// 3.注册路由
	r := routers.SetupRouter()
	r.Run()
}
