package main

import (
	"qin/configs/consts"
	"qin/dao"
	"qin/pkg/jwt"
)

func main() {
	jwt.InitRedis()
	dao.Init()
	r := InitRouter()
	err := r.Run("127.0.0.1:" + consts.WebServerPort)
	if err != nil {
		panic(err)
	}
}
