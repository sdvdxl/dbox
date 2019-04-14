package ex

import "fmt"

type ErrMsg struct {
	Desc string
	Code int
}

type Error interface {
	Arg(args ...interface{}) error
	Error() string
	MarshalJSON() ([]byte, error)
}

func addDesc(oldDesc string, args ...interface{}) string {
	if len(args) == 0 {
		return oldDesc
	}

	var str string
	for _, v := range args {
		str += fmt.Sprint(v)
	}

	return oldDesc + str
}

func Check(err error) {
	if err != nil {
		panic(err.Error())
	}
}
