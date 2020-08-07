package core

import (
	"log"
	"testing"
)

func TestGetImgList(t *testing.T) {
	imgpath := "E:\\go\\src\\yomo\\html\\img"
	imgs,err := GetImgList(imgpath)
	if err != nil{
		t.Fail()
		return
	}
	log.Print(imgs)
}

