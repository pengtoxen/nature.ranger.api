package constant

var FieldValidateInfo = map[string]string{
	"Required":     "必填",
	"Min":          "最小值为 %d",
	"Max":          "最大值为 %d",
	"Range":        "必须在区间 %d 到 %d",
	"MinSize":      "Minimum size is %d",
	"MaxSize":      "Maximum size is %d",
	"Length":       "长度为 %d",
	"Alpha":        "必须是字母",
	"Numeric":      "必须是数字",
	"AlphaNumeric": "必须是数字或者字母",
	"Match":        "必须匹配 %s",
	"NoMatch":      "必须不匹配 %s",
	"AlphaDash":    "必须是数字或者字母或者下划线",
	"Email":        "必须是有效的邮箱地址",
	"IP":           "必须是有效的IP地址",
	"Base64":       "必须是有效的base64字符",
	"Mobile":       "必须是有效的手机号码",
	"Tel":          "必须是有效的座机号码",
	"Phone":        "必须是有效的手机号码或者座机号码",
	"ZipCode":      "必须是有效的邮编号码",
}

var MessageTmpls = map[string]map[string]string{
	"status": {
		"status":  "状态",
		"public":  "公开",
		"private": "私有",
		"enable":  "开启",
		"disable": "关闭",
		"yes":     "是",
		"no":      "否",
	},
	"message": {
		"success":         "成功",
		"error":           "失败",
		"submit_success":  "提交成功",
		"submit_error":    "提交错误",
		"operate_success": "操作成功",
		"operate_fail":    "操作失败",
		"operate_cancel":  "操作取消",
		"info":            "信息",
		"cancel":          "取消",
		"delete":          "删除",
		"confirm":         "取消",
		"illegal_form":    "非法的格式",
		"duplicate":       "重复",
	},
	"common": {
		"id_is_required":                 "id不能为空",
		"username_or_password_incorrect": "用户名或密码不正确",
		"login_fail":                     "登录失败",
		"invalid_access_token":           "无效的token",
	},
}
