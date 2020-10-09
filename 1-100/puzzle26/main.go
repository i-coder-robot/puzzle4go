package main

import "fmt"

//[intermediate] 如果Add函数的调用代码为：
//func main() {
//	var a Integer = 1
//	var b Integer = 2
//	var i interface{} = &a
//	sum := i.(*Integer).Add(b)
//	fmt.Println(sum)
//}
//则Add函数定义正确的是（）
//A.
//
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
//
//type Integer int
//func (a *Integer) Add(b Integer) Integer {
//	return *a + b
//}

//D.
//type Integer int
//func (a *Integer) Add(b *Integer) Integer {
//	return *a + *b
//}
//参考答案：AC

//B和D，都是要接受Integer指针类型的参数，所以不对

//如果将一个接口类型变量断言成一个指针类型的变量，在断言成功的前提下，两个变量将共享内存空间,
//而且add方法，可以是struct和它的指针类型调用，所以AC是可以的

func main() {
		var a Integer = 1
		var b Integer = 2
		var i interface{} = &a
		sum := i.(*Integer).Add(b)
		fmt.Println(sum)
}
