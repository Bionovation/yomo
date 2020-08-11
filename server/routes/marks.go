package routes

import (
	"fmt"
	"github.com/kataras/iris/v12/context"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"yomo/server/config"
	"yomo/server/models"
	"yomo/utils"
)

// 提交标注信息
func PutMark(ctx context.Context)  {
	res := utils.NewRes(ctx)

	imgname := ctx.Params().Get("imgname")
	wsname := ctx.Params().Get("wsname")
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

	// 提取标注信息
	var mks []models.MarkItem
	err := ctx.ReadJSON(&mks)
	if err != nil{
		res.FailErr(err)
		return
	}

	allines := ""
	for _,mk := range mks{
		line, err := mk.ToString()
		if err != nil{
			res.FailErr(err)
			return
		}
		allines = allines + line + "\n"
	}


	// 保存
	txt := strings.TrimRight(imgname,filepath.Ext(imgname)) + ".txt"
	txt = filepath.Join(imgfolder,txt)
	os.Remove(txt)

	err = ioutil.WriteFile(txt,[]byte(allines), os.ModePerm)
	if err != nil{
		res.FailErr(err)
		return
	}

	res.DoneMsg("ok")
}
