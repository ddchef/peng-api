package global

type CustomError struct {
	ErrorCode int
	ErrorMsg  string
}

type CustomErrors struct {
	BusinessError   CustomError
	ValidateError   CustomError
	TokenError      CustomError
	PermissionError CustomError
}

var Errors = CustomErrors{
	BusinessError:   CustomError{40000, "业务错误"},
	ValidateError:   CustomError{42200, "请求参数错误"},
	TokenError:      CustomError{40100, "登录授权失效"},
	PermissionError: CustomError{40101, "无操作权限"},
}
