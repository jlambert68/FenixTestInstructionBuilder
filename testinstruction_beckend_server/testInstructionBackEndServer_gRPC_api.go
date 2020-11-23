package testinstruction_beckend_server

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"jlambert/FenixInception3/FenixTestInstructionBuilder/grpc_api/backend_server_grpc_api"
)

// Used for checking if Backen Server is alive
func (s *TestInstructionBackendServer) AreYouAlive(ctx context.Context, emptyParameter *backend_server_grpc_api.EmptyParameter) (*backend_server_grpc_api.AckNackResponse, error) {

	testInstructionBackendObject.logger.WithFields(logrus.Fields{}).Info("Incoming: 'AreYouAlive'")

	// Temp-solution for handling that QML-server is up and running
	testInstructionBackendObject.qmlServerHasConnected = true

	testInstructionBackendObject.logger.WithFields(logrus.Fields{}).Info("Leaving 'AreYouAlive'")
	return &backend_server_grpc_api.AckNackResponse{Acknack: true, Comments: "I'am alive, from " + testInstructionBackendObject.uuid}, nil
}

// *********************************************************************
// Generates a guid in string format to be sent front end
func (s *TestInstructionBackendServer) GenerateGuid(ctx context.Context, emptyParameter *backend_server_grpc_api.EmptyParameter) (*backend_server_grpc_api.GuidResponse, error) {

	var returnGuid *backend_server_grpc_api.GuidResponse
	newGuid, _ := uuid.NewUUID()

	returnGuid = &backend_server_grpc_api.GuidResponse{
		Guid:                 newGuid.String(),
		Acknack:              true,
		Comments:             "",
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}

	return returnGuid, nil

}

// *********************************************************************
// Load stored data about Plugins and send towards frontend, via QML-server
func (s *TestInstructionBackendServer) LoadPluginModelFromServer(ctx context.Context, emptyParameter *backend_server_grpc_api.EmptyParameter) (*backend_server_grpc_api.PluginQmlModelFromServerResponse, error) {

	var returnMessage *backend_server_grpc_api.PluginQmlModelFromServerResponse

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

	returnMessage = &backend_server_grpc_api.PluginQmlModelFromServerResponse{
		JsonStringForPluginQmlModel: jsonToSendAsString,
		Acknack:                     false,
		Comments:                    "",
		XXX_NoUnkeyedLiteral:        struct{}{},
		XXX_unrecognized:            nil,
		XXX_sizecache:               0,
	}

	return returnMessage, nil

}

// *********************************************************************
// Load stored data about Domains and send towards frontend, via QML-server
func (s *TestInstructionBackendServer) LoadDomainModelFromServer(ctx context.Context, emptyParameter *backend_server_grpc_api.EmptyParameter) (*backend_server_grpc_api.DomainQmlModelFromServerResponse, error) {

	var returnMessage *backend_server_grpc_api.DomainQmlModelFromServerResponse

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

	returnMessage = &backend_server_grpc_api.DomainQmlModelFromServerResponse{
		JsonStringForDomainQmlModel: jsonToSendAsString,
		Acknack:                     false,
		Comments:                    "",
		XXX_NoUnkeyedLiteral:        struct{}{},
		XXX_unrecognized:            nil,
		XXX_sizecache:               0,
	}

	return returnMessage, nil

}
