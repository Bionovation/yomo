package server

import (
	"github.com/kataras/iris/v12/context"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"yomo/server/core"
	"yomo/utils"
)

// 获取图像列表
func handleImgList(ctx context.Context)  {
	res := utils.NewRes(ctx)
	imgfolder,_ := filepath.Abs(IMGPATH)

	ctx.Application().Logger().Info(imgfolder)
	data,err := core.GetImgList(imgfolder)
	if err != nil{
		res.FailErr(err)
	}

	res.DoneData(data)
}

// 获取单张图片
func handleGetImg(ctx context.Context)  {
	imgname := ctx.Params().Get("imgname")
	ext := strings.ToLower(filepath.Ext(imgname))

	if ext != ".jpg" && ext != ".bmp" && ext != ".png"{
		ctx.NotFound()
		return
	}

	imgfolder,_ := filepath.Abs(IMGPATH)
	imgpath := filepath.Join(imgfolder,imgname)

	f,err := os.Open(imgpath)
	if err != nil{
		ctx.NotFound()
		return
	}
	defer f.Close()

	data,err := ioutil.ReadAll(f)


	ctx.Write(data)
}