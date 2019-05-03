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
	"github.com/sdvdxl/dbox/api/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	endpoint        string
	accessKeyID     string
	accessKeySecret string
	bucket          string
)
// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "配置OSS",
	Run: func(cmd *cobra.Command, args []string) {
		hasParams := false
		if endpoint != "" {
			hasParams = true
			config.Cfg.AliOss.Endpoint = endpoint
		}

		if accessKeyID != "" {
			hasParams = true
			config.Cfg.AliOss.AccessKeyID = accessKeyID
		}

		if accessKeySecret != "" {
			hasParams = true
			config.Cfg.AliOss.AccessKeySecret = accessKeySecret
		}

		if bucket != "" {
			hasParams = true
			config.Cfg.AliOss.Bucket = bucket
		}

		if !hasParams {
			cmd.Help()
			os.Exit(1)
		}

		config.Update()
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.PersistentFlags().StringVarP(&endpoint, "endpoint", "e", "", "")
	configCmd.PersistentFlags().StringVarP(&accessKeyID, "accessKeyID", "k", "", "")
	configCmd.PersistentFlags().StringVarP(&accessKeySecret, "accessKeySecret", "s", "", "")
	configCmd.PersistentFlags().StringVarP(&bucket, "bucket", "b", "", "")

	viper.BindPFlag("cloud.aliOss.endPoint", configCmd.PersistentFlags().Lookup("endpoint"))
	viper.BindPFlag("cloud.aliOss.accessKeyID", configCmd.PersistentFlags().Lookup("accessKeyID"))
	viper.BindPFlag("cloud.aliOss.accessKeySecret", configCmd.PersistentFlags().Lookup("accessKeySecret"))
	viper.BindPFlag("cloud.aliOss.bucket", configCmd.PersistentFlags().Lookup("bucket"))

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
