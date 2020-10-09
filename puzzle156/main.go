package main

import "fmt"

//以下代码打印出来什么内容，说出为什么。
type People interface {
	Show()
}
type Student struct {}

func(stu * Student) Show() {}

func live() People {
	var stu *Student
	return stu
}

func main() {

	if live() == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}
}
//
//	BBBBBBB

//interface 在 golang 中是一个非常重要的特性。它相对于其它语言有很多优势：
//
//duck typing。大多数的静态语言需要显示的声明类型的继承关系。
//而 golang 通过 interface 实现了 duck typing， 使得我们无需显示的类型继承。
//不像其它实现了 duck typing 的动态语言那样，只能在运行时才能检查到类型的转换错误。
//而 golang 的 interface 特性可以让我们在编译时就能发现错误。

//考点：interface内部结构

//接口分为下面2种

var name interface{}

type Life interface {
	DoWork()
	Eat()
	Happy()
}


//他们的底层结构如下：

//type eface struct { empty
//	//空接口
//	_type *_type         //类型信息
//	data  unsafe.Pointer //指向数据的指针(go语言中特殊的指针类型unsafe.Pointer类似于c语言中的void*)
//}

//	带有方法的接口
//type iface struct {
//	tab  *itab           //存储type信息还有结构实现方法的集合
//	data unsafe.Pointer  //指向数据的指针(go语言中特殊的指针类型unsafe.Pointer类似于c语言中的void*)}

//_type表示实际类型信息.每个类型的_type信息由编译器在编译时生成。
//	type _type struct {
	//	size       uintptr  该类型所占用的字节数量
	//	ptrdata    uintptr
	//	hash       uint32   数据hash值
	//	tflag      tflag
	//	fieldalign uint8
	//	kind       uint8	表示类型的种类，如bool,int ,float,string ,struct,interface等。
	//	alg        *typeAlg
	//	gcdata    *byte
	//	str       nameOff	表示类型的名字信息，他是一个NameOFF（int32），通过这个nameOff，
	//	ptrToThis typeOff
//	}
//	type itab struct {
	//	inter  *interfacetype  //接口类型    inter 指向对应的interface的类型信息。
	//	_type  *_type          //结构类型	type和eface钟的一样，指向的是实际类型的描述信息。
	//	link   *itab
	//	bad    int32
	//	inhash int32
	//	fun    [1]uintptr      //方法集合,表示对于该特定的实际类型而言，interface中所有函数的地址。
//	}

//type interfacetype struct {
//	typ     _type
//	pkgpath name        接口包名
//	mhdr    []imethod   接口定义的函数列表
//}

