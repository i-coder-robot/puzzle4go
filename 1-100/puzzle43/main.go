package main

//[intermediate] 关于cap函数的适用类型，下面说法正确的是（）
//A. array
//B. slice
//C. map
//D. channel
//参考答案：ABD


// The cap built-in function returns the capacity of v, according to its type:
//	Array: the number of elements in v (same as len(v)).
//	Pointer to array: the number of elements in *v (same as len(v)).
//	Slice: the maximum length the slice can reach when resliced;
//	if v is nil, cap(v) is zero.
//	Channel: the channel buffer capacity, in units of elements;
//	if v is nil, cap(v) is zero.

func main() {

}
