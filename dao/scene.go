package dao

import (
	"errors"
	"log"
	"qin/model"
	"reflect"
)

func AddScene(sceneInfo *model.Scene) (err error) {
	s := &model.Scene{}
	DB.Model(&model.Scene{}).Where("name = ?", sceneInfo.Name).Limit(1).Find(s)
	if s.Id > 0 {
		return errors.New("添加失败，景点名称重复")
	}
	s.Name = sceneInfo.Name
	s.Description = sceneInfo.Description
	return DB.Model(&model.Scene{}).Create(s).Error
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
func GetScene(name string) (resp *model.GetSceneResponse, err error) {
	var s []*model.Scene
	log.Println(name, " &s.type = ", reflect.TypeOf(&s))
	DB.Model(&model.Scene{}).Where("name like ?", "%"+name+"%").Find(&s)
	log.Println(s)
	resp = &model.GetSceneResponse{}
	resp.Scene = make([]*model.Scenes, 0)
	for _, v := range s {
		scenes := &model.Scenes{}
		scenes.Id = v.Id
		scenes.Name = v.Name
		scenes.Description = v.Description
		scenes.Urls, err = GetImages(v.Id)
		if err != nil {
			log.Println("GetImages err = ", err)
			return resp, err
		}
		scenes.Comments, err = GetComment(v.Id)
		if err != nil {
			log.Println("GetComment err = ", err)
			return resp, err
		}
		resp.Scene = append(resp.Scene, scenes)
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
