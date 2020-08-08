package models

// 图像列表个体，图片名称以及图片是否标注过
type ImgItem struct {
	Index  int
	Name   string
	Marked bool
}
