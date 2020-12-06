package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/quick"
	"jlambert/FenixInception3/FenixTestInstructionBuilder/qml_server"
	"os"
)

type QmlBridge struct {
	core.QObject

	//	_ func(data bool)          `signal:"sendToQml"`
	//	_ func(data string) string `slot:"sendToGo"`
	_ func() bool   `slot:"checkIfServerIsOnline"`
	_ func() string `slot:"generateGuid"`
	_ func() string `slot:"loadPluginModelFromServer"`
	_ func() string `slot:"loadDomainModelFromServer"`
}

type qml_struct struct {
}

var qmlLogger *logrus.Logger

// Main function that starts everything
// Designed this way becasue of compilation process of "qtrecipe"
func main() {

	// Init logger
	qmlLogger = InitLogger("")

	// Set QML parts
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	gui.NewQGuiApplication(len(os.Args), os.Args)

	var view = quick.NewQQuickView(nil)
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)

	var qmlBridge = NewQmlBridge(nil)
	qmlBridge.ConnectGenerateGuid(qml_server.GenerateGuid)
	qmlBridge.ConnectCheckIfServerIsOnline(checkIfServerIsOnline)
	qmlBridge.ConnectLoadPluginModelFromServer(qml_server.LoadPluginModelFromServer)
	qmlBridge.ConnectLoadDomainModelFromServer(qml_server.LoadDomainModelFromServer)

	view.RootContext().SetContextProperty("QmlBridge", qmlBridge)
	view.SetSource(core.NewQUrl3("qrc:/qml/qtProject/FenixTestInstructionBuilder/main.qml", 0))

	// Cleanup when closing
	defer qml_server.Cleanup()

	fmt.Println("Starting QML-server-part")
	qml_server.Start_qml_server(qmlLogger)

	view.Show()

	gui.QGuiApplication_Exec()
}

// *********************************************************************
// Used by QML to verify that the QML-code was started from server and not from QML-editor
func checkIfServerIsOnline() bool {

	return qml_server.CheckIfServerIsOnline(qmlLogger)
}

// *********************************************************************
// Forward a call from frontend to backend to generate a guid in string format
//func generateGuid() string {

//	return "qml_server.GenerateGuid()"
//}

// *********************************************************************
// Forward a call from frontend to backend to load stored data about Plugins
//func loadPluginModelFromServer() string {

// Send back response to frontend
//	return qml_server.LoadPluginModelFromServer()
//}

// *********************************************************************
// Used by QML to load stored data about Domains
//func loadDomainModelFromServer() string {

// Send back response to frontend
//	return qml_server.LoadDomainModelFromServer()
//}
