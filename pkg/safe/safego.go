package safe

import (
	"fmt"
	"github.com/xxfasu/urlshortener/pkg/logs"
	"go.uber.org/zap"
	"time"
)

// Go 函数用于在新的 goroutine 中执行传入的函数 fn，并捕获可能发生的 panic。
func Go(fn func()) {
	go func() {
		// 使用 defer 确保在函数退出前执行恢复逻辑
		defer func() {
			if rv := recover(); rv != nil { // 捕获 panic
				// 获取当前时间，格式化为易读的字符串
				timestamp := time.Now().Format("2006/01/02 - 15:04:05")

				// 记录错误日志，包含时间戳和 panic 信息
				logs.Log.Error(fmt.Sprintf("[safego] %s panic recovered: %v", timestamp, rv),
					zap.Stack("stack")) // 添加堆栈信息
			}
		}()

		// 执行传入的函数 fn
		fn()
	}()
}
