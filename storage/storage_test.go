package storage

import (
	"Mock-API-Data/config"
	"Mock-API-Data/model"
	"Mock-API-Data/util"
	"testing"
	"time"
)

func TestNewStorage(t *testing.T) {

	confPath := "../cmd/conf.yaml"
	conf, err := config.Load(confPath)
	if err != nil {
		t.Fatal(err)
		t.FailNow()
	}

	storage, err := NewStorage(conf)
	if err != nil {
		t.Fatal(err)
		t.FailNow()
	}
	user := model.User{
		Name:     "测试",
		CreateAt: time.Now(),
	}
	storage.DB().FirstOrCreate(&user)

	project := model.Project{
		Name:               "测试",
		UserId:             user.Id,
		Key:                util.NewProjectKey(user.Id),
		Host:               "http://www.baidu.com",
		InsecureSkipVerify: true,
	}

	storage.DB().FirstOrCreate(&project)
	var newProject model.Project
	storage.DB().Where("key = ?", project.Key).Find(&newProject)
	t.Log(newProject)
}
