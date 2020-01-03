【go语言专题】JAVA&GO通过GRPC互相调用  [https://blog.csdn.net/qq_37362891/article/details/103809482](https://blog.csdn.net/qq_37362891/article/details/103809482)



protoc -I helloworld/ helloworld/helloworld.proto --go_out=plugins=grpc:helloworld
protoc -I common/ common/service.proto --go_out=plugins=grpc:common


