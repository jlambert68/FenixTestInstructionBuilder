import QtQuick 2.10				//Item
import QtQuick.Controls 2.3		//Button
import QtQuick.Layouts 1.3		//ColumnLayout
import CustomQmlTypes 1.0		//CustomLabel and BridgeTemplate

import QtQuick 2.2
import QtQuick.Controls 2.0
import QtQuick.Controls 1.2

import QtQuick.Dialogs 1.0
import QtQml.Models 2.2
import CustomQmlTypes 1.0		//CustomLabel and BridgeTemplate




Item {
	id: root

	property BridgeTemplate template: BridgeTemplate{}

	width: 250
	height: 200


    ColumnLayout {
        id: mainLayout
        RowLayout {
            id: layout

            spacing: 6
            Layout.alignment : Qt.AlignTop

 /*Loader {
       sourceComponent : myExtraComponents
       }
*/

            	ColumnLayout {
            	    id: myExtraComponents
            		anchors.fill: parent

            		Label {
            			property CustomLabel label: CustomLabel{}
            			Layout.alignment: Qt.AlignCenter

            			text: label.text
            		}

            		Label {
            			property CustomLabel label: CustomLabel{}
            			Layout.alignment: Qt.AlignCenter

            			text: label.text
            		}

            		Label {
            			property CustomLabel label: CustomLabel{}
            			Layout.alignment: Qt.AlignCenter

            			text: label.text
            		}

            		Button {
            			Layout.fillWidth: true
            			Layout.alignment: Qt.AlignBottom

            			text: "start!"

            			onClicked: {
            				enabled : false
            				root.template.clicked()
            			}
            		}
            		Button {
            			Layout.fillWidth: true
            			property var globalForJs: 10

            			text: "sendVariantListMap"
            			onClicked: {
            				text = "look into the console"
            				root.template.sendVariantListMap(True, [1.23, "hello"],  {"json": {"A": true, "B": 2.23, "C": "hello", "D": {"DA": 3.14, "DB": "Pi", "DC": {"DCA": 1.73, "DCB": "Roten ur 2"}}}})
            			    root.template.sendglobalForJs(globalForJs)

                            var JsonString = '{"a":"A whatever, run","b":"B fore something happens"}';
                            console.log(JsonString);
                            var JsonObject= JSON.parse(JsonString);
                            JsonObject.a = "Sidan 77?";
                            JsonObject.b = "Detta är en banan!"

                            var JsonString2 = JSON.stringify(JsonObject);

                            root.template.sendJsonObject(JsonString2)


                            //retrieve values from JSON again
                            var aString = JsonObject.a;
                            var bString = JsonObject.b;

                            console.log(aString);
                            console.log(bString);
                            console.log(JsonString2);

                            var myJsonStruct = {"Name": "Jonas", "Age": 51};
                            console.log(myJsonStruct);

                            var sammy = {
                              "first_name"  :  "Sammy",
                              "last_name"   :  "Shark",
                              "online"      :  True
                            };

                            console.log(sammy);
                            var sammyStr = JSON.stringify(sammy);
                            root.template.sendJsonObject(sammyStr)

                            var guiData = {
                                            "metadata": {
                                              "isSaved": False,
                                              "isEditedByUser": False ,
                                              "isChangedByDB": False,
                                              "lastSaved": "0"
                                            },
                                            "pluginGuidTextField":
                                              {
                                                "dataWhenLoadedOrSaved": "0",
                                                "dataAfterEdit": "0"
                                              },
                                              "pluginNameCombobox":
                                              {
                                                "dataWhenLoadedOrSaved": "0",
                                                "dataAfterEdit": "0"
                                              }
                                              };




            			}
            		}
            	}
 //   QmlToBeImported{}


            // ************ Test Instruction ************

            ColumnLayout {

                 spacing: 2
                 Layout.minimumWidth : 250
                 Layout.alignment : Qt.AlignTop


                 GroupBox {
                    id: rowBoxTestInstruction
                    title: "Test Instructions"
                    Layout.fillWidth: true
                    Layout.fillHeight: true


                //Column {
                //    id: column0
                //    width: parent.width
                    //heigt: parent.width

            ScrollView {
                anchors.fill: parent

                width: parent.width
                height : columnTestInstruction.height //parent.height
                //contentWidth: column.width    // The important part
                // contentHeight: column.height  // Same
                //clip : true                   // Prevent drawing column outside the scrollview borders

                Column {
                    id: columnTestInstruction
                    width: parent.width

                    // Your items here

                    ColumnLayout {
                        id: rowBoxTestInstructionLayout
                        spacing: 12


                        GroupBox {
                            id: pluginGroupBox
                            title: "Plugin"
                            Layout.fillWidth: true

                            ColumnLayout {
                                id: pluginGroupBoxLayout

                                // -- pluginGuid --
                                ColumnLayout {
                                    id: pluginGuidLayout
                                    //spacing: 2
                                    //Layout.alignment : Qt.AlignTop

                                    Label {
                                        text: "Guid"
                                    }
                                    TextField {
                                        id: pluginGuidTextField
                                        placeholderText: "guid based on name"
                                        //Layout.fillWidth: true

                                        enabled: false

                                        ToolTip.delay: 1000
                                        ToolTip.timeout: 5000
                                        ToolTip.visible: hovered
                                        ToolTip.text: qsTr("Used as unique id for the plugin.")



                                    }
                                }
                                // -- pluginName --
                                ColumnLayout {
                                    id: pluginNameLayput
                                    //spacing: 2
                                    //Layout.alignment : Qt.AlignTop

                                    Label {
                                        text: "Name"
                                    }
                                    ComboBox {
                                        id: pluginNameCombobox

                                        Layout.fillWidth: true

                                        ToolTip.delay: 1000
                                        ToolTip.timeout: 5000
                                        ToolTip.visible: hovered
                                        ToolTip.text: qsTr("Name for the plugin.")

                                        currentIndex: 0
                                        model: ListModel {
                                            id: pluginNameItems
                                            ListElement { text: "- Choose plugin -"; color: "" }
                                            ListElement { text: "Banana"; color: "Yellow" }
                                            ListElement { text: "Apple"; color: "Green" }
                                            ListElement { text: "Coconut"; color: "Brown" }
                                        }
                                        onCurrentIndexChanged: {
                                            pluginGuidTextField.text = pluginNameItems.get(currentIndex).color
                                            //console.log(systemDomainNameTextField.text)

                                            // Update JSON model
                                            guiData.pluginNameCombobox.dataAfterDeit = pluginNameItems.get(currentIndex).text
                                            // Update JSOn model with indirect changes
                                            guiData.pluginGuidTextField.dataAfterEdit = pluginGuidTextField.text

                                            console.log(guiData);
                                            var guiDataStr = JSON.stringify(guiData);
                                            root.template.sendJsonObject(sammyguiDataStrStr)

                                        }
                                    }
                                }

                            }
                        }

                        GroupBox {
                            id: systemDomainGroupBox
                            title: "System Domain"
                            Layout.fillWidth: true

                            ColumnLayout {
                                id: systemDomainGroupBoxLayout

                                // -- systemDomainId --
                                ColumnLayout {
                                    id: systemDomainIdLayout
                                    //spacing: 2
                                    //Layout.alignment : Qt.AlignTop

                                    Label {
                                        text: "Guid"
                                    }
                                    TextField {
                                        id: systemDomainIdField
                                        placeholderText: "guid based on Domain Name"
                                        //Layout.fillWidth: true

                                        ToolTip.delay: 1000
                                        ToolTip.timeout: 5000
                                        ToolTip.visible: hovered
                                        ToolTip.text: qsTr("The Domain/system's Id where the Sender operates.")



                                        }
                                }

                                // -- systemDomainName --
                                ColumnLayout {
                                    id: systemDomainNameLayput
                                    //spacing: 2
                                    //Layout.alignment : Qt.AlignTop

                                    Label {
                                        text: "Name"
                                    }
                                    ComboBox {
                                        id: systemDomainNameCombobox
                                        Layout.fillWidth: true

                                        ToolTip.delay: 1000
                                        ToolTip.timeout: 5000
                                        ToolTip.visible: hovered
                                        ToolTip.text: qsTr("The Domain/system's Name where the Sender operates.")

                                        currentIndex: 0
                                        model: ListModel {
                                            id: systemDomainNameItems
                                            ListElement { text: "- Choose Domain -"; color: "" }
                                            ListElement { text: "Banana"; color: "Yellow" }
                                            ListElement { text: "Apple"; color: "Green" }
                                            ListElement { text: "Coconut"; color: "Brown" }
                                        }
                                        onCurrentIndexChanged: {
                                            systemDomainIdField.text = systemDomainNameItems.get(currentIndex).color
                                            //console.log(systemDomainNameTextField.text)

                                        }
                                    }
                                }
                            }
                        }

                        // -- Test Instruction Type --
                       GroupBox {
                            id: stestInstructionTypeGroupBox
                            title: "Test Instruction Type"
                            Layout.fillWidth: true

                            ColumnLayout {
                                id: testInstructionTypeGroupBoxLayout

                                // -- testInstructionTypeGuid --
                                ColumnLayout {
                                    id: testInstructionTypeGuidLayout
                                    //spacing: 2

                                    Label {
                                        text: "Guid"
                                    }
                                    TextField {
                                        id: testInstructionTypeGuidTextField
                                        placeholderText: "guid based on Test Instruction Type Name"

                                        enabled: false

                                        ToolTip.delay: 1000
                                        ToolTip.timeout: 5000
                                        ToolTip.visible: hovered
                                        ToolTip.text: qsTr("The unique guid for the Type of TestInstruction. Set by Client system.")

                                        }
                                }

                                // -- testInstructionTypeName --
                                ColumnLayout {
                                    id: testInstructionTypeNameLayput
                                    //spacing: 2

                                    Label {
                                        text: "Name"
                                    }
                                    ComboBox {
                                        id: testInstructionTypeNameCombobox
                                        Layout.fillWidth: true

                                        ToolTip.delay: 1000
                                        ToolTip.timeout: 5000
                                        ToolTip.visible: hovered
                                        ToolTip.text: qsTr("The name for the Type of TestInstruction. Set by Client system.")

                                        currentIndex: 0
                                        model: ListModel {
                                            id: testInstructionTypeNameItems
                                            ListElement { text: "- Choose Domain -"; color: "" }
                                            ListElement { text: "Banana"; color: "Yellow" }
                                            ListElement { text: "Apple"; color: "Green" }
                                            ListElement { text: "Coconut"; color: "Brown" }
                                        }
                                        onCurrentIndexChanged: {
                                            testInstructionTypeGuidTextField.text = testInstructionTypeNameItems.get(currentIndex).color
                                            //console.log(systemDomainNameTextField.text)

                                        }
                                    }
                                }
                            }
                        }


                        // -- Test Instruction --
                       GroupBox {
                            id: testInstructionGroupBox
                            title: "Test Instruction"
                            Layout.fillWidth: true

                            ColumnLayout {
                                id: testInstructionGroupBoxLayout

                                // -- testInstructionGuid --
                                ColumnLayout {
                                    id: testInstructionGuidLayout
                                    //spacing: 2

                                    Label {
                                        text: "Guid"
                                    }
                                    TextField {
                                        id: testInstructionGuidTextField
                                        placeholderText: "Unique guid"

                                        enabled: false

                                        ToolTip.delay: 1000
                                        ToolTip.timeout: 5000
                                        ToolTip.visible: hovered
                                        ToolTip.text: qsTr("The unique guid for the TestInstruction.")


                                        }


                                }

                                // -- testInstructionName --
                                ColumnLayout {
                                    id: testInstructionNameLayout
                                    //spacing: 2

                                    Label {
                                        text: "Name"
                                    }
                                    TextField {
                                        id: testInstructionNameTextField
                                        placeholderText: "Test Instruction Name"

                                        ToolTip.delay: 1000
                                        ToolTip.timeout: 5000
                                        ToolTip.visible: hovered
                                        ToolTip.text: qsTr("The name of the TestInstruction.")



                                        onEditingFinished: {
                                            console.log("Denna är ändrad")

                                        }
                                        }
                                }
                            }
                        }
                        }
                        // -- testInstructionDescription
                        ColumnLayout {
                            id: testInstructionDescriptionLayout
                            //spacing: 2

                            Label {
                                text: "Test Instruction Description"
                            }
                            TextArea {
                                id: testInstructionDescriptionTextArea
                                //placeholderText: qsTr("Enter description")
                                //placeholderText: "Test Instruction Name"

                                ToolTip.delay: 1000
                                ToolTip.timeout: 5000
                                //ToolTip.visible: hovered
                                ToolTip.text: qsTr("The description of the TestInstruction.")

                                }
                        }



                        // -- testInstructionMouseOver -->
                        ColumnLayout {
                            id: testInstructionMouseOverLayout
                            //spacing: 2

                            Label {
                                text: "Test Instruction Mouse Over"
                            }
                            TextField {
                                id: testInstructionMouseOverTextField
                                placeholderText: "Mouse over text"

                                ToolTip.delay: 1000
                                ToolTip.timeout: 5000
                                ToolTip.visible: hovered
                                ToolTip.text: qsTr("The mouse over text for the TestInstruction.")

                                }
                        }

                        // -- testInstructionVisible --
                        ColumnLayout {
                            id: testInstructionVisibleLayout

                            CheckBox {
                                id: testInstructionVisibleCheckBox
                                checked: true
                                text: qsTr("Test Instruction Visible")

                                ToolTip.delay: 1000
                                ToolTip.timeout: 5000
                                ToolTip.visible: hovered
                                ToolTip.text: qsTr("Should the TestInstruction be visible in GUI or not.")

                                }
                        }

                        // -- testInstructionEnable --
                        ColumnLayout {
                            id: testInstructionEnableLayout

                            CheckBox {
                                id: testInstructionEnableCheckBox
                                checked: true
                                text: qsTr("Test Instruction Enabled")

                                ToolTip.delay: 1000
                                ToolTip.timeout: 5000
                                ToolTip.visible: hovered
                                ToolTip.text: qsTr("Should the TestInstruction be enabled or not in GUI.")

                                }
                        }

                        // -- TestInstructionColor --
                        ColumnLayout {
                            id: testInstructionColorVLayout
                            //spacing: 2

                           GroupBox {
                                id: textBoxInputMaskGroupBox
                                title: "Test Instruction Color"
                                Layout.fillWidth: true

                                ColumnLayout {
                                    id: testInstructionColorVLayout2

                                    CheckBox {
                                        id: testInstructionColorCheckBox
                                        checked: false
                                        text: qsTr("Use Color")

                                        ToolTip.delay: 1000
                                        ToolTip.timeout: 5000
                                        ToolTip.visible: hovered
                                        ToolTip.text: qsTr("The color used for presenting the TestInstructionBlock, e.g. #FAF437.")

                                        }

                                        RowLayout {
                                            id: testInstructionColorHLayout
                                            //spacing: 2
                                            Rectangle {
                                                id: testInstructionColorColorArea
                                                width: 20
                                                height: 20
                                                color: "#777777"
                                                border.color: "black"
                                                border.width: 2

                                                MouseArea {
                                                    anchors.fill: parent
                                                    acceptedButtons: Qt.LeftButton
                                                    onClicked: {
                                                        parent.color = 'red'
        /*


                                                         */
                                                    }
                                                }
                                            }

                                                TextField {
                                                        id: testInstructionColorHLayoutTextField
                                                        placeholderText: "Eg. #F3A566"

                                                        text: "777777"
                                                        inputMask: ">HHHHHH"

                                                        ToolTip.delay: 1000
                                                        ToolTip.timeout: 5000
                                                        ToolTip.visible: hovered
                                                        ToolTip.text: qsTr("The color used for presenting the TestInstructionBlock, e.g. #FAF437..")

                                                        onTextChanged: acceptableInput ? testInstructionColorColorArea.color = '#' + testInstructionColorHLayoutTextField.text : print("Input not acceptable")
                                                        }
                                        }
                                    }
                                }

                        }


                        // -- ChoosenLandingZone --

                        // -- Version --
                        GroupBox {
                            id: versionGroupBox
                            title: "Test Instruction Version"
                            Layout.fillWidth: true

                            ColumnLayout {
                                id: versionGroupBoxkLayout

                                // -- MajorVersion --
                                ColumnLayout {
                                    id: majorVersionLayout
                                    //spacing: 2

                                    Label {
                                        text: "Major"
                                    }
                                    TextField {
                                        id: majorVersionTextField
                                        placeholderText: "2"

                                        ToolTip.delay: 1000
                                        ToolTip.timeout: 5000
                                        ToolTip.visible: hovered
                                        ToolTip.text: qsTr("The version numbering system consists of major and minor. A change in major version forces the user to upgrade the TestInstruction.")

                                        }
                                }

                                // -- MinorVersion --
                                ColumnLayout {
                                    id: minorVersionLayout
                                    //spacing: 2

                                    Label {
                                        text: "Minor"
                                    }
                                    TextField {
                                        id: minorVersionTextField
                                        placeholderText: "3"

                                        ToolTip.delay: 1000
                                        ToolTip.timeout: 5000
                                        ToolTip.visible: hovered
                                        ToolTip.text: qsTr("The version numbering system consists of major and minor. A change in minor version makes it optional for  the user to upgrade the TestInstruction.")

                                        }
                                }
                            }
                        }

                    }
                 }


}
            }

            // ************ Test Instructions Details ************
            ColumnLayout {
                id: testInstructionDetailsLayout
        //       anchors.fill: parent
        //       anchors.margins: marginUsed as unique n
                 spacing: 2
                 Layout.minimumWidth : 300
                 Layout.alignment : Qt.AlignTop

                 GroupBox {
                     id: rowBoxTestInstructionDetails
                     title: "Test Instructions Details"
                     Layout.fillWidth: true
                     Layout.fillHeight: true

            ScrollView {
                anchors.fill: parent

                width: parent.width
                height : columnTestInstructionSDetails.height //parent.height
                //contentWidth: column.width    // The important part
                // contentHeight: column.height  // Same
                clip : true                   // Prevent drawing column outside the scrollview borders

                Column {
                    id: columnTestInstructionSDetails
                    width: parent.width

                    // Your items here

                     ColumnLayout {
                         id: rowBoxTestInstructioDetailsnLayout
                         spacing: 12

                       GroupBox {
                            id: testInstructionAttributeGroupBox
                            title: "Test Instruction Attribute"
                            Layout.fillWidth: true

                            ColumnLayout {
                                id: testInstructionAttributeGroupBoxLayout

                                 // -- TestInstructionAttributeGuid --
                                 ColumnLayout {
                                     id: testInstructionAttributeGuidLayout
                                     //spacing: 2
                                     //Layout.alignment : Qt.AlignTop

                                     Label {
                                         text: "Guid"
                                     }
                                     TextField {
                                         id: testInstructionAttributeGuidTextField
                                         placeholderText: "guid"
                                         //Layout.fillWidth: true

                                         enabled: false

                                         ToolTip.delay: 1000
                                         ToolTip.timeout: 5000
                                         ToolTip.visible: hovered
                                         ToolTip.text: qsTr("The unique guid for the TestInstructionAttribute, set by plugin.")

                                     }
                                 }

                                 // -- TestInstructionAttributeName --
                                 ColumnLayout {
                                     id: testInstructionAttributeNameLayout
                                     //spacing: 2
                                     //Layout.alignment : Qt.AlignTop

                                     Label {
                                         text: "Name"
                                     }
                                     TextField {
                                         id: testInstructionAttributeNameTextField
                                         placeholderText: "Name"
                                         //Layout.fillWidth: true

                                         enabled: true

                                         ToolTip.delay: 1000
                                         ToolTip.timeout: 5000
                                         ToolTip.visible: hovered
                                         ToolTip.text: qsTr("The name of the TestInstructionAttribute.")

                                     }
                                 }
                            }
                         }

                         // -- TestInstructionAttributeTypeGuid --
                         ColumnLayout {
                             id: testInstructionAttributeTypeGuidLayout
                             //spacing: 2
                             //Layout.alignment : Qt.AlignTop

                             Label {
                                 text: "Test Instruction Attribute Type Guid"
                             }
                             TextField {
                                 id: testInstructionAttributeTypeGuidTextField
                                 placeholderText: "guid"
                                 //Layout.fillWidth: true

                                 enabled: false

                                 ToolTip.delay: 1000
                                 ToolTip.timeout: 5000
                                 ToolTip.visible: hovered
                                 ToolTip.text: qsTr("The unique guid for the TestInstructionAttribute-type, set by plugin.")

                             }
                         }



                         // -- TestInstructionAttributeDescription --
                        ColumnLayout {
                            id: testInstructionAttributeDescriptionLayout
                            //spacing: 2

                            Label {
                                text: "Test Instruction Attribute Description"
                            }
                            TextArea {
                                id: testInstructionAttributeDescriptionTextArea
                                //placeholderText: qsTr("Enter description")
                                //placeholderText: "Test Instruction Name"

                                ToolTip.delay: 1000
                                ToolTip.timeout: 5000
                                //ToolTip.visible: hovered
                                ToolTip.text: qsTr("The description of the TestInstructionAttribute.")

                                }
                        }
                         // -- TestInstructionAttributeMouseOver --
                         ColumnLayout {
                             id: testInstructionAttributeMouseOverLayout
                             //spacing: 2
                             //Layout.alignment : Qt.AlignTop

                             Label {
                                 text: "Test Instruction Attribute MouseOver"
                             }
                             TextField {
                                 id: testInstructionAttributeMouseOverTextField
                                 placeholderText: "Mouse Over Text"
                                 //Layout.fillWidth: true

                                 ToolTip.delay: 1000
                                 ToolTip.timeout: 5000
                                 ToolTip.visible: hovered
                                 ToolTip.text: qsTr("The mouse over text for the TestInstructionAttribute.")

                             }
                         }

                         // -- TestInstructionAttributeVisible --
                        ColumnLayout {
                            id: testInstructionAttributeVisibleLayout

                            CheckBox {
                                id: estInstructionAttributeVisibleCheckBox
                                checked: true
                                text: qsTr("Test Instruction Attribute Visible")

                                ToolTip.delay: 1000
                                ToolTip.timeout: 5000
                                ToolTip.visible: hovered
                                ToolTip.text: qsTr("Should the TestInstructionAttribute be visible in attributes list in GUI or not.")

                                }
                        }

                         // -- TestInstructionAttributeEnable --
                        ColumnLayout {
                            id: testInstructionAttributeEnableLayout

                            CheckBox {
                                id: testInstructionAttributeEnableCheckBox
                                checked: true
                                text: qsTr("Test Instruction Attribute Enabled")

                                ToolTip.delay: 1000
                                ToolTip.timeout: 5000
                                ToolTip.visible: hovered
                                ToolTip.text: qsTr("Should the TestInstructionAttribute be enabled or not.")

                                }
                        }

                         // -- TestInstructionAttributeMandatory --
                        ColumnLayout {
                            id: testInstructionAttributeMandatoryLayout

                            CheckBox {
                                id: testInstructionAttributeMandatoryCheckBox
                                checked: true
                                text: qsTr("Test Instruction Attribute Mandatory")

                                ToolTip.delay: 1000
                                ToolTip.timeout: 5000
                                ToolTip.visible: hovered
                                ToolTip.text: qsTr("Should the Test Instruction Attribute be enabled or not.")

                                }
                        }

                         // -- TestInstructionAttributeVisibleInBlockArea --
                        ColumnLayout {
                            id: testInstructionAttributeVisibleInBlockAreaLayout

                            CheckBox {
                                id: testInstructionAttributeVisibleInBlockAreaCheckBox
                                checked: false
                                text: qsTr("Test Instruction Attribute Visible In Block Area")

                                ToolTip.delay: 1000
                                ToolTip.timeout: 5000
                                ToolTip.visible: hovered
                                ToolTip.text: qsTr("Should the TestInstructionAttribute be visible in Test Instruction Block Area in GUI or not.")

                                }
                        }

Rectangle {
    width: 200; height: 100

    DelegateModel {
        id: visualModel
        model: ListModel {
            ListElement { name: "Apple" }
            ListElement { name: "Orange" }
        }

        groups: [
            DelegateModelGroup { name: "selected" }
        ]

        delegate: Rectangle {
            id: item
            height: 25
            width: 200
            Text {
                text: {
                    var text = "Name: " + name
                    if (item.DelegateModel.inSelected)
                        text += " (" + item.DelegateModel.selectedIndex + ")"
                    return text;
                }
            }
            MouseArea {
                anchors.fill: parent
                onClicked: item.DelegateModel.inSelected = !item.DelegateModel.inSelected
            }
        }
    }

    ListView {
        anchors.fill: parent
        model: visualModel
    }
}

}

                 }
            }
            }
            }


            // ************ Test Instructions Attribute Type Details ************
             ColumnLayout {
                id: testInstructionAttributeLayout
                //        anchors.fill: parent
                //        anchors.margins: margin
                spacing: 2
                Layout.minimumWidth : 200
                Layout.alignment : Qt.AlignTop

                // ******************** Textbox ********************
                 GroupBox {
                    id: rowBoxTestInstructionsAttributeTypeDetailsTextBox
                    title: "Test Instructions Attribute Type Details - TextBox"
                    Layout.fillWidth: true
                    Layout.fillHeight: true

                    // Layout holding all object inside groupbox
                    ColumnLayout {
                        id: testInstructionsAttributeTypeDetailsTextboxLayout
                        spacing: 12

                        // -- TestInstructionAttributeInputTextBoxGuid --
                        ColumnLayout {
                            id: testInstructionsAttributeTypeDetailsTextBoxLayout
                            //spacing: 2
                            //Layout.alignment : Qt.AlignTop

                            Label {
                                text: "Test Instruction Attribute Type Guid"
                            }
                            TextField {
                                id: testInstructionsAttributeTypeDetailsTextBoxTextField
                                placeholderText: "guid"
                                //Layout.fillWidth: true
                                enabled: false

                                ToolTip.delay: 1000
                                ToolTip.timeout: 5000
                                ToolTip.visible: hovered
                                ToolTip.text: qsTr("The unique guid for the TestInstructionAttributeInputTextBoxProperties, set by plugin, set by plugin.")

                            }
                        }

                        // -- TestInstructionAttributeInputTextBoxName --
                        ColumnLayout {
                            id: testInstructionAttributeInputTextBoxNameLayout
                            //spacing: 2
                            //Layout.alignment : Qt.AlignTop

                            Label {
                                text: "Test Instruction Attribute Type Guid"
                            }
                            TextField {
                                id: testInstructionAttributeInputTextBoxNameTextField
                                placeholderText: "Attribute Name"
                                //Layout.fillWidth: true

                                ToolTip.delay: 1000
                                ToolTip.timeout: 5000
                                ToolTip.visible: hovered
                                ToolTip.text: qsTr("The name of the TestInstructionAttributeInputTextBoxProperties.")

                            }
                        }

                        // -- TextBoxEditable --
                            ColumnLayout {
                                id: textBoxEditableLayout

                                CheckBox {
                                    id: textBoxEditableCheckBox
                                    checked: true
                                    text: qsTr("TextBox is Editable")

                                    ToolTip.delay: 1000
                                    ToolTip.timeout: 5000
                                    ToolTip.visible: hovered
                                    ToolTip.text: qsTr("Should the the TextBox be editable or not, in the GUI.")

                                    }
                            }
                        // -- TextBoxInputMask --
                        ColumnLayout {
                            id: textBoxInputMaskLayout

                           GroupBox {
                                id: textBoxInputMaskLayoutGroupBox
                                title: "TextBox Input Mask"
                                Layout.fillWidth: true

                                ColumnLayout {
                                    id: textBoxInputMaskLayout2

                                    CheckBox {
                                        id: textBoxInputMaskCheckBox
                                        checked: false
                                        text: qsTr("Should inputmaks be used")

                                        ToolTip.delay: 1000
                                        ToolTip.timeout: 5000
                                        ToolTip.visible: hovered
                                        ToolTip.text: qsTr("Should the the TextBox be editable or not, in the GUI.")

                                    }

                                    TextField {
                                            id: textBoxInputMaskTextField
                                            placeholderText: "Eg. #999 999"

                                            ToolTip.delay: 1000
                                            ToolTip.timeout: 5000
                                            ToolTip.visible: hovered
                                            ToolTip.text: qsTr("Inputmask for the TextBoxs.")

                                            onTextChanged: acceptableInput ? testInstructionColorColorArea.color = '#' + testInstructionColorHLayoutTextField.text : print("Input not acceptable")
                                            }
                                    }
                                }
                            }


                        // -- TextBoxAttributeTypeGuid --
                        ColumnLayout {
                            id: textBoxAttributeTypeGuidLayout
                            //spacing: 2

                            Label {
                                text: "Test Instruction Type Guid"
                            }
                            TextField {
                                id: textBoxAttributeTypeGuidTextField
                                placeholderText: "guid based on Attribute Type Name"

                                enabled: false

                                ToolTip.delay: 1000
                                ToolTip.timeout: 5000
                                ToolTip.visible: hovered
                                ToolTip.text: qsTr("The unique guid for the Type of the TextBox. Used for datamanupulation.")

                                }
                        }

                        // -- TextboxAttributeTypeName --
                        ColumnLayout {
                            id: textboxattributeTypeNameLayput
                            //spacing: 2

                            Label {
                                text: "Textbox Attribute Type Name"
                            }
                            ComboBox {
                                id: textboxAttributeTypeNameCombobox
                                Layout.fillWidth: true

                                ToolTip.delay: 1000
                                ToolTip.timeout: 5000
                                ToolTip.visible: hovered
                                ToolTip.text: qsTr("The name for the Type of the ComboBox. Used for datamanupulation.")

                                currentIndex: 0
                                model: ListModel {
                                    id: textboxAttributeTypeNameItems
                                    ListElement { text: "- Choose Domain -"; color: "" }
                                    ListElement { text: "Banana"; color: "Yellow" }
                                    ListElement { text: "Apple"; color: "Green" }
                                    ListElement { text: "Coconut"; color: "Brown" }
                                }
                                onCurrentIndexChanged: {
                                    textBoxAttributeTypeGuidTextField.text = testInstructionTypeNameItems.get(currentIndex).color
                                    //console.log(systemDomainNameTextField.text)

                                }
                            }
                        }

                        // -- TextBoxAttributeValue --
                        ColumnLayout {
                            id: textBoxAttributeValueLayout
                            //spacing: 2

                            Label {
                                text: "TextBoxAttributeValue"
                            }
                            TextField {
                                id: textBoxAttributeValueTextField
                                placeholderText: "Preset value"

                                enabled: true

                                ToolTip.delay: 1000
                                ToolTip.timeout: 5000
                                ToolTip.visible: hovered
                                ToolTip.text: qsTr("The value for the the TextBox, used for preset values.")

                                }
                            }

                        }
                }

                // ******************** Conbobox ********************
                GroupBox {
                    id: rowBoxTestInstructionsAttributeTypeDetailsComboBox
                    title: "Test Instructions Attribute Type Details - ComboBox"
                    Layout.fillWidth: true
                    Layout.fillHeight: true

                    // Layout holding all object inside groupbox
                    ColumnLayout {
                        id: testInstructionsAttributeTypeDetailsComboBoxLAllayout
                        spacing: 12

                        // -- TestInstructionAttributeComboBoxGuid --
                        ColumnLayout {
                            id: testInstructionsAttributeTypeDetailsComboBoxLayout
                            //spacing: 2
                            //Layout.alignment : Qt.AlignTop

                            Label {
                                text: "Test Instruction Attribute Type Guid"
                            }
                            TextField {
                                id: testInstructionsAttributeTypeDetailsComboBoxTextField
                                placeholderText: "guid"
                                //Layout.fillWidth: true
                                enabled: false

                                ToolTip.delay: 1000
                                ToolTip.timeout: 5000
                                ToolTip.visible: hovered
                                ToolTip.text: qsTr("The unique guid for the TestInstructionAttributeInputComboBoxProperties, set by plugin, set by plugin.")

                            }
                        }

                        // -- TestInstructionAttributeComboBoxName --
                        ColumnLayout {
                            id: testInstructionAttributeInputComboBoxNameLayout
                            //spacing: 2
                            //Layout.alignment : Qt.AlignTop

                            Label {
                                text: "Test Instruction Attribute Type Guid"
                            }
                            TextField {
                                id: testInstructionAttributeInputComboBoxNameTextField
                                placeholderText: "Attribute Name"
                                //Layout.fillWidth: true

                                ToolTip.delay: 1000
                                ToolTip.timeout: 5000
                                ToolTip.visible: hovered
                                ToolTip.text: qsTr("The name of the TestInstructionAttributeInputComboBoxProperties.")

                            }
                        }

                        // -- ComboBoxEditable --
                            ColumnLayout {
                                id: comboBoxEditableLayout

                                CheckBox {
                                    id: comboBoxEditableCheckBox
                                    checked: true
                                    text: qsTr("ComboBox is Editable")

                                    ToolTip.delay: 1000
                                    ToolTip.timeout: 5000
                                    ToolTip.visible: hovered
                                    ToolTip.text: qsTr("Should the the ComboBox be editable or not, in the GUI.")

                                    }
                            }
                        // -- ComboBoxInputMask --
                        ColumnLayout {
                            id: comboBoxInputMaskLayout

                            GroupBox {
                                id: comboBoxInputMaskGroupBox
                                title: "ComboBox Input Mask"
                                Layout.fillWidth: true

                                ColumnLayout {
                                    id: comboBoxInputMaskLayout2

                                    CheckBox {
                                        id: comboBoxInputMaskCheckBox
                                        checked: false
                                        text: qsTr("Should inputmaks be used")

                                        ToolTip.delay: 1000
                                        ToolTip.timeout: 5000
                                        ToolTip.visible: hovered
                                        ToolTip.text: qsTr("Should the the ComboBox be editable or not, in the GUI.")

                                    }

                                    TextField {
                                        id: comboBoxInputMaskTextField
                                        placeholderText: "Eg. #999 999"

                                        ToolTip.delay: 1000
                                        ToolTip.timeout: 5000
                                        ToolTip.visible: hovered
                                        ToolTip.text: qsTr("Inputmask for the ComboBoxs.")

                                        onTextChanged: acceptableInput ? testInstructionColorColorArea.color = '#' + testInstructionColorHLayoutTextField.text : print("Input not acceptable")
                                            }
                                        }
                                    }
                            }
                        // -- ComboBoxAttributeTypeGuid --
                        ColumnLayout {
                            id: comboBoxAttributeTypeGuidLayout
                            //spacing: 2

                            Label {
                                text: "Test Instruction Type Guid"
                            }
                            TextField {
                                id: comboBoxAttributeTypeGuidTextField
                                placeholderText: "guid based on Attribute Type Name"

                                enabled: false

                                ToolTip.delay: 1000
                                ToolTip.timeout: 5000
                                ToolTip.visible: hovered
                                ToolTip.text: qsTr("The unique guid for the Type of the ComboBox. Used for datamanupulation.")

                                }
                        }

                        // -- ComboBoxAttributeTypeName --
                        ColumnLayout {
                            id: comboBoxattributeTypeNameLayput
                            //spacing: 2

                            Label {
                                text: "ComboBox Attribute Type Name"
                            }
                            ComboBox {
                                id: comboBoxAttributeTypeNameCombobox
                                Layout.fillWidth: true

                                ToolTip.delay: 1000
                                ToolTip.timeout: 5000
                                ToolTip.visible: hovered
                                ToolTip.text: qsTr("The name for the Type of the ComboBox. Used for datamanupulation.")

                                currentIndex: 0
                                model: ListModel {
                                    id: comboBoxAttributeTypeNameItems
                                    ListElement { text: "- Choose Domain -"; color: "" }
                                    ListElement { text: "Banana"; color: "Yellow" }
                                    ListElement { text: "Apple"; color: "Green" }
                                    ListElement { text: "Coconut"; color: "Brown" }
                                }
                                onCurrentIndexChanged: {
                                    comboBoxAttributeTypeGuidTextField.text = testInstructionTypeNameItems.get(currentIndex).color
                                    //console.log(systemDomainNameTextField.text)

                                }
                            }
                        }

                        // -- ComboBoxAttributeValueGuid --
                        ColumnLayout {
                            id: comboBoxAttributeValueGuidLayout
                            //spacing: 2

                            Label {
                                text: "ComboBox Attribute Value Guid"
                            }
                            TextField {
                                id: comboBoxAttributeValueGuidTextField
                                placeholderText: "Preset guid"

                                enabled: false

                                ToolTip.delay: 1000
                                ToolTip.timeout: 5000
                                ToolTip.visible: hovered
                                ToolTip.text: qsTr("The guid of the value for the the ComboBox, used for preset values.")

                                }
                            }

                            // -- ComboBoxAttributeValue --
                            ColumnLayout {
                                id: comboBoxAttributeValueLayout
                                //spacing: 2

                                Label {
                                    text: "ComboBox Attribute Value"
                                }
                                TextField {
                                    id: comboBoxAttributeValueTextField
                                    placeholderText: "Preset value"

                                    enabled: true

                                    ToolTip.delay: 1000
                                    ToolTip.timeout: 5000
                                    ToolTip.visible: hovered
                                    ToolTip.text: qsTr("The value for the the ComboBox, used for preset values.")

                                    }
                                }

                        }
                    }

                }








        }
    }
}



