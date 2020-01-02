protoc -I helloworld/ helloworld/helloworld.proto --go_out=plugins=grpc:helloworld



protoc -I common/ common/service.proto --go_out=plugins=grpc:common
