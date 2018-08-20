package routers

import (
	"github.com/kooksee/kstore/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	// 通过配置文件加载存储后端,oss存储,本地leveldb存储,minio,aws,ipfs,kfs存储等
	// 用户权限管理，以及授权管理，就是对相关的资源添加白名单
	// 资源的访问需要签名才行,通过签名可以获取用户的地址
	// 如果资源公开，那么就不需要授权访问了，可以直接进行访问
	// 每个用户可以配置自己独有的oss
	// 用户云上面可以使用自己的存储也可以使用云上的存储
	// 关于数据访问，每一条数据都带有一个metadata,所有的metadata都有一个license以及收费的标准
	// 用户通过对某一资源进行转账操作，如果校验通过，那么，该用户拥有资源的访问权限
	// 提供统一的对外的访问接口，每一次的访问都有时效性以及缓存，在缓存中，那么
}
