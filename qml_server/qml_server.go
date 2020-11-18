package qml_server

import (
	//	"encoding/json"
	//	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	//"jlambert/FenixInception3/FenixTestInstructionBuilder/grpc_api/qml_server_grpc_api"
	"golang.org/x/net/context"
	"jlambert/FenixInception3/FenixTestInstructionBuilder/grpc_api/backend_server_grpc_api"
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
		qmlServerObject.logger.WithFields(logrus.Fields{
			"ID": "ab1abcdb-d786-4bcc-b274-b52bd931f43d",
		}).Info("Clean up and shut down servers")

		qmlServerObject.logger.WithFields(logrus.Fields{
			"ID": "a22e8de2-c38c-4483-95e0-e9dc21dd7ca8",
		}).Info("Gracefull stop for: registerQmlGrpcServer")

		registerQmlGrpcServer.GracefulStop()

		qmlServerObject.logger.WithFields(logrus.Fields{
			"common_config.QmlServer_port: ": common_config.QmlServer_port,
			"ID":                             "19593393-3ac9-45ac-96e9-cdf911b167c7",
		}).Info("Close net.Listing")
		lis.Close()

		//log.Println("Close DB_session: %v", DB_session)
		//DB_session.Close()
	}
}

// Main function that starts everything
// Designed this way becasue of compilation process of "qtrecipe"
func Start_qml_server() {

	// Cleanup when closing
	defer cleanup()

	// Set up QmlServerObject
	qmlServerObject = &QmlServerObject_struct{}

	// Init logger
	qmlServerObject.InitLogger("")

	// Inititate and start QML Server GRPC-engine
	InitiateAndStartQmlGrpcServer()

	// Start Ping Alive towards TestInstructionBackendServer

	var err error

	// Set up connection to TestInstructionBackend Server
	remoteTestInstructionBackendServerConnection, err = grpc.Dial(testInstructionBackendServer_address_to_dial, grpc.WithInsecure())
	if err != nil {
		qmlServerObject.logger.WithFields(logrus.Fields{
			"Id": "a415ceff-0cd2-4f10-ba72-499ce06e1eea",
			"testInstructionBackendServer_address_to_dial": testInstructionBackendServer_address_to_dial,
			"error message": err,
		}).Error("Did not connect to TestInstruction Backend Server!")
		os.Exit(0)
	} else {
		qmlServerObject.logger.WithFields(logrus.Fields{
			"testInstructionBackendServer_address_to_dial": testInstructionBackendServer_address_to_dial,
		}).Info("gRPC connection OK to TestInstruction Backend Server!")

		// Creates a new Clients
		testInstructionBackendServerGrpcClient = backend_server_grpc_api.NewTestInstructionBackendGrpcServicesClient(remoteTestInstructionBackendServerConnection)

		ctx := context.Background()
		go func() {
			time.Sleep(30 * time.Second) // or runtime.Gosched() or similar per @misterbee

			returnMessage, err := testInstructionBackendServerGrpcClient.AreYouAlive(ctx, &backend_server_grpc_api.EmptyParameter{})

			if err != nil {
				qmlServerObject.logger.WithFields(logrus.Fields{
					"client": testInstructionBackendServerGrpcClient,
					"error":  err,
				}).Fatal("Problem to connect to TestInstruction Backend Server")
			}

			qmlServerObject.logger.WithFields(logrus.Fields{
				"returnMessage: ": returnMessage,
			}).Info("Receive I am Alive from TestInstruction Backend Server")

		}()

		// Initate and Start QML-engine
		// ***** LAST PART OF THE PROGRAM *****
		initiateAndStartQmlEngine()

	}
}
