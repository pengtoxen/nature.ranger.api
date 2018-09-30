package services

import (
	// "github.com/astaxie/beego/context"
	"github.com/astaxie/beego/session"
	"nature.ranger.api/utils"
	"strconv"
	"time"
)

var cruSession session.Store

type Token struct {
	Token  string `json:"token"`
	Expire int64  `json:"expire"`
}

func InitData(handler session.Store) {
	if cruSession == nil {
		cruSession = handler
	}
}

func GenerateToken(uid int64) error {
	tokenstr := tokenString(uid)
	token := Token{
		Token:  tokenstr,
		Expire: expireTime(),
	}
	return cruSession.Set("access_token", token)
}

func GetToken() Token {
	return cruSession.Get("access_token").(Token)
}

func IsValided(tokenStr string) bool {
	token := GetToken()
	if token.Token != tokenStr {
		return false
	}
	return true
}

func IsExpired() bool {
	token := GetToken()
	expire := token.Expire
	t := time.Now()
	now := t.UTC().UnixNano()
	if expire < now {
		return false
	}
	return true
}

func Delete() error {
	return cruSession.Delete("access_token")
}

func expireTime() int64 {
	t := time.Now()
	timestamp := t.UTC().UnixNano() + 1*24*3600
	return timestamp
}

func tokenString(uid int64) string {
	var str []string
	str = append(str, strconv.FormatInt(uid, 10))
	t := time.Now()
	timestamp := strconv.FormatInt(t.UTC().UnixNano(), 10)
	str = append(str, timestamp)
	ret := utils.StrConcat(str)
	return utils.Sha1(ret)
}
