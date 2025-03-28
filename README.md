# CipT
> 该项目启发自 《随波逐流》、"厨师" <br>
> 其中 Core/NoKey/BaseFamily/codec修改自[junjun-cai](https://github.com/junjun-cai/codec)的[codec](https://github.com/junjun-cai/codec)<br>
> Go Version: 1.23.2

## 目前开发中...

### 暂且完成

1. `Logger`: 日志管理,分为信息日志和错误日志,且都可以指定多个输出流
2. `Task`: 任务管理,高并发执行任务(可指定任务函数)
3. `Core/NoKey/BaseFamily`: Base家族编码解码
4. `Proc/input.go`: 识别输入并提取文件内容, 且实现"分页查询"
5. 正则表达式的扩展: 
   1. \x, \y, \z指定种类数量的字符，\X, \Y, \Z除\x, \y, \z以外的所有字符（Python实现）

### 正在进行

* 正则表达式库的扩展的 Python实现 转为 Go实现 并与当前程序对接