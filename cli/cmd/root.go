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
	"github.com/sdvdxl/dbox/api/config"
	"github.com/sdvdxl/dbox/api/dao"
	"github.com/sdvdxl/dbox/api/ex"
	"github.com/sdvdxl/dbox/api/log"
	"github.com/sdvdxl/dbox/api/service"
	"github.com/spf13/cobra"
	"os"
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
	if err := os.Mkdir(config.GetBasePath(), 0766); err != nil && !os.IsExist(err) {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if fi, err := os.Stat(config.GetConfigFile()); os.IsNotExist(err) {
		file, _ := os.Create(config.GetConfigFile())
		_, err = file.WriteString(config.Content)
		ex.Check(err)
		ex.Check(file.Close())

		printInitInfoAndExit()
	} else if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	} else if fi.IsDir() {
		printInitInfoAndExit()
	}

	ex.Check(config.Parse(config.GetConfigFile()))

	log.Init()

	dao.Init()
}

func checkConfig() {
	if config.Cfg.AliOss.AccessKeyID == "" || config.Cfg.AliOss.AccessKeySecret == "" ||
		config.Cfg.AliOss.Endpoint == "" || config.Cfg.AliOss.Bucket == "" {
		printInitInfoAndExit()
	}
	service.UseCloudFileManager(&service.AliOssFileManager{})
}

func printInitInfoAndExit() {
	fmt.Println("OSS未配置，使用config配置")
	ex.Check(configCmd.Help())
	os.Exit(1)
}
