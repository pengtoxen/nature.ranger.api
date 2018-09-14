package users

import (
	"nature.ranger.api/constant"
	"nature.ranger.api/controllers"
	"nature.ranger.api/models/users"
	// "nature.ranger.api/utils"
	// "github.com/astaxie/beego"
	// "github.com/astaxie/beego/utils/pagination"
)

// Operations about Users
type UserController struct {
	controllers.BaseController
}

/**
 * 新增用户
 * @author pwt
 * @date 2018-8-21
 * @param
 * @return
 */
func (u *UserController) Login() {
	username := u.GetString("username")
	password := u.GetString("password")
	openid := u.GetString("openid")
	nickname := u.GetString("nickname")
	avatar := u.GetString("avatar")
	phone := u.GetString("phone")
	user := &(users.User{
		Username: username,
		Password: password,
		Openid:   openid,
		Nickname: nickname,
		Avatar:   avatar,
		Phone:    phone,
	})
	//字段验证
	if ret, msg := u.FieldValidate(user); !ret {
		u.JsonResultError(msg)
		return
	}
	//名称唯一性验证
	if !users.IsUserNameUnique(username, 0) {
		msg := "用户名" + constant.MessageTmpls["message"]["duplicate"]
		u.JsonResultError(msg)
		return
	}
	//手机号唯一性验证
	if !users.IsPhoneUnique(phone, 0) {
		msg := "手机号" + constant.MessageTmpls["message"]["duplicate"]
		u.JsonResultError(msg)
		return
	}
	uid, err := users.AddUser(user)
	if err != nil {
		u.JsonResultError(err.Error())
		return
	}
	userInfo, err := users.GetUser(uid)
	if err != nil {
		u.JsonResultError(err.Error())
		return
	}
	u.JsonResultOk(userInfo)
}
