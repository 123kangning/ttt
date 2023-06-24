package dao

import (
	"errors"
	"log"
	"qin/model"
)

func AddScene(sceneInfo *model.Scene) (err error) {
	s := &model.Scene{}
	DB.Where("name = ?", sceneInfo.Name).Limit(1).Find(s)
	if s.Id > 0 {
		return errors.New("添加失败，景点名称重复")
	}
	s.Name = sceneInfo.Name
	s.Description = sceneInfo.Description
	return DB.Model("scenes").Create(s).Error
}
func AddImageToScene(url string, sid int) (err error) {
	i := &model.Images{}
	DB.Where("sid = ? and url = ?", sid, url).Limit(1).Find(i)
	if i.Id > 0 {
		return errors.New("添加失败，相同图片已经上传")
	}
	i.Sid = sid
	i.Url = url
	return DB.Create(i).Error
}
func GetScene(sid int) (resp *model.GetSceneResponse, err error) {
	s := model.Scene{}
	DB.Where("id = ?", sid).Find(&s)
	log.Println(s)
	resp = &model.GetSceneResponse{}
	if s.Id == 0 {
		return resp, errors.New("没有该景区")
	}
	resp.Id = s.Id
	resp.Name = s.Name
	resp.Description = s.Description
	resp.Urls, err = GetImages(sid)
	if err != nil {
		log.Println("GetImages err = ", err)
		return resp, err
	}
	resp.Comments, err = GetComment(sid)
	if err != nil {
		log.Println("GetComment err = ", err)
		return resp, err
	}
	return resp, nil
}
func GetImages(sid int) (urls []string, err error) {
	var images []model.Images
	DB.Where("sid = ?", sid).Find(&images)
	urls = make([]string, 0)
	for _, v := range images {
		urls = append(urls, v.Url)
	}
	log.Println(urls)
	return urls, nil
}
func GetComment(sid int) (comments []*model.Comment, err error) {
	DB.Where("sid = ?", sid).Find(&comments)
	return comments, nil
}
func AddComments(comment *model.Comment) error {
	return DB.Create(comment).Error
}
