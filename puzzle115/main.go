package main

//说说go语言的channel特性？
//
//A. 给一个 nil channel 发送数据，造成永远阻塞
//
//B. 从一个 nil channel 接收数据，造成永远阻塞
//
//C. 给一个已经关闭的 channel 发送数据，引起 panic
//
//D. 从一个已经关闭的 channel 接收数据，如果缓冲区中为空，则返回一个零值
//
//E. 无缓冲的channel是同步的，而有缓冲的channel是非同步的


func main() {
	
}
