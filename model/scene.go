package model

type Scene struct {
	Id          int
	Name        string
	Description string
}
type Images struct {
	Id  int
	Sid int
	Url string
}
type GetSceneResponse struct {
	BaseResp
	Id          int
	Name        string
	Description string
	Urls        []string
	Comments    []*Comment
}
type Comment struct {
	Id, Uid, Sid int
	Content      string
}
