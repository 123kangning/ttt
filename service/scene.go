package service

import (
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"qin/dao"
	"qin/web"
	"strconv"
)

func AddScene(scene *web.Scene) (err error) {
	return dao.AddScene(scene)
}
func AddImageToScene(name string, file multipart.File, sid int) (err error) {
	url := "./public/" + strconv.Itoa(sid) + "_" + name
	f, err := os.OpenFile(url, os.O_WRONLY|os.O_CREATE, 0644)
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
	return dao.AddImageToScene(name, url, sid)
}
