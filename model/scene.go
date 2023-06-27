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
	Urls []string
}
type GetSceneResponse struct {
	BaseResp
	Scene []*Scenes
}
type Comment struct {
	Id                int
	Username, Content string
}
type GetCommentResponse struct {
	BaseResp
	Comments []*Comment
}
