package conf

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

func loadLocal(localPath ...string) error {

	// 设置默认值
	viper.SetDefault("server.port", ":8000")
	viper.SetDefault("server.host", "127.0.0.1")

	// 设置配置文件的名称（不带扩展名）
	viper.SetConfigName("local")
	// 设置配置文件的类型
	viper.SetConfigType("toml")
	// 添加配置文件所在的路径
	if len(localPath) != 0 {
		viper.AddConfigPath(localPath[0]) // 当前目录
	} else {
		viper.AddConfigPath("./config") // 当前目录
	}

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("读取配置文件失败: %v", err)
		return err
	}

	// 将配置文件内容反序列化到 Config 结构体
	if err := viper.Unmarshal(&Config); err != nil {
		log.Fatalf("解析配置文件失败: %v", err)
		return err
	}

	// 打印解析后的配置内容
	fmt.Printf("服务器配置: %+v\n", Config.System)
	fmt.Printf("Mysql配置: %+v\n", Config.Mysql)
	fmt.Printf("Reids配置: %+v\n", Config.Redis)
	fmt.Printf("OSS配置: %+v\n", Config.AliyunOSS)

	// 监听配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置文件已修改:", e.Name)
		if err := viper.Unmarshal(&Config); err != nil {
			log.Fatalf("重新解析配置文件失败: %v", err)
		}
		fmt.Printf("更新后的服务器配置: %+v\n", Config.System)
		fmt.Printf("更新后的Mysql配置: %+v\n", Config.Mysql)
		fmt.Printf("更新后的Reids配置: %+v\n", Config.Redis)
		fmt.Printf("更新后的OSS配置: %+v\n", Config.AliyunOSS)
	})
	return nil
}
