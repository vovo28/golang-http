package golang_http

import (
	"log"
	"reflect"
	"testing"
)

type args[T any] struct {
	message string
	data    T
}
type CurTest struct {
	name string
	args args[string]
	want Response[string]
}

func TestErrorResponse(t *testing.T) {

	var tests []CurTest

	tests = append(tests, CurTest{
		name: "First test",
		args: args[string]{
			message: "test",
			data:    "update failed",
		}, want: Response[string]{
			Status:  "1",
			Message: "test",
			Data:    "update failed",
		},
	})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ErrorResponse(tt.args.message, tt.args.data); !reflect.DeepEqual(got, tt.want) {

				log.Printf("status: %s, message: %s, data: %s", got.Status, got.Message, got.Data)
				log.Printf("status: %s, message: %s, data: %s", tt.want.Status, tt.want.Message, tt.want.Data)
				t.Errorf("ErrorResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSuccessResponse(t *testing.T) {
	var tests []CurTest

	tests = append(tests, CurTest{
		name: "First test",
		args: args[string]{
			message: "test",
			data:    "update success",
		}, want: Response[string]{
			Status:  "0",
			Message: "test",
			Data:    "update success",
		},
	})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SuccessResponse(tt.args.message, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SuccessResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
