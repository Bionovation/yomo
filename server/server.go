package server

import (
	"github.com/kataras/iris/v12"
	"os"
	"path/filepath"
)

const IMGPATH  = "html/img/"

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
	// set img
	app.Get("/imgs",handleImgList)
	app.Get("/img/:imgname",handleGetImg)
	app.Get("/img/:imgname/info",handleGetImgInfo)
}