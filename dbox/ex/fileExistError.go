package ex

import (
	"encoding/json"
	"fmt"
)

var FileExistErr Error = FileExist{desc: "file already exist", code: CodeFileNotExist}

// FileNotExist 文件不存在
type FileExist struct {
	desc string
	code int
}

func (e FileExist) MarshalJSON() ([]byte, error) {
	return json.Marshal(ErrMsg{Code: e.code, Desc: e.desc})
}

func (e FileExist) Arg(args ...interface{}) error {
	e.desc = addDesc(e.desc+", file: ", args...)
	return e
}

func (e FileExist) Error() string {
	return fmt.Sprintf("desc:%s, Code:%v", e.desc, e.code)
}
