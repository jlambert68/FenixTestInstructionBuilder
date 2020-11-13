package main

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"jlambert/FenixInception3/FenixTestInstructionBuilder/common_config"
	"net"
	//"jlambert/AllCombinations/grpc_api/worker_server_grpc_api"
	//"database/sql"
)

type MotherObject_struct struct {
	logger *logrus.Logger
	//workers []workerList_struct
	//workerIdChannel chan mother_server_grpc_api.WorkerResult
	//motherDB *sql.DB
	//workerIdToProcessChannel chan int64
	//currentTaskName string
	//messageToWorkerChannel chan worker_server_grpc_api.MessageToWorkerStruct
	//fullTaskStructureObject worker_server_grpc_api.CombinationObjectStruct
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
	workerInformation mother_server_grpc_api.WorkerInformation
}


*/
var motherObject *MotherObject_struct

var (
	registerMotherServer *grpc.Server
	lis                  net.Listener
)

var (
	// Standard Worker gRPC Server
	remoteWorkerServerConnection *grpc.ClientConn
	//workerClient                 worker_server_grpc_api.WorkerServerClient

	worker_address_to_dial string = common_config.LocalWorkerServer_address + ":6660" //common_config.MotherServer_port
)

// Server used for register clients Name, Ip and Por and Clients Test Enviroments and Clients Test Commandst
type MotherServer struct{}

// The data stored in DB to keep verisons of a project
type projectsObjectInDB_struct struct {
	version_key         string
	time_nanoint        int64
	date_time_string    string
	project_name        string
	project_description string
	deleted             bool
	project_key         string
}

/*
// Object used in transformation from/to GUI-Object to/from DB
type GUIObjectToSave_struct struct {
	Ref         string                                    `json:"$ref"`
	Headers     []worker_server_grpc_api.HeaderStruct     `json:"headers"`
	ValueSets   []worker_server_grpc_api.ValueSetStruct   `json:"valueSets"`
	Rules       []worker_server_grpc_api.RuleStruct       `json:"rules"`
	HeaderTypes []worker_server_grpc_api.HeaderTypeStruct `json:"headerTypes"`
	RuleTypes   []worker_server_grpc_api.RuleTypeStruct   `json:"ruleTypes"`
}

*/

// Variables used for dicide what kind of object to save in DB
// 1) Ordinary Combination Object
// 2) HeadersAndHeaderValues received from jExcel
type objectTypeToSave int

const (
	OrdinaryCombinationObject    objectTypeToSave = 1
	HeadersAndHeaderValuesObject objectTypeToSave = 2
	NewOrdinaryCombinationObject objectTypeToSave = 3
)

// Used to set HeaderTypes
const (
	HeaderTypeStandard_id  = 1
	HeaderTypeStandardText = "Standard"
)
