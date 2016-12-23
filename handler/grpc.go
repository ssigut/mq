package handler

import (
	"fmt"

	"github.com/asim/mq/broker"
	"github.com/asim/mq/proto/grpc/mq"
	"golang.org/x/net/context"
)

type GRPC struct{}

func (g *GRPC) Pub(ctx context.Context, req *mq.PubRequest) (*mq.PubResponse, error) {
	if err := broker.Publish(req.Topic, req.Payload); err != nil {
		return nil, fmt.Errorf("pub error: %v", err)
	}
	return new(mq.PubResponse), nil
}

func (g *GRPC) Sub(req *mq.SubRequest, stream mq.MQ_SubServer) error {
	ch, err := broker.Subscribe(req.Topic)
	if err != nil {
		return fmt.Errorf("could not subscribe: %v", err)
	}
	defer broker.Unsubscribe(req.Topic, ch)

	for p := range ch {
		if err := stream.Send(&mq.SubResponse{Payload: p}); err != nil {
			return fmt.Errorf("failed to send payload: %v", err)
		}
	}

	return nil
}
