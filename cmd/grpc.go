package cmd

import (
	"log"
	"net"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func ServeGRPC(config *viper.Viper) {
	port := config.GetString("grpc.port")
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal("failed to serve grpc: ", err)
	}

	s := grpc.NewServer()

	logrus.Info("listening on ", port)
	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to serve grpc: ", err)
	}
}
