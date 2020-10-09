package main

//在go语言中，new和make的区别？
// New
// 1.new的作用是初始化一个指向类型的指针(*T)
// 2.new函数是内建函数，函数定义：func new(Type) *Type
// 3.使用new函数来分配空间。传递给new 函数的是一个类型，不是一个值。返回值是 指向这个新分配的零值的指针。

//make 的作用是为 slice，map 或 chan 初始化并返回引用(T)。
//make函数是内建函数，函数定义：func make(Type, size IntegerType) Type
//第一个参数是一个类型，第二个参数是长度,返回值是一个类型
//make(T, args)函数的目的与new(T)不同。
//它仅仅用于创建 Slice, Map 和 Channel，并且返回类型是 T（不是*T）的一个初始化的（不是零值）的实例。

func main() {
	
}
