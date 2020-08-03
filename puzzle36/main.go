package main
import "C"

//[primary] value是整型变量，下面if表达式符合编码规范的是（）
//A. if value == 0
//B. if value
//C. if value != 0
//D. if !value
//参考答案：AC

//BD都是布尔类型直接判断的

var value int
func main() {
	if value == 0{}
	//if value{}
	if value != 0{}
	//if !value{}
}
