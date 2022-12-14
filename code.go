/**
    @author: potten
    @since: 2022/12/5
    @desc: //TODO 定义常用响应状态
**/
package httpd

//  错误码规则：
//  (1) 错误码需为大于0的数
//  (2) 错误码为5位数
//  	---------------------------------------
//			第1位		2、3位		4、5位
//		---------------------------------------
//	  	  服务级错误码   模块级错误码   具体错误码
//		---------------------------------------
const (
	// 通用成功
	OK = 200 // // OK
	// 通用错误
	Err          = 500 // 通用错误
	BadRequest   = 400 // HTTP请求失败
	NotFound     = 404 // 不存在或未找到
	TokenInvalid = 420 // token验证失败
	SignInvalid  = 403 // sign验证失败

	// 服务级错误码
	ErrParam      = 10100 // 参数有误
	ErrSignParam  = 10101 // 签名参数有误
	ErrPermission = 10402 // api权限错误

	ErrMongoConnect    = 10201 // mongo连接出错
	ErrRedisConnect    = 10202 // redis连接出错
	ErrRabbitMQConnect = 10203 // rabbit连接出错

	// 模块级错误码 - 用户模块
	ErrUserService = 20100 // 用户服务异常
	ErrUserPhone   = 20101 // 用户手机号不合法
	ErrUserCaptcha = 20102 // 用户验证码有误

	// 模块级错误码 - 库存模块
	ErrOrderService = 20200 // 订单服务异常
	ErrOrderOutTime = 20201 // 订单超时
)
