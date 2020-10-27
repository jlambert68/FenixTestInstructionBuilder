import QtQuick 2.0

Item {



    TableStuff {
        id: tabs
        //anchors.fill: rootWindow //parent

        width: rootTable.rootWindowWidth
        height: rootTable.rootWindowHeight

        Rectangle {
            id: rootForOverAll
            property string title: "Red"


            anchors.fill: parent
            color: "#e3e3e3"

            Rectangle {
                id: firstRec
                anchors.fill: parent; anchors.margins: 20
                color: "#ff7f7f"
                Text {
                    width: parent.width - 20
                    anchors.centerIn: parent; horizontalAlignment: Qt.AlignHCenter
                    text: "Roses are red"
                    font.pixelSize: 20
                    wrapMode: Text.WordWrap
                }

                OverAllAdmin{
                    id: myOverAll
                    property int overAllWidth: firstRec.width
                    property int overAllHeight: firstRec.height
                    anchors.fill: firstRec
                }

                /*

                Rectangle {
                    id: overAllAdmin2


                    //width: myOverAll.overAllWidth
                    //height: myOverAll.overAllHeight
                    anchors {
                        fill: firstRec
                        margins: 20
                    }
                     color: "#f1237f"

                }
                */

            }
        }

        Rectangle {
            property string title: "Green"
            anchors.fill: parent
            color: "#e3e3e3"

            Rectangle {
                id: secondRec
                anchors.fill: parent; anchors.margins: 20
                color: "#7fff7f"
                Text {
                    width: parent.width - 20
                    anchors.centerIn: parent; horizontalAlignment: Qt.AlignHCenter
                    text: "Flower stems are green"
                    font.pixelSize: 20
                    wrapMode: Text.WordWrap
                }

                PluginsTable{
                    id: rootPluginsTable
                    property int overAllWidth: secondRec.width
                    property int overAllHeight: secondRec.height
                    anchors.fill: secondRec
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
}
