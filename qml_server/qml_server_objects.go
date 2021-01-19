package qml_server

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"jlambert/FenixInception3/FenixTestInstructionBuilder/common_config"
	"jlambert/FenixInception3/FenixTestInstructionBuilder/grpc_api/backend_server_grpc_api"
	"net"
	//"jlambert/AllCombinations/grpc_api/backend_server_grpc_api"
	//"database/sql"
)

type QmlServerObject_struct struct {
	logger                      *logrus.Logger
	dialedBackendGrpcServer     bool
	backendGrpcServerIsAlive    bool
	OnlyForTest_SwitchOfBackend string
	//	qmlBridge                *QmlBridge
	//workers []workerList_struct
	//workerIdChannel chan qml_server_grpc_api.WorkerResult
	//motherDB *sql.DB
	//workerIdToProcessChannel chan int64
	//currentTaskName string
	//messageToWorkerChannel chan backend_server_grpc_api.MessageToWorkerStruct
	//fullTaskStructureObject backend_server_grpc_api.CombinationObjectStruct
}

/*
type clientConnectionInformation_struct struct {
	ip string
	prt string
	uuId	string
}
*/

/*
type workerList_struct struct {
	workerHasGotTask bool
	workerInformation qml_server_grpc_api.WorkerInformation
}


*/
var QmlServerObject *QmlServerObject_struct

var (
	registerQmlGrpcServer *grpc.Server
	lis                   net.Listener
)

var (
	// Standard Worker gRPC Server
	remoteTestInstructionBackendServerConnection *grpc.ClientConn
	//workerClient                 backend_server_grpc_api.WorkerServerClient

	testInstructionBackendServer_address_to_dial string = common_config.TestInstructionBackendServer_address + ":6660" //common_config.QmlServer_port
)

// Client for Connection towards backend server
var testInstructionBackendServerGrpcClient backend_server_grpc_api.TestInstructionBackendGrpcServicesClient

// Server used for register clients Name, Ip and Por and Clients Test Enviroments and Clients Test Commandst
type QMLgRrpServer_struct struct{}
