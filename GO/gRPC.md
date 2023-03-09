#go #gRPC 

---
# Protobuf
Google Protocol Buffer(简称Protobuf)
轻便高效的序列化数据结构的协议，而可以用于网络通信和数据存储。
特点：性能高、传输快、维护方便
一些第三方的rpc库都会支持protobuf
* github地址:https://github.com/protocolbuffers/protobuf
* golang库所属地址:https://github.com/golang/protobuf
# proto文件介绍
1. message
`message`:`protobuf`中定义一个消息类型是通过关键字`message`字段指定的。
消息就是要传输的数据格式的定义。
例如：
```protobuf
// 指定当前proto语法的版本，有2和3  
syntax="proto3";  
// 指定等会生出来成的package  
package service;  
// option go_package="path;name" path表示生成的go文件的存放地址，会自动生成目录 name表示生成的go文件所属的包名  
option go_package="../service";  
//消息 传输的对象  
message User{  
  string username=1;  
  int32 age=2;  
}
```
在消息中承载的数据分别对应于每一个字段。
其中每个字段都有一个名字和一种类型。
2. 字段规则
* `required`:消息体中的必填字段，不设置会导致编解码异常。（例如位置1）
* `option`:消息体中的可选字段。（例如位置2）
* `repeated`:消息体中的可重复字段，重复的值的顺序会被保留（例如位置3）在go中重复的会被定义为切片。
```protobuf
message User{
	string username=1;
	int32 age=2;
	optional string password=3;
	repeated string addresses=4;
}
```