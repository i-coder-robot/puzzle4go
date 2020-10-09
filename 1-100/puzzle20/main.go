package main

import "fmt"

//[intermediate] 下面赋值正确的是（）
//A. var x = nil
//B. var x interface{} = nil
//C. var x string = nil
//D. var x error = nil
//参考答案：BD



func main() {
	//var x = nil
	//var x interface{} = nil
	//var x string = nil
	var x error = nil

	//A 不能用nil给一个没有类型的变量
	//C string不能nil赋值
	fmt.Println(x)
}
