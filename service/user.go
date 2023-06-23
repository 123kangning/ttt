package service

import (
	"qin/dao"
	"qin/web"
)

func SignIn(user *web.User) (err error) {
	return dao.SignIn(user)
}
func Login(user *web.User) (err error) {
	return dao.Login(user)
}
func SignOut(user *web.User) (err error) {
	return dao.SignOut(user)
}
