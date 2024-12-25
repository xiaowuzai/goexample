package conf

import (
	"github.com/spf13/viper"
)

// Config 存储应用的所有配置, 使用 viper 来管理
type Config struct {
	Environment      string `mapstructure:"ENVIRONMENT"`
	CasbinModelPath  string `mapstructure:"CASBIN_MODEL_PATH"`  // 模型文件路径
	CasbinPolicyPath string `mapstructure:"CASBIN_POLICY_PATH"` // 策略文件路径
}

// LoadConfig  从指定路径中加载配置。confType: env,json,yml
func LoadConfig(path string, confType string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType(confType) //json yml

	viper.AutomaticEnv() // 用 ENV 替换默认的环境变量值

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
