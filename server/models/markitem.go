package models

import (
	"fmt"
	"strconv"
	"strings"
)

type MarkItem struct {
	ClassId int        // class id
	Rect    [4]float64 // yolo的标注格式
}

// 0 0.498366 0.836914 0.034314  0.044922
func (m *MarkItem)LoadFromString(line string) error {
	var err error
	nums := strings.Split(strings.TrimSpace(line), " ")
	if len(nums) < 5 {
		return fmt.Errorf("LoadFromString failed, error string to parse")
	}
	m.ClassId, err = strconv.Atoi(nums[0])
	if err != nil{
		return fmt.Errorf("LoadFromString failed, classid can not parse")
	}

	posm := 0
	posn := 1

	for {
		if posn >= len(nums){
			break
		}
		if len(strings.TrimSpace(nums[posn])) == 0{
			posn = posn + 1
			continue
		}
		//log.Println(posn, len(nums))
		m.Rect[posm], err = strconv.ParseFloat(nums[posn],64)
		if err != nil{
			return fmt.Errorf("LoadFromString failed, rect can not parse")
		}
		posm = posm + 1
		posn = posn + 1
	}

	if posm != 4{
		return fmt.Errorf("LoadFromString failed, rect values can not parse")
	}

	return nil
}

// 转换成一行数据
func (m *MarkItem)ToString() (string, error) {
	if m == nil {
		return "", fmt.Errorf("empty marked item.")
	}

	line := fmt.Sprintf("%d %v %v %v %v", m.ClassId, m.Rect[0], m.Rect[1], m.Rect[2], m.Rect[3])
	return line, nil
}