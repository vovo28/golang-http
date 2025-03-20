package golang_http

type Response[T any] struct {
  Status  string `json:"status"`
  Message string `json:"message"`
  Data    T      `json:"data"`
}

type Request[T any] struct {
  Data T `json:"data"`
}

func NewRequest[T any](data T) Request[T] {
  return Request[T]{
    Data: data,
  }
}

func NewResponse[T any](status string, message string, data T) Response[T] {
  return Response[T]{
    Status:  status,
    Message: message,
    Data:    data,
  }
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
func (r *Response[T]) isSuccess() bool {
  return r.Status == "0"
}

func EmptyResponse[T any]() Response[T] {
  return Response[T]{Status: "", Message: "", Data: *new(T)}
}
