package core

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
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
	rd := bufio.NewReader(mf)
	for{
		line,err := rd.ReadString('\n')
		if err != nil || io.EOF == err{
			break
		}
		line = strings.TrimSpace(line)

		mk := models.MarkItem{}
		err = mk.LoadFromString(line)
		if err != nil{
			log.Println(err)
			continue
		}else {
			mks = append(mks, mk)
		}
	}

	return mks, nil
}