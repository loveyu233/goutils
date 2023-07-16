package log

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type Zap struct {
	Mode       string //开发模式
	Level      zapcore.Level
	FileName   string // 日志存放路径
	MaxSize    int    // 最大日志存储大小
	MaxAge     int    // 最大存储天数
	MaxBackups int    // 最大备份数量
}

func (z Zap) InitZapLogger() (err error) {
	// 创建Core三大件，进行初始化
	writeSyncer := getLogWriter(z.FileName, z.MaxSize, z.MaxAge, z.MaxBackups)
	encoder := getEncoder()
	// 创建核心-->如果是dev模式，就在控制台和文件都打印，否则就只写到文件中
	var core zapcore.Core
	if z.Mode == "dev" {
		// 开发模式，日志输出到终端
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		// NewTee创建一个核心，将日志条目复制到两个或多个底层核心中。
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, writeSyncer, z.Level),
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), z.Level),
		)
	} else {
		core = zapcore.NewCore(encoder, writeSyncer, z.Level)
	}
	// 创建 logger 对象
	log := zap.New(core, zap.AddCaller())
	// 替换全局的 logger, 后续在其他包中只需使用zap.L()调用即可
	zap.ReplaceGlobals(log)
	return
}

// 获取Encoder，给初始化logger使用的
func getEncoder() zapcore.Encoder {
	// 使用zap提供的 NewProductionEncoderConfig
	encoderConfig := zap.NewProductionEncoderConfig()
	// 设置时间格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 时间的key
	encoderConfig.TimeKey = "time"
	// 级别
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// 显示调用者信息
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	// 返回json 格式的 日志编辑器
	return zapcore.NewJSONEncoder(encoderConfig)
}

// 获取切割的问题，给初始化logger使用的
func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	// 使用 lumberjack 归档切片日志
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}
