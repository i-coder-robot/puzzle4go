package main
import "fmt"

//[primary] 关于字符串连接，下面语法正确的是（）
//A. str := ‘abc’ + ‘123’
//B. str := "abc" + "123"
//C. str ：= '123' + "abc"
//D. fmt.Sprintf("abc%d", 123)
//
//参考答案：BD

func main() {
	 //str := ‘abc’ + ‘123’
	str := "abc" + "123"
	//str := '123' + "abc"
	 s:=fmt.Sprintf("abc%d", 123)

	fmt.Println(str)
	fmt.Println(s)
}
