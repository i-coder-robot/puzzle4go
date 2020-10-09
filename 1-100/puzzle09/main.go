package main
import "C"

//[intermediate] 对于函数定义：
//
//func add(args ...int) int {
//	sum := 0
//	for _, arg := range args {
//		sum += arg
//	}
//	return sum
//}
//下面对add函数调用正确的是（）
//A. add(1, 2)
//B. add(1, 3, 7)
//C. add([]int{1, 2})
//D. add([]int{1, 3, 7}...)
//参考答案：ABD

func main() {
	//add(1, 2)
	//add(1, 3, 7)
	add([]int{1, 2})
	//add([]int{1, 3, 7}...)
}

//可以看到C是错误的提示，人家要的是int,你给一个int切片，那不就完犊子了嘛
//D 这个调用时...三个小点点，把int切片解构成每一个int，所以是正确的

func add(args ...int) int {
	sum := 0
	for _, arg := range args {
		sum += arg
	}
	return sum
}
