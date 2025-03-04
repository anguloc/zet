package jmcomic

type FileData struct {
	Name       string // html文件名，带.html后缀
	NameNotExt string // html文件名，不带.html后缀
	Id         uint64 // 章节id
	Path       string // 文件绝对路径
}

type PhotoData struct {
	Name       string // 文件名称，带后缀
	NameNotExt string // 文件名称，不带后缀
	Id         uint64 // 图片id
	IdStr      string // 图片id 字符串类型
	Path       string // 图片绝对路径
}
