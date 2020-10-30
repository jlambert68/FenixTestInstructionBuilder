// Javascript-functions for calling Golang
// All functions can be executed in a mocked version for use when started from QT-designer

function jsGenerateGuid() {
    // If QML is not started from Golang then mock value
    if (rootTable.startedByGolang === true) {
    return QmlBridge.generateGuid();
    } else {
        return "aaabbbcc-2ad1-4677-8e1d-909a7ecabaaa"
    }

}
