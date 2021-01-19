package qml_server

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"jlambert/FenixInception3/FenixTestInstructionBuilder/grpc_api/backend_server_grpc_api"
	"strconv"
)

var onlyForTest_SwitchOfBackend string = "Off"

// *********************************************************************
// Used by QML to verify that the QML-code was started from server and not from QML-editor
func CheckIfServerIsOnline(qmlLogger *logrus.Logger) bool {

	var returnMessage bool

	qmlLogger.WithFields(logrus.Fields{
		"ID": "a2bd8237-97f9-4c34-b057-1c64bf61be4c",
	}).Debug("Incoming call from QML-GUI to 'CheckIfServerIsOnline()'")

	if qml_qGRPC_bridges_CanStart == true {

		switch onlyForTest_SwitchOfBackend {
		case "On":
			returnMessage = true
		case "Off":
			returnMessage = false
		default:
			returnMessage = false
		}

	} else {
		returnMessage = false
	}

	qmlLogger.WithFields(logrus.Fields{
		"ID": "5a5d8452-a927-490e-8764-dd20fbd57605",
	}).Debug("Will answer: " + strconv.FormatBool(returnMessage))

	return returnMessage
}

// *********************************************************************
// Forward a call from frontend to backend to generate a guid in string format
func GenerateGuid() string {

	QmlServerObject.logger.WithFields(logrus.Fields{
		"ID": "ec435580-98a3-42be-865b-bc59234648f4",
	}).Debug("Incoming call from QML-GUI to 'GenerateGuid()'")

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
		QmlServerObject.logger.WithFields(logrus.Fields{
			"Id":             "b847e2e5-4665-4dd6-8492-200b1a2d48a3",
			"err":            err,
			"backendMessage": backendMessage,
		}).Error("Backend couldn't generate GUID for frontend.")

		returnGuid = "EROOR in Backend"

	} else {
		if backendMessage.Acknack == false {
			// Backend couldn't generate GUID
			QmlServerObject.logger.WithFields(logrus.Fields{
				"Id":             "5d075469-d78e-436a-a798-51374891e183",
				"err":            err,
				"backendMessage": backendMessage,
			}).Error("Backend couldn't generate GUID for frontend.")

			returnGuid = "EROOR in Backend"

		} else {
			// OK, forwarding GUID to front end
			QmlServerObject.logger.WithFields(logrus.Fields{
				"Id":             "89b79ad2-c559-49be-84e6-02eb61ba9993",
				"err":            err,
				"backendMessage": backendMessage,
			}).Error("Forwarding new GUID from backend to frontend.")

			returnGuid = backendMessage.Guid

		}
	}

	// Send back response to frontend
	QmlServerObject.logger.WithFields(logrus.Fields{
		"ID":         "25f1d06c-7395-461f-b3ff-e554502f6471",
		"returnGuid": returnGuid,
	}).Debug("Leaving 'GenerateGuid()'")

	return returnGuid
}

// *********************************************************************
// Initiate channels used for decoupling QML and gRPC

// LoadPluginModel
var LoadPluginModelFromServerChannel_FromGrpc chan string
var LoadPluginModelFromServerChannel_ToGrpc chan bool

// LoadDomainMode
var LoadDomainModelFromServerChannel_FromGrpc chan string
var LoadDomainModelFromServerChannel_ToGrpc chan bool

// Trigger for when QML-gRPC-bridge is ready to use
var qml_qGRPC_bridges_CanStart = false

func initiateChannels() {

	// Define channels  used for communication between QML och gRPC
	QmlServerObject.logger.WithFields(logrus.Fields{
		"ID": "a2dfc565-e025-4cb2-a2b5-ab6739378c5a",
	}).Debug("Initiate channels used for communication between QML och gRPC")

	LoadPluginModelFromServerChannel_FromGrpc = make(chan string, 1)
	LoadPluginModelFromServerChannel_ToGrpc = make(chan bool, 1)

	LoadDomainModelFromServerChannel_FromGrpc = make(chan string, 1)
	LoadDomainModelFromServerChannel_ToGrpc = make(chan bool, 1)

	// Start up gRPC services as goroutines
	QmlServerObject.logger.WithFields(logrus.Fields{
		"ID": "3e74d4b7-585d-421b-8c51-38cf809d17a5",
	}).Debug("Start functions, as goroutines, used by QML that calls backend via gRPC ")

	// Start up up gRPC-parts of QML-gRPC-bridge as goroutines
	go LoadPluginModelFromServer_gRPC()
	go LoadDomainModelFromServer_gRPC()

}

// *********************************************************************
// Turn on/off the function to talk to gRPC via the QML-gRPC bridge
func StartStopQmlGrpcBridgr(canBeRun bool) {
	qml_qGRPC_bridges_CanStart = canBeRun
}

// *********************************************************************
// Forward a call from frontend to backend to load stored data about Plugins
// Using channels for async connection

func LoadPluginModelFromServer() string {

	if qml_qGRPC_bridges_CanStart == true {
		var returnMessage string

		QmlServerObject.logger.WithFields(logrus.Fields{
			"ID": "b35ab173-c624-401e-abd5-cda8f54eabe8",
		}).Debug("Call from QML to 'LoadPluginModelFromServer()'")

		// Trigger Call to gROC
		LoadPluginModelFromServerChannel_ToGrpc <- true

		// Wait for return message from gRPC
		returnMessage = <-LoadPluginModelFromServerChannel_FromGrpc

		QmlServerObject.logger.WithFields(logrus.Fields{
			"ID":            "8cde35f4-209a-43d4-a977-d2d0520d445b",
			"returnMessage": returnMessage,
		}).Debug("Message got back, via channel, from LoadPluginModelFromServer_gRPC-call")

		return returnMessage
	} else {
		return "Backend not ready"
	}

}

// *********************************************************************
// Forward a call from frontend to backend to load stored data about Plugins
// Using gRPC to backend

func LoadPluginModelFromServer_gRPC() {

	// Variable to be sent back to frontend
	var jsonToSendAsString string
	jsonToSendAsString = "NOT SET"

	for {
		// Wait for trigger message to start processing
		<-LoadPluginModelFromServerChannel_ToGrpc
		QmlServerObject.logger.WithFields(logrus.Fields{
			"ID": "1936d25f-7edc-474d-8203-cb6ed094dbc3",
		}).Debug("Message from qml, via channel, to do 'LoadPluginModelFromServer_gRPC'-call")

		// Message received from backend server
		var backendMessage *backend_server_grpc_api.PluginQmlModelFromServerResponse
		var err error

		// Call backend
		ctx := context.Background()
		backendMessage, err = testInstructionBackendServerGrpcClient.LoadPluginModelFromServer(ctx, &backend_server_grpc_api.EmptyParameter{})

		if err != nil {
			// Something went wrong in backend
			jsonToSendAsString = "EROOR in Backend when trying to generate list of Plugins"
			QmlServerObject.logger.WithFields(logrus.Fields{
				"Id":             "b4e2b4d6-2471-4a54-8cb9-3523ba34251e",
				"err":            err,
				"backendMessage": backendMessage,
			}).Error("Backend couldn't generate list of Plugins for frontend.")

			jsonToSendAsString = "ERROR in Backend"

		} else {
			if backendMessage.Acknack == false {
				// Backend couldn't generate GUID
				QmlServerObject.logger.WithFields(logrus.Fields{
					"Id":             "36a17a2c-9c1a-41aa-9405-c050436f8b6b",
					"err":            err,
					"backendMessage": backendMessage,
				}).Error("Backend couldn't generate list of Plugins for frontend.")

				jsonToSendAsString = "ERROR in Backend"

			} else {
				// OK, forwarding GUID to front end
				QmlServerObject.logger.WithFields(logrus.Fields{
					"Id":             "c462b8f4-8291-4d01-8ec3-e913acd25ebe",
					"err":            err,
					"backendMessage": backendMessage,
				}).Error("Forwarding listy of Plugins from backend to frontend.")

				jsonToSendAsString = backendMessage.JsonStringForPluginQmlModel

			}
		}

		// Send back response to frontend
		LoadPluginModelFromServerChannel_FromGrpc <- jsonToSendAsString

	}

}

// *********************************************************************
// Forward a call from frontend to backend to load stored data about Domains
// Using channels for async connection

func LoadDomainModelFromServer() string {

	if qml_qGRPC_bridges_CanStart == true {
		var returnMessage string

		QmlServerObject.logger.WithFields(logrus.Fields{
			"ID": "2181a9c3-dfd4-4207-ba1c-8cf56884e7cd",
		}).Debug("Call from QML to 'LoadDomainModelFromServer()'")

		// Trigger Call to gROC
		LoadDomainModelFromServerChannel_ToGrpc <- true

		// Wait for return message from gRPC
		returnMessage = <-LoadDomainModelFromServerChannel_FromGrpc

		QmlServerObject.logger.WithFields(logrus.Fields{
			"ID":            "63abc3d4-55ce-4ebb-8426-8f72ab0cd2d5",
			"returnMessage": returnMessage,
		}).Debug("Message got back, via channel, from LoadDomainModelFromServer_gRPC-call")

		return returnMessage
	} else {
		return "Backend not ready"
	}

}

// *********************************************************************
// Forward a call from frontend to backend to load stored data about Domains
// Using gRPC to backend
func LoadDomainModelFromServer_gRPC() string {

	// Variable to be sent back to frontend
	var jsonToSendAsString string
	jsonToSendAsString = "NOT SET"

	for {
		// Wait for trigger message to start processing
		<-LoadDomainModelFromServerChannel_ToGrpc
		QmlServerObject.logger.WithFields(logrus.Fields{
			"ID": "1936d25f-7edc-474d-8203-cb6ed094dbc3",
		}).Debug("Message from qml, via channel, to do 'LoadPluginModelFromServer_gRPC'-call")

		// Message received from backend server
		var backendMessage *backend_server_grpc_api.PluginQmlModelFromServerResponse
		var err error

		// Call backend
		ctx := context.Background()
		backendMessage, err = testInstructionBackendServerGrpcClient.LoadPluginModelFromServer(ctx, &backend_server_grpc_api.EmptyParameter{})

		if err != nil {
			// Something went wrong in backend
			jsonToSendAsString = "EROOR in Backend when trying to generate list of Domains"
			QmlServerObject.logger.WithFields(logrus.Fields{
				"Id":             "e461c7ea2-d704-4b6e-a5a4-7f445c65e6fa",
				"err":            err,
				"backendMessage": backendMessage,
			}).Error("Backend couldn't generate list of Domains for frontend.")

			jsonToSendAsString = "EROOR in Backend"

		} else {
			if backendMessage.Acknack == false {
				// Backend couldn't generate GUID
				QmlServerObject.logger.WithFields(logrus.Fields{
					"Id":             "c3544c25-7a92-413c-99ac-895da968cf6e",
					"err":            err,
					"backendMessage": backendMessage,
				}).Error("Backend couldn't generate list of Domains for frontend.")

				jsonToSendAsString = "EROOR in Backend"

			} else {
				// OK, forwarding GUID to front end
				QmlServerObject.logger.WithFields(logrus.Fields{
					"Id":             "d76404d2-8497-49ea-ab2f-8295e597f5d5",
					"err":            err,
					"backendMessage": backendMessage,
				}).Info("Forwarding list of Domains from backend to frontend.")

				jsonToSendAsString = backendMessage.JsonStringForPluginQmlModel

			}
		}

		// Send back response to frontend
		LoadDomainModelFromServerChannel_FromGrpc <- jsonToSendAsString

	}

}
