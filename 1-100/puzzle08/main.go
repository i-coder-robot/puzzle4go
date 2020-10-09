package main

import "fmt"

//[primary] 关于循环语句，下面说法正确的有（）
//A. 循环语句既支持for关键字，也支持while和do-while
//B. 关键字for的基本使用方法与C/C++中没有任何差异
//C. for循环支持continue和break来控制循环，但是它提供了一个更高级的break，可以选择中断哪一个循环
//D. for循环不支持以逗号为间隔的多个赋值语句，必须使用平行赋值的方式来初始化多个变量
//
//参考答案：CD

//A while ,do-while 什么鬼
//B 知道C/C++可以把答案打到屏幕上

func main() {
	for i:=0;i<5;i++{
		if i==3{
			//滚出，循环结束，后面i=4...都不执行
			break
		}
	}

	for i:=0;i<5;i++{
		if i==2{
			//跳出，循环继续执行后面的i=3...继续执行
			continue
		}
	}
	x:=1
	for x<100{
		fmt.Println("来啊，快活啊~")
		x++
	}
}
