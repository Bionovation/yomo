package models

type MarkItem struct {
	ClassId int        // class id
	Rect    [4]float64 // yolo的标注格式
}
