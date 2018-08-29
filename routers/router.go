package routers

import (
	"github.com/kooksee/kstore/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	// 关于人本身的的知识图谱
	// 添加一个人的信息
	// 修改一个人的信息
	// 删除一个人的信息
	// 查询一个人的信息
}
