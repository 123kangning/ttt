package service

import (
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"qin/configs/consts"
	"qin/dao"
	"qin/model"
	"strconv"
)

func AddScene(scene *model.Scene) (err error) {
	return dao.AddScene(scene)
}
func AddImageToScene(name string, file multipart.File, sid int) (err error) {
	url := consts.PrePath + strconv.Itoa(sid) + "_" + name
	f, err := os.OpenFile(url, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Println("Error creating file: ", err)
		return errors.New(fmt.Sprintln("Error creating file: ", err))
	}
	defer f.Close()
	_, err = io.Copy(f, file)
	if err != nil {
		log.Println("Error writing file to disk: ", err)
		return errors.New(fmt.Sprintln("Error writing file to disk: ", err))
	}
	log.Println("url = ", url)
	return dao.AddImageToScene(url, sid)
}
func GetScene(name string) (resp *model.GetSceneResponse, err error) {
	return dao.GetScene(name)
}
func AddComments(comment *model.Comment) (err error) {
	return dao.AddComments(comment)
}
