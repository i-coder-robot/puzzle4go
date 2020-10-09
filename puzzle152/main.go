package main

//如何获取 go 程序运行时的协程数量, gc 时间, 对象数, 堆栈信息?

//调用接口 runtime.ReadMemStats 可以获取以上所有信息,
//注意: 调用此接口会触发 STW(Stop The World)
//https://golang.org/pkg/runtime/#ReadMemStats

//如果需要打入到日志系统, 可以使用 go 封装好的包, 输出 json 格式.
//https://golang.org/pkg/expvar/

func main() {
	
}
