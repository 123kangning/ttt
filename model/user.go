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
	Id, Uid       int
	name, content string
}
