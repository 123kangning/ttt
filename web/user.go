package web

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"qin/model"
	"qin/pkg/jwt"
	"qin/service"
)

func SignIn(c *gin.Context) {
	req := &model.User{}
	if err := c.Bind(req); err != nil {
		log.Println("err = ", err, " req = ", req)
		c.JSON(http.StatusBadRequest, "bind error")
		return
	}
	err := service.SignIn(req)
	resp := &model.BaseResp{}
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMessage = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}
func Login(c *gin.Context) {
	req := &model.User{}
	if err := c.Bind(req); err != nil {
		log.Println("err = ", err, " req = ", req)
		c.JSON(http.StatusBadRequest, "bind error")
		return
	}
	err := service.Login(req)
	resp := &model.UserLoginResponse{}
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMessage = err.Error()
	} else {
		resp.Token, _ = jwt.GenToken(req.Username)
	}
	c.JSON(http.StatusOK, resp)
}
func SignOut(c *gin.Context) {
	req := &model.User{}
	username, _ := c.Get("username")
	req.Username = username.(string)
	err := service.SignOut(req)
	resp := &model.BaseResp{}
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMessage = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}
func AddJourney(c *gin.Context) {
	req := &model.Journey{}
	if err := c.Bind(req); err != nil {
		log.Println("err = ", err, " req = ", req)
		c.JSON(http.StatusBadRequest, "bind error")
		return
	}
	u, _ := c.Get("username")
	req.Username = u.(string)
	err := service.AddJourney(req)
	resp := &model.BaseResp{}
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMessage = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}
func GetJourneys(c *gin.Context) {
	u, _ := c.Get("username")
	username := u.(string)
	js, err := service.GetJourneys(username)
	resp := &model.GetJourneys{}
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMessage = err.Error()
	}
	resp.Journeys = js
	c.JSON(http.StatusOK, resp)
}
func GetJourney(c *gin.Context) {
	u, _ := c.Get("username")
	username := u.(string)
	name := c.Query("Name")
	js, err := service.GetJourney(username, name)
	resp := &model.GetJourney{}
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMessage = err.Error()
	}
	resp.Journey = js
	c.JSON(http.StatusOK, resp)
}
