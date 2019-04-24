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
	"github.com/sdvdxl/dbox/dbox/ex"
	"github.com/sdvdxl/dbox/dbox/log"
	"github.com/sdvdxl/dbox/dbox/service/file"
	"github.com/spf13/cobra"
	"os"
)

var (
	category string
	file     string
	fileName string
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "上传文件",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		checkConfig()
		if file == "" {
			fmt.Println("请填写文件名")
			ex.Check(cmd.Help())
			os.Exit(1)
		} else if f, err := os.Stat(file); err != nil {
			fmt.Println(err.Error() + ": " + file)
			os.Exit(1)
		} else {
			if f.IsDir() {
				fmt.Println("不能是文件夹")
				os.Exit(1)
			}
		}

		if err := fileService.UploadLocalFile(file,fileName, category); err != nil {
			fmt.Println(err)
			log.Log.Error("upload file error:", err, "file:", file, "category:", category)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)
	uploadCmd.PersistentFlags().StringVarP(&file, "file", "f", "", "* 文件路径（不能是文件夹）")

	uploadCmd.PersistentFlags().StringVarP(&category, "category", "c", "default", "文件夹")
	uploadCmd.PersistentFlags().StringVarP(&fileName, "name", "n", "", "文件名")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uploadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uploadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
