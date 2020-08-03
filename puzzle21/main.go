package main

import "fmt"

//[intermediate] 关于整型切片的初始化，下面正确的是（）
//A. s := make([]int)
//B. s := make([]int, 0)
//C. s := make([]int, 5, 10)
//D. s := []int{1, 2, 3, 4, 5}
//参考答案：BCD

//A 也没有长度

func main() {
	//s := make([]int)
	//s := make([]int, 0)
	//s := make([]int, 5, 10)
	s := []int{1, 2, 3, 4, 5}

	fmt.Println(s)
}
