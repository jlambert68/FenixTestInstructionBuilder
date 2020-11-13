package main

import (
	//	"encoding/json"
	//	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	//"jlambert/FenixInception3/FenixTestInstructionBuilder/grpc_api/mother_server_grpc_api"
	"golang.org/x/net/context"
	"jlambert/FenixInception3/FenixTestInstructionBuilder/grpc_api/worker_server_grpc_api"
	"os"
	"time"
	//"google.golang.org/grpc"
	"jlambert/FenixInception3/FenixTestInstructionBuilder/common_config"
)

// *********************************************************************
//Clean up when Server should be closed down
var cleanupProcessed bool = false

func cleanup() {

	if cleanupProcessed == false {

		cleanupProcessed = true

		// Cleanup before close down application
		motherObject.logger.WithFields(logrus.Fields{
			"ID": "ab1abcdb-d786-4bcc-b274-b52bd931f43d",
		}).Info("Clean up and shut down servers")

		motherObject.logger.WithFields(logrus.Fields{
			"ID": "a22e8de2-c38c-4483-95e0-e9dc21dd7ca8",
		}).Info("Gracefull stop for: registerMotherServer")

		registerMotherServer.GracefulStop()

		motherObject.logger.WithFields(logrus.Fields{
			"common_config.MotherServer_port: ": common_config.MotherServer_port,
			"ID":                                "19593393-3ac9-45ac-96e9-cdf911b167c7",
		}).Info("Close net.Listing")
		lis.Close()

		//log.Println("Close DB_session: %v", DB_session)
		//DB_session.Close()
	}
}

// Main function that starts everything
// Designed this way becasue of compilation process of "qtrecipe"
func main() {

	// Cleanup when closing
	defer cleanup()

	// Set up MotherObject
	motherObject = &MotherObject_struct{}

	// Init logger
	motherObject.InitLogger("")

	// Inititate and start Mother GRPC-engine
	initiateAndStartMotherGrpcServer()

	// Start Ping Alive towards Worker

	var err error

	// Set up connection to Mother Server
	remoteWorkerServerConnection, err = grpc.Dial(worker_address_to_dial, grpc.WithInsecure())
	if err != nil {
		motherObject.logger.WithFields(logrus.Fields{
			"worker_address_to_dial": worker_address_to_dial,
			"error message":          err,
		}).Error("Did not connect to Mother Server!")
		os.Exit(0)
	} else {
		motherObject.logger.WithFields(logrus.Fields{
			"worker_address_to_dial": worker_address_to_dial,
		}).Info("gRPC connection OK to Mother Server!")

		// Creates a new Clients
		workerClient := worker_server_grpc_api.NewWorkerServerClient(remoteWorkerServerConnection)

		ctx := context.Background()
		go func() {
			time.Sleep(30 * time.Second) // or runtime.Gosched() or similar per @misterbee

			returnMessage, err := workerClient.AreYouAlive(ctx, &worker_server_grpc_api.EmptyParameter{})

			if err != nil {
				motherObject.logger.WithFields(logrus.Fields{
					"client": workerClient,
					"error":  err,
				}).Fatal("Problem to connect to Worker Server")
			}

			motherObject.logger.WithFields(logrus.Fields{
				"returnMessage: ": returnMessage,
			}).Info("Receive I am Alive from Worker")

		}()

		// Initate and Start QML-engine
		// ***** LAST PART OF THE PROGRAM *****
		initiateAndStartQmlEngine()

	}
}
