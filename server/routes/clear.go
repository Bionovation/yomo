package routes

import (
	"fmt"
	"github.com/kataras/iris/v12/context"
	"os"
	"path/filepath"
	"strings"
	"yomo/server/config"
	"yomo/server/core"
	"yomo/utils"
)

func ClearUnMarked(ctx context.Context)  {
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
		return
	}

	if len(data) < 1{
		res.DoneData("ok")
		return
	}

	// 删除没有.txt 的图像
	for _, img := range data {
		txt := strings.TrimRight(img.Name,filepath.Ext(img.Name)) + ".txt"
		txt = filepath.Join(imgfolder,txt)

		if !utils.Exist(txt){
			os.Remove(filepath.Join(imgfolder, img.Name))
		}
	}

	res.DoneData(data)
}