package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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
	//AddCat(client)
	AddCats(client)
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
	AddCatVerbose(client)
}

func AddCatVerbose (client pb.CatServiceClient) {
	req := &pb.Cat{
		Name: "Frajola",
		Color: "Black and White",
		Age: "12",
	}

	responseStream, err := client.AddCatVerbose(context.Background(), req)
	if err != nil {
		log.Fatalf("Erro to call GRPC request stream: %v", err)
	}

	for {
		stream, err := responseStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Could not recive de messag stream: %v", err)
		}
		fmt.Println("Status: ", stream.Status)
	}
}

func AddCats (client pb.CatServiceClient) {
	reqs := []*pb.Cat{
		&pb.Cat{
			Name: "Frajola",
			Color: "White and Black",
			Age: "12",
		},
		&pb.Cat{
			Name: "Thor",
			Color: "White",
			Age: "8",
		},
		&pb.Cat{
			Name: "Lucyfer",
			Color: "Yallow",
			Age: "4",
		},
		&pb.Cat{
			Name: "Zeuzz",
			Color: "White and Black",
			Age: "1",
		},
	}

	stream, err := client.AddCats(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receveing response: %v", err)
	}
	fmt.Println(res)
}