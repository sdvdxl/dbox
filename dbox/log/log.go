package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.SugaredLogger

func init() {
	Init()
}

func Init() {
	l, _ := zap.NewDevelopment(zap.AddCaller(), zap.AddStacktrace(zapcore.WarnLevel))
	Log = l.Sugar()
}
