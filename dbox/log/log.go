package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
)

var Log *zap.SugaredLogger

func init() {
	Init()
}

func Init() {
	l, _ := zap.NewDevelopment(zap.AddCaller(), zap.AddStacktrace(zapcore.WarnLevel))
	Log = l.Sugar()
}

type Logger zap.SugaredLogger

func (Logger) Print(v ...interface{}) {

	strs := make([]string, len(v))
	for i, v := range v {
		strs[i] = fmt.Sprint(v)
	}
	Log.Info(strings.Join(strs, ", "))
}
