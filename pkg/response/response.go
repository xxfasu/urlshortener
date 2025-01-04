package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func result(c *gin.Context, code int, data interface{}, msg string) {
	// 开始时间
	c.JSON(code, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func Ok(c *gin.Context) {
	result(c, http.StatusOK, map[string]interface{}{}, "操作成功")
}

func OkWithMessage(c *gin.Context, message string) {
	result(c, http.StatusOK, map[string]interface{}{}, message)
}

func OkWithData(c *gin.Context, data interface{}) {
	result(c, http.StatusOK, data, "操作成功")
}

func OkWithDetailed(c *gin.Context, data interface{}, message string) {
	result(c, http.StatusOK, data, message)
}

func Fail(c *gin.Context) {
	result(c, http.StatusInternalServerError, map[string]interface{}{}, "操作失败")
}

func FailWithMessage(c *gin.Context, message string) {
	result(c, http.StatusInternalServerError, map[string]interface{}{}, message)
}
