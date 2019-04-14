package ex

import (
	"encoding/json"
	"fmt"
)

var FileNotExistErr Error = FileNotExist{desc: "file not exist", code: CodeFileNotExist}

// FileNotExist 文件不存在
type FileNotExist struct {
	desc string
	code int
}

func (e FileNotExist) MarshalJSON() ([]byte, error) {
	return json.Marshal(ErrMsg{Code: e.code, Desc: e.desc})
}

func (e FileNotExist) Arg(args ...interface{}) error {
	e.desc = addDesc(e.desc+", file: ", args...)
	return e
}

func (e FileNotExist) Error() string {
	return fmt.Sprintf("desc:%s, Code:%v", e.desc, e.code)
}
