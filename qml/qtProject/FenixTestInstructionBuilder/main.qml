import QtQuick 2.12
import QtQuick.Window 2.12
import QtQuick 2.0
import QtQuick.Controls 1.4
import QtQuick.Controls 2.3
import QtQuick.Dialogs 1.1
import "jsFunctions.js" as JServer


Window {
    visible: true
    width: 640
    height: 480
    title: qsTr("Hello World")
    id: rootWindow

    Component.onCompleted: {
        //popUpNoConnectionToBackendMain.open()
        JServer.jsCheckIfStartedByGolang()
    }

    MainTable {
        id: rootTable
        // anchors.fill parent
        property int rootWindowWidth: parent.width
        property int rootWindowHeight: parent.height
        property bool startedByGolang: false
        property bool checkHasBeenDone: false
        property double backEndCheckIfAliveStartTime: 0
        property int backEndChrootTableeckIfAlivesecondsElapsed: 0
        property bool backEndCheckIfAliveMyServerResponse: false
        property bool backendIsAlive: false

        Connections {
            target: QmlBridge
            onSendToQml: rootTable.startedByGolang = data
        }
        Component.onCompleted: {
            //popUpNoConnectionToBackendMain.open()
            //console.log("Qt.resolvedUrl('.')" + Qt.resolvedUrl("."))

        }

    }



    // *********************************************************************************
    // Check if backend is running
    // PopUp that locks application when there is no connection to backend
    Popup {
        id: popUpNoConnectionToBackendMain
        //x: 100
        //y: pluginsTableView.y + pluginsTableView.height + 50
        width: rootWindow.width - 10
        height: rootWindow.height - 10
        modal: true
        focus: true
        closePolicy: Popup.CloseOnEscape //| Popup.CloseOnPressOutsideParent
        opacity: 0.5
        anchors.centerIn: parent
        background: Rectangle {
            color: "blue"
            border.color: "red"
        }

        exit: Transition {
            NumberAnimation {
                property: "opacity"
                from: 1.0
                to: 0.0
                duration: 1200
            }
        }
        enter: Transition {

            NumberAnimation {
                property: "opacity"
                from: 0.0
                to: 1.0
                duration: 800
            }
        }

        Label {
            text: "No connection to backend"
        }

        Item {
            id: timerItemForBackendCheck
            width: parent.width
            height: parent.height


            property double startTime: 0
            property int secondsElapsed: 0
            property bool myServerResponse: false
            /*
            function restartCounter(doCheckOfBackend) {

                timerItemForBackendCheck.startTime = 0
            }

            function timeChanged() {
                if (timerItemForBackendCheck.startTime == 0) {
                    timerItemForBackendCheck.startTime = new Date().getTime(
                                ) //returns the number of milliseconds since the epoch (1970-01-01T00:00:00Z);
                }
                var currentTime = new Date().getTime()
                timerItemForBackendCheck.secondsElapsed = (currentTime - startTime) / 1000

                if ((timerItemForBackendCheck.secondsElapsed > 1) || (doCheckOfBackend == true)){
                    timerItemForBackendCheck.myServerResponse = QmlBridge.checkIfServerIsOnline()
                    console.log("QmlBridge.checkIfServerIsOnline(): " + timerItemForBackendCheck.myServerResponse)
                    restartCounter()

                    if (timerItemForBackendCheckmyServerResponse == true) {
                        popUpNoConnectionToBackendMain.close()
                        rootTable.backendIsAlive = true
                    } else {
                         popUpNoConnectionToBackendMain.open()
                         rootTable.backendIsAlive = false
                    }
                }
            }

            */

            Timer {
                id: elapsedTimer
                interval: 1000
                running: true
                repeat: true
                onTriggered: JServer.jsBackenTimerTimeChanged(false)
            }

            Text {
                id: counterText
                text: timerItemForBackendCheck.secondsElapsed
                color: 'steelblue';
                font.pixelSize: 30;
                anchors {
                    horizontalCenter: parent.horizontalCenter
                    verticalCenter: parent.verticalCenter;
                }
            }
        }
    }



}




