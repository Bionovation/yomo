package core

import (
	"os"
	"yomo/server/models"
	"yomo/utils"
)

func LoadMarks(txt string) ([]models.MarkItem, error) {
	mks := make([]models.MarkItem, 0)
	if !utils.Exist(txt){
		return mks, nil
	}

	mf, err := os.Open(txt)
	if err != nil{
		return nil, err
	}
	defer  mf.Close()

	// 按行读取

	return nil, nil
}