package main

import (
	"fmt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/quick"
	"jlambert/FenixInception3/FenixTestInstructionBuilder/qml_server"
	"os"
)

// Main function that starts everything
// Designed this way becasue of compilation process of "qtrecipe"
//author: https://github.com/longlongh4

/*
type QmlBridge struct {
	core.QObject

	_ func(data string)        `signal:"sendToQml"`
	_ func(data string) string `slot:"sendToGo"`
}

*/

type QmlBridge struct {
	core.QObject

	//	_ func(data bool)          `signal:"sendToQml"`
	//	_ func(data string) string `slot:"sendToGo"`
	_ func() bool   `slot:"checkIfServerIsOnline"`
	_ func() string `slot:"generateGuid"`
	_ func() string `slot:"loadPluginModelFromServer"`
	_ func() string `slot:"loadDomainModelFromServer"`
}

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	gui.NewQGuiApplication(len(os.Args), os.Args)

	var view = quick.NewQQuickView(nil)
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)

	var qmlBridge = NewQmlBridge(nil)
	qmlBridge.ConnectGenerateGuid(qml_server.GenerateGuid)
	qmlBridge.ConnectCheckIfServerIsOnline(qml_server.CheckIfServerIsOnline)
	qmlBridge.ConnectLoadPluginModelFromServer(loadPluginModelFromServer)
	qmlBridge.ConnectLoadDomainModelFromServer(loadDomainModelFromServer)

	/*
		qmlBridge.ConnectSendToGo(func(data string) string {
			fmt.Println("go:", data)
			return "hello from go"
		})

	*/
	view.RootContext().SetContextProperty("QmlBridge", qmlBridge)
	view.SetSource(core.NewQUrl3("qrc:/qml/qtProject/FenixTestInstructionBuilder/main.qml", 0))

	/*go func() {
		for t := range time.NewTicker(time.Second * 1).C {
			qmlBridge.SendToQml(t.Format(time.ANSIC))
		}
	}()
	*/
	fmt.Println("Starting QML-server")
	qml_server.Start_qml_server()

	view.Show()

	gui.QGuiApplication_Exec()
}

// *********************************************************************
// Used by QML to verify that the QML-code was started from server and not from QML-editor
//func checkIfServerIsOnline() bool {

//	return qml_server.CheckIfServerIsOnline()
//}

// *********************************************************************
// Forward a call from frontend to backend to generate a guid in string format
//func generateGuid() string {

//	return "qml_server.GenerateGuid()"
//}

// *********************************************************************
// Forward a call from frontend to backend to load stored data about Plugins
func loadPluginModelFromServer() string {

	// Send back response to frontend
	return qml_server.LoadPluginModelFromServer()
}

// *********************************************************************
// Used by QML to load stored data about Domains
func loadDomainModelFromServer() string {

	// Send back response to frontend
	return "qml_server.LoadDomainModelFromServer()"
}
