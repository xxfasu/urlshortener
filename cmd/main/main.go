package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kardianos/service"
	"github.com/xxfasu/urlshortener/cmd/main/wire"
	"github.com/xxfasu/urlshortener/internal/conf"
	"github.com/xxfasu/urlshortener/pkg/logs"
	"log"
	"sync"
)

// 定义程序结构体
type program struct {
	server    *gin.Engine
	svc       service.Service
	once      sync.Once
	clearFunc func()
}

// Start 方法在服务启动时调用
func (p *program) Start(s service.Service) error {
	// 在一个新的 goroutine 中启动服务
	go p.run()
	return nil
}

// Stop 方法在服务停止时调用
func (p *program) Stop(s service.Service) error {
	// 这里可以添加清理资源的代码
	p.once.Do(func() {
		if p.clearFunc != nil {
			p.clearFunc()
		}
	})
	return nil
}

// 运行 Gin 服务器
func (p *program) run() {
	err := conf.InitConfig()
	if err != nil {
		panic(err)
	}

	logs.InitLog()
	wire, fn, err := wire.NewWire()
	p.clearFunc = fn
	if err != nil {
		panic(err)
	}
	p.server = wire

	if err = p.server.Run(conf.Config.System.Port); err != nil {
		panic(err)
	}
}

func main() {
	// 定义服务配置
	svcConfig := &service.Config{
		Name:        "URLShortener",
		DisplayName: "URLShortener",
		Description: "This is db.sql short link generator project",
	}
	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	prg.svc = s

	// 设置日志
	errs := make(chan error)
	logger, err := s.Logger(errs)
	if err != nil {
		log.Fatal(err)
	}

	// 启动服务
	if err := s.Run(); err != nil {
		logger.Error(err)
	}

	// 处理错误
	go func() {
		for {
			select {
			case err := <-errs:
				log.Println("Error:", err)
			}
		}
	}()

	log.Println("Shutting down...")
}
