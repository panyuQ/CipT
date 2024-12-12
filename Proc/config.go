package Proc

type Config struct {
	PageSize          int // 页面大小
	Workers           int
	TasksBufferSize   int
	ResultsBufferSize int
}

func NewConfig() *Config {
	return &Config{
		PageSize:          1000,
		Workers:           10,
		TasksBufferSize:   1000,
		ResultsBufferSize: 100000,
	}
}
