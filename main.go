package main

import (
	_ "github.com/kooksee/kstore/routers"
	"github.com/astaxie/beego"
	"github.com/kooksee/kstore/conf"
)

func main() {
	cfg := conf.DefaultCfg()
	cfg.InitLogs()
	cfg.InitDb()

	beego.Run()
}

