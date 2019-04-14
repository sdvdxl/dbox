package ex

import (
	"encoding/json"
	"fmt"
)

var FileErr Error = FileError{desc: "file error", code: CodeFile}

type FileError struct {
	code int
	desc string
}

func (e FileError) Arg(args ...interface{}) error {
	e.desc = addDesc(e.desc+", file: ", args...)
	return e
}

func (e FileError) Error() string {
	return fmt.Sprintf("desc:%s, Code:%v", e.desc, e.code)
}

func (e FileError) MarshalJSON() ([]byte, error) {
	return json.Marshal(ErrMsg{Code: e.code, Desc: e.desc})
}
