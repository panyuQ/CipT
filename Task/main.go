package Task

import (
	"CipT/Logger"
	"runtime"
	"sync"
)

var (
	// 默认缓冲大小
	maxWorkers           = runtime.NumCPU()
	minWorkers           = 1
	maxTasksBufferSize   = maxWorkers * 100
	minTasksBufferSize   = minWorkers * 10
	maxResultsBufferSize = maxTasksBufferSize * 100
	minResultsBufferSize = minTasksBufferSize * 10
)

// Task 是一个任务结构体
type Task struct {
	ID     int           // 这个任务的ID
	Args   []interface{} // 这个任务的参数
	Result []interface{} // 这个任务的结果
}

func NewTask(id int, args []interface{}) Task {
	return Task{ID: id, Args: args, Result: nil}
}

// WorkPool 是一个工作池结构体
type WorkPool struct {
	function  func([]interface{}) []interface{} // 任务执行函数
	tasks     chan Task                         // 任务缓冲区
	workers   int                               // 工作协程数量
	waitGroup sync.WaitGroup                    // 用于等待任务完成
	stop      chan struct{}                     // 用于广播停止信号
	stopOnce  sync.Once                         // 确保只执行一次停止操作

	Logger            *Logger.Logger // 日志
	Results           chan Task      // 任务结果缓冲区
	TasksBufferSize   int            // 任务缓冲区大小
	ResultsBufferSize int            // 结果缓冲区大小
}

// NewWorkPool 创建一个新的工作池
func NewWorkPool(workers, tasksBufferSize, resultsBufferSize int, taskFunc func([]interface{}) []interface{}, logger *Logger.Logger) *WorkPool {
	if workers < minWorkers {
		workers = maxWorkers
	}
	if tasksBufferSize < minTasksBufferSize {
		tasksBufferSize = maxTasksBufferSize
	}
	if resultsBufferSize < minResultsBufferSize {
		resultsBufferSize = maxResultsBufferSize
	}

	// 初始化工作池
	wp := &WorkPool{
		function:          taskFunc,
		workers:           workers,
		Logger:            logger,
		TasksBufferSize:   tasksBufferSize,
		ResultsBufferSize: resultsBufferSize,
	}

	return wp
}

// AddTask 添加任务到工作池
func (wp *WorkPool) AddTask(task Task) {
	select {
	case wp.tasks <- task:
		//wp.logInfo("Task", task.ID, "added to the queue")
	case <-wp.stop:
		//wp.logError("Work pool is stopped, cannot add new tasks")
	}
}

// Start 启动工作池，初始化并启动 worker
func (wp *WorkPool) Start() {
	// 初始化通道
	wp.tasks = make(chan Task, wp.TasksBufferSize)
	wp.Results = make(chan Task, wp.ResultsBufferSize)
	wp.stop = make(chan struct{})

	wp.logInfo("Starting with", wp.workers, "worker(s) ( TasksBufferSize", wp.TasksBufferSize, "ResultsBufferSize", wp.ResultsBufferSize, ")")
	// 启动工作goroutines
	for i := 0; i < wp.workers; i++ {
		wp.waitGroup.Add(1)
		go wp.worker()
	}
}

// Stop 停止工作池，等待所有任务完成
func (wp *WorkPool) Stop(mandatory bool) {
	wp.stopOnce.Do(func() {
		if mandatory { // 强制结束
			close(wp.stop)      // 广播停止信号
			wp.waitGroup.Wait() // 等待所有 worker 完成任务
			close(wp.Results)   // 关闭结果通道
			close(wp.tasks)     // 关闭任务通道
		} else { // 非强制结束
			close(wp.tasks)     // 只禁止接收新的任务，等待所有任务完成
			wp.waitGroup.Wait() // 等待所有 worker 完成当前任务
			close(wp.stop)      // 广播停止信号
			close(wp.Results)   // 关闭结果通道
		}
		wp.logInfo("All tasks completed. ( mandatory:", mandatory, ")")
	})
}

// worker 是一个工作goroutine，负责处理任务
func (wp *WorkPool) worker() {
	defer wp.waitGroup.Done()
	for {
		select {
		case task, ok := <-wp.tasks:
			if !ok { // 通道关闭
				//wp.logInfo("Worker", id, "exiting due to channel close")
				return
			}
			//wp.logInfo("Worker", id, "is processing task", task.ID)
			if wp.function != nil {
				task.Result = wp.function(task.Args)
			}
			wp.Results <- task // 将任务结果发送到结果通道
		case <-wp.stop: // 收到停止信号
			//wp.logInfo("Worker", id, "received stop signal, exiting")
			return
		}
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
