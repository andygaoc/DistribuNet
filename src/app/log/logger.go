package log

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"time"
)

var Logger *zap.Logger

func init() {
	SetLogger("log/")
}
func SetLogger(LogDir string) {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	//encoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	encoderConfig.LevelKey = ""
	encoderConfig.TimeKey = "time"
	encoderConfig.MessageKey = "prefix"

	level := zap.NewAtomicLevel()
	level.SetLevel(zapcore.DebugLevel)

	writer := getWriter(LogDir, "emperor")
	jsonEncoder := zapcore.NewJSONEncoder(encoderConfig)
	core := zapcore.NewTee(
		zapcore.NewCore(jsonEncoder, zapcore.AddSync(writer), level),
	)
	Logger = zap.New(core, zap.AddCaller())
	defer Logger.Sync()
}

func Any(k string, v interface{}) zapcore.Field {
	return zap.Any(k, v)
}
func String(k string, v string) zapcore.Field {
	return zap.String(k, v)
}
func Int64(k string, v int64) zapcore.Field {
	return zap.Int64(k, v)
}
func Error(err error) zapcore.Field {
	return zap.Error(err)
}

func getWriter(log_dir, prefix string) io.Writer {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	p := log_dir + prefix + "_%Y%m%d_" + hostname + ".log"

	hook, err := rotatelogs.New(
		p,
		rotatelogs.WithMaxAge(time.Hour*24*7),
	)

	if err != nil {
		panic(err)
	}
	return hook
}
