package xlog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strings"
)

// error logger
var errorLogger *zap.SugaredLogger

var levelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func getLoggerLevel(lvl string) zapcore.Level {
	if level, ok := levelMap[lvl]; ok {
		return level
	}
	return zapcore.InfoLevel
}

func Init() {
	filePath := "log.log"
	// configFilePath 文件夹位置
	configFilePath := "./logs"
	if configFilePath != "" {
		if strings.HasSuffix(configFilePath, "/") { //判断是否以/结尾
			filePath = configFilePath + filePath
		} else {
			filePath = configFilePath + "/" + filePath
		}
	}
	configLogLevel := "info" //等级

	level := getLoggerLevel(configLogLevel) //日志等级
	hook := lumberjack.Logger{
		Filename: filePath, // 日志文件路径
		MaxSize:  128,      // 每个日志文件保存的最大尺寸 单位：M
		//LocalTime: true,
		MaxBackups: 100,  // 日志文件最多保存多少个备份
		MaxAge:     30,   // 文件最多保存多少天
		Compress:   true, // 是否压缩
	}
	syncWriter := zapcore.AddSync(&hook)
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleErrors := zapcore.Lock(os.Stderr)
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewJSONEncoder(encoder), syncWriter, zap.NewAtomicLevelAt(level)),
		zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
		zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	errorLogger = logger.Sugar()
}

func Debug(args ...interface{}) {
	if errorLogger == nil {
		Init()
	}
	errorLogger.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	if errorLogger == nil {
		Init()
	}
	errorLogger.Debugf(template, args...)
}

func Info(args ...interface{}) {
	if errorLogger == nil {
		Init()
	}
	errorLogger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	if errorLogger == nil {
		Init()
	}
	errorLogger.Infof(template, args...)
}

func Warn(args ...interface{}) {
	if errorLogger == nil {
		Init()
	}
	errorLogger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	if errorLogger == nil {
		Init()
	}
	errorLogger.Warnf(template, args...)
}

func Error(args ...interface{}) {
	if errorLogger == nil {
		Init()
	}
	errorLogger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	if errorLogger == nil {
		Init()
	}
	errorLogger.Errorf(template, args...)
}

func DPanic(args ...interface{}) {
	if errorLogger == nil {
		Init()
	}
	errorLogger.DPanic(args...)
}

func DPanicf(template string, args ...interface{}) {
	if errorLogger == nil {
		Init()
	}
	errorLogger.DPanicf(template, args...)
}

func Panic(args ...interface{}) {
	if errorLogger == nil {
		Init()
	}
	errorLogger.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	if errorLogger == nil {
		Init()
	}
	errorLogger.Panicf(template, args...)
}

func Fatal(args ...interface{}) {
	if errorLogger == nil {
		Init()
	}
	errorLogger.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	if errorLogger == nil {
		Init()
	}
	errorLogger.Fatalf(template, args...)
}
