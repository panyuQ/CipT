package logger

import (
	"testing"
)

func TestGLogger(t *testing.T) {
	sss := "Hello, World!"
	GLogger.Info.Println(sss)
	GLogger.Warn.Println(sss)
	GLogger.Error.Println(sss)
}

func TestDefaultLogger_ClearEmptyLogFile(t *testing.T) {
	LLL, _ := NewLogger(nil, nil, nil)
	defer func(LLL *Logger) {
		err := LLL.Close()
		if err != nil {
			t.Error(err)
		}
		err = LLL.ClearEmptyLogFile()
		if err != nil {
			t.Error(err)
		}
	}(LLL)

}
