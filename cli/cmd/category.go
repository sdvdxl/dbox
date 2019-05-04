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
	"github.com/sdvdxl/dbox/api/service"
	"os"

	"github.com/spf13/cobra"
)

// categoryCmd represents the category command
var categoryCmd = &cobra.Command{
	Use:   "category",
	Short: "manage category",
	Run: func(cmd *cobra.Command, args []string) {
		checkConfig()
		categoryService := service.Category{}
		rows := make([][]string, 0, 0)
		for _, v := range categoryService.FindAll() {
			row := make([]string, 3)
			row[0] = fmt.Sprint(v.ID)
			row[1] = fmt.Sprint(v.Name)
			rows = append(rows, row)
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Category"})
		//table.SetFooter([]string{"", "", "Total", "$146.93"}) // Add Footer
		table.SetBorder(false)
		// Set Border to false
		table.SetAutoWrapText(false)
		table.AppendBulk(rows)
		// Add Bulk Data
		//table.SetFooter()
		table.Render()
		fmt.Println()

	},
}

func init() {
	rootCmd.AddCommand(categoryCmd)
}
