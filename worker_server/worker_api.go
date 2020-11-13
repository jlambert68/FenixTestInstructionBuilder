package worker_server

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"jlambert/FenixInception3/FenixTestInstructionBuilder/grpc_api/worker_server_grpc_api"
)

func (s *WorkerServer) WorkerTask(ctx context.Context, messageToWorker *worker_server_grpc_api.MessageToWorkerStruct) (*worker_server_grpc_api.AckNackResponse, error) {

	var returnMessage *worker_server_grpc_api.AckNackResponse

	workerObject.logger.WithFields(logrus.Fields{
		"messageToWorker: ": messageToWorker,
	}).Info("Incoming: 'WorkerTask'")

	if workerObject.iAmBusy == true {
		workerObject.logger.WithFields(logrus.Fields{
			"workerObject.iAmBusy: ": workerObject.iAmBusy,
		}).Info("Worker allready busy do calculations, can't new work load")
		returnMessage = &worker_server_grpc_api.AckNackResponse{
			Acknack:  false,
			Comments: "Worker allready busy do calculations, can't new work load",
		}
	} else {
		workerObject.logger.WithFields(logrus.Fields{
			"workerObject.iAmBusy: ": workerObject.iAmBusy,
		}).Info("Take new work load")
		workerObject.iAmBusy = true
		workerObject.currentTaskuuid = messageToWorker.TaskUuid
		workerObject.currentTaskName = messageToWorker.TaskName
		returnMessage = &worker_server_grpc_api.AckNackResponse{Acknack: true, Comments: "Accepted Task"}
		//go workerObject.ExecuteTaskFromMother(messageToWorker)
	}

	workerObject.logger.WithFields(logrus.Fields{}).Info("Leaving 'WorkerTask'")

	return returnMessage, nil
}

func (s *WorkerServer) AreYouAlive(ctx context.Context, emptyParameter *worker_server_grpc_api.EmptyParameter) (*worker_server_grpc_api.AckNackResponse, error) {

	workerObject.logger.WithFields(logrus.Fields{}).Info("Incoming: 'AreYouAlive'")

	workerObject.logger.WithFields(logrus.Fields{}).Info("Leaving 'AreYouAlive'")
	return &worker_server_grpc_api.AckNackResponse{Acknack: true, Comments: "I'am alive, from " + workerObject.uuid}, nil
}
