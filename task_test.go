package main

import (
	"log"
	"runtime"
	"testing"
)

// 示例任务处理函数
func exampleTaskHandler(args []interface{}) interface{} {
	sum := 0
	for _, arg := range args {
		if num, ok := arg.(int); ok {
			sum += num
		}
	}
	return sum
}

func TestTasks(t *testing.T) {
	logger, err := NewLogger(nil, nil)
	if err != nil {
		log.Fatalf("Failed to create logger: %v\n", err)
		return
	}

	// 获取系统CPU核心数
	workers := runtime.NumCPU()
	bufferSize := 100 // 设置任务队列的缓冲区大小

	logger.Info.Printf("Starting with %d workers and buffer size %d\n", workers, bufferSize)

	// 创建工作池并设置任务处理函数
	wp := NewWorkPool(workers, bufferSize, bufferSize, exampleTaskHandler, logger)

	wp.Start()

	// 模拟添加任务
	for i := 1; i <= 100; i++ {
		task := NewTask(i, []interface{}{i, i + 1, i + 2})
		wp.AddTask(task)
	}

	// 停止工作池并等待所有任务完成
	wp.Stop()
	wp.Wait()

	// 输出任务结果
	for result := range wp.Results {
		wp.logInfo("Task", result.ID, "is done(", result.Result, ")")
	}
	wp.logInfo("All tasks completed")
}
