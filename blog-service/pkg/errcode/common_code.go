package errcode

var (
	Success                  = NewError(0, "成功")
	ServeError               = NewError(100000000, "服务器内部错误")
	InValidParams            = NewError(100000001, "入参错误")
	NotFound                 = NewError(100000002, "找不到")
	UnauthorizedAuthNotExist = NewError(100000003, "鉴权失败,找不到对应的appkey & appSecret")

	UnauthorizedTokenError    = NewError(100000004, "鉴权失败,Token错误")
	UnauthorizedTokenTimeout  = NewError(100000005, "鉴权失败,Token超时")
	UnauthorizedTokenGenerate = NewError(100000006, "鉴权失败,Token生成失败")
	TooManyRequests           = NewError(100000007, "请求过多")
)
