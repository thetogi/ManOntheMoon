package model

import (
	"ManOnTheMoonReviewService/db"
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v8"
	"time"
)

type User struct {
	Password string `json:"password"`
	Username string `json:"username"`
	UserId   string `json:"userId"`
}

func (u *User) Logout(c context.Context, token *jwt.Token) error {
	err := db.RCache.Set(c, token.Raw, "LOGGED_OUT", time.Hour*72).Err()
	return err
}

func (u *User) IsTokenBlacklisted(c context.Context, token string) bool {
	_, err := db.RCache.Get(c, token).Result()
	if err == redis.Nil {
		return false
	}
	return true
}
