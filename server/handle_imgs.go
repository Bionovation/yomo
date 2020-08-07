package server

import (
	"fmt"
	"github.com/kataras/iris/v12/context"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"yomo/server/core"
	"yomo/server/models"
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


// 获取图片信息
func handleGetImgInfo(ctx context.Context)  {
	imgname := ctx.Params().Get("imgname")
	ext := strings.ToLower(filepath.Ext(imgname))

	res := utils.NewRes(ctx)

	if ext != ".jpg" && ext != ".bmp" && ext != ".png"{
		res.FailErr(fmt.Errorf("format not supported."))
		return
	}

	imgfolder,_ := filepath.Abs(IMGPATH)
	imgpath := filepath.Join(imgfolder,imgname)
	if !utils.Exist(imgpath){
		res.FailErr(fmt.Errorf("image file not exist."))
		return
	}

	info := models.ImgInfo{}
	info.Name = imgname

	txt := strings.TrimRight(imgname,filepath.Ext(imgname)) + ".txt"
	txt = filepath.Join(imgfolder,txt)
	info.Marked = utils.Exist(txt)

	marks,err := core.LoadMarks(txt)
	if err != nil {
		res.FailErr(err)
		return
	}

	info.Marks = marks
	res.DoneData(info)
}