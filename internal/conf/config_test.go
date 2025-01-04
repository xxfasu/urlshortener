package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"testing"
)

func TestInitConfig(t *testing.T) {
	initNacos()

	v := viper.New()

	dataId := "ofcloud-backend"
	group := "ofcloud-backend"

	err := loadConfigFromNacos(dataId, group, v)
	if err != nil {
		log.Fatalf("Failed to load config from Nacos: %v", err)
	}

	// 读取配置
	err = v.Unmarshal(&Config)
	if err != nil {
		t.Fatalf("Failed to unmarshal config: %v", err)
	}
	fmt.Println("Initial config value:", Config)

	// 监听配置变化
	err = watchConfig(dataId, group, v, func() {
		// 配置变化后的回调逻辑
		err = v.Unmarshal(&Config)
		if err != nil {
			t.Fatalf("Failed to unmarshal config: %v", err)
		}
		fmt.Println("Updated config value:", Config)
	})
	if err != nil {
		log.Fatalf("Failed to watch config: %v", err)
	}
	select {}
}
