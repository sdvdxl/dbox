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
	"github.com/sdvdxl/dbox/api/ex"
	"github.com/sdvdxl/dbox/api/log"
	"github.com/sdvdxl/dbox/api/model"
	"github.com/sdvdxl/dbox/api/service"
	"github.com/spf13/cobra"
	"os"
)

var (
	category string
	file     string
	fileName string
)

// putCmd represents the upload command
var putCmd = &cobra.Command{
	Use:   "put",
	Short: "upload file to cloud",
	Example:"put beauty.jpg -c Pictures",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		checkConfig()
		file = args[0]
		if f, err := os.Stat(file); err != nil {
			fmt.Println(err.Error() + ": " + file)
			os.Exit(1)
		} else {
			if f.IsDir() {
				fmt.Println("can not be folder")
				os.Exit(1)
			}
		}

		fileService := &service.FileService{}
		if f, err := fileService.UploadLocalFile(file, fileName, category); err != nil {
			if _, ok := err.(ex.FileExist); ok {
				fmt.Println("file already exist")
				fDTO := model.FileDTO{}
				fDTO.Name = f.Name
				printTables(fileService.FindByFuzz(fDTO))
				os.Exit(1)
			}

			fmt.Println(err)
			log.Log.Error("upload file error:", err, "file:", file, "category:", category)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(putCmd)
	putCmd.PersistentFlags().StringVarP(&category, "category", "c", "default", "文件夹")
	putCmd.PersistentFlags().StringVarP(&fileName, "filename", "f", "", "file name")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// putCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// putCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
