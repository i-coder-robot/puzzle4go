package main

//[intermediate] 关于函数返回值的错误设计，下面说法正确的是（）
//A. 如果失败原因只有一个，则返回bool
//B. 如果失败原因超过一个，则返回error
//C. 如果没有失败原因，则不返回bool或error
//D. 如果重试几次可以避免失败，则不要立即返回bool或error
//参考答案：ABCD

//函数返回值，尽量详细告诉调用者，函数执行后的结果，应该有错误返回值，例如channel中

func main() {

}
