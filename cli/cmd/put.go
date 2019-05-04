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
	"github.com/olekukonko/tablewriter"
	"github.com/sdvdxl/dbox/api/log"
	"github.com/sdvdxl/dbox/api/model"
	"github.com/sdvdxl/dbox/api/service"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	category    string
	file        string
	fileName    string
	isAllFile   bool
	isRecursion bool
)

// putCmd represents the upload command
var putCmd = &cobra.Command{
	Use:     "put",
	Short:   "upload file to cloud",
	Example: "put beauty.jpg -c Pictures",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		checkConfig()
		file = args[0]
		f, err := os.Stat(file)
		if err != nil {
			fmt.Println(err.Error() + ": " + file)
			os.Exit(1)
		} else {
			if f.IsDir() {
				if !isAllFile && !isRecursion {
					fmt.Println("can not be folder, if you want to upload all files, must be with flag `-a` or `-r`")
					os.Exit(1)
				}

				// 如果是目录，不能指定文件名，设置为空，使用默认值
				filename = ""
			}
		}

		fileService := &service.FileService{}
		files := make([]string, 0, 0)
		if !f.IsDir() {
			files = append(files, file)
		} else {
			if isRecursion {
				err = filepath.Walk(file, func(path string, info os.FileInfo, err error) error {
					if info.IsDir() {
						return nil
					}

					files = append(files, path)
					return nil
				})
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}

			} else {
				fs, err := ioutil.ReadDir(file)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}

				for _, v := range fs {
					if v.IsDir() {
						continue
					}
					
					fpath := file + string(filepath.Separator) + v.Name()
					log.Log.Info("recursion -> file:", fpath, " isDir:", v.IsDir())
					files = append(files, fpath)
				}
			}

		}

		result := make(map[*model.File]error, 0)
		for _, file := range files {
			f, err := fileService.UploadLocalFile(file, fileName, category)
			if f == nil {
				f = &model.File{Name: file}
			}
			result[f] = err
		}

		fmt.Println("\nupload result")
		rows := make([][]string, 0, 0)
		for f, e := range result {
			row := make([]string, 3)
			r := "success"
			if e != nil {
				r = e.Error()
			}
			row[0] = fmt.Sprint(f.ID)
			row[1] = fmt.Sprint(f.Name)
			row[2] = fmt.Sprint(r)
			rows = append(rows, row)
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"File ID", "File Name", "Result"})
			//table.SetFooter([]string{"", "", "Total", "$146.93"}) // Add Footer
			table.SetBorder(false)
			// Set Border to false
			table.SetAutoWrapText(false)
			table.AppendBulk(rows)
			// Add Bulk Data
			//table.SetFooter()
			table.Render()
			fmt.Println()
		}

	},
}

func init() {
	rootCmd.AddCommand(putCmd)
	putCmd.PersistentFlags().StringVarP(&category, "category", "c", "default", "文件夹")
	putCmd.PersistentFlags().StringVarP(&fileName, "filename", "f", "", "file name")
	putCmd.PersistentFlags().BoolVarP(&isAllFile, "all", "a", false, "upload all files of the folder")
	putCmd.PersistentFlags().BoolVarP(&isRecursion, "recursion", "r", false, "upload all files of the folder and all files of sub folders")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// putCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// putCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
