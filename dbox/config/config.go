package config

import (
	"github.com/sdvdxl/dbox/dbox/log"
	"github.com/sdvdxl/dbox/dbox/model"
	"github.com/spf13/viper"
	"os"
)

var Cfg *Config

func Parse(configFile string) error {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	Cfg = &Config{}
	Cfg.MetaDB = viper.GetString("metaDB")
	if Cfg.MetaDB == "" {
		udir, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}
		Cfg.MetaDB = udir + "/meta.db"
		log.Log.Info("metaDB not set, will use use home path:", Cfg.MetaDB)
	}

	Cfg.AliOss.Endpoint = viper.GetString("cloud.aliOss.endPoint")
	Cfg.AliOss.AccessKeySecret = viper.GetString("cloud.aliOss.accessKeySecret")
	Cfg.AliOss.AccessKeyID = viper.GetString("cloud.aliOss.accessKeyID")
	Cfg.AliOss.Bucket = viper.GetString("cloud.aliOss.bucket")

	return nil
}

type Config struct {
	MetaDB string
	AliOss model.AliOss
}
