package server

import (
	"github.com/kataras/iris/v12"
	"os"
	"path/filepath"
)

const IMGPATH  = "F:\\Workdatas\\红细胞标注\\5\\5\\"

func Run()  {
	app := newApp()
	app.Run(iris.Addr(":8080"))
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

	app.Get("/imgs",handleImgList)
	app.Get("/img/:imgname",handleGetImg)
	app.Get("/info/:imgname",handleGetImgInfo)

	// 提交标注信息
	app.Put("/mark/:imgname",handlePutMark)
}