package main
//实现消息队列（多生产者，多消费者）
//使用切片加锁可以实现


//大文件排序
//归并排序，分而治之,拆分为小文件，在排序

//http get跟head
//
//HEAD和GET本质是一样的，区别在于HEAD不含有呈现数据，而仅仅是HTTP头信息。
//有的人可能觉得这个方法没什么用，其实不是这样的。
//想象一个业务情景：欲判断某个资源是否存在，我们通常使用GET，但这里用HEAD则意义更加明确。


//http 401,403
//
//400 bad request，请求报文存在语法错误
//401 unauthorized，表示发送的请求需要有通过 HTTP 认证的认证信息
//403 forbidden，表示对请求资源的访问被服务器拒绝
//404 not found，表示在服务器上没有找到请求的资源
//http keep-alive
//
//client发出的HTTP请求头需要增加Connection:keep-alive字段
//Web-Server端要能识别Connection:keep-alive字段，并且在http的response里指定Connection:keep-alive字段，
//告诉client，我能提供keep-alive服务，并且"应允"client我暂时不会关闭socket连接

//http能不能一次连接多次请求，不等后端返回
//
//http本质上市使用socket连接，因此发送请求，接写入tcp缓冲，是可以多次进行的，这也是http是无状态的原因

//tcp与udp区别，udp优点，适用场景
//
//tcp传输的是数据流，而udp是数据包，tcp会进过三次握手，udp不需要
func main() {
	
}
