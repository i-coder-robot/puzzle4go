package main

import (
	"fmt"
	"runtime"
	"sync"
)

//下面的代码会输出什么，并说明原因
func main() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("A: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("B: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

//考点：go执行的随机性和闭包
//
//解答：谁也不知道执行后打印的顺序是什么样的，所以只能说是随机数字。
//但是A:均为输出10，B:从0~9输出(顺序不定)。
//第一个go func中i是外部for的一个变量，地址不变化。遍历完成后，最终i=10。故go func执行时，i的值始终是10。
//第二个go func中i是函数参数，与外部for中的i完全是两个变量。尾部(i)将发生值拷贝，go func内部指向值拷贝地址。