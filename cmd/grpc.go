package cmd

import (
	"github.com/spf13/viper"
)

func ServeGRPC(config *viper.Viper) {
	// var (
	// 	log = helpers.Logger
	// )
	// port := config.GetString("grpc.port")
	// lis, err := net.Listen("tcp", ":"+port)
	// if err != nil {
	// 	log.Fatal("failed to serve grpc: ", err)
	// }

	// s := grpc.NewServer()

	// logrus.Info("listening on ", port)
	// if err := s.Serve(lis); err != nil {
	// 	log.Fatal("failed to serve grpc: ", err)
	// }
}
