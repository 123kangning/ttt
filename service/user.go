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
func AddJourney(journey *model.Journey) (err error) {
	return dao.AddJourney(journey)
}
func GetJourneys(username string) (journeys []*model.Journey, err error) {
	return dao.GetJourneys(username)
}
