package config

import (
	"github.com/mitchellh/go-homedir"
	"github.com/sdvdxl/dbox/api/ex"
	"github.com/sdvdxl/dbox/api/model"
	"github.com/spf13/viper"
	"os"
)

var Cfg *Config

func Update() {
	ex.Check(viper.WriteConfig())
}

func Parse(configFile string) error {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	Cfg = &Config{}

	Cfg.AliOss.Endpoint = viper.GetString("cloud.aliOss.endPoint")
	Cfg.AliOss.AccessKeySecret = viper.GetString("cloud.aliOss.accessKeySecret")
	Cfg.AliOss.AccessKeyID = viper.GetString("cloud.aliOss.accessKeyID")
	Cfg.AliOss.Bucket = viper.GetString("cloud.aliOss.bucket")
	return nil
}

type Config struct {
	AliOss model.AliOss
}

const (
	cfgPath = string(os.PathSeparator) + ".dbox"
	Content = `
# 更改文件名为 cfg.yml
# 阿里云OSS配置
cloud:
  #  use: "aliOss"
  aliOss:
    endpoint: ""
    accessKeyID: ""
    accessKeySecret: ""
    bucket: ""

# 默认文件夹
defaultFolder: "默认"
`
)

func GetDBFile() string {
	return GetBasePath() + "meta.db"
}
func GetConfigFile() string {
	return GetBasePath() + "cfg.yml"
}

// GetBasePath 获取配置路径，最后带目录分隔符
func GetBasePath() string {
	h, err := homedir.Dir()
	ex.Check(err)
	return h + cfgPath + string(os.PathSeparator)
}

// GetHomeDir 获取配置路径，最后带目录分隔符
func GetHomeDir() string {
	h, err := homedir.Dir()
	ex.Check(err)
	return h + string(os.PathSeparator)
}
