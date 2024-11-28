package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"time"
)

type LoggerConfig struct {
	EnableStdOutput  bool // 是否启用标准输出
	EnableFileOutput bool // 是否启用文件输出

	Prefix   string // 日志前缀
	LogFlags int    // 日志标志
	FileDir  string // 日志文件所在目录
	FileName string // 日志文件名称
}

func NewLoggerConfig(err bool) *LoggerConfig {
	if err {
		return &LoggerConfig{
			EnableStdOutput:  true,
			EnableFileOutput: true,
			Prefix:           "ERROR\t",
			LogFlags:         log.LstdFlags | log.Lshortfile,
			FileDir:          "log/error",
			FileName:         "error_" + time.Now().Format("2006-01-02_15-04-05") + ".log",
		}
	} else {
		return &LoggerConfig{
			EnableStdOutput:  true,
			EnableFileOutput: false,
			Prefix:           "INFO\t",
			FileDir:          "log/info",
			LogFlags:         log.LstdFlags | log.Lshortfile,
		}
	}

}

func (loggerConfig *LoggerConfig) GetLogger() (*log.Logger, error) {
	var writers []io.Writer
	if loggerConfig.EnableStdOutput {
		writers = append(writers, os.Stdout)
	}
	if loggerConfig.EnableFileOutput {
		filePath := path.Join(loggerConfig.FileDir, loggerConfig.FileName)
		err := createDirIfNotExist(filePath)
		if err != nil {
			return nil, err
		}
		file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return nil, fmt.Errorf("failed to open log file: %w", err)
		}
		writers = append(writers, file)
	}
	if len(writers) < 1 {
		return nil, fmt.Errorf("no logger configured for \"%s\"", loggerConfig.Prefix)
	}
	multiWriter := io.MultiWriter(writers...)
	return log.New(multiWriter, loggerConfig.Prefix, loggerConfig.LogFlags), nil
}

type Logger struct {
	Info  *log.Logger
	Error *log.Logger
}

// NewLogger 根据配置创建日志记录器
func NewLogger(infoLoggerConfig *LoggerConfig, errorLoggerConfig *LoggerConfig) (*Logger, error) {
	if infoLoggerConfig == nil {
		infoLoggerConfig = NewLoggerConfig(false)
	}
	if errorLoggerConfig == nil {
		errorLoggerConfig = NewLoggerConfig(true)
	}
	infoLogger, err := infoLoggerConfig.GetLogger()
	if err != nil {
		return nil, err
	}
	errorLogger, err := errorLoggerConfig.GetLogger()
	if err != nil {
		return nil, err
	}
	return &Logger{
		Info:  infoLogger,
		Error: errorLogger,
	}, nil
}

// createDirIfNotExist 检查并创建目录路径
func createDirIfNotExist(filePath string) error {
	dir := filepath.Dir(filePath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		// 如果目录不存在，则创建
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}
	}
	return nil
}
