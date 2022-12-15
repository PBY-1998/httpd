/**
    @author: potten
    @since: 2022/12/5
    @desc: //TODO
**/
package httpd

import (
	"github.com/google/uuid"
)

type response struct {
	RequestId string      `json:"requestId"`       // 请求uuid
	Code      int         `json:"code"`            // 错误码
	Msg       string      `json:"msg"`             // 提示信息
	Data      interface{} `json:"data"`            // 返回数据
	Error     interface{} `json:"error,omitempty"` // 错误描述
}

// 构造函数
func newResponse() *response {
	var id, _ = uuid.NewUUID() // request 唯一uuid
	return &response{
		RequestId: id.String(),
		Code:      200,
		Msg:       "",
		Data:      nil,
		Error:     nil,
	}
}
