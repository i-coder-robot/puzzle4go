package main

import "fmt"

//[primary] 定义一个包内全局字符串变量，下面语法正确的是 （）
//A. var str string
//B. str := ""
//C. str = ""
//D. var str = ""
//参考答案 AD

var str string
//var str = ""
//str := ""
//str = ""
func main() {
	fmt.Println(str)
}
