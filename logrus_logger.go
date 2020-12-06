package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"jlambert/FenixInception3/FenixTestInstructionBuilder/common_config"
	"log"
	"os"
	"time"
)

func InitLogger(filename string) *logrus.Logger {
	var logger *logrus.Logger

	a := logrus.StandardLogger()
	fmt.Println("*** Logrus ***", a)
	fmt.Println("*** Logrus2 ***")
	logger = logrus.StandardLogger()

	switch common_config.LoggingLevel {

	case logrus.DebugLevel:
		log.Println("'common_config.LoggingLevel': ", common_config.LoggingLevel)

	case logrus.InfoLevel:
		log.Println("'common_config.LoggingLevel': ", common_config.LoggingLevel)

	case logrus.WarnLevel:
		log.Println("'common_config.LoggingLevel': ", common_config.LoggingLevel)

	default:
		log.Println("Not correct value for debugging-level, this was used: ", common_config.LoggingLevel)
		os.Exit(0)

	}

	logrus.SetLevel(common_config.LoggingLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339Nano,
		DisableSorting:  true,
	})

	//If no file then set standard out

	if filename == "" {
		logger.Out = os.Stdout

	} else {
		file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
		if err == nil {
			logger.Out = file
		} else {
			log.Println("Failed to log to file, using default stderr")
		}
	}

	return logger

	// Should only be done from init functions
	//grpclog.SetLoggerV2(grpclog.NewLoggerV2(logger.Out, logger.Out, logger.Out))

}
