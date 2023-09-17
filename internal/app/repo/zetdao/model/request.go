package model

const (
	MarkDmHyRss = "dmhy_rss"
)

const (
	requestParseInit    = 0 // 初始化
	requestParsing      = 1 // 解析中
	requestParseSuccess = 2 // 解析完成
	requestParseFail    = 3 // 解析失败
)

func (r *Request) IsInit() bool {
	return r.ParseStatus == requestParseInit
}

func (r *Request) IsParsing() bool {
	return r.ParseStatus == requestParsing
}

func (r *Request) IsSuccess() bool {
	return r.ParseStatus == requestParseSuccess
}

func (r *Request) IsFail() bool {
	return r.ParseStatus == requestParseFail
}
