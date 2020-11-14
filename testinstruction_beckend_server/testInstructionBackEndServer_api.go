package testinstruction_beckend_server

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"jlambert/FenixInception3/FenixTestInstructionBuilder/grpc_api/backend_server_grpc_api"
)

func (s *TestInstructionBackendServer) WorkerTask(ctx context.Context, messageToWorker *backend_server_grpc_api.MessageToWorkerStruct) (*backend_server_grpc_api.AckNackResponse, error) {

	var returnMessage *backend_server_grpc_api.AckNackResponse

	testInstructionBackendObject.logger.WithFields(logrus.Fields{
		"messageToWorker: ": messageToWorker,
	}).Info("Incoming: 'WorkerTask'")

	if testInstructionBackendObject.iAmBusy == true {
		testInstructionBackendObject.logger.WithFields(logrus.Fields{
			"testInstructionBackendObject.iAmBusy: ": testInstructionBackendObject.iAmBusy,
		}).Info("Worker allready busy do calculations, can't new work load")
		returnMessage = &backend_server_grpc_api.AckNackResponse{
			Acknack:  false,
			Comments: "Worker allready busy do calculations, can't new work load",
		}
	} else {
		testInstructionBackendObject.logger.WithFields(logrus.Fields{
			"testInstructionBackendObject.iAmBusy: ": testInstructionBackendObject.iAmBusy,
		}).Info("Take new work load")
		testInstructionBackendObject.iAmBusy = true
		testInstructionBackendObject.currentTaskuuid = messageToWorker.TaskUuid
		testInstructionBackendObject.currentTaskName = messageToWorker.TaskName
		returnMessage = &backend_server_grpc_api.AckNackResponse{Acknack: true, Comments: "Accepted Task"}
		//go testInstructionBackendObject.ExecuteTaskFromMother(messageToWorker)
	}

	testInstructionBackendObject.logger.WithFields(logrus.Fields{}).Info("Leaving 'WorkerTask'")

	return returnMessage, nil
}

func (s *TestInstructionBackendServer) AreYouAlive(ctx context.Context, emptyParameter *backend_server_grpc_api.EmptyParameter) (*backend_server_grpc_api.AckNackResponse, error) {

	testInstructionBackendObject.logger.WithFields(logrus.Fields{}).Info("Incoming: 'AreYouAlive'")

	testInstructionBackendObject.logger.WithFields(logrus.Fields{}).Info("Leaving 'AreYouAlive'")
	return &backend_server_grpc_api.AckNackResponse{Acknack: true, Comments: "I'am alive, from " + testInstructionBackendObject.uuid}, nil
}
