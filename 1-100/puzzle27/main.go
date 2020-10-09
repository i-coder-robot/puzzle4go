package main

//[intermediate] 如果Add函数的调用代码为：
//func main() {
//	var a Integer = 1
//	var b Integer = 2
//	var i interface{} = a
//	sum := i.(Integer).Add(b)
//	fmt.Println(sum)
//}
//则Add函数定义正确的是（）
//A.
//type Integer int
//func (a Integer) Add(b Integer) Integer {
//	return a + b
//}
//B.
//type Integer int
//func (a Integer) Add(b *Integer) Integer {
//	return a + *b
//}
//C.
//type Integer int
//func (a *Integer) Add(b Integer) Integer {
//	return *a + b
//}
//D.
//
//type Integer int
//func (a *Integer) Add(b *Integer) Integer {
//	return *a + *b
//}
//参考答案：A

//A 看定义是可以的，很顺眼，
//C,定义是指针类型，不能调用
//BD参数都是指针，直接完犊子了

func main() {
}
