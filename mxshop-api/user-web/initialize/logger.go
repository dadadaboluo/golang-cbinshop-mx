package initialize

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func InitLogger() {

	//logger, _ := zap.NewProduction()
	//zap.ReplaceGlobals(logger)

	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	core := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel)

	// 创建 logger 并替换全局的 zap 记录器
	logger := zap.New(core)
	zap.ReplaceGlobals(logger)
}
