package main

import (
	"context"
	"fmt"
	//"io"
	"log"
	//"time"

	"github.com/codeedu/fc2-grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Erro to connect GRPC server: %v", err)
	}
	defer connection.Close()

	client := pb.NewCatServiceClient(connection)
	AddCat(client)
}

func AddCat(client pb.CatServiceClient) {
	req := &pb.Cat{
		Name: "Frajola",
		Color: "Black and White",
		Age: "12",
	}

	res, err := client.AddCat(context.Background(), req)
	if err != nil {
		log.Fatalf("Erro to call GRPC request: %v", err)
	}
	fmt.Println(res)
}