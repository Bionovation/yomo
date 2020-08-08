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
	log.Println("===========TestGetImgList over============")
}

func TestLoadMarks(t *testing.T) {
	txt := "E:\\go\\src\\yomo\\html\\img\\2020073001440.txt"
	mks,err := LoadMarks(txt)
	if err != nil{
		t.Fail()
		return
	}

	log.Print(len(mks))
	log.Println("===========TestLoadMarks over============")
}