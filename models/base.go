package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func TableName(name string) string {
	return beego.AppConfig.String("mysqlpre") + name
}

func TablePre() string {
	return beego.AppConfig.String("mysqlpre")
}

func Pagination(pagi []int) (offset int, start int) {
	var len = len(pagi)
	var page int = 1
	if len == 0 {
		offset, _ = beego.AppConfig.Int("pageoffset")
	} else if len == 1 {
		page = pagi[0]
		offset, _ = beego.AppConfig.Int("pageoffset")
	} else {
		page = pagi[0]
		offset = pagi[1]
	}
	start = (page - 1) * offset
	return offset, start
}

func QueryCond(condArr map[string]string) (cond *orm.Condition) {
	cond = orm.NewCondition()
	for k, v := range condArr {
		cond = cond.And(k, v)
	}
	return cond
}
