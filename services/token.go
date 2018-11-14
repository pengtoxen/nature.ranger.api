package services

import (
	"github.com/astaxie/beego/cache"
	"nature.ranger.api/utils"
	"strconv"
	"time"
)

var cacheObj cache.Cache

func init() {
	if cacheObj == nil {
		cacheObj, _ = cache.NewCache("file", `{"CachePath":"./tmp/cache","FileSuffix":".cache","DirectoryLevel":2,"EmbedExpiry":3600}`)
	}
}

func GenerateToken(uid int64) (string, error) {
	tokenStr := tokenString(uid)
	err := cacheObj.Put(tokenStr, uid, 3600*time.Second)
	if err == nil {
		return tokenStr, nil
	}
	return "", err
}

func GetToken(tokenStr string) int64 {
	return cacheObj.Get(tokenStr).(int64)
}

func IsValided(tokenStr string) bool {
	return cacheObj.IsExist(tokenStr)
}

func Delete(tokenStr string) error {
	return cacheObj.Delete(tokenStr)
}

// func expireTime() int64 {
// 	t := time.Now()
// 	timestamp := t.UTC().UnixNano() + 1*24*3600
// 	return timestamp
// }

func tokenString(uid int64) string {
	var str []string
	str = append(str, strconv.FormatInt(uid, 10))
	t := time.Now()
	timestamp := strconv.FormatInt(t.UTC().UnixNano(), 10)
	str = append(str, timestamp)
	ret := utils.StrConcat(str)
	return utils.Sha1(ret)
}
