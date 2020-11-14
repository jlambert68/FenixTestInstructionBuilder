package testinstruction_beckend_server

import (
	"fmt"
	"google.golang.org/grpc"
	"jlambert/FenixInception3/FenixTestInstructionBuilder/grpc_api/backend_server_grpc_api"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"jlambert/FenixInception3/FenixTestInstructionBuilder/common_config"
	"jlambert/FenixInception3/FenixTestInstructionBuilder/grpc_api/qml_server_grpc_api"
	"strconv"
)

func (testInstructionBackendObject *TestInstructionBackendObject_struct) SendMotherIpAndPortForWorker() {

	var err error

	// Set up connection to Mother Server
	remoteQmlServerConnection, err = grpc.Dial(qmlServer_address_to_dial, grpc.WithInsecure())
	if err != nil {
		testInstructionBackendObject.logger.WithFields(logrus.Fields{
			"qmlServer_address_to_dial": qmlServer_address_to_dial,
			"error message":             err,
		}).Error("Did not connect to Mother Server!")
		os.Exit(0)
	} else {
		testInstructionBackendObject.logger.WithFields(logrus.Fields{
			"qmlServer_address_to_dial": qmlServer_address_to_dial,
		}).Info("gRPC connection OK to Mother Server!")

		// Creates a new Clients
		workerClient := qml_server_grpc_api.NewMotherServerClient(remoteQmlServerConnection)

		//messageToMother := &qml_server_grpc_api.WorkerInformation{testInstructionBackendObject.ip, testInstructionBackendObject.port, testInstructionBackendObject.uuid, ""}
		messageToMother := &qml_server_grpc_api.WorkerInformation{
			WorkerIp:             testInstructionBackendObject.ip,
			WorkerPort:           testInstructionBackendObject.port,
			WorkerUuid:           testInstructionBackendObject.uuid,
			WorkerTaskUuid:       "",
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		}
		ctx := context.Background()
		returnMessage, err := workerClient.WorkerIPandPortIs(ctx, messageToMother)
		if err != nil {
			testInstructionBackendObject.logger.WithFields(logrus.Fields{
				"client": workerClient,
				"error":  err,
			}).Fatal("Problem to connect to Mother Server")
		}

		testInstructionBackendObject.logger.WithFields(logrus.Fields{
			"returnMessage: ":                     returnMessage,
			"testInstructionBackendObject.ip: ":   testInstructionBackendObject.ip,
			"testInstructionBackendObject.port: ": testInstructionBackendObject.port,
			"testInstructionBackendObject.uuid":   testInstructionBackendObject.uuid,
		}).Info("Sent IP and Port to Mother Server")

	}

}

// Used for only process cleanup once
var cleanupProcessed bool = false

func cleanup() {

	if cleanupProcessed == false {

		cleanupProcessed = true

		// Cleanup before close down application
		testInstructionBackendObject.logger.WithFields(logrus.Fields{}).Info("Clean up and shut down servers")

		testInstructionBackendObject.logger.WithFields(logrus.Fields{}).Info("Gracefull stop for: registerTaxiHardwareStreamServer")
		registerTestInstructionBackendServer.GracefulStop()

		testInstructionBackendObject.logger.WithFields(logrus.Fields{
			"localServerEngineLocalPort: ": localServerEngineLocalPort,
		}).Info("Close net.Listing")
		_ = lis.Close()

		//log.Println("Close DB_session: %v", DB_session)
		//DB_session.Close()
	}
}

func Worker_main() {

	var err error

	defer cleanup()

	// Set up WorkerObject
	testInstructionBackendObject = &TestInstructionBackendObject_struct{iAmBusy: false}

	// Create unique id for this worker
	uuId, _ := uuid.NewUUID()
	fmt.Println(uuId)
	testInstructionBackendObject.uuid = uuId.String()

	//

	// Init logger
	testInstructionBackendObject.InitLogger("")

	// Find first non allocated port from defined start port
	testInstructionBackendObject.logger.WithFields(logrus.Fields{}).Info("Worker Server tries to start")
	for counter := 0; counter < 10; counter++ {
		localServerEngineLocalPort = localServerEngineLocalPort + counter
		testInstructionBackendObject.logger.WithFields(logrus.Fields{
			"localServerEngineLocalPort: ": localServerEngineLocalPort,
		}).Info("Start listening on:")
		lis, err = net.Listen("tcp", ":"+strconv.Itoa(localServerEngineLocalPort))

		if err != nil {
			testInstructionBackendObject.logger.WithFields(logrus.Fields{
				"err: ": err,
			}).Error("failed to listen:")
		} else {
			testInstructionBackendObject.logger.WithFields(logrus.Fields{
				"localServerEngineLocalPort: ": localServerEngineLocalPort,
			}).Info("Success in listening on port:")
			testInstructionBackendObject.port = strconv.Itoa(localServerEngineLocalPort)
			testInstructionBackendObject.ip = common_config.LocalWorkerServer_address

			break
		}
	}

	// Creates a new RegisterWorkerServer gRPC server
	go func() {
		testInstructionBackendObject.logger.WithFields(logrus.Fields{}).Info("Starting Worker Server")
		registerTestInstructionBackendServer = grpc.NewServer()
		backend_server_grpc_api.RegisterTestInstructionBackendGrpcServicesServer(registerTestInstructionBackendServer, &TestInstructionBackendServer{})

		testInstructionBackendObject.logger.WithFields(logrus.Fields{
			"localServerEngineLocalPort: ": localServerEngineLocalPort,
		}).Info("registerTestInstructionBackendServer for TestInstruction Backend Server started")
		registerTestInstructionBackendServer.Serve(lis)
	}()

	// Register at Mother Server
	testInstructionBackendObject.SendMotherIpAndPortForWorker()

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanup()
		os.Exit(0)
	}()

	for {
		fmt.Println("sleeping...for another 5 minutes")
		time.Sleep(300 * time.Second) // or runtime.Gosched() or similar per @misterbee
	}

	//Wait until user exit
	/*
		   for {
			   time.Sleep(10)
		   }
	*/
}
