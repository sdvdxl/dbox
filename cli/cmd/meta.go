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
	"github.com/spf13/cobra"
	"os"
)

// metaCmd represents the sync command
var metaCmd = &cobra.Command{
	Use:   "meta",
	Short: "sync db file between local and cloud",
	Long: `upload local db file to cloud, or download db file from cloud to local
available arg is one of upload, download, merge.
`,
	ValidArgs: []string{"upload", "download", "merge"},
	Example:   "sync upload",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 ||
			(args[0] != "upload" &&
				args[0] != "download" &&
				args[0] != "merge") {
			fmt.Println("args must be upload, download or merge")
			os.Exit(1)
		}

		checkConfig()

		fileService := &service.FileService{}
		if err := fileService.SyncDBFile(args[0]); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(metaCmd)
}
