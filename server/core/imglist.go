package core

import (
	"io/ioutil"
	"path/filepath"
	"strings"
	"yomo/server/models"
	"yomo/utils"
)

// 获取图片列表
func GetImgList(imgpath string) ([]models.ImgItem, error) {
	dir, err := ioutil.ReadDir(imgpath)
	if err != nil {
		return nil, err
	}

	imgnames := make([]string,0)
	for _, fi := range dir {
		if fi.IsDir() {
			continue
		}
		ext := strings.ToLower(filepath.Ext(fi.Name()))
		if(ext == ".jpg" || ext == ".bmp" || ext == ".png"){
			imgnames = append(imgnames,fi.Name())
		}
	}

	// 读取状态，如果有txt文件，则说明标注过了，前端打勾


	imgs := make([]models.ImgItem,0, len(imgnames))
	for i, imgname :=  range imgnames{
		//fmt.Println(imgname)
		name := strings.TrimRight(imgname,filepath.Ext(imgname))
		txtfile := filepath.Join(imgpath,name + ".txt")
		checkfile := filepath.Join(imgpath,name + ".ok")

		img := models.ImgItem{
			Index:i,
			Name:imgname,
			Marked:utils.Exist(txtfile),
			Checked:utils.Exist(checkfile),
		}
		imgs = append(imgs,img)
	}


	return imgs, nil
}