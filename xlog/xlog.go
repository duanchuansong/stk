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

type Config struct {
	LogPath string
	LogFileName string
	LogLevel string
	LogOutFile bool
}

var LogConfig = &Config{
	LogPath:     "",
	LogFileName: "",
	LogLevel:    "info",
	LogOutFile:false,
}

func Init(cfg *Config) {
	configLogLevel := cfg.LogLevel //等级
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
	level := getLoggerLevel(configLogLevel) //日志等级
	var core zapcore.Core
	if cfg.LogOutFile {
		filePath := cfg.LogPath
		// configFilePath 文件夹位置
		configFilePath := cfg.LogFileName
		if configFilePath != "" {
			if strings.HasSuffix(configFilePath, "/") { //判断是否以/结尾
				filePath = configFilePath + filePath
			} else {
				filePath = configFilePath + "/" + filePath
			}
		}
		hook := lumberjack.Logger{
			Filename: filePath, // 日志文件路径
			MaxSize:  128,      // 每个日志文件保存的最大尺寸 单位：M
			//LocalTime: true,
			MaxBackups: 100,  // 日志文件最多保存多少个备份
			MaxAge:     30,   // 文件最多保存多少天
			Compress:   true, // 是否压缩
		}
		syncWriter := zapcore.AddSync(&hook)
		core = zapcore.NewTee(
			zapcore.NewCore(zapcore.NewJSONEncoder(encoder), syncWriter, zap.NewAtomicLevelAt(level)),
			zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
			zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority),
		)
	}else{
		core = zapcore.NewTee(
			zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
			zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority),
		)
	}
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	errorLogger = logger.Sugar()
}

func Debug(args ...interface{}) {
	if errorLogger == nil {
		Init(LogConfig)
	}
	errorLogger.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	if errorLogger == nil {
		Init(LogConfig)
	}
	errorLogger.Debugf(template, args...)
}

func Info(args ...interface{}) {
	if errorLogger == nil {
		Init(LogConfig)
	}
	errorLogger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	if errorLogger == nil {
		Init(LogConfig)
	}
	errorLogger.Infof(template, args...)
}

func Warn(args ...interface{}) {
	if errorLogger == nil {
		Init(LogConfig)
	}
	errorLogger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	if errorLogger == nil {
		Init(LogConfig)
	}
	errorLogger.Warnf(template, args...)
}

func Error(args ...interface{}) {
	if errorLogger == nil {
		Init(LogConfig)
	}
	errorLogger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	if errorLogger == nil {
		Init(LogConfig)
	}
	errorLogger.Errorf(template, args...)
}

func DPanic(args ...interface{}) {
	if errorLogger == nil {
		Init(LogConfig)
	}
	errorLogger.DPanic(args...)
}

func DPanicf(template string, args ...interface{}) {
	if errorLogger == nil {
		Init(LogConfig)
	}
	errorLogger.DPanicf(template, args...)
}

func Panic(args ...interface{}) {
	if errorLogger == nil {
		Init(LogConfig)
	}
	errorLogger.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	if errorLogger == nil {
		Init(LogConfig)
	}
	errorLogger.Panicf(template, args...)
}

func Fatal(args ...interface{}) {
	if errorLogger == nil {
		Init(LogConfig)
	}
	errorLogger.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	if errorLogger == nil {
		Init(LogConfig)
	}
	errorLogger.Fatalf(template, args...)
}
