package golang_http

import (
  "bytes"
  "encoding/json"
  "fmt"
  "io"
  "net/http"
)

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

// Post 发送POST请求 泛型参数和泛型响应
func Post[T any, V any](url string, request Request[T]) (Response[V], error) {

  // 将payload序列化为JSON
  jsonData, err := json.Marshal(request)
  if err != nil {
    return EmptyResponse[V](), fmt.Errorf("failed to marshal payload: %w", err)
  }

  // 将JSON数据转换为bytes.Buffer
  buffer := bytes.NewBuffer(jsonData)

  resp, err := http.Post(url, "application/json", buffer)
  if err != nil {
    return EmptyResponse[V](), fmt.Errorf("failed to send request: %w", err)
  }

  body, err := io.ReadAll(resp.Body)

  var response Response[V]
  err = json.Unmarshal(body, &response)
  if err != nil {
    return EmptyResponse[V](), fmt.Errorf("failed to unmarshal response: %w", err)
  }

  defer resp.Body.Close()
  return response, nil
}
