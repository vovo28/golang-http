package golang_http

type Response[T any] struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type Request[T any] struct {
	Data T `json:"data"`
}

// ErrorResponse 返回错误信息
/* 函数后跟着 T 标明泛型类型 */
func ErrorResponse[T any](message string, data T) Response[T] {
	return Response[T]{
		Status:  "1",
		Message: message,
		Data:    data,
	}
}

func SuccessResponse[T any](message string, data T) Response[T] {
	return Response[T]{
		Status:  "0",
		Message: message,
		Data:    data,
	}
}
