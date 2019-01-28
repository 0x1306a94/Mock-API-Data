package util

import (
	"Mock-API-Data/model"
	"time"

	"testing"
)

func TestGenerateAuthorizationToken(t *testing.T) {

	tt := time.Now()
	user := model.User{
		Id:       1,
		Name:     "king",
		Email:    "0x1306a94@gmail.com",
		CreateAt: tt,
		UpdateAt: tt,
	}

	tokenStr, err := GenerateAuthorizationToken(user)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(tokenStr)
	}
}

func TestParseUserWithToken(t *testing.T) {

	tt := time.Now()
	user := model.User{
		Id:       1,
		Name:     "king",
		Email:    "0x1306a94@gmail.com",
		CreateAt: tt,
		UpdateAt: tt,
	}
	t.Log(user)
	tokenStr, err := GenerateAuthorizationToken(user)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(tokenStr)
	}

	newUser, err := ParseUserWithToken(tokenStr)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(newUser)
	}
}
