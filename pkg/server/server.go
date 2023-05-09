package server

import (
	"flag"
	"fmt"
	"net"

	pb "github.com/mrbryside/pkg/proto/kbtg/kbtg1000"
	"github.com/mrbryside/pkg/service"
	"google.golang.org/grpc"
)

type Config struct {
	// port for gRPC service on HTTP/2
	GRPCPort string
}

func RunServer() error {
	var cfg Config
	flag.StringVar(&cfg.GRPCPort, "grpc-port", "9090", "port to serve gRPC service (default 9090)")
	flag.Parse()

	return RunServerWithConfig(cfg)
}

func RunServerWithConfig(cfg Config) error {
	// ==================== start gRPC server ====================
	gs := grpc.NewServer()

	grpcPort := ":" + cfg.GRPCPort
	grpcL, err := net.Listen("tcp", grpcPort)
	if err != nil {
		fmt.Println("Cannot bind to port", grpcPort)
		return err
	}
	// TODO register all proto api that we created in the gRPC server here
	pb.RegisterEchoServiceServer(gs, service.NewEchoServerImpl())
	// ---------------------------------------------------------

	fmt.Println("Start gRPC server on port", grpcPort)
	if err := gs.Serve(grpcL); err != nil {
		return err
	}

	return nil
}
