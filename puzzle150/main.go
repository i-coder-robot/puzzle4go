package main

//介绍下 golang 的 runtime 机制?
//Runtime 负责管理任务调度，垃圾收集及运行环境。同时，Go提供了一些高级的功能，
//如goroutine, channel, 以及Garbage collection。
//这些高级功能需要一个runtime的支持. runtime和用户编译后的代码被linker静态链接起来，
//形成一个可执行文件。这个文件从操作系统角度来说是一个用户态的独立的可执行文件。
//从运行的角度来说，这个文件由2部分组成，一部分是用户的代码，另一部分就是runtime。
//runtime通过接口函数调用来管理goroutine, channel及其他一些高级的功能。
//从用户代码发起的调用操作系统API的调用都会被runtime拦截并处理。
//
//Go runtime的一个重要的组成部分是goroutine scheduler。
//他负责追踪，调度每个goroutine运行，
//实际上是从应用程序的process所属的thread pool中分配一个thread来执行这个goroutine。
//因此，和java虚拟机中的Java thread和OS thread映射概念类似，
//每个goroutine只有分配到一个OS thread才能运行。
func main() {
	
}
