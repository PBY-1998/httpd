/**
    @author: potten
    @since: 2022/12/5
    @desc: //TODO
**/
package httpd

type respHandler struct {
	*resp
}

type respHandlerInterface interface {
	Success(code int, data interface{}) *resp
	Error(code int, msg interface{}) *resp
}

// NewRespHandler 构造函数
func NewRespHandler() respHandlerInterface {
	return &respHandler{}
}

//Success 通用返回成功
func (h *respHandler) Success(code int, data interface{}) *resp {
	return newResp(code, "success", data, nil)
	//res, _ := json.Marshal(response)
	//return string(res)
}

// Error 通用返回失败
func (h *respHandler) Error(code int, msg interface{}) *resp {
	return newResp(code, "error", nil, msg)
}
