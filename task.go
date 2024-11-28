package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var (
	maxWorkers              = runtime.NumCPU()
	defaultTaskBufferSize   = 10
	defaultResultBufferSize = 50
)

// Task 是一个任务结构体
type Task struct {
	ID     int64
	Args   []interface{}
	Result interface{}
}

func NewTask(args []interface{}) Task {
	return Task{ID: time.Now().Unix(), Args: args, Result: nil}
}

// WorkPool 是一个工作池结构体
type WorkPool struct {
	Results chan Task

	function  func([]interface{}) interface{}
	tasks     chan Task
	workers   int
	waitGroup sync.WaitGroup
	stop      chan struct{}
	stopOnce  sync.Once
}

// NewWorkPool 创建一个新的工作池 \\\\
// asd
func NewWorkPool(workers int, taskBufferSize int, resultBufferSize int, taskFunc func([]interface{}) interface{}) *WorkPool {
	if workers < 1 {
		workers = maxWorkers
	}
	if taskBufferSize < 0 {
		taskBufferSize = defaultTaskBufferSize
	}
	if resultBufferSize < 0 {
		resultBufferSize = defaultResultBufferSize
	}
	return &WorkPool{
		function: taskFunc,
		tasks:    make(chan Task, taskBufferSize),
		Results:  make(chan Task, resultBufferSize),
		workers:  workers,
		stop:     make(chan struct{}),
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
			fmt.Printf("Worker %d is processing task %d\n", id, task.ID)
			if wp.function != nil {
				task.Result = wp.function(task.Args)
			}
			wp.Results <- task
		case <-wp.stop:
			return
		}
	}
}

// AddTask 添加任务到工作池
func (wp *WorkPool) AddTask(task Task) {
	select {
	case wp.tasks <- task:
		fmt.Printf("Task %d added to the queue\n", task.ID)
	case <-wp.stop:
		fmt.Println("Work pool is stopped, cannot add new tasks")
	}
}

func (wp *WorkPool) Wait() {
	wp.waitGroup.Wait()
}

// Stop 停止工作池
func (wp *WorkPool) Stop() {
	go func() {
		wp.stopOnce.Do(func() {
			close(wp.stop)
			close(wp.tasks)
		})
	}()
}

// 示例任务处理函数
func exampleTaskHandler(args []interface{}) interface{} {
	// 这里可以自定义任务处理逻辑
	sum := 0
	for _, arg := range args {
		if num, ok := arg.(int); ok {
			sum += num
		}
	}
	return sum
}

func task() {
	// 获取系统CPU核心数
	workers := runtime.NumCPU()
	bufferSize := 100 // 设置任务队列的缓冲区大小

	fmt.Printf("Starting with %d workers and buffer size %d\n", workers, bufferSize)

	// 创建工作池并设置任务处理函数
	wp := NewWorkPool(workers, bufferSize, bufferSize, exampleTaskHandler)

	wp.Start()

	// 模拟添加任务
	for i := 1; i <= 100; i++ {
		task := NewTask([]interface{}{i, i + 1, i + 2})
		wp.AddTask(task)
	}

	// 等待所有任务完成
	wp.Stop()

	wp.Wait()

	close(wp.Results)
	for result := range wp.Results {
		fmt.Printf("Task %d is done(%d)\n", result.ID, result.Result)
	}
	fmt.Println("All tasks completed")
}
