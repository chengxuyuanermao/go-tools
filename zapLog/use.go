package zapLog

import (
	"fmt"
	"gopkg.in/natefinch/lumberjack.v2"
	"net/http"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

/**
https://www.liwenzhou.com/posts/Go/zap/
*/

var sugarLogger *zap.SugaredLogger

func Use() {
	InitLogger()
	defer sugarLogger.Sync()

	// 请求示例
	simpleHttpGet("www.sogo.com")
	simpleHttpGet("http://www.sogo.com")
}

func InitLogger() {
	writeSyncer := getLogWriter("./log/zapLog/", "zapLog")
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel) // debug级别的都将被写入

	// 使用zap.New(…)方法来手动传递所有配置
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	// 实例化赋值给全局变量
	sugarLogger = logger.Sugar()
}

// zap提供了几个快速创建logger的方法，zap.NewExample()、zap.NewDevelopment()、zap.NewProduction()，还有高度定制化的创建方法zap.New()
func getEncoder() zapcore.Encoder {
	// 编码器(如何写入日志)。
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder   // 修改时间编码器
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // 在日志文件中使用大写字母记录日志级别
	return zapcore.NewConsoleEncoder(encoderConfig)         // 或使用json格式：NewJSONEncoder()。前面是类似nginx的格式
}

func getLogWriter(filepath, filename string) zapcore.WriteSyncer {
	fileName := filepath + filename + "_" + time.Now().Format("20060102150405") + ".log" //生成文件的路径
	/*
		Filename: 日志文件的位置
		MaxSize：在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups：保留旧文件的最大个数
		MaxAges：保留旧文件的最大天数
		Compress：是否压缩/归档旧文件
	*/
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func simpleHttpGet(url string) {
	a := "多读书"
	sugarLogger.Debug("aa %v", "vv", "我哦", "dd")
	sugarLogger.Errorf(fmt.Sprintf("未找到%v玩家", a))

	sugarLogger.Debugf("Trying to hit GET request for %s", url)
	resp, err := http.Get(url)
	if err != nil {
		sugarLogger.Errorf("Error fetching URL %s : Error = %s", url, err)
	} else {
		sugarLogger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
		resp.Body.Close()
	}
}
