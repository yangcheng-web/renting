package subscriber

import (
	"context"
	"github.com/asim/go-micro/v3/util/log"
	pb "renting/GetSmsCd/proto"
)

type GetUserInfo struct{}

func (e *GetUserInfo) Handle(ctx context.Context, msg *pb.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *pb.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
