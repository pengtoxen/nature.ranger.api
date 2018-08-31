package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"nature.ranger.api/constant"
	//"strconv"
	"strings"
)

type BaseController struct {
	beego.Controller
	IsLogin bool
	//UserInfo string
	UserUserId   int64
	UserUsername string
	UserAvatar   string
}

/**
 * 初始化方法
 * @author pwt
 * @date 2018-8-29
 * @param
 * @return
 */
func init() {
	validation.SetDefaultMessage(constant.FieldValidateInfo)
}

/**
 * 字段验证
 * @author pwt
 * @date 2017-9-19
 * @param
 * @return
 */
func (this *BaseController) FieldValidate(obj interface{}) (bool, string) {
	valid := validation.Validation{}
	isValid, err := valid.Valid(obj)
	if err != nil {
		return false, constant.MessageTmpls["message"]["submit_error"]
	}
	if !isValid {
		for _, err := range valid.Errors {
			var keyName string
			if strings.Contains(err.Key, ".") {
				keyName = strings.Split(err.Key, ".")[0]
			} else {
				keyName = ""
			}
			return false, keyName + err.Message
		}
	}
	return true, constant.MessageTmpls["message"]["operate_success"]
}

func (this *BaseController) JsonResult(arg ...interface{}) {
	json := map[string]interface{}{"code": 0, "msg": "操作成功", "data": ""}
	len := len(arg)
	if len == 1 {
		json["data"] = arg[0]
	} else if len == 2 {
		json["code"] = arg[0]
		json["data"] = arg[1]
		if json["code"] == 1 {
			json["msg"] = "操作失败"
		}
	} else if len == 3 {
		json["code"] = arg[0]
		json["msg"] = arg[1]
		json["data"] = arg[2]
	}
	this.Data["json"] = json
	this.ServeJSON()
}

func (this *BaseController) JsonResultOk(data interface{}) {
	this.JsonResult(data)
}

func (this *BaseController) JsonResultError(data interface{}) {
	this.JsonResult(1, data)
}

func (this *BaseController) dump(data interface{}) {
	fmt.Printf("%+v\n", data)
}
