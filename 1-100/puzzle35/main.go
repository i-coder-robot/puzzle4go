package main

//[primary] flag是bool型变量，下面if表达式符合编码规范的是（）
//A. if flag == 1
//B. if flag
//C. if flag == false
//D. if !flag
//参考答案：BCD

var flag bool
func main() {
	//if flag == 1{}
	if flag{}
	if flag == false{}
	if !flag{}
}
