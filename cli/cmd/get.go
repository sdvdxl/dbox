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
	"github.com/sdvdxl/dbox/api/service"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	filename string
	dir      string
)
// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "down load file with file id",
	Run: func(cmd *cobra.Command, args []string) {
		checkConfig()

		fileService := &service.FileService{}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("file id is number")
			os.Exit(1)
		}

		f, err := fileService.Download(id, dir, filename)
		if err != nil {
			fmt.Println(err, f)
			os.Exit(1)
		}

		fmt.Println()
		fmt.Println("file download success: ", f)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.PersistentFlags().StringVarP(&dir, "dir", "d", "", "dir to save")
	getCmd.PersistentFlags().StringVarP(&filename, "filename", "f", "", "file name")

}
