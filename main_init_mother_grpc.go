package main

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"jlambert/FenixInception3/FenixTestInstructionBuilder/common_config"
	"jlambert/FenixInception3/FenixTestInstructionBuilder/grpc_api/mother_server_grpc_api"
	"net"
)

// *********************************************************************
// Inititate and starts Mother GRPC-engine
func initiateAndStartMotherGrpcServer() {
	var err error

	motherObject.logger.WithFields(logrus.Fields{
		"ID": "56b6419f-d714-4ab0-be62-f3c7f08b9558",
	}).Info("Mother Server tries to start")

	// Find first non allocated port from defined start port
	// TODO

	motherObject.logger.WithFields(logrus.Fields{
		"common_config.MotherServer_port): ": common_config.MotherServer_port,
		"ID":                                 "8f904ace-9d24-452b-891a-5b8e5c247ba2",
	}).Info("Start listening on:")
	lis, err = net.Listen("tcp", string(common_config.MotherServer_port))

	if err != nil {
		motherObject.logger.WithFields(logrus.Fields{
			"err: ": err,
			"ID":    "0fbf0f08-6114-4cd8-9992-6a387955fb5f",
		}).Fatal("failed to listen:")

	} else {
		motherObject.logger.WithFields(logrus.Fields{
			"common_config.MotherServer_port): ": common_config.MotherServer_port,
			"ID":                                 "93496c07-2b6c-4edc-a1f9-3fd555fa1201",
		}).Info("Success in listening on port:")

	}

	// Creates a new RegisterMotherServer gRPC server
	go func() {
		motherObject.logger.WithFields(logrus.Fields{
			"ID": "ebc26735-9d13-4b13-91b8-20999cd5e254",
		}).Info("Starting Mother Server")

		registerMotherServer = grpc.NewServer()
		mother_server_grpc_api.RegisterMotherServerServer(registerMotherServer, &MotherServer{})

		motherObject.logger.WithFields(logrus.Fields{
			"common_config.MotherServer_port): ": common_config.MotherServer_port,
			"ID":                                 "cfed87c4-55aa-4cd1-980a-a15eb75ab6fb",
		}).Info("registerMotherServer for Mother Server started")

		registerMotherServer.Serve(lis)
	}()
}
