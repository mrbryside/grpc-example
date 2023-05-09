package service

import (
	"context"

	pb "github.com/mrbryside/pkg/proto/kbtg/kbtg1000"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type EchoServerImpl struct{}

func NewEchoServerImpl() *EchoServerImpl {
	return &EchoServerImpl{}
}

func (e *EchoServerImpl) Query(ctx context.Context, pl *pb.Echo) (*pb.EchoList, error) {
	var eList []*pb.Echo
	eList = append(eList, &pb.Echo{
		Id:          "1",
		Name:        pl.Name,
		Description: "Desc",
		Status:      "SUCCESS",
		CreateDate:  timestamppb.Now(),
		UpdateDate:  timestamppb.Now(),
	})

	resp := &pb.EchoList{
		Api:      "1",
		EchoList: eList,
	}

	return resp, nil
}
