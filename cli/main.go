package main

import (
	"github.com/sdvdxl/dbox/cli/cmd"
	"github.com/sdvdxl/dbox/dbox/dao"
	"github.com/sdvdxl/dbox/dbox/log"
)

func main() {
	defer log.Close()
	defer dao.Close()
	cmd.Execute()
}
