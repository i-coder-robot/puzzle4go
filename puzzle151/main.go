package main

//Go语言局部变量分配在栈还是堆？
//Go语言编译器会自动决定把一个变量放在栈还是放在堆，
//编译器会做逃逸分析，当发现变量的作用域没有跑出函数范围，
//就可以在栈上，反之则必须分配在堆。


//参考go FAQ里面的一段话：
//How do I know whether a variable is allocated on the heap or the stack?
//
//From a correctness standpoint, you don't need to know.
//Each variable in Go exists as long as there are references to it.
//The storage location chosen by the implementation is irrelevant to the semantics of the language.
//
//The storage location does have an effect on writing efficient programs.
//When possible, the Go compilers will allocate variables that are local to a function in that function's stack frame.
//However, if the compiler cannot prove that the variable is not referenced after the function returns,
//then the compiler must allocate the variable on the garbage-collected heap to avoid dangling pointer errors.
//Also, if a local variable is very large, it might make more sense to store it on the heap rather than the stack.
//
//In the current compilers, if a variable has its address taken,
//that variable is a candidate for allocation on the heap.
//However, a basic escape analysis recognizes some cases
//when such variables will not live past the return from the function and can reside on the stack.



func main() {
	
}
