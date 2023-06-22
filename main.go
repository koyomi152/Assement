package main

import (
	"ginHomework/api"
	"ginHomework/dao"
)

func main() {
	dao.InitDB()
	api.InitRedis()
	api.InitRouter()
}
