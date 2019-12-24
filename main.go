// Package main implements a server for Greeter service.
package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	sd "github.com/umsu2/testing_grpc/hellosvc"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

func (*server) SayHello(ctx context.Context, req *sd.HelloRequest) (*sd.HelloReply, error) {
	return &sd.HelloReply{Message: fmt.Sprintf("hello mr {%s}", req.Name)}, nil
}
func (*server) SayBye(ctx context.Context, req *sd.ByeRequest) (*empty.Empty, error) {
	fmt.Println(req.Name)
	return &empty.Empty{}, nil
}
func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	t1 := time.Now()
	s := &server{}
	name := "yang"
	ctx := context.Background()
	for i := 0; i < 1000; i++ {
		r, err := s.SayHello(ctx, &sd.HelloRequest{Name: fmt.Sprintf("%s : %d", name, i)})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.GetMessage())
		_, err = s.SayBye(ctx, &sd.ByeRequest{Name: fmt.Sprintf("%s : %d", name, i)})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
	}
	fmt.Println(time.Since(t1))
	gs := grpc.NewServer()
	sd.RegisterHelloWorldSvcServer(gs, s)
	fmt.Println("staring to serve...")
	if err := gs.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
