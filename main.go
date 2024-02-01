package main

import (
	"ginBlog/model"
	"ginBlog/routes"
)

func main() {
	//引用数据库
	model.InitDB()

	routes.InitRouter()
}
