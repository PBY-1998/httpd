/**
    @author: potten
    @since: 2022/12/5
    @desc: //TODO
**/
package httpd

import (
	"github.com/google/uuid"
)

var id, _ = uuid.NewUUID() // request 唯一uuid

type resp struct {
	RequestId string      `json:"requestId"`        // 请求uuid
	Code      int         `json:"code"`             // 错误码
	Msg       string      `json:"msg"`              // 提示信息
	Data      interface{} `json:"data"`             // 返回数据
	Errors    interface{} `json:"errors,omitempty"` // 错误描述
}

// 构造函数
func newResp(code int, msg string, data interface{}, errors interface{}) *resp {
	return &resp{
		RequestId: id.String(),
		Code:      code,
		Msg:       msg,
		Data:      data,
		Errors:    errors,
	}
}
