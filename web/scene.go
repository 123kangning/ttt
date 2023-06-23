package web

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"qin/service"
	"strconv"
)

type Scene struct {
	Name        string
	Description string
}

func AddScene(c *gin.Context) {
	req := &Scene{}
	if err := c.Bind(req); err != nil {
		log.Println("err = ", err, " req = ", req)
		c.JSON(http.StatusBadRequest, "bind error")
		return
	}
	err := service.AddScene(req)
	resp := &BaseResp{}
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMessage = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}
func AddImageToScene(c *gin.Context) {
	pic, err := c.FormFile("file")
	if err != nil {
		log.Println("err = ", err)
		c.JSON(http.StatusBadRequest, "FromFile error")
		return
	}
	name := pic.Filename
	file, err := pic.Open()
	defer file.Close()
	if err != nil {
		log.Println("open pic error = ", err)
		c.JSON(http.StatusBadRequest, "open pic error")
		return
	}
	sid, _ := strconv.Atoi(c.PostForm("sid"))
	err = service.AddImageToScene(name, file, sid)
	resp := &BaseResp{}
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMessage = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}
