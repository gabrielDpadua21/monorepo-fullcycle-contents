package services

import (
	"context"
	//"fmt"
	//"io"
	//"log"
	//"time"

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