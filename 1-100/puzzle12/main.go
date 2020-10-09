package main

//[primary] 关于const常量定义，下面正确的使用方式是（）
//A.
//const Pi float64 = 3.14159265358979323846
//const zero = 0.0

//B.
//const (
//	size int64 = 1024
//	eof = -1
//)

//C.
//const (
//	ERR_ELEM_EXIST error = errors.New("element already exists")
//	ERR_ELEM_NT_EXIST error = errors.New("element not exists")
//)

//D.
//const u, v float32 = 0, 3
//const a, b, c = 3, 4, "foo"
//参考答案：ABD


//A.
const Pi float64 = 3.14159265358979323846
const zero = 0.0
//B
const (
	size int64 = 1024
	eof = -1
)

//C.
//const (
//	ERR_ELEM_EXIST error = errors.New("element already exists")
//	ERR_ELEM_NT_EXIST error = errors.New("element not exists")
//)

//D.
const u, v float32 = 0, 3
const a, b, c = 3, 4, "foo"

//const英文本身就是  常量,常数  的意思，那么C ERR_ELEM_EXIST，提示报错  这是什么鬼

func main() {
	
}
