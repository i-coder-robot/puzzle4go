package main

//[intermediate] 关于GetPodAction定义，下面赋值正确的是（）
type Fragment interface {
	Exec(transInfo *TransInfo) error
}
type GetPodAction struct {
}
func (g GetPodAction) Exec(transInfo *TransInfo) error {
	...
	return nil
}
//A. var fragment Fragment = new(GetPodAction)
//B. var fragment Fragment = GetPodAction
//C. var fragment Fragment = &GetPodAction{}
//D. var fragment Fragment = GetPodAction{}
//
//参考答案：ACD

//GetPodAction实例，谁可以赋值给接口 Fragment
//A new可以，因为它返回指针类型
//C 很明显了，人家都用&符号了
//B 搞不懂，你把一个结构体赋值过去， 想干啥，什么鬼，完犊子
//D 一个结构体对象赋值给接口变量 ok 啊

func main() {

}
