package main

import "fmt"

//[primary] 关于局部变量的初始化，下面正确的使用方式是（）
//A. var i int = 10
//B. var i = 10
//C. i := 10
//D. i = 10
//参考答案：ABC

func main() {
	var i int = 10
	//var i = 10
	//i := 10
	//i = 10

	fmt.Println(i)
}
