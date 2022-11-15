// Package logger
// descr 处理日志相关逻辑
// author fm
// date 2022/11/15 17:12
package logger

import (
	"os"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gohub-lesson/pkg/app"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Logger 全局 Logger 对象
var Logger *zap.Logger

// InitLogger 日志初始化
func InitLogger(
	filename string,
	maxSize, maxBackup, maxAge int,
	compress bool,
	logType, level string,
) {

	// 1. 获取日志写入介质
	writerSync := getLogWriter(filename, maxSize, maxBackup, maxAge, compress, logType)

	// 2. 设置日志等级，具体请见 config/log.go 文件
	logLevel := new(zapcore.Level)

	if err := logLevel.UnmarshalText([]byte(level)); err != nil {
		panic("日志初始化错误，日志级别设置有误。请修改 config/log.go 文件中的 log.level 配置项")
	}

	// 3. 初始化 core
	core := zapcore.NewCore(getZapEncoder(), writerSync, logLevel)

	// 4. 初始化 Logger
	Logger = zap.New(
		core,
		// 调用文件和行号，内部使用 runtime.Caller
		zap.AddCaller(),
		// 封装了一层，调用文件去除一层(runtime.Caller(1))
		zap.AddCallerSkip(1),
		// Error 时才会显示 stacktrace
		zap.AddStacktrace(zap.ErrorLevel),
	)

	// 5. 将自定义的 logger 替换为全局的 logger
	// zap.L().Fatal() 调用时，就会使用我们自定的 Logger
	zap.ReplaceGlobals(Logger)
}

// getLogWriter 日志记录介质。项目中使用了两种介质，os.Stdout和文件
func getLogWriter(
	filename string,
	maxSize, maxBackup, maxAge int,
	compress bool,
	logType string,
) zapcore.WriteSyncer {

	// 如果配置了按照日期记录日志文件
	if logType == "daily" {
		logName := time.Now().Format("2006-01-02.log")
		filename = strings.ReplaceAll(filename, "logs.log", logName)
	}

	// 滚动日志，详见 config/log.go
	lumberjackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxAge:     maxAge,
		MaxBackups: maxBackup,
		Compress:   compress,
	}

	if app.IsLocal() {
		// 本地开发终端打印和记录文件
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberjackLogger))
	}

	// 生产环境只记录文件
	return zapcore.AddSync(lumberjackLogger)
}

// getZapEncoder 获取日志存储格式
func getZapEncoder() zapcore.Encoder {

	// 初始化配置
	encodeConfig := zapcore.EncoderConfig{
		MessageKey:    "message",
		LevelKey:      "level",
		TimeKey:       "time",
		NameKey:       "logger",
		CallerKey:     "caller",
		FunctionKey:   zapcore.OmitKey,
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		// 日志级别名称大写，如 ERROR、INFO
		EncodeLevel: zapcore.CapitalLevelEncoder,
		// 时间格式，自定义为 2006-01-02 15:04:05
		EncodeTime: customTimeEncoder,
		// 执行时间，以秒为单位
		EncodeDuration: zapcore.SecondsDurationEncoder,
		// Caller 短格式，如：types/converter.go:17，长格式为绝对路径
		EncodeCaller: zapcore.ShortCallerEncoder,
	}

	// 本地环境
	if app.IsLocal() {
		// 终端输出的关键词高亮
		encodeConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		// 本地设置内置的 console解码器（支持 stacktrace 换行）
		return zapcore.NewConsoleEncoder(encodeConfig)
	}

	return zapcore.NewJSONEncoder(encodeConfig)
}

// customTimeEncoder 自定义友好的时间格式
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}
