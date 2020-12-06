package qml_server

import (
	//	"encoding/json"
	//	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"time"
	//"google.golang.org/grpc"
	"jlambert/FenixInception3/FenixTestInstructionBuilder/common_config"
)

// *********************************************************************
//Clean up when Server should be closed down
var cleanupProcessed bool = false

func Cleanup() {

	if cleanupProcessed == false {

		cleanupProcessed = true

		// Cleanup before close down application
		QmlServerObject.logger.WithFields(logrus.Fields{
			"ID": "ab1abcdb-d786-4bcc-b274-b52bd931f43d",
		}).Info("Clean up and shut down servers")

		QmlServerObject.logger.WithFields(logrus.Fields{
			"ID": "a22e8de2-c38c-4483-95e0-e9dc21dd7ca8",
		}).Info("Gracefull stop for: registerQmlGrpcServer")

		registerQmlGrpcServer.GracefulStop()

		QmlServerObject.logger.WithFields(logrus.Fields{
			"common_config.QmlServer_port: ": common_config.QmlServer_port,
			"ID":                             "19593393-3ac9-45ac-96e9-cdf911b167c7",
		}).Info("Close net.Listing")
		lis.Close()

		//log.Println("Close DB_session: %v", DB_session)
		//DB_session.Close()
	}
}

// *********************************************************************
// Main function that starts everything
// Designed this way becasue of compilation process of "qtrecipe"
func Start_qml_server(mylogger *logrus.Logger) {

	// Set up QmlServerObject
	QmlServerObject = &QmlServerObject_struct{
		dialedBackendGrpcServer:  false,
		backendGrpcServerIsAlive: false,
		//qmlBridge : qmlBridge,
		logger: mylogger,
	}

	// Init logger
	//QmlServerObject.InitLogger("")

	// Initiate channels for communication between QML and gRPC
	initiateChannels()

	// Inititate and start QML Server GRPC-engine
	InitiateAndStartQmlGrpcServer()

	// Start Ping Alive towards TestInstructionBackendServer
	go QmlServerObject.verifyConnectionTowardsBackendGrpcServer()

	// Wait until backend server is up and running
	//
	for {
		if QmlServerObject.dialedBackendGrpcServer == true && QmlServerObject.backendGrpcServerIsAlive == true {
			// Only start qml server if a gRPC connection has been done
			QmlServerObject.logger.WithFields(logrus.Fields{
				"id": "5199753b-2f2b-4e7e-83ef-6581fcf686bb",
			}).Info("QML-gRPC bridge is ready for use")

			// Activate QML-gRPC-bridge to be  ready for use
			StartStopQmlGrpcBridgr(true)
			break
			// Due to qtreceipt compiling style, the program doesn't continue from here

		} else {
			QmlServerObject.logger.WithFields(logrus.Fields{
				"id": "44692e25-625e-4857-aec2-6d617489de9d",
				"QmlServerObject.dialedBackendGrpcServer":  QmlServerObject.dialedBackendGrpcServer,
				"QmlServerObject.backendGrpcServerIsAlive": QmlServerObject.backendGrpcServerIsAlive,
			}).Info("sleeping...for another 15 seconds wail waiting for Backend to start")
		}
		//fmt.Println("sleeping...for another 15 seconds wail waiting for Backend to start")
		time.Sleep(15 * time.Second) // or runtime.Gosched() or similar per @misterbee
	}

}
