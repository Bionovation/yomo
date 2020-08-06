package core

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

func GetImgList(imgpath string) ([]string, error) {
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

	return imgnames, nil
}