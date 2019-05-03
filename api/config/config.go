package config

import (
	"github.com/sdvdxl/dbox/api/ex"
	"github.com/sdvdxl/dbox/api/model"
	"github.com/spf13/viper"
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
	LogFile string
	MetaDB  string
	AliOss  model.AliOss
}
