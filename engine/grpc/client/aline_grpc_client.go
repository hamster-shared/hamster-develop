package client

import (
	"context"
	"github.com/hamster-shared/a-line/engine/grpc/api"
	"io"
	"log"
	"time"
)

func runChat(client api.AlineRPCClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.AlineChat(ctx)
	if err != nil {
		log.Fatalf("client.RouteChat failed: %v", err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("client.RouteChat failed: %v", err)
			}
			log.Printf("Got message %s at point(%d, %d)", in.String(), in.Type, in.GetType())
		}
	}()

	notes := []api.AlineMessage{
		api.AlineMessage{
			Type: 1,
		},
	}

	for _, note := range notes {
		if err := stream.Send(&note); err != nil {
			log.Fatalf("client.RouteChat: stream.Send(%v) failed: %v", note, err)
		}
	}
	stream.CloseSend()
	<-waitc
}
