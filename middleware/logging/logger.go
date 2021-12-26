package logging

import (
	"encoding/json"

	"os"
	"time"

	"go.uber.org/zap"
)

const (
	Debug = iota + 1
	Info
	Waring
	Error
)

var HostName, _ = os.Hostname()

type Logger struct {
	z       *zap.Logger
	options *Options
	f       []Filed
	Msg     string
	Level   int
}

// Default 默认logger
func Default() Logger {
	logger, _ := zap.NewProduction()
	return Logger{z: logger}
}

// 根据配置生成logger实例
func NewLogger(options ...Option) *Logger {
	logger := new(Logger)
	op := &Options{}
	for _, o := range options {
		o(op)
	}
	b, err := json.Marshal(op)
	if err != nil {
		panic(err)
	}
	var cfg zap.Config
	if err := json.Unmarshal(b, &cfg); err != nil {
		panic(err)
	}
	z, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	logger.z = z
	return logger
}

type Filed zap.Field

func (t *Logger) MustAppend() *Logger {
	t.f = []Filed{}
	// 初始化调用时间戳
	now := time.Now().UnixNano()
	t.f = append(t.f, Filed(zap.Int64("timestamp", now)))
	// 时间
	date := time.Now().Format("2006-01-02 15:04:05")
	t.f = append(t.f, Filed(zap.String("date", date)))
	// tracId
	t.f = append(t.f, Filed(zap.String("traceId", "")))

	// 执行host
	t.f = append(t.f, Filed(zap.String("host", HostName)))
	return t
}

// string 将数据转为string
func (t *Logger) String(key, value string) *Logger {
	t.f = append(t.f, Filed(zap.String(key, value)))
	return t
}

// Int 类型的值
func (t *Logger) Int(key string, value int) *Logger {
	t.f = append(t.f, Filed(zap.Int(key, value)))
	return t
}

// 刷新
func (t *Logger) Flush() {
	zf := []zap.Field{}
	for _, f := range t.f {
		zf = append(zf, zap.Field(f))
	}
	t.z.Info(t.Msg, zf...)
	t.z.Sync()
}
