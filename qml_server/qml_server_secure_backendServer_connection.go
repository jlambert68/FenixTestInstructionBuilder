package qml_server

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"jlambert/FenixInception3/FenixTestInstructionBuilder/grpc_api/backend_server_grpc_api"
	"time"
)

// Very that a connection to Backend Servers gRPC-api is up and open
func (qmlServerObject *QmlServerObject_struct) verifyConnectionTowardsBackendGrpcServer() {

	var err error
	ctx := context.Background()

	for {
		// Set up connection to TestInstructionBackend Server
		remoteTestInstructionBackendServerConnection, err = grpc.Dial(testInstructionBackendServer_address_to_dial, grpc.WithInsecure())
		if err != nil {
			// Couldn't dial backend Server
			qmlServerObject.logger.WithFields(logrus.Fields{
				"Id": "a415ceff-0cd2-4f10-ba72-499ce06e1eea",
				"testInstructionBackendServer_address_to_dial": testInstructionBackendServer_address_to_dial,
				"error message": err,
			}).Warning("Did not connect to TestInstruction Backend Server!")
			qmlServerObject.dialedBackendGrpcServer = false
		} else {
			// Connection to Backend Server was a success
			qmlServerObject.dialedBackendGrpcServer = true
			qmlServerObject.logger.WithFields(logrus.Fields{
				"Id": "b8942280-1334-43f9-8d50-f86f583454af",
				"testInstructionBackendServer_address_to_dial": testInstructionBackendServer_address_to_dial,
			}).Info("gRPC connection OK to TestInstruction Backend Server!")

			// Creates a new gRPC-Client towards Backend server
			testInstructionBackendServerGrpcClient = backend_server_grpc_api.NewTestInstructionBackendGrpcServicesClient(remoteTestInstructionBackendServerConnection)

			// Start up alive check regarding backend gRPC server
			for {
				// If client hasn't dialed Backend Server we can call backend
				if qmlServerObject.dialedBackendGrpcServer == true {
					returnMessage, err := testInstructionBackendServerGrpcClient.AreYouAlive(ctx, &backend_server_grpc_api.EmptyParameter{})

					if err != nil {
						// Backend is not responding as intended
						qmlServerObject.backendGrpcServerIsAlive = false

						qmlServerObject.logger.WithFields(logrus.Fields{
							"Id":     "a2541c7e-046a-4a4f-a688-3649168259d6",
							"client": testInstructionBackendServerGrpcClient,
							"error":  err,
						}).Info("Problem to connect to TestInstruction Backend Server")

						// Can't Connect Backend server, exit loop and try to reconnect from the beginning
						break

					} else if returnMessage.Acknack == false {
						// Backend is not responding as intended
						qmlServerObject.backendGrpcServerIsAlive = false
						qmlServerObject.logger.WithFields(logrus.Fields{
							"id":              "c8b56354-3bef-4bd7-b3ba-ff1b2ff7a179",
							"returnMessage: ": returnMessage,
						}).Debug("TestInstruction Backend Server is not responing as intended")

					} else {
						// Backend is Alive and well
						qmlServerObject.backendGrpcServerIsAlive = true
						qmlServerObject.logger.WithFields(logrus.Fields{
							"id":              "00ded43d-d0fb-4bbd-8d76-20f8136bb32d",
							"returnMessage: ": returnMessage,
						}).Debug("Receive 'I am Alive' from TestInstruction Backend Server")
					}

					// Wait before next "Is Backend Server Alive" test
					qmlServerObject.logger.WithFields(logrus.Fields{
						"id": "9cea845e-1ce8-404a-8a05-8d8c16b7e04e",
					}).Debug("Waiting 30 seconds before next 'Is Backend Server Alive'-test")
					time.Sleep(30 * time.Second) // or runtime.Gosched() or similar per @misterbee

				}
			}
			// Wait before next "Dail-attempt to Backend Server
			qmlServerObject.logger.WithFields(logrus.Fields{
				"id": "7bb43f78-ffdb-4c20-b884-0c2e0aca01ab",
			}).Info("Waiting 30 seconds before next 'Dial to Backend Server'")
			time.Sleep(30 * time.Second)
		}
	}
}
