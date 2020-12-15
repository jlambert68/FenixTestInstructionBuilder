import QtQuick 2.12
import QtQuick.Window 2.12
import QtQuick 2.0
import QtQuick.Controls 1.4
import "jsFunctions.js" as JServer
import QtQuick.Controls 1.4

Window {
    visible: true
    width: 640
    height: 480
    title: qsTr("Hello World")
    id: rootWindow

    MainTable {
        id: rootTable
        anchors.fill: parent
        property int rootWindowWidth: parent.width
        property int rootWindowHeight: parent.height
        property bool startedByGolang: false
        property bool checkHasBeenDone: false

        Connections {
            target: QmlBridge
            onSendToQml: rootTable.startedByGolang = data
        }
//        Component.onCompleted: {
//        }


    }
}


/*

    TableStuff {
        id: tabs
        anchors.fill: parent
        //width: 640; height: 480

        Rectangle {
            property string title: "Red"
            anchors.fill: parent
            color: "#e3e3e3"

            Rectangle {
                anchors.fill: parent; anchors.margins: 20
                color: "#ff7f7f"
                Text {
                    width: parent.width - 20
                    anchors.centerIn: parent; horizontalAlignment: Qt.AlignHCenter
                    text: "Roses are red"
                    font.pixelSize: 20
                    wrapMode: Text.WordWrap
                }
            }
        }

        Rectangle {
            property string title: "Green"
            anchors.fill: parent
            color: "#e3e3e3"

            Rectangle {
                anchors.fill: parent; anchors.margins: 20
                color: "#7fff7f"
                Text {
                    width: parent.width - 20
                    anchors.centerIn: parent; horizontalAlignment: Qt.AlignHCenter
                    text: "Flower stems are green"
                    font.pixelSize: 20
                    wrapMode: Text.WordWrap
                }
            }
        }

        Rectangle {
            property string title: "Blue"
            anchors.fill: parent; color: "#e3e3e3"

            Rectangle {
                anchors.fill: parent; anchors.margins: 20
                color: "#7f7fff"
                Text {
                    width: parent.width - 20
                    anchors.centerIn: parent; horizontalAlignment: Qt.AlignHCenter
                    text: "Violets are blue"
                    font.pixelSize: 20
                    wrapMode: Text.WordWrap
                }
            }
        }
    }
    */

