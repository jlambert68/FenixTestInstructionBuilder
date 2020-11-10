package main

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"jlambert/FenixInception3/FenixTestInstructionBuilder/common_config"
	"jlambert/FenixInception3/FenixTestInstructionBuilder/grpc_api/mother_server_grpc_api"
	"net"
	"os"
	"path/filepath"
	"github.com/google/uuid"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/widgets"
)

xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxSKRIVA KOMMENTARER VID VARJE DEL

type QmlBridge struct {
	core.QObject

	_ func(data bool)          `signal:"sendToQml"`
	_ func(data string) string `slot:"sendToGo"`
	_ func() bool              `slot:"checkIfServerIsOnline"`
	_ func() string            `slot:"generateGuid"`
	_ func() string            `slot:"loadPluginModelFromServer"`
	_ func() string            `slot:"loadDomainModelFromServer"`
}


// *********************************************************************
func checkIfServerIsOnline() bool {

	return true
}

// *********************************************************************
func generateGuid() string {
	var returnGuid string
	newGuid, _ := uuid.NewUUID()

	returnGuid = newGuid.String()

	return returnGuid
}

// *********************************************************************
func loadPluginModelFromServer() string {

	type ListElementStruct struct {
		Name        string `json:"name"`
		GUID        string `json:"guid"`
		ReadyForUse bool   `json:"readyForUse"`
		Activated   bool   `json:"activated"`
		Description string `json:"description"`
	}

	type dataForPlugInStruct struct {
		ListElement []ListElementStruct `json:"ListElement"`
	}

	var dataForPlugin dataForPlugInStruct
	var listElementData []ListElementStruct

	// Append First Mock-ELement
	listElement1 := &ListElementStruct{
		Name:        "Send UTR",
		GUID:        "2243085a-feee-4ae7-8ccf-03f69c0704a4",
		ReadyForUse: true,
		Activated:   true,
		Description: "Sends UTRs via MQ towards Custody Cash",
	}
	listElementData = append(listElementData, *listElement1)

	// Appned Second Mock-ELement
	listElement2 := &ListElementStruct{
		Name:        "Validate New Pacs008",
		GUID:        "d456499c-2ad1-4677-8e1d-909a7ecab560",
		ReadyForUse: true,
		Activated:   true,
		Description: "Validates that a newly created Pacs008 has been sent to CMaaS",
	}
	listElementData = append(listElementData, *listElement2)

	dataForPlugin.ListElement = listElementData

	jsonToSend, _ := json.MarshalIndent(dataForPlugin, "", " ")
	jsonToSendAsString := string(jsonToSend)
	//fmt.Println(jsonToSendAsString)

	return jsonToSendAsString
}

// *********************************************************************
func loadDomainModelFromServer() string {

	type ListElementStruct struct {
		Name        string `json:"name"`
		GUID        string `json:"guid"`
		ReadyForUse bool   `json:"readyForUse"`
		Activated   bool   `json:"activated"`
		Description string `json:"description"`
	}

	type dataForDomainStruct struct {
		ListElement []ListElementStruct `json:"ListElement"`
	}

	var dataForDomain dataForDomainStruct
	var listElementData []ListElementStruct

	// Append First Mock-ELement
	listElement1 := &ListElementStruct{
		Name:        "Custody Cash",
		GUID:        "2243085a-feee-4ae7-8ccf-03f69c0704a4",
		ReadyForUse: true,
		Activated:   true,
		Description: "All test regarding Custody Cash",
	}
	listElementData = append(listElementData, *listElement1)

	// Appned Second Mock-ELement
	listElement2 := &ListElementStruct{
		Name:        "Cusotdy Arrangement",
		GUID:        "d456499c-2ad1-4677-8e1d-909a7ecab560",
		ReadyForUse: false,
		Activated:   false,
		Description: "All tests regarding Custody Arrangement",
	}
	listElementData = append(listElementData, *listElement2)

	dataForDomain.ListElement = listElementData

	jsonToSend, _ := json.MarshalIndent(dataForDomain, "", " ")
	jsonToSendAsString := string(jsonToSend)
	//fmt.Println(jsonToSendAsString)

	return jsonToSendAsString
}

// *********************************************************************



func initQQuickView(path string) *quick.QQuickView {

	var view = quick.NewQQuickView(nil)

	println(filepath.Dir(path))

	var watcher = core.NewQFileSystemWatcher2([]string{filepath.Dir(path)}, nil)

	var reload = func(p string) {
		println("changed:", p)
		view.SetSource(core.NewQUrl())
		view.Engine().ClearComponentCache()
		view.SetSource(core.NewQUrl3(path, 0))
	}

	//watcher.ConnectFileChanged(reload)
	watcher.ConnectDirectoryChanged(reload)

	return view
}

// Used for only process cleanup once
var cleanupProcessed bool = false

func cleanup() {

	if cleanupProcessed == false {

		cleanupProcessed = true

		// Cleanup before close down application
		motherObject.logger.WithFields(logrus.Fields{
			"ID": "ab1abcdb-d786-4bcc-b274-b52bd931f43d",
		}).Info("Clean up and shut down servers")

		motherObject.logger.WithFields(logrus.Fields{
			"ID": "a22e8de2-c38c-4483-95e0-e9dc21dd7ca8",
		}).Info("Gracefull stop for: registerTaxiHardwareStreamServer")

		registerMotherServer.GracefulStop()

		motherObject.logger.WithFields(logrus.Fields{
			"common_config.MotherServer_port: ": common_config.MotherServer_port,
			"ID":                                "19593393-3ac9-45ac-96e9-cdf911b167c7",
		}).Info("Close net.Listing")
		lis.Close()

		//log.Println("Close DB_session: %v", DB_session)
		//DB_session.Close()
	}
}

func main() {

	// ******************************************************
	var err error
	defer cleanup()

	// Set up WorkerObject
	motherObject = &MotherObject_struct{}

	// Init logger
	motherObject.InitLogger("")

	// Find first non allocated port from defined start port
	motherObject.logger.WithFields(logrus.Fields{
		"ID": "56b6419f-d714-4ab0-be62-f3c7f08b9558",
	}).Info("Mother Server tries to start")

	motherObject.logger.WithFields(logrus.Fields{
		"common_config.MotherServer_port): ": common_config.MotherServer_port,
		"ID":                                 "8f904ace-9d24-452b-891a-5b8e5c247ba2",
	}).Info("Start listening on:")
	lis, err = net.Listen("tcp", string(common_config.MotherServer_port))

	if err != nil {
		motherObject.logger.WithFields(logrus.Fields{
			"err: ": err,
			"ID":    "0fbf0f08-6114-4cd8-9992-6a387955fb5f",
		}).Fatal("failed to listen:")

	} else {
		motherObject.logger.WithFields(logrus.Fields{
			"common_config.MotherServer_port): ": common_config.MotherServer_port,
			"ID":                                 "93496c07-2b6c-4edc-a1f9-3fd555fa1201",
		}).Info("Success in listening on port:")

	}

	// Creates a new RegisterMotherServer gRPC server
	go func() {
		motherObject.logger.WithFields(logrus.Fields{
			"ID": "ebc26735-9d13-4b13-91b8-20999cd5e254",
		}).Info("Starting Mother Server")

		registerMotherServer = grpc.NewServer()
		mother_server_grpc_api.RegisterMotherServerServer(registerMotherServer, &MotherServer{})

		motherObject.logger.WithFields(logrus.Fields{
			"common_config.MotherServer_port): ": common_config.MotherServer_port,
			"ID":                                 "cfed87c4-55aa-4cd1-980a-a15eb75ab6fb",
		}).Info("registerMotherServer for Mother Server started")

		registerMotherServer.Serve(lis)
	}()
	// ******************************************************

	var path = filepath.Join(os.Getenv("GOPATH"), "src", "jlambert", "FenixInception3", "FenixTestInstructionBuilder", "qml", "qtProject", "FenixTestInstructionBuilder", "main.qml")
	println(path)
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	app := widgets.NewQApplication(len(os.Args), os.Args)

	//view := quick.NewQQuickView(nil)
	var view = initQQuickView(path)
	var qmlBridge = NewQmlBridge(nil)
	qmlBridge.ConnectSendToGo(func(data string) string {
		fmt.Println("go:", data)
		return "hello from go"
	})
	qmlBridge.ConnectGenerateGuid(generateGuid)
	qmlBridge.ConnectCheckIfServerIsOnline(checkIfServerIsOnline)
	qmlBridge.ConnectLoadPluginModelFromServer(loadPluginModelFromServer)
	qmlBridge.ConnectLoadDomainModelFromServer(loadDomainModelFromServer)

	view.RootContext().SetContextProperty("QmlBridge", qmlBridge)
	view.SetSource(core.NewQUrl3(path, 0))

	view.SetTitle("goroutine Example")
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	//view.SetSource(core.NewQUrl3("qrc:/qml/main.qml", 0))
	view.Show()

	// Tell QML that it should use Go-server for instead of local mocks for certain functions
	qmlBridge.SendToQml(true)

	app.Exec()
}
