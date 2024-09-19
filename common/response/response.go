package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// SuccessWithData 成功且返回数据
func SuccessWithData(c *gin.Context, data interface{}) {
	response := &Response{
		Code: 0,
		Msg:  "success",
		Data: data,
	}
	c.JSON(http.StatusOK, response)
}

// SuccessWithoutData 成功但不返回数据
func SuccessWithoutData(c *gin.Context) {
	SuccessWithData(c, nil)
}

type ErrorCode struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

var (
	ParamsError          = ErrorCode{904030, "请求参数错误"}
	TooManyRequestsError = ErrorCode{904290, "请求频率过多"}
	NoAuthError          = ErrorCode{904010, "未认证!"}
	OperationError       = ErrorCode{905011, "操作错误"}
	UnKnownError         = ErrorCode{999999, "未知错误!"}
)

// Failed 失败返回错误对象
func Failed(c *gin.Context, err ErrorCode) {
	response := &Response{
		Code: err.Code,
		Msg:  err.Msg,
		Data: nil,
	}

	c.JSON(http.StatusOK, response)

}

// FailedWithMsg 失败返回错误字符串
func FailedWithMsg(c *gin.Context, msg string) {
	response := &Response{
		Code: 999999,
		Msg:  msg,
		Data: nil,
	}

	c.JSON(http.StatusOK, response)
}
