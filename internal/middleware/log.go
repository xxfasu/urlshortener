package middleware

import (
	"bytes"
	"github.com/google/uuid"
	"github.com/xxfasu/urlshortener/pkg/logs"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

type LogM struct {
}

func NewLogM() *LogM {
	return &LogM{}
}

func (m *LogM) Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if strings.HasPrefix(ctx.Request.URL.String(), "/crmebimage") {
			ctx.Next()
			return
		}
		start := time.Now()
		trace := uuid.NewString()
		logs.Log.WithValue(ctx, zap.String("trace", trace))

		ctx.Next()
		requestParams := m.getRequestParams(ctx)
		method := ctx.Request.Method

		logs.Log.WithContext(ctx).Info("Request",
			zap.String("request_method", method),
			zap.String("request_url", ctx.Request.URL.String()),
			zap.String("client_ip", ctx.ClientIP()),
			zap.Int("status", ctx.Writer.Status()),
			zap.String("latency", time.Since(start).String()),
			zap.String("request_params", requestParams),
		)
	}
}

func (m *LogM) getRequestParams(ctx *gin.Context) string {
	method := ctx.Request.Method

	// 准备一个变量来存所有请求参数的字符串
	var requestParams string

	// 1. 处理 Query 参数（常见于 GET）
	if ctx.Request.URL.RawQuery != "" {
		if len(requestParams) > 0 {
			requestParams += " | "
		}
		requestParams += "Query: " + ctx.Request.URL.RawQuery
	}

	// 2. 如果是 POST 或 PUT、PATCH 等，尝试读取请求体
	//    - 根据 Content-Type 决定如何处理
	if method == http.MethodPost || method == http.MethodPut || method == http.MethodPatch {
		contentType := ctx.GetHeader("Content-Type")
		if strings.Contains(contentType, "application/json") {
			// 读取原始 JSON body
			rawData, err := ctx.GetRawData()
			if err == nil {
				if len(requestParams) > 0 {
					requestParams += " | "
				}
				requestParams += "JSON: " + string(rawData)
				// 重新放回 body，以便后续 Handler 还能 Bind 到
				ctx.Request.Body = io.NopCloser(bytes.NewBuffer(rawData))
			}
		} else if strings.Contains(contentType, "application/x-www-form-urlencoded") ||
			strings.Contains(contentType, "multipart/form-data") {
			// 表单
			_ = ctx.Request.ParseForm() // 解析表单
			formString := ctx.Request.Form.Encode()
			if formString != "" {
				if len(requestParams) > 0 {
					requestParams += " | "
				}
				requestParams += "Form: " + formString
			}
		}
		// 若还有别的 Content-Type，可以再添加分支逻辑
	}

	// 3. Path Param
	//    若路径定义了 :id 等动态参数，这里可以获取
	//    演示：取所有已知的路由参数再拼接
	// c.Params 本质上是 []gin.Param
	if len(ctx.Params) > 0 {
		var pathParams []string
		for _, p := range ctx.Params {
			pathParams = append(pathParams, p.Key+"="+p.Value)
		}
		if len(pathParams) > 0 {
			if len(requestParams) > 0 {
				requestParams += " | "
			}
			requestParams += "PathParam: " + strings.Join(pathParams, "&")
		}
	}
	return requestParams
}
