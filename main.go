package main

import (
	"qin/configs/consts"
	"qin/dao"
)

func main() {
	dao.Init()
	r := InitRouter()
	err := r.Run("0.0.0.0:" + consts.WebServerPort)
	if err != nil {
		panic(err)
	}
}
