package logger

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
)

var GLogger, _ = NewLogger(nil, nil, nil)

type Logger struct {
	Info        *log.Logger
	Warn        *log.Logger
	Error       *log.Logger
	infoFile    *os.File
	warnFile    *os.File
	errorFile   *os.File
	infoConfig  *Config
	warnConfig  *Config
	errorConfig *Config
}

// NewLogger 根据配置创建日志记录器
func NewLogger(infoLoggerConfig *Config, warnLoggerConfig *Config, errorLoggerConfig *Config) (*Logger, error) {
	if infoLoggerConfig == nil {
		infoLoggerConfig = NewDefaultConfig(LogoLevelInfo)
	}
	if warnLoggerConfig == nil {
		warnLoggerConfig = NewDefaultConfig(LogoLevelWarn)
	}
	if errorLoggerConfig == nil {
		errorLoggerConfig = NewDefaultConfig(LogoLevelError)
	}
	infoLogger, infoFile, err := infoLoggerConfig.NewLogger()
	if err != nil {
		return nil, err
	}
	warnLogger, warnFile, err := warnLoggerConfig.NewLogger()
	if err != nil {
		return nil, err
	}
	errorLogger, errorFile, err := errorLoggerConfig.NewLogger()
	if err != nil {
		return nil, err
	}
	return &Logger{
		Info:        infoLogger,
		Warn:        warnLogger,
		Error:       errorLogger,
		infoFile:    infoFile,
		warnFile:    warnFile,
		errorFile:   errorFile,
		infoConfig:  infoLoggerConfig,
		warnConfig:  warnLoggerConfig,
		errorConfig: errorLoggerConfig,
	}, nil
}

func (logger *Logger) Close() error {
	if logger.infoConfig.EnableFileOutput {
		err := logger.infoFile.Close()
		if err != nil {
			return fmt.Errorf("close info log file fail: %w", err)
		}
	}
	if logger.warnConfig.EnableFileOutput {
		err := logger.warnFile.Close()
		if err != nil {
			return fmt.Errorf("close warn log file fail: %w", err)
		}
	}
	if logger.errorConfig.EnableFileOutput {
		err := logger.errorFile.Close()
		if err != nil {
			return fmt.Errorf("close error log file fail: %w", err)
		}
	}
	return nil
}

func (logger *Logger) ClearEmptyLogFile() error {
	if logger.infoConfig.EnableFileOutput {
		err := clearEmptyFile(path.Join(logger.infoConfig.FileDir, logger.infoConfig.FileName))
		if err != nil {
			return fmt.Errorf("clear info log file fail: %w", err)
		}
	}
	if logger.warnConfig.EnableFileOutput {
		err := clearEmptyFile(path.Join(logger.warnConfig.FileDir, logger.warnConfig.FileName))
		if err != nil {
			return fmt.Errorf("clear warn log file fail: %w", err)
		}
	}
	if logger.errorConfig.EnableFileOutput {
		err := clearEmptyFile(path.Join(logger.errorConfig.FileDir, logger.errorConfig.FileName))
		if err != nil {
			return fmt.Errorf("clear error log file fail: %w", err)
		}
	}
	return nil
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

// fileEmpty 检查指定文件是否为空
func fileEmpty(filePath string) (bool, error) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return false, err
	}
	return fileInfo.Size() == 0, nil
}

func clearEmptyFile(filePath string) error {
	isEmpty, err := fileEmpty(filePath)
	if err != nil {
		return err
	}
	if isEmpty {
		// 文件为空，可以执行清除操作，例如删除文件
		if err := os.Remove(filePath); err != nil {
			return fmt.Errorf("failed to remove empty log file: %w", err)
		}
	}
	return nil
}
