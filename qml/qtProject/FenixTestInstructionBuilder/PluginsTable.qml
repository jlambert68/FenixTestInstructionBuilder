import QtQuick 2.12
import QtQuick 2.0
import QtQuick.Controls 1.4
import QtQuick.Controls 2.3		//Button
import QtQuick.Layouts 1.3		//ColumnLayout
import QtQuick.Dialogs 1.1
import QtQuick.Controls.Styles 1.4

Item {
    id: pluginsTableItem

    width: myOverAll.overAllWidth
    height: myOverAll.overAllHeight
    anchors {
        fill: rootPluginsTable
        margins: 5
    }
    property string newOrEditStr: "none"
    property bool newOrEditIsProcessedBool: true

    //Component.onCompleted: {
    // var JsonString = '{"newOrEdet":"none","commsndProcessed":"true"}';
    // var JsonObject= JSON.parse(JsonString);

    //}







    ListModel {
        id: pluginsModel

        ListElement {
            name: "Send UTR"
            guid: "2243085a-feee-4ae7-8ccf-03f69c0704a4"
            readyForUse: true
            activated: true

            attributes: [
                ListElement { description: "Sends UTRs via MQ towards Custody Cash" }
            ]
        }
        ListElement {
            name: "Validate New Pacs008"
            guid: "d456499c-2ad1-4677-8e1d-909a7ecab560"
            readyForUse: false
            activated: false

            attributes: [
                ListElement { description: "Validates that a newly created Pacs008 has been sent to CMaaS" }
            ]
        }
    }

    Component {
        id: fruitDelegate
        Item {
            width: 200; height: 50
            Text { id: nameField; text: name }
            Text { text: '$' + cost; anchors.left: nameField.right }
            Row {
                anchors.top: nameField.bottom
                spacing: 5
                Text { text: "Attributes:" }
                Repeater {
                    model: attributes
                    Text { text: description }
                }
            }
        }
    }



    GroupBox {
        id: pluginGroupBox
        title: "Supported Plugins"
        anchors.fill: parent


        ColumnLayout {
            id: mainLayout
            anchors.fill: parent
            Layout.fillWidth: true
            Layout.alignment: Qt.AlignTop

            TableView {
                id: pluginsTableView
                //anchors.fill: parent
                Layout.fillWidth: parent
                Layout.alignment: Qt.AlignTop
                //width: 500


                TableViewColumn {
                    role: "name"
                    title: "Name"
                    width: 100
                }

                TableViewColumn {
                    role: "guid"
                    title: "GUID"
                    width: 200
                }

                TableViewColumn {
                    role: "readyForUse"
                    title: "Ready for use"
                    width: 50
                }

                TableViewColumn {
                    role: "activated"
                    title: "Activated"
                    width: 50
                }

                model: pluginsModel

                onActivated: {
                    console.log("User clicked on row: "+ row)
                }
            }


            // *************** Buttons for New, Edit and Delete
            RowLayout {
                id: buttonlayout

                spacing: 6
                Layout.alignment : Qt.AlignTop




                // ** New **
                Button {
                    Layout.alignment : Qt.AlignLeft

                    text: "New Plugin"

                    onClicked: {
                        console.log("Creare New Plugin")
                        newOrEditStr = "New"
                        newOrEditIsProcessedBool = false
                        popUpTestInstruction.open()


                    }
                }
                // ** Edit **
                Button {
                    Layout.alignment : Qt.AlignLeft

                    text: "Edit Plugin"

                    onClicked: {
                        // check that a row is selected before executing command
                        if (pluginsTableView.currentRow === -1) {
                            myMessageDialog.visible = true

                        } else {
                            console.log("Edit Plugin on row: "+ pluginsTableView.currentRow)
                            //enabled : false
                            //root.template.clicked()
                            newOrEditStr = "Edit"
                            newOrEditIsProcessedBool = false
                            popUpTestInstruction.open()

                        }
                    }
                }

                // ** Delete **
                Button {
                    Layout.alignment : Qt.AlignLeft

                    text: "Delete Plugin"

                    onClicked: {
                        // check that a row is selected before executing command
                        if (pluginsTableView.currentRow === -1) {
                            myMessageDialog.visible = true

                        } else {
                            console.log("Delete Plugin on row: " + pluginsTableView.currentRow)
                            pluginsModel.remove(pluginsTableView.currentRow);
                            //enabled : false
                            //root.template.clicked()
                        }
                    }
                }
            }
            // Buttons for New, Edit and Delete ***************
        }
    }



    // *************** MessageBoix shown when no row has been selected in the table before trying to do some stuff
    MessageDialog {
        id: myMessageDialog
        title: "No Row selected"
        text: "A row has to be seleced before executing command."

        standardButtons: StandardButton.Close
        icon: StandardIcon.Warning
        Component.onCompleted: visible = false
    }
    // MessageBox shown when no row has been selected in the table before trying to do some stuff ***************



    Popup {
        id: popUpTestInstruction
        //x: 100
        //y: pluginsTableView.y + pluginsTableView.height + 50
        width: 400
        height: popUptestInstructioLastButton.y + popUptestInstructioLastButton.height +30
        modal: true
        focus: true
        closePolicy: Popup.CloseOnEscape //| Popup.CloseOnPressOutsideParent
        opacity: 0.7
        anchors.centerIn: parent
        // Variabled for handling if values are changed by user
        property string inComingGuid: ""
        property bool inComingGuidChanged: false
        property string inComingName : ""
        property bool inComingNameChanged: false
        property string incomingDescription: ""
        property bool incomingDescriptionChanged: false
        property bool inComingActivated: false
        property bool inComingActivatedChanged: false
        property bool inComingReadeForUse: false
        property bool inComingReadeForUseChanged: false



        exit: Transition {
            NumberAnimation { property: "opacity"; from: 1.0; to: 0.0; duration: 1200 }



        }
        enter: Transition {

            NumberAnimation { property: "opacity"; from: 0.0; to: 1.0; duration: 800}

        }

        onAboutToShow: {
            switch(pluginsTableItem.newOrEditStr ) {
            case "New":
                popUptestInstructionGuid.text = ""
                popUptestInstructionName.text = ""
                poUpTestInstructionIsReadyForUsedCheckBox.checked = false
                poUpTestInstructionIsActivatedCheckBox.checked = false
                popUptestInstructionDescription.text = ""
                break;

            case "Edit":
                // Copy values from current row in pluginsTableItem
                popUptestInstructionGuid.text = pluginsModel.get(pluginsTableView.currentRow).guid
                popUptestInstructionName.text = pluginsModel.get(pluginsTableView.currentRow).name
                poUpTestInstructionIsReadyForUsedCheckBox.checked = pluginsModel.get(pluginsTableView.currentRow).readyForUse
                poUpTestInstructionIsActivatedCheckBox.checked = pluginsModel.get(pluginsTableView.currentRow).activated
                popUptestInstructionDescription.text = pluginsModel.get(pluginsTableView.currentRow).attributes.get(0).description

                break;

            default:
                console.log("'pluginsTableItem.newOcheckedcheckedrEditStr' in PopUp has an unexpected value: " + pluginsTableItem.newOrEditStr)
            };

            // Clear initial values and flags for change-control
            popUpTestInstruction.inComingGuidChanged = false
            popUpTestInstruction.inComingGuid = popUptestInstructionGuid.text

            popUpTestInstruction.inComingNameChanged = false
            popUpTestInstruction.inComingName = popUptestInstructionName.text

            popUpTestInstruction.inComingReadeForUseChanged = false
            popUpTestInstruction.inComingReadeForUse = poUpTestInstructionIsReadyForUsedCheckBox.checked

            popUpTestInstruction.inComingActivatedChanged = false
            popUpTestInstruction.inComingActivated = poUpTestInstructionIsActivatedCheckBox.checked

            popUpTestInstruction.incomingDescriptionChanged  = false
            popUpTestInstruction.incomingDescription = popUptestInstructionDescription.text


        }

        onAboutToHide: {

        }



        // signal popUpSignalReceived(string txt)
        // popUpSignalReceived:console.debug(txt)

        ColumnLayout {
            spacing: 12
            width: popUpTestInstruction.width - 20

            ColumnLayout {
                Label {
                    text: pluginsTableItem.newOrEditStr + " Test Instruction"
                    font.pixelSize: 22
                    font.bold: true
                    color: "steelblue"
                }
            }

            ColumnLayout {
                Label {
                    text: "Guid"
                }
                TextField {
                    id: popUptestInstructionGuid
                    placeholderText: "Autogenerated guid"
                    //Layout.fillWidth: true
                    enabled: false

                    ToolTip.delay: 1000
                    ToolTip.timeout: 5000
                    ToolTip.visible: hovered
                    ToolTip.text: qsTr("Used as unique id for the plugin.")

                    onAcceptableInputChanged: {
                        // Do control if value has change from inconming value
                        if (popUptestInstructionGuid.text != popUpTestInstruction.inComingGuid) {
                            popUpTestInstruction.inComingGuidChanged = true
                        } else {
                            popUpTestInstruction.inComingGuidChanged = false
                        }
                    }
                }
            }

            ColumnLayout {
                Label {
                    text: "Name"
                }
                TextField {
                    id: popUptestInstructionName
                    placeholderText: "Plugin Name"
                    //Layout.fillWidth: true

                    ToolTip.delay: 1000
                    ToolTip.timeout: 5000
                    ToolTip.visible: hovered
                    ToolTip.text: qsTr("The Name for the plugin.")


                    onEditingFinished: {
                        // Do control if value has change from inconming value
                        if (popUptestInstructionName.text != popUpTestInstruction.inComingName) {
                            popUpTestInstruction.inComingNameChanged = true
                        } else {
                            popUpTestInstruction.inComingNameChanged = false
                        }
                    }
                }
            }

            // -- poUpTestInstructionIsActivatedCheckBox --
            ColumnLayout {
                CheckBox {
                    id: poUpTestInstructionIsActivatedCheckBox
                    checked: true
                    text: qsTr("Test Instruction is Activated")

                    ToolTip.delay: 1000
                    ToolTip.timeout: 5000
                    ToolTip.visible: hovered
                    ToolTip.text: qsTr("Should the TestInstruction be activated for use in Fenix.")

                    onClicked: {
                        // Do control if value has change from inconming value
                        if (poUpTestInstructionIsActivatedCheckBox.checked != popUpTestInstruction.inComingActivated) {
                            popUpTestInstruction.inComingActivatedChanged = true
                        } else {
                            popUpTestInstruction.inComingActivatedChanged = false
                        }
                    }

                }
            }

            // -- poUpTestInstructionIsReadyForUsedCheckBox --
            ColumnLayout {
                CheckBox {
                    id: poUpTestInstructionIsReadyForUsedCheckBox
                    checked: true
                    text: qsTr("Test Instruction is Ready for use")

                    ToolTip.delay: 1000
                    ToolTip.timeout: 5000
                    ToolTip.visible: hovered
                    ToolTip.text: qsTr("Does the Test Instruction has all values XXXXXXXXXX")

                    onClicked: {
                        // Do control if value has change from inconming value
                        if (poUpTestInstructionIsReadyForUsedCheckBox.checked != popUpTestInstruction.inComingReadeForUse) {
                            popUpTestInstruction.inComingReadeForUseChanged = true
                        } else {
                            popUpTestInstruction.inComingReadeForUseChanged = false
                        }
                    }

                }
            }

            // -- testInstructionDescription
            ColumnLayout {
                id: textBoxColumnLayoutId
                Label {
                    text: "Test Instruction Description"
                }


                Item {
                    id: textBoxItem
                    height: focusScope.height +15


                    FocusScope {
                        id: focusScope;
                        width: popUpTestInstruction.width - 20;
                        height: popUptestInstructionDescription.paintedHeight + (2 * popUptestInstructionDescription.anchors.topMargin) +130;

                        property alias  value                          : popUptestInstructionDescription.text;
                        property alias  fontSize                       : popUptestInstructionDescription.font.pointSize;
                        property alias  textColor                      : popUptestInstructionDescription.color;
                        property alias  placeHolder                    : typeSomething.text;


                        onHeightChanged: {
                            console.log("focusScope.height: " + focusScope.height)
                            textBoxItem.height = focusScope.height +15
                            console.log("textBoxItem.height: " + textBoxItem.height)



                        }

                        Rectangle {
                            id: background;
                            anchors.fill: parent;
                            color: "#AAAAAA";
                            radius: 5;
                            antialiasing: true;
                            border {
                                width: 3;
                                color: (focusScope.activeFocus ? "red" : "steelblue");
                            }

                        }
                        TextEdit {
                            id: popUptestInstructionDescription;
                            focus: true
                            selectByMouse: true
                            font.pointSize: popUptestInstructionName.font.pointSize;
                            font.family: popUptestInstructionName.font.family
                            wrapMode: TextEdit.WrapAtWordBoundaryOrAnywhere;

                            color: "black";
                            anchors {
                                top: parent.top;
                                topMargin: 10;
                                left: parent.left;
                                leftMargin: 10;
                                right: parent.right;
                                rightMargin: 10;
                            }

                            onTextChanged: {
                                // Do control if value has change from inconming value
                                if (popUptestInstructionDescription.text != popUpTestInstruction.incomingDescription) {
                                    popUpTestInstruction.incomingDescriptionChanged = true
                                } else {
                                    popUpTestInstruction.incomingDescriptionChanged = false
                                }
                            }
                        }
                        Text {
                            id: typeSomething;
                            text: "Type something...";
                            color: "gray";
                            opacity: (value === "" ? 1.0 : 0.0);
                            font {
                                italic: true
                                pointSize: fontSize;
                            }
                            anchors {
                                left: parent.left;
                                right: parent.right;
                                leftMargin: 10;
                                rightMargin: 10;
                                verticalCenter: parent.verticalCenter;
                            }

                            Behavior on opacity { NumberAnimation { duration: 250; } }
                        }
                        MouseArea {
                            visible: (!focusScope.activeFocus);
                            anchors.fill: parent
                            onClicked: { popUptestInstructionDescription.forceActiveFocus(); }
                        }
                        Text {
                            id: clear;
                            text: "\u2717" // 'x' //"\u2713"
                            color: 'steelblue';
                            font.pixelSize: 30;
                            opacity: (value !== "" ? 1.0 : 0.0);
                            anchors {
                                right: parent.right;
                                bottom: parent.bottom;
                                margins: 5;
                            }

                            Behavior on opacity { NumberAnimation { duration: 250; } }

                            MouseArea {
                                anchors.fill: parent
                                onClicked: {
                                    value = "";
                                    focusScope.focus = false;
                                }
                            }
                        }
                    }
                }
            }



            Button {
                id: popUptestInstructioSaveButton
                text: "Save"
                onClicked: {
                    // Verify that user has changed anything before continue to save any changes
                    if (popUpTestInstruction.inComingGuidChanged || popUpTestInstruction.inComingNameChanged ||
                            popUpTestInstruction.inComingActivatedChanged || popUpTestInstruction.inComingReadeForUseChanged ||
                            popUpTestInstruction.incomingDescriptionChanged) {

                        // Name and Description must have values
                        if (popUptestInstructionName.text.length === 0 ||
                                popUptestInstructionDescription.text.length === 0) {
                            console.log("Can't Save NEW because Name or Description is missing data")

                        } else {

                            // Do the popup process a New object or Edit an existing object
                            switch(pluginsTableItem.newOrEditStr ) {
                            case "New":

                                // Create a new Plugin in table
                                pluginsModel.append({name: popUptestInstructionName.text,
                                                        guid: popUptestInstructionGuid.text,
                                                        readyForUse: poUpTestInstructionIsReadyForUsedCheckBox.checked,
                                                        activated: poUpTestInstructionIsActivatedCheckBox.checked});
                                pluginsModel.set(pluginsTableView.rowCount-1).attributes.setproperty(0,"description",popUptestInstructionDescription.text)

                                // Close popUp
                                popUpTestInstruction.close()

                                break;
                            case "Edit":
                                // Copy values back to current row in pluginsTableItem
                                pluginsModel.get(pluginsTableView.currentRow).guid = popUptestInstructionGuid.text
                                pluginsModel.get(pluginsTableView.currentRow).name = popUptestInstructionName.text
                                pluginsModel.get(pluginsTableView.currentRow).readyForUse = poUpTestInstructionIsReadyForUsedCheckBox.checked
                                pluginsModel.get(pluginsTableView.currentRow).activated = poUpTestInstructionIsActivatedCheckBox.checked
                                pluginsModel.get(pluginsTableView.currentRow).attributes.get(0).description = popUptestInstructionDescription.text

                                console.log("Saving changes")

                                // Close popUp
                                popUpTestInstruction.close()


                                break;
                            default:
                                console.log("'pluginsTableItem.newOcheckedcheckedrEditStr' in PopUp has an unexpected value: " + pluginsTableItem.newOrEditStr)
                            };
                        }

                    } else {
                        console.log("Nothing is changed and therby nothing is to be saved!!!")
                    }

                }


            }

            Button {
                id: popUptestInstructioLastButton
                Layout.alignment : Qt.AlignLeft

                text: "Cancel"

            }

        }

    }
}





/*

  TO be used for editing a row
popup.open()
    }

    Popup {
        id: popup
        x: 100
        y: 100
        width: 200
        height: 300
        modal: true
        focus: true
        closePolicy: Popup.CloseOnEscape | Popup.CloseOnPressOutsideParent
    }
  */