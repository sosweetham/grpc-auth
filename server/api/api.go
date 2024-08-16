package api

import (
	"net"
	"os"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/sohamjaiswal/grpc-ftp/pkg/pb"
)

type server struct {
	pb.UnimplementedGrpcFtpServer
}

func Start() {
	hostAddress := os.Getenv("HOST") + ":" + os.Getenv("PORT")

	lis, err := net.Listen("tcp", hostAddress)

	if err != nil {
		log.Fatalf("Failed to start listener: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGrpcFtpServer(grpcServer, &server{})
	reflection.Register(grpcServer)
	log.Printf("FTP Service registered at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed tp serve: %v", err)
	}
}