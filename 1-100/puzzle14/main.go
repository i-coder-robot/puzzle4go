package main

import "fmt"

//[intermediate] 下面的程序的运行结果是（）
//func main() {
//	if (true) {
//		defer fmt.Printf("1")
//	} else {
//		defer fmt.Printf("2")
//	}
//	fmt.Printf("3")
//}
//A. 321
//B. 32
//C. 31
//D. 13
//
//参考答案：C

//A 人家有一个if，怎么可能是3个
//defer是在函数执行后，再执行，所以先执行3，然后是true，那么就是1

func main() {
		if (true) {
			defer fmt.Printf("1")
		} else {
			defer fmt.Printf("2")
		}
		fmt.Printf("3")
}

