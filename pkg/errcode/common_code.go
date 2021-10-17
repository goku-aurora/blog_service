package errcode

var (
	Success = NewError(200,"成功")
	ServerError = NewError(100000,"服务器内部错误")
	InvalidParams = NewError(100001,"入参错误")
	NotFound = NewError(100002,"找不到")
	UnauthorizeAuthNotExist = NewError(100003,"鉴权失败，找不到对应的AppKey和AppSecret")
	UnauthorizeTokenError = NewError(100004,"鉴权失败，Token错误")
	UnauthorizeTokenTimeout = NewError(100005,"鉴权失败，Token超时")
	UnauthorizeTokenGenerate = NewError(100006,"鉴权失败，Token生成失败")
	TooManyRequests = NewError(100007,"请求过多")
)

