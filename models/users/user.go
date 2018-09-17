package users

import (
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"nature.ranger.api/models"
)

var (
	tableName string = "user"
)

type User struct {
	Id       int64  `orm:"pk;" json:"id"`
	Username string `json:"username" valid:"Required;AlphaDash"`
	Password string `json:"password" valid:"Required;AlphaDash"`
	Openid   string `json:"openid"`
	Nickname string `json:"nickname" valid:"Required;AlphaDash"`
	Avatar   string `json:"avatar"`
	Phone    string `json:"phone" valid:"Required;Phone"`
}

func init() {
	orm.RegisterModelWithPrefix(models.TablePre(), new(User))
}

func AddUser(user *User) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(user)
	return id, err
}

func UpdateUser(user *User) error {
	o := orm.NewOrm()
	_, err := o.Update(user, "username", "password", "nickname", "avatar", "phone")
	return err
}

func DeleteUser(user *User) error {
	o := orm.NewOrm()
	_, err := o.Delete(user)
	return err
}

/**
 * 手机号唯一性验证
 * @author pwt
 * @date 2018-8-17
 * @param
 * @return
 */
func IsPhoneUnique(phone string, id int64) bool {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName(tableName))
	qs = qs.Filter("phone", phone)
	if id > 0 {
		qs = qs.Filter("id", id)
	}
	cnt, _ := qs.Count()
	if cnt > 0 {
		return false
	}
	return true
}

/**
 * 用户名唯一性验证
 * @author pwt
 * @date 2018-8-29
 * @param
 * @return
 */
func IsUserNameUnique(username string, id int64) bool {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName(tableName))
	qs = qs.Filter("phone", username)
	if id > 0 {
		qs = qs.Filter("id", id)
	}
	cnt, _ := qs.Count()
	if cnt > 0 {
		return false
	}
	return true
}

/**
 * 获取单条记录
 * @author pwt
 * @date 2018-8-17
 * @param
 * @return
 */
func GetUser(id int64) (user User, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName(tableName))
	qs = qs.Filter("id", id)
	qs = qs.OrderBy("-id")
	err = qs.One(&user, "id", "username", "password", "openid", "avatar", "phone")
	if err == nil {
		return user, nil
	}
	return user, err
}

/**
 * 获取所有记录
 * @author pwt
 * @date 2018-8-17
 * @param
 * @return
 */
func GetAllUser(condArr map[string]string, pagi ...int) (user []User, num int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName(tableName))
	cond := orm.NewCondition()
	if condArr["keywords"] != "" {
		cond = cond.AndCond(cond.And("username__icontains", condArr["keywords"]).Or("nickname__icontains", condArr["keywords"]))
	}
	if condArr["phone"] != "" {
		cond = cond.And("phone", condArr["phone"])
	}
	qs = qs.SetCond(cond)
	qs = qs.Limit(models.Pagination(pagi))
	num, err = qs.All(&user)
	return user, num, err
}

func Login(condArr map[string]string) (user User, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(models.TableName(tableName))
	cond := models.QueryCond(condArr)
	qs = qs.SetCond(cond)
	err = qs.One(&user, "id", "username", "password", "openid", "avatar", "phone")
	if err == nil {
		return user, nil
	}
	return user, err
}
