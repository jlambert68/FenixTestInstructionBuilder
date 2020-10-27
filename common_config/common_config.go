package common_config

import "github.com/sirupsen/logrus"

// gRPC-ports
const MotherServer_address = "127.0.0.1"
const MotherServer_port = ":5998"

const LocalWorkerServer_address = "127.0.0.1"
const LocalWorkerServer_initial_port = 6660

// Logrus debug level

//const LoggingLevel = logrus.DebugLevel
//const LoggingLevel = logrus.InfoLevel
const LoggingLevel = logrus.InfoLevel
