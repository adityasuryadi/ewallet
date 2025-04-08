package main

import (
	"github.com/adityasuryadi/ewallet/bootstrap"
	"github.com/adityasuryadi/ewallet/cmd"
	"github.com/adityasuryadi/ewallet/helpers"
)

func main() {

	config := bootstrap.NewViper("./config")
	// logger := bootstrap.SetupLogger()

	// db := bootstrap.NewDatabase(config, logger)
	go cmd.ServeGRPC(config)

	helpers.SetupLogger()

	cmd.ServeHttp(config)
}
