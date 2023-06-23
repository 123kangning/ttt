package dao

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
	"qin/web"
)

type user struct {
	id       int
	username string
	password string
}

func getPassword(Password string) []byte {
	hash, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("bcrypt.GenerateFromPassword error")
		return nil
	}
	return hash
}
func PasswordEqual(Password1, Password2 string) error {
	//hash, _ := bcrypt.GenerateFromPassword([]byte(Password1), bcrypt.DefaultCost)
	err := bcrypt.CompareHashAndPassword([]byte(Password2), []byte(Password1))
	if err == nil { //密码正确
		return nil
	}
	return err
}
func SignIn(userInfo *web.User) (err error) {
	u := &user{}
	DB.Where("username = ?", userInfo.Username).Limit(1).Find(u)
	if u.id > 0 {
		return errors.New("注册失败，用户名重复")
	}
	u.username = userInfo.Username
	u.password = string(getPassword(userInfo.Password))
	return DB.Create(u).Error
}
func Login(userInfo *web.User) (err error) {
	u := &user{}
	DB.Where("username = ?", userInfo.Username).Limit(1).Find(u)
	if u.id > 0 {
		if PasswordEqual(userInfo.Password, u.password) == nil {
			return nil
		}
		return errors.New("密码错误")
	}
	return errors.New("用户名错误")
}
func SignOut(userInfo *web.User) (err error) {
	return DB.Where("username = ?", userInfo.Username).Delete(&user{}).Error
}
