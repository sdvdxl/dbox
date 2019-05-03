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
	"github.com/olekukonko/tablewriter"
	"github.com/sdvdxl/dbox/api/log"
	"github.com/sdvdxl/dbox/api/model"
	"github.com/sdvdxl/dbox/api/service/file"
	"github.com/spf13/pflag"
	"golang.org/x/exp/errors/fmt"
	"os"

	"github.com/spf13/cobra"
)

// findCmd represents the find command
var findCmd = &cobra.Command{
	Use:   "find",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.PersistentFlags().VisitAll(func(flag *pflag.Flag) {
			log.Log.Info(flag.Name, " ", flag.Value)
		})

		checkConfig()
		if findCondition.Name == "" {
			cmd.Help()
			os.Exit(1)
		}

		rows := make([][]string, 0, 0)
		for _, v := range fileService.FindByFuzz(findCondition) {
			row := make([]string, 3)
			row[0] = fmt.Sprint(v.ID)
			row[1] = fmt.Sprint(v.Category)
			row[2] = fmt.Sprint(v.Name)
			rows = append(rows, row)
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"文件序号ID", "文件夹", "文件名"})
		//table.SetFooter([]string{"", "", "Total", "$146.93"}) // Add Footer
		table.SetBorder(false) // Set Border to false
		table.SetAutoWrapText(false)
		table.AppendBulk(rows) // Add Bulk Data
		//table.SetFooter()
		table.Render()
		fmt.Println()
	},
}

var (
	findCondition model.FileDTO
)

func init() {
	rootCmd.AddCommand(findCmd)
	findCmd.PersistentFlags().StringVarP(&findCondition.Category, "category", "c", "", "文件夹")
	findCmd.PersistentFlags().StringVarP(&findCondition.Name, "file", "f", "", "* 文件名字")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// findCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// findCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
