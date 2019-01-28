package util

import (
	"Mock-API-Data/model"
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type tokenClaims struct {
	UserId     int64
	UserName   string
	Admin      bool
	Management bool
	Email      string
	CreateAt   time.Time
	UpdateAt   time.Time
	jwt.StandardClaims
}

var tokenSecret = []byte("mock-secret")

func GenerateAuthorizationToken(user model.User) (string, error) {
	expiresAt := time.Now().Add(time.Hour * (24 * 10)).UnixNano()
	claims := &tokenClaims{
		UserId:     user.Id,
		UserName:   user.Name,
		Admin:      user.Admin,
		Management: user.Management,
		Email:      user.Email,
		CreateAt:   user.CreateAt,
		UpdateAt:   user.UpdateAt,
	}
	claims.ExpiresAt = expiresAt

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(tokenSecret)
}

func ParseUserWithToken(tokenStr string) (*model.User, error) {
	var claims tokenClaims
	token, err := jwt.ParseWithClaims(tokenStr, &claims, func(*jwt.Token) (interface{}, error) {
		return tokenSecret, nil
	})
	if claims, ok := token.Claims.(*tokenClaims); ok && token.Valid {
		if time.Now().UnixNano() >= claims.ExpiresAt {
			return nil, errors.New("token 已过期")
		}
		user := &model.User{
			Id:         claims.UserId,
			Name:       claims.UserName,
			Admin:      claims.Admin,
			Management: claims.Management,
			Email:      claims.Email,
			CreateAt:   claims.CreateAt,
			UpdateAt:   claims.UpdateAt,
		}
		return user, nil
	}
	return nil, err
}
