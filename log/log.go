package log

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger        *zap.SugaredLogger
	once          sync.Once
	filePath      string
	errorFilePath string
)

var devOptions = []zap.Option{
	zap.WithCaller(true),
	zap.AddCallerSkip(1),
}

func fullPathEncodeCaller(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(caller.FullPath())
}

func makeDevConfig() zap.Config {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.DisableStacktrace = true
	config.Level = zap.NewAtomicLevelAt(getLogLevel())
	config.EncoderConfig.EncodeCaller = fullPathEncodeCaller
	return config
}

func makeDevConfigWriteOutputToFile() zap.Config {
	config := makeDevConfig()
	if os.Getenv("LOG_OUTPUT") == "ALL" && os.Getenv("LOG_OUTPUT_FILE") != "" {
		config.OutputPaths = []string{filePath}
	} else if os.Getenv("LOG_OUTPUT") == "ERROR" && os.Getenv("LOG_OUTPUT_ERROR_FILE") != "" {
		config.Level = zap.NewAtomicLevelAt(zapcore.ErrorLevel) // prints only error level messages
		config.OutputPaths = []string{errorFilePath}
	}
	return config
}

var prodOptions = []zap.Option{
	// add production log settings here(if any)
}

func makeProdConfig() zap.Config {
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(getLogLevel())
	return config
}

func getLogLevel() zapcore.Level {
	switch strings.ToLower(os.Getenv("LOG_LEVEL")) {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warning":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}

func GetLogger() *zap.SugaredLogger {
	Init()
	return logger
}

func createLogFile() error {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("error getting work directory path")
		return err
	}

	dirPath := wd + "/" + "log_output"

	err = os.MkdirAll(dirPath, 0750)
	if err != nil {
		fmt.Printf("error creating directory, err: %v\n", err)
		return err
	}

	if os.Getenv("LOG_OUTPUT") == "ALL" && os.Getenv("LOG_OUTPUT_FILE") != "" {
		filePath = fmt.Sprintf("%v/%v.log", dirPath, os.Getenv("LOG_OUTPUT_FILE"))
		if _, err = os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644); err != nil {
			fmt.Printf("error opening file, err: %v\n", err)
			return err
		}
	}

	if os.Getenv("LOG_OUTPUT") == "ERROR" && os.Getenv("LOG_OUTPUT_ERROR_FILE") != "" {
		errorFilePath = fmt.Sprintf("%v/%v.log", dirPath, os.Getenv("LOG_OUTPUT_ERROR_FILE"))
		if _, err = os.OpenFile(errorFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644); err != nil {
			fmt.Printf("error opening error file, err: %v\n", err)
			return err
		}
	}

	return nil
}

func makeLogger() (*zap.Logger, error) {
	if strings.ToLower(os.Getenv("ENV")) == "production" {
		return makeProdConfig().Build(prodOptions...)
	}
	if os.Getenv("LOG_OUTPUT") != "" {
		err := createLogFile()
		if err == nil {
			return makeDevConfigWriteOutputToFile().Build(devOptions...)
		}
	}
	return makeDevConfig().Build(devOptions...)
}

func Init() {
	once.Do(func() {
		zap_logger, err := makeLogger()
		if err != nil {
			panic(err)
		}

		logger = zap_logger.Sugar()
		logger.Infof("Initialized logger at log level: %s", zap_logger.Level())
		go func() {
			ticker := time.NewTicker(time.Second)
			for range ticker.C {
				_ = logger.Sync()
			}
		}()
	})
}

func Debugf(template string, args ...interface{}) {
	GetLogger().Debugf("\n"+template, args...)
}

func Infof(template string, args ...interface{}) {
	GetLogger().Infof("\n"+template, args...)
}

func Warnf(template string, args ...interface{}) {
	GetLogger().Warnf("\n"+template, args...)
}

func Errorf(template string, args ...interface{}) {
	GetLogger().Errorf("\n"+template, args...)
}

func Panicf(template string, args ...interface{}) {
	GetLogger().Panicf("\n"+template, args...)
}

func Fatalf(template string, args ...interface{}) {
	GetLogger().Fatalf("\n"+template, args...)
}
