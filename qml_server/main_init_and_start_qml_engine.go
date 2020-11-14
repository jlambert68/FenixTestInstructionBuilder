package qml_server

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/widgets"
	//"jlambert/FenixInception3/FenixTestInstructionBuilder"
	"os"
	"path/filepath"
)

// *********************************************************************
// Struct that handles communication to and from QML
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
// Initiate view for QML and start a directory change watcher
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

// *********************************************************************
// Initiate and start QML-wngine
func initiateAndStartQmlEngine() {

	// ******************************************************
	// Set up QML connection
	var path = filepath.Join(os.Getenv("GOPATH"), "src", "jlambert", "FenixInception3", "FenixTestInstructionBuilder", "qml", "qtProject", "FenixTestInstructionBuilder", "main.qml")
	println(path)
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	app := widgets.NewQApplication(len(os.Args), os.Args)

	//view := quick.NewQQuickView(nil)
	var view = initQQuickView(path)
	var qmlBridge = NewQmlBridge(nil)
	/*CLEANUP
	qmlBridge.ConnectSendToGo(func(data string) string {
		fmt.Println("go:", data)
		return "hello from go"
	})
	CLEANUP*/

	// Connect services exposed to QML
	qmlBridge.ConnectGenerateGuid(generateGuid)
	qmlBridge.ConnectCheckIfServerIsOnline(checkIfServerIsOnline)
	qmlBridge.ConnectLoadPluginModelFromServer(loadPluginModelFromServer)
	qmlBridge.ConnectLoadDomainModelFromServer(loadDomainModelFromServer)

	view.RootContext().SetContextProperty("QmlBridge", qmlBridge)
	view.SetSource(core.NewQUrl3(path, 0))

	view.SetTitle("Test Instruction Builder")
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	//view.SetSource(core.NewQUrl3("qrc:/qml/main.qml", 0))
	view.Show()

	// Tell QML that it should use Go-server for instead of local mocks for certain functions
	qmlBridge.SendToQml(true)

	app.Exec()
}
