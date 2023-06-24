package dao

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
	"qin/model"
)

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
func SignIn(userInfo *model.User) (err error) {
	u := &model.User{}
	DB.Where("username = ?", userInfo.Username).Limit(1).Find(u)
	if u.Id > 0 {
		return errors.New("注册失败，用户名重复")
	}
	u.Username = userInfo.Username
	u.Password = string(getPassword(userInfo.Password))
	return DB.Create(u).Error
}
func Login(userInfo *model.User) (err error) {
	u := &model.User{}
	DB.Where("username = ?", userInfo.Username).Limit(1).Find(u)
	if u.Id > 0 {
		if PasswordEqual(userInfo.Password, u.Password) == nil {
			return nil
		}
		return errors.New("密码错误")
	}
	return errors.New("用户名错误")
}
func SignOut(userInfo *model.User) (err error) {
	return DB.Where("username = ?", userInfo.Username).Delete(&model.User{}).Error
}
func AddJourney(journey *model.Journey) (err error) {
	j := model.Journey{}
	DB.Where("username = ? and name = ?", journey.Username, journey.Name).Limit(1).Find(&j)
	if j.Id > 0 {
		return errors.New("添加失败，有相同规划存在")
	}
	return DB.Create(journey).Error
}
func GetJourneys(username string) (journeys []string, err error) {
	var js []*model.Journey
	err = DB.Where("username = ?", username).Find(&js).Error
	journeys = make([]string, 0)
	for _, v := range js {
		journeys = append(journeys, v.Name)
	}
	return journeys, err
}
func GetJourney(username, name string) (journey *model.Journey, err error) {
	err = DB.Where("username = ? and name = ?", username, name).Find(&journey).Error
	return journey, err
}
