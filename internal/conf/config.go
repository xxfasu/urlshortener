// admin_routes.go

package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

// system 配置结构体
type system struct {
	Port string `mapstructure:"port"`
	Host string `mapstructure:"host"`
}

// CrmebConfig 表示 Crmeb 的配置
type CrmebConfig struct {
	Version                    string `toml:"version"`                        // 当前代码版本
	Domain                     string `toml:"domain"`                         // 配合swagger使用，待部署域名
	WechatApiURL               string `toml:"wechat_api_url"`                 // 请求微信接口的专用服务器
	WechatJsAPIDebug           bool   `toml:"wechat_js_api_debug"`            // 微信js api系列是否开启调试模式
	WechatJsAPIBeta            bool   `toml:"wechat_js_api_beta"`             // 微信js api是否是beta版本
	AsyncConfig                bool   `toml:"async_config"`                   // 是否同步config表数据到redis
	AsyncWeChatProgramTempList bool   `toml:"async_wechat_program_temp_list"` // 是否同步小程序公共模板库
	ImagePath                  string `toml:"image_path"`                     // 服务器图片路径配置，斜杠结尾
}

// mysql 配置结构体
type mysql struct {
	Source string `mapstructure:"source"`
}

// redis 配置结构体
type redis struct {
	Addr         string   `mapstructure:"addr"`
	Username     string   `mapstructure:"username"`
	Password     string   `mapstructure:"password"`
	DB           int      `mapstructure:"db.sql"`
	UseCluster   bool     `mapstructure:"use_cluster"`
	ClusterAddrs []string `mapstructure:"cluster_addrs"`
}

// oss 配置结构体
type oss struct {
	Provider string `mapstructure:"provider"`
}

// aliyunOSS 配置结构体
type aliyunOSS struct {
	Endpoint  string `mapstructure:"endpoint"`
	KeyID     string `mapstructure:"key_id"`
	KeySecret string `mapstructure:"key_secret"`
	Bucket    string `mapstructure:"bucket"`
	Region    string `mapstructure:"region"`
	RoleArn   string `mapstructure:"role_arn"`
	Domain    string `mapstructure:"domain"`
}

// qiniuyunOSS 配置结构体
type qiniuyunOSS struct {
	AccessKey string `mapstructure:"access_key"`
	SecretKey string `mapstructure:"secret_key"`
	Bucket    string `mapstructure:"bucket"`
	Zone      string `mapstructure:"zone"`
}

type zapLog struct {
	LogLevel    string `mapstructure:"log_level"`
	Encoding    string `mapstructure:"encoding"`
	LogFileName string `mapstructure:"log_file_name"`
	MaxBackups  int    `mapstructure:"max_backups"`
	MaxAge      int    `mapstructure:"max_age"`
	MaxSize     int    `mapstructure:"max_size"`
	Compress    bool   `mapstructure:"compress"`
}

// Config 总配置结构体
type config struct {
	System      system      `mapstructure:"system"`
	CrmebConfig CrmebConfig `mapstructure:"crmeb_config"`
	Mysql       mysql       `mapstructure:"mysql"`
	Redis       redis       `mapstructure:"redis"`
	OSS         oss         `mapstructure:"oss"`
	AliyunOSS   aliyunOSS   `mapstructure:"aliyun_oss"`
	Log         zapLog      `mapstructure:"zap_log"`
	QiniuyunOSS qiniuyunOSS `mapstructure:"qiniuyun_oss"`
}

var Config *config
var Env *env

func InitConfig(localPath ...string) error {
	err := loadEnv(localPath...)
	if err != nil {
		return err
	}
	switch Env.Environment {
	case "local":
		err = loadLocal(localPath...)
	case "prod":
		err = loadNacos()
	}
	if err != nil {
		return err
	}
	return nil
}

func loadEnv(localPath ...string) error {
	// 设置默认值
	viper.SetDefault("environment", "local")

	// 设置配置文件的名称（不带扩展名）
	viper.SetConfigName("env")
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
	if err := viper.Unmarshal(&Env); err != nil {
		log.Fatalf("解析配置文件失败: %v", err)
		return err
	}

	// 打印解析后的配置内容
	fmt.Printf("Env配置: %+v\n", Env.Environment)
	fmt.Printf("Nacos配置: %+v\n", Env.Nacos)
	return nil
}
