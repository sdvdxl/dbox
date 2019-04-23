// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/sdvdxl/dbox/dbox/config"
	"github.com/sdvdxl/dbox/dbox/dao"
	"github.com/sdvdxl/dbox/dbox/ex"
	"github.com/sdvdxl/dbox/dbox/log"
	"github.com/sdvdxl/dbox/dbox/model"
	"github.com/sdvdxl/dbox/dbox/service/cloud"
	"github.com/spf13/cobra"
	"os"
)

const (
	cfgPath       = string(os.PathSeparator) + ".dbox"
	cfgFile       = cfgPath + string(os.PathSeparator) + "cfg.yml"
	configContent = `
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
defaultFoler: "默认"
`
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dbox",
	Short: "A client of oss with local files index",
	Long:  `可以通过OSS上传文件，并且在本地记录文件索引，配置文件夹。并且可以将本地索引文件再上传到OSS本身进行存储，方便在其他地方使用`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	h, err := homedir.Dir()
	ex.Check(err)

	if err = os.Mkdir(h+cfgPath, 0766); err != nil && !os.IsExist(err) {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	f := h + cfgFile

	if fi, err := os.Stat(f); os.IsNotExist(err) {
		file, _ := os.Create(f)
		_, err = file.WriteString(configContent)
		ex.Check(err)
		ex.Check(file.Close())

		printInitInfoAndExit()
	} else if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	} else if fi.IsDir() {
		printInitInfoAndExit()
	}

	ex.Check(config.Parse(f))
	config.Cfg.LogFile = h + cfgPath + string(os.PathSeparator) + "dbox.log"
	log.Init()
	config.Cfg.MetaDB = h + cfgPath + string(os.PathSeparator) + "meta.db"
	dao.Use(dao.DialectSqlite3, "meta.db")
	dao.Init()
}

func checkConfig() {
	if config.Cfg.AliOss.AccessKeyID == "" || config.Cfg.AliOss.AccessKeySecret == "" ||
		config.Cfg.AliOss.Endpoint == "" || config.Cfg.AliOss.Bucket == "" {
		printInitInfoAndExit()
	}
	cloudService.UseCloudFileManager(&cloudService.AliOssFileManager{})
}

func printInitInfoAndExit() {
	fmt.Println("OSS未配置，使用config配置")
	ex.Check(configCmd.Help())
	os.Exit(1)
}

func printTables(files []model.File) {
	flen := getMaxLen(files)
	fmt.Println(getStr("文件名", flen), "| ", getStr("PATH", flen))
	for _, f := range files {
		fmt.Println(getStr(f.Name, flen), "| ", getStr(f.Path, flen))
	}
}

func getMaxLen(files []model.File) int {
	maxLen := 0
	for _, f := range files {
		l := len(f.Name)
		if l > maxLen {
			maxLen = l
		}
	}

	return maxLen
}

func getStr(str string, num int) string {
	l := len(str)
	if l >= num {
		return str
	}

	for i := 0; i < num-l; i++ {
		str += " "
	}

	return str

}
