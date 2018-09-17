package login

import (
	"nature.ranger.api/constant"
	"nature.ranger.api/controllers"
	"nature.ranger.api/models/users"
	// "nature.ranger.api/utils"
)

// Operations about Login
type LoginController struct {
	controllers.BaseController
}

func (l *LoginController) Login() {
	username := l.GetString("username")
	password := l.GetString("password")
	condArr := make(map[string]string)
	condArr["username"] = username
	condArr["password"] = password
	userInfo, err := users.Login(condArr)
	if err != nil {
		l.JsonResultError(constant.MessageTmpls["common"]["username_or_password_incorrect"])
		return
	}
	l.JsonResultOk(userInfo)
}
