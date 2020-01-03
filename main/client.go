package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "go-grpc/helloworld"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

//'''
//请求远程方法-查询用户
//'''
//def findAll():
//_request = {
//'clazz': 'com.anoyi.grpc.facade.service.UserServiceByFastJSON',
//'method': 'findAll'
//   }
//_request_json = json.dumps(_request, ensure_ascii=False)
//with grpc.insecure_channel('127.0.0.1:6565') as channel:
//stub = service_pb2_grpc.CommonServiceStub(channel)
//response = stub.handle(service_pb2.Request(request=bytes(_request_json), serialize=3))
//print("received: " + response.response)

func main() {
	//建立链接
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	// 1秒的上下文
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
