package controllers

import (
	"nature.ranger.api/constant"
	"nature.ranger.api/models/users"
	"nature.ranger.api/services"
)

type MainController struct {
	BaseController
	IsLogin  bool
	UserInfo users.User
	UserId   int64
}

/**
 * 预处理方法
 * @author pwt
 * @date 2018-11-14
 * @param
 * @return
 */
func (m *MainController) Prepare() {
	token := m.GetString("x-token")
	if token == "" {
		m.JsonResultError(constant.MessageTmpls["common"]["invalid_access_token"])
	}
	if !services.IsValided(token) {
		m.JsonResultError(constant.MessageTmpls["common"]["invalid_access_token"])
	}
	uid := services.GetToken(token)
	user, err := users.GetUser(uid)
	if err != nil {
		m.JsonResultError(constant.MessageTmpls["common"]["invalid_access_token"])
	}
	m.IsLogin = true
	m.UserId = uid
	m.UserInfo = user
}
