package conf

// Nacos 配置结构体
type Nacos struct {
	Addr      string `mapstructure:"addr"`
	Port      uint64 `mapstructure:"port"`
	Namespace string `mapstructure:"namespace"`
	Username  string `mapstructure:"username"`
	Password  string `mapstructure:"password"`
	DataID    string `mapstructure:"data_id"`
	Group     string `mapstructure:"group"`
}

// Env 总配置结构体
type env struct {
	Environment string `mapstructure:"environment"`
	Nacos       Nacos  `mapstructure:"nacos"`
}
