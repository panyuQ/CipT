package main

import (
	"runtime"
	"sync"
)

// Task 是一个任务结构体
type Task struct {
	ID     int
	Args   []interface{}
	Result interface{}
}

func NewTask(id int, args []interface{}) Task {
	return Task{ID: id, Args: args, Result: nil}
}

// WorkPool 是一个工作池结构体
type WorkPool struct {
	Results   chan Task
	function  func([]interface{}) interface{}
	tasks     chan Task
	workers   int
	waitGroup sync.WaitGroup
	stop      chan struct{}
	stopOnce  sync.Once
	Logger    *Logger
}

// NewWorkPool 创建一个新的工作池
func NewWorkPool(workers, taskBufferSize, resultBufferSize int, taskFunc func([]interface{}) interface{}, logger *Logger) *WorkPool {
	if workers < 1 {
		workers = runtime.NumCPU()
	}
	if taskBufferSize < 0 {
		taskBufferSize = 10
	}
	if resultBufferSize < 0 {
		resultBufferSize = 50
	}

	return &WorkPool{
		function: taskFunc,
		tasks:    make(chan Task, taskBufferSize),
		Results:  make(chan Task, resultBufferSize),
		workers:  workers,
		stop:     make(chan struct{}),
		Logger:   logger,
	}
}

// logInfo 输出普通日志
func (wp *WorkPool) logInfo(v ...interface{}) {
	if wp.Logger.Info != nil {
		wp.Logger.Info.Println(v...)
	}
}

// logError 输出错误日志
func (wp *WorkPool) logError(v ...interface{}) {
	if wp.Logger.Error != nil {
		wp.Logger.Error.Println(v...)
	}
}

// Start 启动工作池
func (wp *WorkPool) Start() {
	for i := 0; i < wp.workers; i++ {
		wp.waitGroup.Add(1)
		go wp.worker(i)
	}
}

// worker 是一个工作goroutine
func (wp *WorkPool) worker(id int) {
	defer wp.waitGroup.Done()
	for {
		select {
		case task, ok := <-wp.tasks:
			if !ok {
				return
			}
			wp.logInfo("Worker", id, "is processing task", task.ID)
			if wp.function != nil {
				task.Result = wp.function(task.Args)
			}
			wp.Results <- task
		case <-wp.stop:
			wp.logError("Work pool is stopped, cannot process task")
			return
		}
	}
}

// AddTask 添加任务到工作池
func (wp *WorkPool) AddTask(task Task) {
	select {
	case wp.tasks <- task:
		wp.logInfo("Task", task.ID, "added to the queue")
	case <-wp.stop:
		wp.logError("Work pool is stopped, cannot add new tasks")
	}
}

func (wp *WorkPool) Stop() {
	wp.stopOnce.Do(func() {
		close(wp.stop) // 广播停止信号
	})
}

func (wp *WorkPool) Wait() {
	wp.waitGroup.Wait() // 等待所有 worker 完成任务
	close(wp.Results)   // 等待所有结果写入完成后关闭
	close(wp.tasks)     // 确保所有 goroutine 退出后关闭任务通道
}
