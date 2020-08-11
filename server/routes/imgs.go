package routes

import (
	"fmt"
	"github.com/kataras/iris/v12/context"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"yomo/server/config"
	"yomo/server/core"
	"yomo/server/models"
	"yomo/utils"
)

// 获取图像列表
func ImgList(ctx context.Context)  {

	res := utils.NewRes(ctx)

	wsname := ctx.Params().Get("wsname")
	ws := config.FindWsByName(wsname)
	if ws == nil{
		res.FailMsg(fmt.Sprintf("workspace %v not found.", wsname))
		return
	}

	imgfolder,_ := filepath.Abs(ws.Folder)

	data,err := core.GetImgList(imgfolder)
	if err != nil{
		res.FailErr(err)
	}

	res.DoneData(data)
}

// 获取单张图片
func GetImg(ctx context.Context)  {
	wsname := ctx.Params().Get("wsname")
	imgname := ctx.Params().Get("imgname")

	ws := config.FindWsByName(wsname)
	if ws == nil{
		ctx.NotFound()
		return
	}


	ext := strings.ToLower(filepath.Ext(imgname))

	if ext != ".jpg" && ext != ".bmp" && ext != ".png"{
		ctx.NotFound()
		return
	}

	imgfolder,_ := filepath.Abs(ws.Folder)
	imgpath := filepath.Join(imgfolder,imgname)

	ctx.Application().Logger().Info(imgpath)

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
func GetImgInfo(ctx context.Context)  {
	res := utils.NewRes(ctx)

	wsname := ctx.Params().Get("wsname")
	imgname := ctx.Params().Get("imgname")

	ws := config.FindWsByName(wsname)
	if ws == nil{
		res.FailMsg(fmt.Sprintf("workspace %v not found.", wsname))
		return
	}

	ext := strings.ToLower(filepath.Ext(imgname))



	if ext != ".jpg" && ext != ".bmp" && ext != ".png"{
		res.FailErr(fmt.Errorf("format not supported."))
		return
	}

	imgfolder,_ := filepath.Abs(ws.Folder)
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