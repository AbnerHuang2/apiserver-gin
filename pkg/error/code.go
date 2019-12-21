package error

// 自定义错误码，通常错误由错误码和错误信息两部分组成，便于跟踪和维护错误信息
// 错误码为0表示成功
// 错误码4开头表示业务层面的错误，比如校验等
// 错误码5开头表示服务器错误，比如数组越界等
var (
	// 成功
	SUCCESS = &Code{
		ErrCode: 0,
		Message: "success",
	}
	// 参数校验错误
	ValidateErr = &Code{
		ErrCode: 40001,
		Message: "参数校验错误",
	}
	// 需要授权
	RequireAuth = &Code{
		ErrCode: 40002,
		Message: "请先完成授权",
	}

	// 系统错误（按需求细化）
	SystemErr = &Code{
		ErrCode: 50000,
		Message: "系统错误",
	}
)

type Code struct {
	ErrCode int
	Message string
}