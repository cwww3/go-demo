package main

import (
	"github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

func main() {
	r, _ := rotatelogs.New("ZAP/1.log")
	r.Write([]byte("sss"))
	r.Rotate()
	r.Write([]byte("sss"))
	//log := zap.NewExample()
	//fmt.Println("core-", log.Core())
	//log.Info("hhh")
	w := zapcore.AddSync(os.Stdout)
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	//encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), w, zap.DebugLevel)
	log := zap.New(core, zap.AddCaller()).Sugar()
	log.Info("111")
	strings.HasPrefix()
}
