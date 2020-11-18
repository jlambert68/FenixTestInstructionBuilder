package qml_server

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"jlambert/FenixInception3/FenixTestInstructionBuilder/grpc_api/backend_server_grpc_api"
)

// *********************************************************************
// Used by QML to verify that the QML-code was started from server and not from QML-editor
func checkIfServerIsOnline() bool {

	return true
}

// *********************************************************************
// Forward a call from frontend to backend to generate a guid in string format
func generateGuid() string {

	// Variable to be sent back to frontend
	var returnGuid string
	returnGuid = "NOT SET"

	// Message received from backend server
	var backendMessage *backend_server_grpc_api.GuidResponse
	var err error

	// Call backend
	ctx := context.Background()
	backendMessage, err = testInstructionBackendServerGrpcClient.GenerateGuid(ctx, &backend_server_grpc_api.EmptyParameter{})

	if err != nil {
		// Something went wrong in backend
		returnGuid = "EROOR in Backend when trying to generate a new GUID"
		qmlServerObject.logger.WithFields(logrus.Fields{
			"Id":             "b847e2e5-4665-4dd6-8492-200b1a2d48a3",
			"err":            err,
			"backendMessage": backendMessage,
		}).Error("Backend couldn't generate GUID for frontend.")

		returnGuid = "EROOR in Backend"

	} else {
		if backendMessage.Acknack == false {
			// Backend couldn't generate GUID
			qmlServerObject.logger.WithFields(logrus.Fields{
				"Id":             "5d075469-d78e-436a-a798-51374891e183",
				"err":            err,
				"backendMessage": backendMessage,
			}).Error("Backend couldn't generate GUID for frontend.")

			returnGuid = "EROOR in Backend"

		} else {
			// OK, forwarding GUID to front end
			qmlServerObject.logger.WithFields(logrus.Fields{
				"Id":             "89b79ad2-c559-49be-84e6-02eb61ba9993",
				"err":            err,
				"backendMessage": backendMessage,
			}).Error("Forwarding new GUID from backend to frontend.")

			returnGuid = backendMessage.Guid

		}
	}

	// Send back response to frontend
	return returnGuid
}

// *********************************************************************
// Forward a call from frontend to backend to load stored data about Plugins
func loadPluginModelFromServer() string {

	// Variable to be sent back to frontend
	var jsonToSendAsString string
	jsonToSendAsString = "NOT SET"

	// Message received from backend server
	var backendMessage *backend_server_grpc_api.PluginQmlModelFromServerResponse
	var err error

	// Call backend
	ctx := context.Background()
	backendMessage, err = testInstructionBackendServerGrpcClient.LoadPluginModelFromServer(ctx, &backend_server_grpc_api.EmptyParameter{})

	if err != nil {
		// Something went wrong in backend
		jsonToSendAsString = "EROOR in Backend when trying to generate list of Plugins"
		qmlServerObject.logger.WithFields(logrus.Fields{
			"Id":             "b4e2b4d6-2471-4a54-8cb9-3523ba34251e",
			"err":            err,
			"backendMessage": backendMessage,
		}).Error("Backend couldn't generate list of Plugins for frontend.")

		jsonToSendAsString = "EROOR in Backend"

	} else {
		if backendMessage.Acknack == false {
			// Backend couldn't generate GUID
			qmlServerObject.logger.WithFields(logrus.Fields{
				"Id":             "36a17a2c-9c1a-41aa-9405-c050436f8b6b",
				"err":            err,
				"backendMessage": backendMessage,
			}).Error("Backend couldn't generate list of Plugins for frontend.")

			jsonToSendAsString = "EROOR in Backend"

		} else {
			// OK, forwarding GUID to front end
			qmlServerObject.logger.WithFields(logrus.Fields{
				"Id":             "c462b8f4-8291-4d01-8ec3-e913acd25ebe",
				"err":            err,
				"backendMessage": backendMessage,
			}).Error("Forwarding listy of Plugins from backend to frontend.")

			jsonToSendAsString = backendMessage.JsonStringForPluginQmlModel

		}
	}

	// Send back response to frontend
	return jsonToSendAsString
}

// *********************************************************************
// Used by QML to load stored data about Domainss
func loadDomainModelFromServer() string {

	// Variable to be sent back to frontend
	var jsonToSendAsString string
	jsonToSendAsString = "NOT SET"

	// Message received from backend server
	var backendMessage *backend_server_grpc_api.PluginQmlModelFromServerResponse
	var err error

	// Call backend
	ctx := context.Background()
	backendMessage, err = testInstructionBackendServerGrpcClient.LoadPluginModelFromServer(ctx, &backend_server_grpc_api.EmptyParameter{})

	if err != nil {
		// Something went wrong in backend
		jsonToSendAsString = "EROOR in Backend when trying to generate list of Domains"
		qmlServerObject.logger.WithFields(logrus.Fields{
			"Id":             "e461c7ea2-d704-4b6e-a5a4-7f445c65e6fa",
			"err":            err,
			"backendMessage": backendMessage,
		}).Error("Backend couldn't generate list of Domains for frontend.")

		jsonToSendAsString = "EROOR in Backend"

	} else {
		if backendMessage.Acknack == false {
			// Backend couldn't generate GUID
			qmlServerObject.logger.WithFields(logrus.Fields{
				"Id":             "c3544c25-7a92-413c-99ac-895da968cf6e",
				"err":            err,
				"backendMessage": backendMessage,
			}).Error("Backend couldn't generate list of Domains for frontend.")

			jsonToSendAsString = "EROOR in Backend"

		} else {
			// OK, forwarding GUID to front end
			qmlServerObject.logger.WithFields(logrus.Fields{
				"Id":             "d76404d2-8497-49ea-ab2f-8295e597f5d5",
				"err":            err,
				"backendMessage": backendMessage,
			}).Info("Forwarding list of Domains from backend to frontend.")

			jsonToSendAsString = backendMessage.JsonStringForPluginQmlModel

		}
	}

	// Send back response to frontend
	return jsonToSendAsString
}
