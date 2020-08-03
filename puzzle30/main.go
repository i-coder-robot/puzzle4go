package main

import "fmt"

//[intermediate] 关于接口，下面说法正确的是（）
//A. 只要两个接口拥有相同的方法列表（次序不同不要紧），那么它们就是等价的，可以相互赋值
//B. 如果接口A的方法列表是接口B的方法列表的子集，那么接口B可以赋值给接口A
//C. 接口查询是否成功，要在运行期才能够确定
//D. 接口赋值是否可行，要在运行期才能够确定
//参考答案：ACD

//CD看着很顺眼

//A 没问题，如果你有问题，可以一下，下面的例子
//B这不用说，肯定是不行的啊 可以看下面的例子

type Work interface {
	DoWork()
	Eat()
	Sleep()
}

type Life interface {
	Sleep()
	DoWork()
	Eat()

}

type Son interface {
	Sleep()
	Eat()
}

func main() {
	var w Work
	var l Life
	w=l
	fmt.Println(w)

	var s Son
	l=s
}
