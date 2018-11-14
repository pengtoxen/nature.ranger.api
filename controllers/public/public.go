package public

import (
	"nature.ranger.api/constant"
	"nature.ranger.api/controllers"
	"nature.ranger.api/models/users"
	"nature.ranger.api/services"
	"nature.ranger.api/utils"
)

// Operations about Login
type PublicController struct {
	controllers.BaseController
}

//login
func (this *PublicController) Login() {
	username := this.GetString("username")
	password := this.GetString("password")
	condArr := make(map[string]string)
	condArr["username"] = username
	condArr["password"] = utils.Sha1(password)
	userInfo, err := users.Login(condArr)
	if err != nil {
		this.JsonResultError(constant.MessageTmpls["common"]["username_or_password_incorrect"])
		return
	}
	if err != nil {
		this.JsonResultError(constant.MessageTmpls["common"]["login_fail"])
		return
	}
	var tokenStr string
	tokenStr, err = services.GenerateToken(userInfo.Id)
	if err != nil {
		this.JsonResultError(constant.MessageTmpls["common"]["login_fail"])
		return
	}
	this.JsonResultOk(tokenStr)
}

/**
 * 新增用户
 * @author pwt
 * @date 2018-8-21
 * @param
 * @return
 */
func (this *PublicController) Regist() {
	username := this.GetString("username")
	password := this.GetString("password")
	openid := this.GetString("openid")
	nickname := this.GetString("nickname")
	avatar := this.GetString("avatar")
	phone := this.GetString("phone")
	user := &(users.User{
		Username: username,
		Password: password,
		Openid:   openid,
		Nickname: nickname,
		Avatar:   avatar,
		Phone:    phone,
	})
	//字段验证
	if ret, msg := this.FieldValidate(user); !ret {
		this.JsonResultError(msg)
		return
	}
	//名称唯一性验证
	if !users.IsUserNameUnique(username, 0) {
		msg := "用户名" + constant.MessageTmpls["message"]["duplicate"]
		this.JsonResultError(msg)
		return
	}
	//手机号唯一性验证
	if !users.IsPhoneUnique(phone, 0) {
		msg := "手机号" + constant.MessageTmpls["message"]["duplicate"]
		this.JsonResultError(msg)
		return
	}
	user.Password = utils.Sha1(user.Password)
	uid, err := users.AddUser(user)
	if err != nil {
		this.JsonResultError(err.Error())
		return
	}
	var tokenStr string
	tokenStr, err = services.GenerateToken(uid)
	if err != nil {
		this.JsonResultError(constant.MessageTmpls["common"]["login_fail"])
		return
	}
	this.JsonResultOk(tokenStr)
}
