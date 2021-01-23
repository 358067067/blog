package main

import (
	"blog/model"
	"blog/routes"
)

//main 主入口
func main() {
	model.InitDb()
	routes.InitRouter()
}
