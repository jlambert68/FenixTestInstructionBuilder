
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

            			text: "sendVariantListMap"
            			onClicked: {
            				text = "look into the console"
            				root.template.sendVariantListMap(true, [1.23, "hello"],  {"json": {"A": true, "B": 2.23, "C": "hello", "D": {"DA": 3.14, "DB": "Pi", "DC": {"DCA": 1.73, "DCB": "Roten ur 2"}}}})
            			}
            		}
            	}