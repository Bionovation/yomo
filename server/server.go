package server

import (
	"fmt"
	"github.com/fatedier/frp/utils/log"
	"github.com/kataras/iris/v12"
	"os"
	"path/filepath"
	"yomo/server/config"
	"yomo/server/routes"
)



func Run()  {

	conf,err := config.Load("./yomo.toml")
	if err != nil{
		panic(err)
	}

	go frpLogin()

	log.Info("visit addr : http://%v.%v:8080",conf.Name,conf.Proxy.ServerAddr)

	app := newApp()
	app.Run(iris.Addr(fmt.Sprintf(":%v",conf.Port)))
}


func newApp() *iris.Application {
	app := iris.Default()
	app.Logger().SetLevel("info")

	setRouter(app)

	return app
}

func setRouter(app *iris.Application)  {

	exPath, _ := os.Executable()
	folder := filepath.Dir(exPath)
	app.HandleDir("/", filepath.Join(folder, "./html"))

	app.Get("/workspaces", routes.GetWorkspaceList)
	app.Get("/:wsname/imgs",routes.ImgList)
	app.Get("/:wsname/img/:imgname",routes.GetImg)
	app.Get("/:wsname/info/:imgname",routes.GetImgInfo)

	// 提交标注信息
	app.Put("/:wsname/mark/:imgname",routes.PutMark)
}