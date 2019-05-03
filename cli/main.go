package main

import (
	"github.com/sdvdxl/dbox/cli/cmd"
	"github.com/sdvdxl/dbox/api/dao"
	"github.com/sdvdxl/dbox/api/log"
)

func main() {
	defer log.Close()
	defer dao.Close()
	cmd.Execute()
}
