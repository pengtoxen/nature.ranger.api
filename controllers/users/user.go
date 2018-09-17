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
func (l *UserController) AddUser() {
	username := l.GetString("username")
	password := l.GetString("password")
	openid := l.GetString("openid")
	nickname := l.GetString("nickname")
	avatar := l.GetString("avatar")
	phone := l.GetString("phone")
	user := &(users.User{
		Username: username,
		Password: password,
		Openid:   openid,
		Nickname: nickname,
		Avatar:   avatar,
		Phone:    phone,
	})
	//字段验证
	if ret, msg := l.FieldValidate(user); !ret {
		l.JsonResultError(msg)
		return
	}
	//名称唯一性验证
	if !users.IsUserNameUnique(username, 0) {
		msg := "用户名" + constant.MessageTmpls["message"]["duplicate"]
		l.JsonResultError(msg)
		return
	}
	//手机号唯一性验证
	if !users.IsPhoneUnique(phone, 0) {
		msg := "手机号" + constant.MessageTmpls["message"]["duplicate"]
		l.JsonResultError(msg)
		return
	}
	uid, err := users.AddUser(user)
	if err != nil {
		l.JsonResultError(err.Error())
		return
	}
	userInfo, err := users.GetUser(uid)
	if err != nil {
		l.JsonResultError(err.Error())
		return
	}
	l.JsonResultOk(userInfo)
}

/**
 * 更新用户
 * @author pwt
 * @date 2018-8-21
 * @param
 * @return
 */
func (u *UserController) UpdateUser() {
	uid, _ := u.GetInt64("id", 0)
	if uid == 0 {
		u.JsonResultError(constant.MessageTmpls["common"]["id_is_required"])
		return
	}
	username := u.GetString("username")
	password := u.GetString("password")
	openid := u.GetString("openid")
	nickname := u.GetString("nickname")
	avatar := u.GetString("avatar")
	phone := u.GetString("phone")
	user := &(users.User{
		Id:       uid,
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
	if !users.IsUserNameUnique(username, uid) {
		msg := "用户名" + constant.MessageTmpls["message"]["duplicate"]
		u.JsonResultError(msg)
		return
	}
	//手机号唯一性验证
	if !users.IsPhoneUnique(phone, uid) {
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
	uid, _ := u.GetInt64("id")
	if uid != 0 {
		user := &(users.User{
			Id: uid,
		})
		err := users.DeleteUser(user)
		if err != nil {
			u.JsonResultError(err.Error())
		} else {
			u.JsonResultOk("")
		}
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
	uid, _ := u.GetInt64("uid")
	if uid != 0 {
		user, err := users.GetUser(uid)
		if err != nil {
			u.JsonResultError(err.Error())
		} else {
			u.JsonResultOk(user)
		}
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
