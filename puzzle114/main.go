package main

//说说go语言的同步锁？
//(1) 当一个goroutine获得了Mutex后，其他goroutine就只能乖乖的等待，除非该goroutine释放这个Mutex
//(2) RWMutex在读锁占用的情况下，会阻止写，但不阻止读
//(3) RWMutex在写锁占用情况下，会阻止任何其他goroutine（无论读和写）进来，整个锁相当于由该goroutine独占


func main() {
	
}
