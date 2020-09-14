package worker_server

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"jlambert/FenixInception3/FenixTestInstructionBuilder/common_config"
	"jlambert/FenixInception3/FenixTestInstructionBuilder/grpc_api/mother_server_grpc_api"
	"net"
	"time"
)

type WorkerObject_struct struct {
	logger           *logrus.Logger
	iAmBusy          bool
	uuid             string
	startTime        time.Time
	timeBeforeFinish int32
	currentTaskuuid  string
	currentTaskName  string
	ip               string
	port             string
}

var workerObject *WorkerObject_struct

// Global connection constants
var localServerEngineLocalPort = common_config.LocalWorkerServer_initial_port

var (
	registerWorkerServer *grpc.Server
	lis                  net.Listener
)

var (
	// Standard Mother gRPC Server
	remoteMotherServerConnection *grpc.ClientConn
	motherClient                 mother_server_grpc_api.MotherServerClient

	mother_address_to_dial string = common_config.MotherServer_address + common_config.MotherServer_port
)

// Server used for register clients Name, Ip and Por and Clients Test Enviroments and Clients Test Commandst
type WorkerServer struct{}

const iterationsBetweenAverageTimeCalculationMin = 100
const iterationsBetweenAverageTimeCalculationMax = 1000
