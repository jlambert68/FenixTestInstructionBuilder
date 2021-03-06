package qml_server

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"jlambert/FenixInception3/FenixTestInstructionBuilder/grpc_api/qml_server_grpc_api"
	//"jlambert/AllCombinations/grpc_api/qml_server_grpc_api"
)

func (s *QMLgRrpServer_struct) TestInstructionBackendServerIPandPort(ctx context.Context, workerInformation *qml_server_grpc_api.BackendServerInformation) (*qml_server_grpc_api.AckNackResponse, error) {

	var returnMessage *qml_server_grpc_api.AckNackResponse

	QmlServerObject.backendGrpcServerIsAlive = true

	QmlServerObject.logger.WithFields(logrus.Fields{
		"workerInformation: ": workerInformation,
	}).Info("Incoming: 'WorkerIPandPortIs'")

	// Save Worker Id information
	//QmlServerObject.workers = append(QmlServerObject.workers, workerList_struct{false, *workerInformation})

	//Create return message to worker
	//returnMessage = &qml_server_grpc_api.AckNackResponse{true, "Mother recorded worker Id information"}

	returnMessage = &qml_server_grpc_api.AckNackResponse{
		Acknack:              true,
		Comments:             "QML server recorded worker Id information",
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}
	QmlServerObject.logger.WithFields(logrus.Fields{}).Info("Leaving 'WorkerIPandPortIs'")

	return returnMessage, nil
}

func (s *QMLgRrpServer_struct) OnlyForTestSwitchOfBacken(ctx context.Context, switchOfBackendRequest *qml_server_grpc_api.SwitchOfBackendRequest) (*qml_server_grpc_api.AckNackResponse, error) {

	var returnMessage *qml_server_grpc_api.AckNackResponse

	onlyForTest_SwitchOfBackend = switchOfBackendRequest.BackendStatus

	QmlServerObject.logger.WithFields(logrus.Fields{
		"workerInformation: ": switchOfBackendRequest,
	}).Debug("Incoming: 'OnlyForTestSwitchOfBacken'")

	// Save Worker Id information
	//QmlServerObject.workers = append(QmlServerObject.workers, workerList_struct{false, *workerInformation})

	//Create return message to worker
	//returnMessage = &qml_server_grpc_api.AckNackResponse{true, "Mother recorded worker Id information"}

	returnMessage = &qml_server_grpc_api.AckNackResponse{
		Acknack:  true,
		Comments: "Change response",
	}
	QmlServerObject.logger.WithFields(logrus.Fields{}).Debug("Leaving 'OnlyForTestSwitchOfBacken'")

	return returnMessage, nil
}
