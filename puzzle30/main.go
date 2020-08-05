package main

import "fmt"

//[intermediate] 关于接口，下面说法正确的是（）
//A. 只要两个接口拥有相同的方法列表（次序不同不要紧），那么它们就是等价的，可以相互赋值
//B. 如果接口A的方法列表是接口B的方法列表的子集，那么接口B可以赋值给接口A
//C. 接口查询是否成功，要在运行期才能够确定
//D. 接口赋值是否可行，要在运行期才能够确定
//参考答案：ACD

//A儿子  B爹   父给子转换

type Father interface {
	DoWork()
}

type Son interface {
	Father
	Happy()
}

type Human struct {
	Name string
}

func (h Human) DoWork()  {
	fmt.Println(h.Name+" 工作...")
}

type Person struct {
	Name string
}

func (p Person) DoWork()  {
	fmt.Println(p.Name+"是富二代，无需工作。")
}

func (p Person) Happy()  {
	fmt.Println(p.Name+"在玩")
}

func main() {
	var father Father
	var son Son
	father =Human{Name: "老王"}


	son = father

	fmt.Println(son)
}
