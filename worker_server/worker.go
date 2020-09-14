package worker_server

import (
	"fmt"
	"google.golang.org/grpc"
	"jlambert/FenixInception3/FenixTestInstructionBuilder/grpc_api/worker_server_grpc_api"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"jlambert/FenixInception3/FenixTestInstructionBuilder/common_config"
	"jlambert/FenixInception3/FenixTestInstructionBuilder/grpc_api/mother_server_grpc_api"
	"strconv"
)

func (workerObject *WorkerObject_struct) SendMotherIpAndPortForWorker() {

	var err error

	// Set up connection to Mother Server
	remoteMotherServerConnection, err = grpc.Dial(mother_address_to_dial, grpc.WithInsecure())
	if err != nil {
		workerObject.logger.WithFields(logrus.Fields{
			"mother_address_to_dial": mother_address_to_dial,
			"error message":          err,
		}).Error("Did not connect to Mother Server!")
		os.Exit(0)
	} else {
		workerObject.logger.WithFields(logrus.Fields{
			"mother_address_to_dial": mother_address_to_dial,
		}).Info("gRPC connection OK to Mother Server!")

		// Creates a new Clients
		workerClient := mother_server_grpc_api.NewMotherServerClient(remoteMotherServerConnection)

		//messageToMother := &mother_server_grpc_api.WorkerInformation{workerObject.ip, workerObject.port, workerObject.uuid, ""}
		messageToMother := &mother_server_grpc_api.WorkerInformation{
			WorkerIp:             workerObject.ip,
			WorkerPort:           workerObject.port,
			WorkerUuid:           workerObject.uuid,
			WorkerTaskUuid:       "",
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		}
		ctx := context.Background()
		returnMessage, err := workerClient.WorkerIPandPortIs(ctx, messageToMother)
		if err != nil {
			workerObject.logger.WithFields(logrus.Fields{
				"client": workerClient,
				"error":  err,
			}).Fatal("Problem to connect to Mother Server")
		}

		workerObject.logger.WithFields(logrus.Fields{
			"returnMessage: ":     returnMessage,
			"workerObject.ip: ":   workerObject.ip,
			"workerObject.port: ": workerObject.port,
			"workerObject.uuid":   workerObject.uuid,
		}).Info("Sent IP and Port to Mother Server")

	}

}

// Used for only process cleanup once
var cleanupProcessed bool = false

func cleanup() {

	if cleanupProcessed == false {

		cleanupProcessed = true

		// Cleanup before close down application
		workerObject.logger.WithFields(logrus.Fields{}).Info("Clean up and shut down servers")

		workerObject.logger.WithFields(logrus.Fields{}).Info("Gracefull stop for: registerTaxiHardwareStreamServer")
		registerWorkerServer.GracefulStop()

		workerObject.logger.WithFields(logrus.Fields{
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
	workerObject = &WorkerObject_struct{iAmBusy: false}

	// Create unique id for this worker
	uuId, _ := uuid.NewUUID()
	fmt.Println(uuId)
	workerObject.uuid = uuId.String()

	//

	// Init logger
	workerObject.InitLogger("")

	// Find first non allocated port from defined start port
	workerObject.logger.WithFields(logrus.Fields{}).Info("Worker Server tries to start")
	for counter := 0; counter < 10; counter++ {
		localServerEngineLocalPort = localServerEngineLocalPort + counter
		workerObject.logger.WithFields(logrus.Fields{
			"localServerEngineLocalPort: ": localServerEngineLocalPort,
		}).Info("Start listening on:")
		lis, err = net.Listen("tcp", ":"+strconv.Itoa(localServerEngineLocalPort))

		if err != nil {
			workerObject.logger.WithFields(logrus.Fields{
				"err: ": err,
			}).Error("failed to listen:")
		} else {
			workerObject.logger.WithFields(logrus.Fields{
				"localServerEngineLocalPort: ": localServerEngineLocalPort,
			}).Info("Success in listening on port:")
			workerObject.port = strconv.Itoa(localServerEngineLocalPort)
			workerObject.ip = common_config.LocalWorkerServer_address

			break
		}
	}

	// Creates a new RegisterWorkerServer gRPC server
	go func() {
		workerObject.logger.WithFields(logrus.Fields{}).Info("Starting Worker Server")
		registerWorkerServer = grpc.NewServer()
		worker_server_grpc_api.RegisterWorkerServerServer(registerWorkerServer, &WorkerServer{})

		workerObject.logger.WithFields(logrus.Fields{
			"localServerEngineLocalPort: ": localServerEngineLocalPort,
		}).Info("registerWorkerServer for Worker Server started")
		registerWorkerServer.Serve(lis)
	}()

	// Register at Mother Server
	workerObject.SendMotherIpAndPortForWorker()

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
