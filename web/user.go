package web

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"qin/service"
)

type BaseResp struct {
	StatusCode    int
	StatusMessage string
}
type User struct {
	Username string
	Password string
}

func SignIn(c *gin.Context) {
	req := &User{}
	if err := c.Bind(req); err != nil {
		log.Println("err = ", err, " req = ", req)
		c.JSON(http.StatusBadRequest, "bind error")
		return
	}
	err := service.SignIn(req)
	resp := &BaseResp{}
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMessage = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}
func Login(c *gin.Context) {
	req := &User{}
	if err := c.Bind(req); err != nil {
		log.Println("err = ", err, " req = ", req)
		c.JSON(http.StatusBadRequest, "bind error")
		return
	}
	err := service.Login(req)
	resp := &BaseResp{}
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMessage = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}
func SignOut(c *gin.Context) {
	req := &User{}
	username, _ := c.Get("username")
	req.Username = username.(string)
	err := service.SignOut(req)
	resp := &BaseResp{}
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMessage = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}
