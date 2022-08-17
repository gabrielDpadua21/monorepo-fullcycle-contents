package services

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/codeedu/fc2-grpc/pb"
)

type CatService struct {
	pb.UnimplementedCatServiceServer
}

func NewCatService() *CatService {
	return &CatService{}
}

func (*CatService) AddCat(ctx context.Context, request *pb.Cat) (*pb.Cat, error) {
	id := "123"

	return &pb.Cat{
		Id: &id,
		Name: request.GetName(),
		Color: request.GetColor(),
		Age: request.GetAge(),
	}, nil

}

func (*CatService) AddCatVerbose(req *pb.Cat, stream pb.CatService_AddCatVerboseServer) error {

	stream.Send(&pb.CatResultStream{
		Status: "Init",
		Cat: &pb.Cat{},	
	})

	time.Sleep(time.Second * 3)

	stream.Send(&pb.CatResultStream{
		Status: "Insert",
		Cat: &pb.Cat{},	
	})

	time.Sleep(time.Second * 3)

	stream.Send(&pb.CatResultStream{
		Status: "User insert",
		Cat: &pb.Cat{
			Name: req.GetName(),
			Color: req.GetColor(),
			Age: req.GetAge(),
		},	
	})

	time.Sleep(time.Second * 3)

	stream.Send(&pb.CatResultStream{
		Status: "Completed",
		Cat: &pb.Cat{
			Name: req.GetName(),
			Color: req.GetColor(),
			Age: req.GetAge(),
		},	
	})

	time.Sleep(time.Second * 3)

	return nil
}

func (*CatService) AddCats(stream pb.CatService_AddCatsServer) error {
	cats := []*pb.Cat{}

	for {
		req, err := stream.Recv()
		
		if err == io.EOF {
			return stream.SendAndClose(&pb.Cats{
				Cat: cats,
			})
		}

		if err != nil {
			log.Fatal("Error receving stream: %v", err)
		}

		cats = append(cats, &pb.Cat{
			Name: req.GetName(),
			Color: req.GetColor(),
			Age: req.GetAge(),
		})
		fmt.Println("Adding", req.GetName())
	}
}