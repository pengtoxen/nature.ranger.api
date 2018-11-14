package users

import (
	"nature.ranger.api/constant"
	"nature.ranger.api/controllers"
	"nature.ranger.api/models/users"
)

// Operations about Users
type UserController struct {
	controllers.MainController
}

/**
 * 更新用户
 * @author pwt
 * @date 2018-8-21
 * @param
 * @return
 */
func (u *UserController) UpdateUser() {
	username := u.GetString("username")
	password := u.GetString("password")
	openid := u.GetString("openid")
	nickname := u.GetString("nickname")
	avatar := u.GetString("avatar")
	phone := u.GetString("phone")
	user := &(users.User{
		Id:       u.UserId,
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
	if !users.IsUserNameUnique(username, u.UserId) {
		msg := "用户名" + constant.MessageTmpls["message"]["duplicate"]
		u.JsonResultError(msg)
		return
	}
	//手机号唯一性验证
	if !users.IsPhoneUnique(phone, u.UserId) {
		msg := "手机号" + constant.MessageTmpls["message"]["duplicate"]
		u.JsonResultError(msg)
		return
	}
	err := users.UpdateUser(user)
	if err != nil {
		u.JsonResultError(err.Error())
		return
	}
	u.JsonResultOk("")
}

/**
 * 删除用户
 * @author pwt
 * @date 2018-8-21
 * @param
 * @return
 */
func (u *UserController) DeleteUser() {
	user := &(users.User{
		Id: u.UserId,
	})
	err := users.DeleteUser(user)
	if err != nil {
		u.JsonResultError(err.Error())
	} else {
		u.JsonResultOk("")
	}
	u.JsonResultError(constant.MessageTmpls["common"]["id_is_required"])
}

/**
 * 根据uid获取用户信息
 * @author pwt
 * @date 2018-8-17
 * @param
 * @return
 */
func (u *UserController) GetOne() {
	user, err := users.GetUser(u.UserId)
	if err != nil {
		u.JsonResultError(err.Error())
	} else {
		u.JsonResultOk(user)
	}
	u.JsonResultError(constant.MessageTmpls["common"]["id_is_required"])
}

/**
 * 根据获取所有用户信息
 * @author pwt
 * @date 2018-8-17
 * @param
 * @return
 */
func (u *UserController) GetAll() {
	keywords := u.GetString("keywords")
	phone := u.GetString("phone")
	page, _ := u.GetInt("page", 1)
	condArr := make(map[string]string)
	condArr["keywords"] = keywords
	condArr["phone"] = phone
	users, _, err := users.GetAllUser(condArr, page)
	if err != nil {
		u.JsonResultError(err.Error())
	}
	u.JsonResultOk(users)
}
