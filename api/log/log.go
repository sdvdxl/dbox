package log

import (
	"fmt"
	"github.com/sdvdxl/dbox/api/config"
	"github.com/sdvdxl/dbox/api/ex"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"strings"
)

var Log *zap.SugaredLogger

func Close() {
	if Log != nil {
		ex.Check(Log.Sync())
	}
}

func Init() {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   config.GetBasePath() + "dbox.log",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
	})
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
		w,
		zap.InfoLevel,
	)
	l := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.WarnLevel))
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
