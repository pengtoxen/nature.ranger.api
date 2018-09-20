package initial

import (
	"github.com/astaxie/beego/logs"
)

func InitLog() {
	logs.SetLogger(logs.AdapterFile, `{"filename":"logs/project.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":60}`)
	logs.EnableFuncCallDepth(true)
	logs.Async()
}
