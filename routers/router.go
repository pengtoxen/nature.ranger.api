// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"nature.ranger.api/controllers/Login"
	"nature.ranger.api/controllers/users"
)

func init() {
	//login
	beego.Router("/login/login", &users.LoginController{}, "get:Login")
	//user
	beego.Router("/user/add", &users.UserController{}, "get:AddUser")
	beego.Router("/user/delete", &users.UserController{}, "post:DeleteUser")
	beego.Router("/user/update", &users.UserController{}, "post:UpdateUser")
	beego.Router("/user/info", &users.UserController{}, "get:GetOne")
	beego.Router("/user/list", &users.UserController{}, "get:GetAll")
}
