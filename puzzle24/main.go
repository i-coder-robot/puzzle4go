package main

import "fmt"

//[primary] 关于变量的自增和自减操作，下面语句正确的是（）
//A.
//i := 1
//i++
//B.
//
//i := 1
//j = i++
//C.
//
//i := 1
//++i
//D.
//
//i := 1
//i--
//参考答案：AD



func main() {

	//A
	//i := 1
	//i++
	//fmt.Println(i)

	//B j都没有初始化变量，什么鬼赋值给j
	//i := 1
	//j = i++
	//fmt.Println(j)

	//C这是什么鬼操作 ++报错
	//i := 1
	//++i
	//fmt.Println(i)

	//D
	i := 1
	i--
	fmt.Println(i)
}
