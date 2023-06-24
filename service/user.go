package service

import (
	"qin/dao"
	"qin/model"
)

func SignIn(user *model.User) (err error) {
	return dao.SignIn(user)
}
func Login(user *model.User) (err error) {
	return dao.Login(user)
}
func SignOut(user *model.User) (err error) {
	return dao.SignOut(user)
}
