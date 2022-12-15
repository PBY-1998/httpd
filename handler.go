/**
    @author: potten
    @since: 2022/12/5
    @desc: //TODO
**/
package httpd

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Handler 封装gin.Context
type Context struct {
	*gin.Context
}

func Handler(ctx *gin.Context) Handlerer {
	return &Context{ctx}
}

type Handlerer interface {
	Success(code int, data interface{})
	Error(code int, err interface{})
}

func (c *Context) Success(code int, data interface{}) {
	resp := newResponse()
	resp.Data = data
	resp.Msg = http.StatusText(code)
	c.JSON(code, resp)
}

func (c *Context) Error(code int, err interface{}) {
	resp := newResponse()
	resp.Msg = http.StatusText(code)
	resp.Code = code
	resp.Error = err
	c.JSON(code, resp)
}
