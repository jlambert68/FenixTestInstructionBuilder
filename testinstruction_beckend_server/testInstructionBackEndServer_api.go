package testinstruction_beckend_server

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"jlambert/FenixInception3/FenixTestInstructionBuilder/grpc_api/backend_server_grpc_api"
)

func (s *TestInstructionBackendServer) AreYouAlive(ctx context.Context, emptyParameter *backend_server_grpc_api.EmptyParameter) (*backend_server_grpc_api.AckNackResponse, error) {

	testInstructionBackendObject.logger.WithFields(logrus.Fields{}).Info("Incoming: 'AreYouAlive'")

	testInstructionBackendObject.logger.WithFields(logrus.Fields{}).Info("Leaving 'AreYouAlive'")
	return &backend_server_grpc_api.AckNackResponse{Acknack: true, Comments: "I'am alive, from " + testInstructionBackendObject.uuid}, nil
}
