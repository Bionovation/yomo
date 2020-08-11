package models

import (
	"io/ioutil"
	"os"
	"strings"
)

type Workspace struct {
	Name string
	Folder string
	ClassDef string  `json:"-"`
	ClassNames []string
}

func (ws *Workspace)LoadClassNames() error {
	f,err := os.Open(ws.ClassDef)
	if err != nil {
		return err
	}
	defer f.Close()

	cns := make([]string, 0, 8)
	alltext,err := ioutil.ReadAll(f)
	if err != nil{
		return err
	}

	allines := strings.Split(string(alltext),"\n")

	for _,line := range allines{
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		cns = append(cns, line)
	}

	ws.ClassNames = cns
	return nil
}