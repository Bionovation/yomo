package server

import (
	"fmt"
	"github.com/kataras/iris/v12/context"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"yomo/server/models"
	"yomo/utils"
)

// 提交标注信息
func handlePutMark(ctx context.Context)  {
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
