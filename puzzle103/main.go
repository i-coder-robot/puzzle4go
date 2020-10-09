package main

//在go语言中，Printf()、Sprintf()、Fprintf()函数的区别用法是什么？
//
//都是把格式好的字符串输出，只是输出的目标不一样：
//
//Printf()，   是把格式字符串输出到标准输出（一般是屏幕，可以重定向）。
//Printf()    是和标准输出文件(stdout)关联的,Fprintf 则没有这个限制.
//Sprintf()， 是把格式字符串输出到指定字符串中，所以参数比printf多一个char*。那就是目标字符串地址。
//Fprintf()， 是把格式字符串输出到指定文件设备中，所以参数比printf多一个文件指针FILE*。
//			  主要用于文件操作。Fprintf()是格式化输出到一个stream，通常是到文件。


func main() {
	
}
