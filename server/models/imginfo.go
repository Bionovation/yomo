package models

// 图片的详细信息，主要用于获取标注信息
type ImgInfo struct {
	Name string       // 名称
	Marked bool       // 是否标注过
	Checked bool      // 是否检查过
	Marks  []MarkItem // 标注列表
}
