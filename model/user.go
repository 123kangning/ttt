package model

type User struct {
	Id       int
	Username string
	Password string
}
type BaseResp struct {
	StatusCode    int
	StatusMessage string
}
type UserLoginResponse struct {
	BaseResp
	Token string
}
type Journey struct {
	Id                      int
	Username, Name, Content string
}
type GetJourneys struct {
	BaseResp
	Journeys []*Journey
}
