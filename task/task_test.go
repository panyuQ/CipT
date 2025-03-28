package task

import (
	"CipT/logger"
	"log"
	"runtime"
	"strconv"
	"testing"
)

// 示例任务处理函数
func exampleTaskHandler(args []string) ([]string, error) {
	sum := 0
	for _, arg := range args {
		x, e := strconv.Atoi(arg)
		if e == nil {
			sum += x
		} else {
			return nil, e
		}
	}
	return []string{strconv.Itoa(sum)}, nil
}

func TestTasks(t *testing.T) {
	newLogger, err := logger.NewLogger(nil, nil, nil)
	if err != nil {
		log.Fatalf("Failed to create logger: %v\n", err)
		return
	}

	// 获取系统CPU核心数
	workers := runtime.NumCPU()
	bufferSize := workers * 100 // 设置任务队列的缓冲区大小

	// 创建工作池并设置任务处理函数
	wp := NewWorkPool(workers, bufferSize, bufferSize*100, exampleTaskHandler, newLogger)

	wp.Start()

	// 模拟添加任务
	for i := 1; i <= 100; i++ {
		task := NewTask(i, []string{strconv.Itoa(i), strconv.Itoa(i + 1), strconv.Itoa(i + 2)})
		wp.AddTask(task)
	}

	// 停止工作池并等待所有任务完成
	wp.Stop(false) // 非强制性关闭

	wp.logInfo("Output task results...")
	// 输出任务结果
	for result := range wp.Results {
		wp.logInfo("task", result.ID, "is done(", result.Result, ")")
	}
}
