package dao

import (
	"errors"
	"qin/web"
)

type scene struct {
	id          int
	name        string
	description string
}
type images struct {
	id   int
	sid  int
	name string
	url  string
}

func AddScene(sceneInfo *web.Scene) (err error) {
	s := &scene{}
	DB.Where("name = ?", sceneInfo.Name).Limit(1).Find(s)
	if s.id > 0 {
		return errors.New("添加失败，景点名称重复")
	}
	s.name = sceneInfo.Name
	s.description = sceneInfo.Description
	return DB.Create(s).Error
}
func AddImageToScene(url, name string, sid int) (err error) {
	i := &images{}
	DB.Where("sid = ? and url = ?", sid, url).Limit(1).Find(i)
	if i.id > 0 {
		return errors.New("添加失败，相同图片已经上传")
	}
	i.sid = sid
	i.name = name
	i.url = url
	return DB.Create(i).Error
}
