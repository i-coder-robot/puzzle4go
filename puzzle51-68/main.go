package main

//51 [primary] 声明一个整型变量i__________
//参考答案：var i int
//
//52 [primary] 声明一个含有10个元素的整型数组a__________
//参考答案：var a [10]int
//
//53 [primary] 声明一个整型数组切片s__________
//参考答案：var s []int
//
//54 [primary] 声明一个整型指针变量p__________
//参考答案：var p *int
//
//55 [primary] 声明一个key为字符串型value为整型的map变量m__________
//参考答案：var m map[string]int
//
//56 [primary] 声明一个入参和返回值均为整型的函数变量f__________
//参考答案：var f func(a int) int
//
//57 [primary] 声明一个只用于读取int数据的单向channel变量ch__________
//参考答案：var ch <-chan int
//
//58 [primary] 假设源文件的命名为slice.go，则测试文件的命名为__________
//参考答案：slice_test.go
//
//59 [primary] go test要求测试函数的前缀必须命名为__________
//参考答案：Test
//
//60 [intermediate] 下面的程序的运行结果是__________
//for i := 0; i < 5; i++ {
//defer fmt.Printf("%d ", i)
//}
//参考答案：4 3 2 1 0
//
//61 [intermediate] 下面的程序的运行结果是__________
//func main() {
//	x := 1
//	{
//		x := 2
//		fmt.Print(x)
//	}
//	fmt.Println(x)
//}
//参考答案：21   就近原则，谁离的近，就找谁，和找对象一样，你工作和哪个小姐姐，小哥哥经常接触，就最有机会
//
//62 [intermediate] 下面的程序的运行结果是__________
//func main() {
//	strs := []string{"one", "two", "three"}
//
//	for _, s := range strs {
//		go func() {
//			time.Sleep(1 * time.Second)
//			fmt.Printf("%s ", s)
//		}()
//	}
//	time.Sleep(3 * time.Second)
//}
//参考答案：three three three   闭包问题，当函数执行的时候，s的值都是three
//
//63 [intermediate] 下面的程序的运行结果是__________
//func main() {
//	x := []string{"a", "b", "c"}
//	for v := range x {
//		fmt.Print(v)
//	}
//}
//参考答案：012
//
//64 [intermediate] 下面的程序的运行结果是__________
//func main() {
//	x := []string{"a", "b", "c"}
//	for _, v := range x {
//		fmt.Print(v)
//	}
//}
//参考答案：abc
//
//65 [primary] 下面的程序的运行结果是__________
//func main() {
//	i := 1
//	j := 2
//	i, j = j, i
//	fmt.Printf("%d%d\n", i, j)
//}
//参考答案：21
//
//66 [primary] 下面的程序的运行结果是__________
//func incr(p *int) int {
//	*p++
//	return *p
//}
//func main() {
//	v := 1
//	incr(&v)
//	fmt.Println(v)
//}
//参考答案：2
//
//67 [primary] 启动一个goroutine的关键字是__________
//参考答案：go
//
//68 [intermediate] 下面的程序的运行结果是__________
//type Slice []int
//func NewSlice() Slice {
//	return make(Slice, 0)
//}
//func (s* Slice) Add(elem int) *Slice {
//	*s = append(*s, elem)
//	fmt.Print(elem)
//	return s
//}
//func main() {
//	s := NewSlice()
//	defer s.Add(1).Add(2)
//	s.Add(3)
//}
//参考答案：132

func main() {

}
