package web

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"qin/model"
	"qin/service"
	"strconv"
)

func AddScene(c *gin.Context) {
	req := &model.Scene{}
	if err := c.Bind(req); err != nil {
		log.Println("err = ", err, " req = ", req)
		c.JSON(http.StatusBadRequest, "bind error")
		return
	}
	err := service.AddScene(req)
	resp := &model.BaseResp{}
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMessage = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}
func AddImageToScene(c *gin.Context) {
	pic, err := c.FormFile("picture")
	if err != nil {
		log.Println("err = ", err)
		log.Println("pic = ", pic)
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
	resp := &model.BaseResp{}
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMessage = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}
func GetScene(c *gin.Context) {
	name := c.Query("Name")
	resp, err := service.GetScene(name)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMessage = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}
func AddComments(c *gin.Context) {
	req := &model.Comment{}
	if err := c.Bind(req); err != nil {
		log.Println("err = ", err, " req = ", req)
		c.JSON(http.StatusBadRequest, "bind error")
		return
	}
	u, _ := c.Get("username")
	req.Username = u.(string)
	err := service.AddComments(req)
	resp := &model.BaseResp{}
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMessage = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}
func GetCommon(c *gin.Context) {
	resp := &model.GetCommentResponse{}
	var err error
	resp.Comments, err = service.GetCommon()
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMessage = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}
