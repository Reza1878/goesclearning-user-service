package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type RPCConfig struct {
	Port string
}

func RPCDial(cfg RPCConfig) (*grpc.ClientConn, error) {
	address := fmt.Sprintf("localhost:%s", cfg.Port)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(1024*1024*64),
			grpc.MaxCallSendMsgSize(1024*1024*64),
		),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC client: %w", err)
	}

	log.Printf("[Success] - Connected to RPC server at %s", address)
	return conn, nil
}
