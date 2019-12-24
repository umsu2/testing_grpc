// Package main implements a server for Greeter service.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	sd "github.com/umsu2/testing_grpc/hellosvc"
	"google.golang.org/grpc"
)

const defaultName = "yang"

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := sd.NewHelloWorldSvcClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	t1 := time.Now()
	for i := 0; i < 1000; i++ {
		r, err := c.SayHello(ctx, &sd.HelloRequest{Name: fmt.Sprintf("%s : %d", name, i)})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.GetMessage())
		_, err = c.SayBye(ctx, &sd.ByeRequest{Name: fmt.Sprintf("%s : %d", name, i)})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
	}
	fmt.Println(time.Since(t1))
}
