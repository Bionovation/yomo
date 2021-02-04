package routes

import (
	"fmt"
	"github.com/kataras/iris/v12/context"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"yomo/server/config"
	"yomo/utils"
)

// 标记为检查过
func PutCheck(ctx context.Context)  {
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

	// 标记为保存，生成同名 .ok 文件
	check := strings.TrimRight(imgname,filepath.Ext(imgname)) + ".ok"
	check = filepath.Join(imgfolder,check)
	if utils.Exist(check){
		res.DoneData("ok")
		return
	}

	err := ioutil.WriteFile(check,[]byte("ok"), os.ModePerm)
	if err != nil{
		res.FailErr(err)
		return
	}

	res.DoneMsg("ok")
}

