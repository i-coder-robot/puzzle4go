package main

//69 [primary] 数组是一个值类型（）
//参考答案：对的
//
//70 [primary] 使用map不需要引入任何库（）
//参考答案：对的
//
//71 [intermediate] 内置函数delete可以删除数组切片内的元素（）
//参考答案：错的
//
//72 [primary] 指针是基础类型（）
//参考答案：错的
//
//73 [primary] interface{}是可以指向任意对象的Any类型（）
//参考答案：对的
//
//74 [intermediate] 下面关于文件操作的代码可能触发异常（）
//file, err := os.Open("test.go")
//defer file.Close()
//if err != nil {
//fmt.Println("open file failed:", err)
//return
//}
//...
//参考答案：对的

// Open opens the named file for reading. If successful, methods on
// the returned file can be used for reading; the associated file
// descriptor has mode O_RDONLY.
// If there is an error, it will be of type *PathError.



//
//75 [primary] Golang不支持自动垃圾回收（）
//参考答案：错的，支持的
//
//76 [primary] Golang支持反射，反射最常见的使用场景是做对象的序列化（）
//参考答案：对的
//
//77 [primary] Golang可以复用C/C++的模块，这个功能叫Cgo（）
//参考答案：工作中使用go ,没有服用过c/c++模块，这个不是很确定，有知道的小伙伴可以打在弹幕上

//
//78 [primary] 下面代码中两个斜点之间的代码，比如json:"x"，作用是X字段在从结构体实例编码到JSON数据格式的时候，
//使用x作为名字，这可以看作是一种重命名的方式（）
//type Position struct {
//	X int `json:"x"`
//	Y int `json:"y"`
//	Z int `json:"z"`
//}
//参考答案：对的， 输出的json{"x":1,"y":2,"z":3}
//
//79 [primary] 通过成员变量或函数首字母的大小写来决定其作用域（）
//参考答案：对的
//
//80 [primary] 对于常量定义zero(const zero = 0.0)，zero是浮点型常量（）
//参考答案：错的， 是float64
//
//81 [primary] 对变量x的取反操作是~x（）
//参考答案：错的 ， !
//
//82 [primary] 下面的程序的运行结果是xello（）
//func main() {
//	str := "hello"
//	str[0] = 'x'
//	fmt.Println(str)
//}
//参考答案：错的， str[0] = 'x' 这句话编译都不能通过，Go字符串是不可变的
//
//83 [primary] golang支持goto语句（）
//参考答案：对的
//
//84 [primary] 下面代码中的指针p为野指针，因为返回的栈内存在函数结束时会被释放（）
//type TimesMatcher struct {
//	base int
//}
//func NewTimesMatcher(base int) *TimesMatcher{
//	return &TimesMatcher{base:base}
//}
//func main() {
//	p := NewTimesMatcher(3)
//	...
//}
//参考答案：错的，NewTimesMatcher函数后不会被释放
//
//85 [primary] 匿名函数可以直接赋值给一个变量或者直接执行（）
//参考答案：对的
//
//86 [primary] 如果调用方调用了一个具有多返回值的方法，但是却不想关心其中的某个返回值，
//可以简单地用一个下划线“_”来跳过这个返回值，该下划线对应的变量叫匿名变量（）
//参考答案：对的
//
//87 [primary] 在函数的多返回值中，如果有error或bool类型，则一般放在最后一个（）
//参考答案：对的
//
//88 [primary] 错误是业务过程的一部分，而异常不是（）
//参考答案：对的
//
//89 [primary] 函数执行时，如果由于panic导致了异常，则延迟函数不会执行（）
//参考答案：错的，仍然会执行
//
//90 [intermediate] 当程序运行时，如果遇到引用空指针、下标越界或显式调用panic函数等情况，
//则先触发panic函数的执行，然后调用延迟函数。
//调用者继续传递panic，因此该过程一直在调用栈中重复发生：
//函数停止执行，调用延迟执行函数。如果一路在延迟函数中没有recover函数的调用，则会到达该携程的起点，
//该携程结束，然后终止其他所有携程，其他携程的终止过程也是重复发生：函数停止执行，调用延迟执行函数（）
//参考答案：错的，延迟函数会先压栈，然后出栈执行。
//
//91 [primary] 同级文件的包名不允许有多个（）
//参考答案：对的
//
//92 [intermediate] 可以给任意类型添加相应的方法（）
//参考答案：错的
//
//93 [primary] golang虽然没有显式的提供继承语法，但是通过匿名组合实现了继承（）
//参考答案：对的
//
//94 [primary] 使用for range迭代map时每次迭代的顺序可能不一样，因为map的迭代是随机的（）
//参考答案：对的
//
//95 [primary] switch后面可以不跟表达式（）
//参考答案：对的
//
//96 [intermediate] 结构体在序列化时非导出变量（以小写字母开头的变量名）不会被encode，
//因此在decode时这些非导出变量的值为其类型的零值（）
//参考答案：对的
//
//97 [primary] golang中没有构造函数的概念，对象的创建通常交由一个全局的创建函数来完成，以NewXXX来命名（）
//参考答案：对的
//
//98 [intermediate] 当函数deferDemo返回失败时，并不能destroy已create成功的资源（）
//func deferDemo() error {
//	err := createResource1()
//	if err != nil {
//		return ERR_CREATE_RESOURCE1_FAILED
//	}
//	defer func() {
//		if err != nil {
//			destroyResource1()
//		}
//	}()
//
//	err = createResource2()
//	if err != nil {
//		return ERR_CREATE_RESOURCE2_FAILED
//	}
//	defer func() {
//		if err != nil {
//			destroyResource2()
//		}
//	}()
//
//	err = createResource3()
//	if err != nil {
//		return ERR_CREATE_RESOURCE3_FAILED
//	}
//	return nil
//}
//参考答案：错的，有defer啊
//
//99 [intermediate] channel本身必然是同时支持读写的，所以不存在单向channel（）
//参考答案：错的
//
//100 [primary] import后面的最后一个元素是包名（）
//参考答案：错的  import后面跟的是包的路径，而不是包名。

func main() {

}
