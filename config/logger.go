package config

import (
	"fmt"
	"os"

	"github.com/karolpiernikarz/automanage/middlewares"
	"github.com/karolpiernikarz/automanage/utils"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/sirupsen/logrus"
)

func InitLogger() {
	var logfile *os.File
	var logfileApp *os.File
	var awsLogFile *os.File
	os.MkdirAll("/var/log/machhub", os.ModePerm)
	writable, err := utils.IsPathWritable("/var/log/machhub")
	if err != nil {
		fmt.Println("can't write logs to /var/log/machhub, using ./log dir. Error message: " + err.Error())
	}
	if writable {
		logfile, _ = os.OpenFile("/var/log/machhub/access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0660)
		logfileApp, _ = os.OpenFile("/var/log/machhub/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0660)
		awsLogFile, _ = os.OpenFile("/var/log/machhub/aws.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0660)
	} else {
		writable, err := utils.IsPathWritable("log")
		if err != nil {
			fmt.Println("error checking log dir perms")
		}
		if writable {
			logfile, _ = os.OpenFile("log/access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0660)
			logfileApp, _ = os.OpenFile("log/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0660)
			awsLogFile, _ = os.OpenFile("log/aws.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0660)
		}
	}

	middlewares.AwsLogFile = awsLogFile

	log.Logger = zerolog.New(logfile)

	logrus.SetOutput(logfileApp)
	logrus.SetReportCaller(true)

}
