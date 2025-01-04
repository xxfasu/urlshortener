package conf

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
	"log"
	"strings"
	"sync"
)

var (
	configClient config_client.IConfigClient
	once         sync.Once
)

func initNacos() {
	once.Do(func() {
		serverConfigs := []constant.ServerConfig{
			{
				IpAddr: Env.Nacos.Addr,
				Port:   Env.Nacos.Port,
			},
		}

		clientConfig := constant.ClientConfig{
			Username:            Env.Nacos.Username,  // nacos授权的用户名
			Password:            Env.Nacos.Password,  // nacos授权的密码
			NamespaceId:         Env.Nacos.Namespace, // 如果需要支持多namespace，我们可以创建多个client,它们有不同的NamespaceId
			TimeoutMs:           5000,                // 请求服务端的超时时间 5000毫秒
			NotLoadCacheAtStart: true,                // 启动时不从缓存加载服务
			LogDir:              "nacos_log",         // 日志目录
			CacheDir:            "nacos_log/cache",   // 缓存目录
			LogLevel:            "debug",             // 日志等级
			AppendToStdout:      true,                // 打印到控制台
			// 配置日志滚动
			LogRollingConfig: &constant.ClientLogRollingConfig{
				MaxSize:    100,  // 100 MB
				MaxAge:     30,   // 30 天
				MaxBackups: 10,   // 保留10个旧日志文件
				Compress:   true, // 压缩旧日志文件
			},
		}

		var err error
		configClient, err = clients.CreateConfigClient(map[string]interface{}{
			"serverConfigs": serverConfigs,
			"clientConfig":  clientConfig,
		})

		if err != nil {
			log.Fatalf("Failed to create Nacos config client: %v", err)
		}
	})
}

func loadConfigFromNacos(dataId, group string, v *viper.Viper) error {
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
	})
	if err != nil {
		return err
	}

	v.SetConfigType("toml")
	err = v.ReadConfig(strings.NewReader(content))
	if err != nil {
		return err
	}

	return nil
}

func watchConfig(dataId, group string, v *viper.Viper, onChange func()) error {
	err := configClient.ListenConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
		OnChange: func(namespace, group, dataId, data string) {
			log.Println("Config changed, reloading...")
			v.SetConfigType("toml")
			if err := v.ReadConfig(strings.NewReader(data)); err != nil {
				log.Printf("Failed to reload config: %v", err)
				return
			}
			onChange()
		},
	})

	if err != nil {
		return err
	}

	log.Println("Started listening for config changes")
	return nil
}

func loadNacos() error {
	initNacos()

	v := viper.New()

	err := loadConfigFromNacos(Env.Nacos.DataID, Env.Nacos.Group, v)
	if err != nil {
		log.Fatalf("Failed to load config from Nacos: %v", err)
		return err
	}

	// 读取配置
	err = v.Unmarshal(&Config)
	if err != nil {
		log.Fatalf("Failed to unmarshal config: %v", err)
		return err
	}
	fmt.Printf("Initial config value %v:\n", Config)

	// 监听配置变化
	err = watchConfig(Env.Nacos.DataID, Env.Nacos.Group, v, func() {
		// 配置变化后的回调逻辑
		updatedValue := v.GetString("your.config.key")
		fmt.Println("Updated config value:", updatedValue)
	})
	if err != nil {
		log.Fatalf("Failed to watch config: %v", err)
		return err
	}
	return nil
}
