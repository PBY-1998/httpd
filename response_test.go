/**
    @author: potten
    @since: 2022/12/5
    @desc: //TODO
**/
package httpd

import (
	"fmt"
	"testing"
)

var responseHandler = NewRespHandler()

func Test_Success(t *testing.T) {
	type User struct {
		name string `json:"name"`
	}
	res := responseHandler.Success(OK, []User{
		{
			name: "222",
		},
	})
	fmt.Println(res)
}

func Test_Error(t *testing.T) {
	responseHandler.Error(Err, "通用错误返回")
}

func Benchmark_Success(b *testing.B) {
	for i := 0; i < b.N; i++ {
		responseHandler.Success(OK, "api返回成功")
	}
}

func Benchmark_Error(b *testing.B) {
	for i := 0; i < b.N; i++ {
		responseHandler.Error(Err, "500通用错误返回")
	}
}
