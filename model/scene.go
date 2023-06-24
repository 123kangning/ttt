package model

type Scene struct {
	Id                int
	Name, Description string
}
type Images struct {
	Id  int
	Sid int
	Url string
}
type Scenes struct {
	Scene
	Urls     []string
	Comments []*Comment
}
type GetSceneResponse struct {
	BaseResp
	Scene []*Scenes
}
type Comment struct {
	Id, Sid           int
	Username, Content string
}
