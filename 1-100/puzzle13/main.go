package main

import "fmt"

//[primary] 关于布尔变量b的赋值，下面错误的用法是（）
//A. b = true
//B. b = 1
//C. b = bool(1)
//D. b = (1 == 2)
//参考答案：BC

//B go语言人家有false,这个1是什么鬼
//C int 到 bool 不能这么转啊，老铁
var b bool
func main() {
	//b = true
	//b = 1
	//b = bool(1)
	//b = (1 == 2)
	fmt.Println(b)
}
