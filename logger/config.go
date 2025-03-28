package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"time"
)

type LogoLevel uint8

const (
	LogoLevelInfo  LogoLevel = iota // 普通信息
	LogoLevelWarn                   // 警告信息
	LogoLevelError                  // 错误信息
)

type Config struct {
	EnableStdOutput  bool   // 是否启用标准输出
	EnableFileOutput bool   // 是否启用文件输出
	Prefix           string // 日志前缀
	FileDir          string // 日志文件所在目录
	FileName         string // 日志文件名称
	LogFlags         int    // 日志标志
}

// NewDefaultConfig 创建一个默认日志配置
func NewDefaultConfig(logLevel LogoLevel) *Config {
	switch logLevel {
	case LogoLevelInfo:
		return &Config{
			EnableStdOutput:  true,
			EnableFileOutput: false,
			Prefix:           "INFO\t",
			LogFlags:         log.LstdFlags | log.Lshortfile,
		}
	case LogoLevelWarn:
		return &Config{
			EnableStdOutput:  true,
			EnableFileOutput: false,
			Prefix:           "WARN\t",
			LogFlags:         log.LstdFlags | log.Lshortfile,
		}
	case LogoLevelError:
		return &Config{
			EnableStdOutput:  true,
			EnableFileOutput: true,
			Prefix:           "ERROR\t",
			FileDir:          "log/error",
			FileName:         "error_" + time.Now().Format("2006-01-02_15-04-05") + ".log",
			LogFlags:         log.LstdFlags | log.Lshortfile,
		}
	default:
		panic("unhandled default case")
	}

}

// NewLogger 获取日志记录器
func (loggerConfig *Config) NewLogger() (*log.Logger, *os.File, error) {
	var writers []io.Writer
	var file *os.File
	if loggerConfig.EnableStdOutput {
		writers = append(writers, os.Stdout)
	}
	if loggerConfig.EnableFileOutput {
		filePath := path.Join(loggerConfig.FileDir, loggerConfig.FileName)
		err := createDirIfNotExist(filePath)
		if err != nil {
			return nil, nil, err
		}
		file, err = os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to open log file: %w", err)
		}
		writers = append(writers, file)
	}
	if len(writers) < 1 {
		return nil, nil, fmt.Errorf("no logger configured for \"%s\"", loggerConfig.Prefix)
	}
	multiWriter := io.MultiWriter(writers...)
	return log.New(multiWriter, loggerConfig.Prefix, loggerConfig.LogFlags), file, nil
}
