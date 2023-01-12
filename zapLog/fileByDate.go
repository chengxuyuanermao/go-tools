package zapLog

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"time"
)

// mines
var MLog *zap.SugaredLogger // 暴露出去给mines调用记录
var minesFilePath = "./log/mines/"
var minesFileName = "mines"
var minesLogLevel = zapcore.DebugLevel // debug级别的都将被写入

// fruitmatch
var FMLog *zap.SugaredLogger // 暴露出去给mines调用记录
var FruitMatchFilePath = "./log/fruitmatch/"
var FruitMatchFileName = "fruitmatch"
var FruitMatchLogLevel = zapcore.DebugLevel // debug级别的都将被写入

// otherGames。。。

func InitGameLog() {
	// mines
	MLog = initLogger(minesFilePath, minesFileName, minesLogLevel)
	defer MLog.Sync()
	// fruitmatch
	FMLog = initLogger(FruitMatchFilePath, FruitMatchFileName, FruitMatchLogLevel)
	defer FMLog.Sync()

	// otherGames。。。
}

func initLogger(filePath string, fileName string, logLevel zapcore.Level) *zap.SugaredLogger {
	//writeSyncer := getLogWriter(filePath, fileName)
	writeSyncer := getLogWriterByDate(filePath, fileName)
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, logLevel)

	// 生成实例。使用zap.New(…)方法来手动传递所有配置
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(0))
	// 实例化赋值给全局变量
	return logger.Sugar()
}

func getEncoderV2() zapcore.Encoder {
	// 编码器(如何写入日志)。
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder   // 修改时间编码器
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // 在日志文件中使用大写字母记录日志级别
	return zapcore.NewConsoleEncoder(encoderConfig)         // 或使用json格式：NewJSONEncoder()。前面是类似nginx的格式
}

func getLogWriterV2(filepath, filename string) zapcore.WriteSyncer {
	fileName := filepath + filename + "_" + time.Now().Format("2006-01-02-150405") + ".log" //生成文件的路径
	/**
	Filename: 日志文件的位置
	MaxSize：在进行切割之前，日志文件的最大大小（以MB为单位）
	MaxBackups：保留旧文件的最大个数
	MaxAges：保留旧文件的最大天数
	Compress：是否压缩/归档旧文件
	*/
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    20,
		MaxBackups: 20,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

// 使用file-rotatelogs按天切割日志
func getLogWriterByDate(filepath, filename string) zapcore.WriteSyncer {
	path := filepath + filename
	l, _ := rotatelogs.New(
		path+"-%Y%m%d-%H%M",
		rotatelogs.WithMaxAge(15*24*time.Hour),    // 最长保存30天
		rotatelogs.WithRotationTime(time.Hour*24), // 24小时切割一次
	)
	return zapcore.AddSync(l)
}
