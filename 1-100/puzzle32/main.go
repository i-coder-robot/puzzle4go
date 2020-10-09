package main

//[primary] 关于同步锁，下面说法正确的是（）
//A. 当一个goroutine获得了Mutex后，其他goroutine就只能乖乖的等待，除非该goroutine释放这个Mutex
//B. RWMutex在读锁占用的情况下，会阻止写，但不阻止读
//C. RWMutex在写锁占用情况下，会阻止任何其他goroutine（无论读和写）进来，整个锁相当于由该goroutine独占
//D. Lock()操作需要保证有Unlock()或RUnlock()调用与之对应
//参考答案：ABC

//D  RUnlock方法应该和读写锁对应

func main() {
	
}
