package main

import "fmt"

//下面代码输出什么？
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
func main() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}

//考点：defer执行顺序
//
//解答：这道题类似第1题需要注意到defer执行顺序和值传递 index:1肯定是最后执行的，
//但是index:1的第三个参数是一个函数，
//所以最先被调用calc("10",1,2)==>10,1,2,3 执行index:
//2时,与之前一样，需要先调用calc("20",0,2)==>20,0,2,2
//执行到b=1时候开始调用，index:2==>calc("2",0,2)==>2,0,2,2
//最后执行index:1==>calc("1",1,3)==>1,1,3,4

//10 1 2 3
//20 0 2 2
//2 0 2 2
//1 1 3 4
