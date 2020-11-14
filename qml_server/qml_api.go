package qml_server

import (
	"encoding/json"
	"github.com/google/uuid"
)

// *********************************************************************
// Used by QML to verify that the QML-code was started from server and not from QML-editor
func checkIfServerIsOnline() bool {

	return true
}

// *********************************************************************
// Generates a guid in string format
func generateGuid() string {
	var returnGuid string
	newGuid, _ := uuid.NewUUID()

	returnGuid = newGuid.String()

	return returnGuid

}

// *********************************************************************
// Used by QML to load stored data about Plugins
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
// Used by QML to load stored data about Domainss
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
