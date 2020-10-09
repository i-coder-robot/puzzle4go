package main
//简述一下golang的协程调度原理?
//
//M(machine): 代表着真正的执行计算资源，可以认为它就是os thread（系统线程）。
//P(processor): 表示逻辑processor，是线程M的执行的上下文。
//G(goroutine): 调度系统的最基本单位goroutine，存储了goroutine的执行stack信息、
//				goroutine状态以及goroutine的任务函数等。


//G-M-P三者的关系与特点：
//P的个数取决于设置的GOMAXPROCS，go新版本默认使用最大内核数，比如你有8核处理器，那么P的数量就是8
//M的数量和P不一定匹配，可以设置很多M，M和P绑定后才可运行，多余的M处于休眠状态。
//P包含一个LRQ（Local Run Queue）本地运行队列，这里面保存着P需要执行的协程G的队列
//除了每个P自身保存的G的队列外，调度器还拥有一个全局的G队列GRQ（Global Run Queue），
//这个队列存储的是所有未分配的协程G。

func main() {
	
}
