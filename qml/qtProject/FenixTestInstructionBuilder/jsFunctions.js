// Javascript-functions for calling Golang
// All functions can be executed in a mocked version for use when started from QT-designer

// *********************************************************************************
// Check if Golang-server exist if not then force use Mock-data
function jsCheckIfServerIsOnline() {
    var serverResponse = false;

    //Only do this test once
    if (rootTable.checkHasBeenDone === false) {
        rootTable.checkHasBeenDone = true

        serverResponse = QmlBridge.checkIfServerIsOnline();
        if (serverResponse === true) {
            rootTable.startedByGolang = true
        } else {
            rootTable.startedByGolang = false
        }
    }

}

// *********************************************************************************
// Generate a GUID
function jsGenerateGuid() {
    // If QML is not started from Golang then mock value
    if (rootTable.startedByGolang === true) {
        return QmlBridge.generateGuid();
    } else {
        return "aaabbbcc-2ad1-4677-8e1d-909a7ecabaaa"
    }
}

// *********************************************************************************
// Load PluginModel from server
function jsLoadPluginModelFromServer() {
    // Check if server is running
    jsCheckIfServerIsOnline()

    // If QML is not started from Golang then mock value
    var pluginModeToProcess;

    if (rootTable.startedByGolang === true) {
        // Receive PluginModel from Server
        var receivedmodelString = QmlBridge.loadPluginModelFromServer();

        // Convert received model into json-objectName
        pluginModeToProcess = JSON.parse(receivedmodelString);

    } else {

        // JSON-mockObject for PluginTable
        var modelMock = {
            "ListElement": [
                {
                    "name": "Send UTR",
                    "guid": "aa43085a-feee-4ae7-8ccf-03f69c0704aa",
                    "readyForUse": true,
                    "activated": true,
                    "description": "Sends UTRs via MQ towards Custody Cash"
                },
                {
                    "name": "Validate New Pacs008",
                    "guid": "aa56499c-2ad1-4677-8e1d-909a7ecab5aa",
                    "readyForUse": false,
                    "activated": false,
                    "description": "Validates that a newly created Pacs008 has been sent to CMaaS"
                }
            ]
        };

        pluginModeToProcess = modelMock;

    }

    // Process JSON object for content
    pluginModeToProcess["ListElement"].forEach((item) => {
                                                   // Create a new Plugin in table
                                                   pluginsModel.append({name: item.name,
                                                                           guid: item.guid,
                                                                           readyForUse: item.readyForUse,
                                                                           activated: item.activated,
                                                                           description: item.description})
                                               });
}

// *********************************************************************************
// Load DomainModel from server
function jsLoadDomainModelFromServer() {
    // Check if server is running
    jsCheckIfServerIsOnline()

    // If QML is not started from Golang then mock value
    var domainModeToProcess;

    if (rootTable.startedByGolang === true) {
        // Receive DomainModel from Server
        var receivedmodelString = QmlBridge.loadDomainModelFromServer();

        // Convert received model into json-objectName
        domainModeToProcess = JSON.parse(receivedmodelString);

    } else {

        // JSON-mockObject for DomainTable
        var modelMock = {
            "ListElement": [
                {
                    "name": "Custody Cash",
                    "guid": "bb43085a-feee-4ae7-8ccf-03f69c0704bb",
                    "readyForUse": true,
                    "activated": true,
                    "description": "All test regarding Custody Cash"
                },
                {
                    "name": "Cusotdy Arrangement",
                    "guid": "bb56499c-2ad1-4677-8e1d-909a7ecab5bb",
                    "readyForUse": false,
                    "activated": false,
                    "description": "All tests regarding Custody Arrangement"
                }
            ]
        };

        domainModeToProcess = modelMock;

    }

    // Process JSON object for content
    domainModeToProcess["ListElement"].forEach((item) => {
                                                   // Create a new Plugin in table
                                                   domainsModel.append({name: item.name,
                                                                           guid: item.guid,
                                                                           readyForUse: item.readyForUse,
                                                                           activated: item.activated,
                                                                           description: item.description})
                                               });
}

// *********************************************************************************
