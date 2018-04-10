package config

import (
	"github.com/ic2hrmk/fib/common"
)

type configuration struct {
	//	Generator settings
	generationSpeed int

	//	Logger communication settings
	nodeId        string
	loggerAddress string
}

const (
	nodeIdLength    = 8
	defaultLoggerIP = "127.0.0.1:10000"
)

func init() {
	appConfiguration.nodeId = common.GenerateRandomString(nodeIdLength)
	appConfiguration.loggerAddress = defaultLoggerIP
}

var appConfiguration configuration

func GetGenerationSpeed() int {
	return appConfiguration.generationSpeed
}

func GetNodeId() string {
	return appConfiguration.nodeId
}

func GetLoggerAddress() string {
	return appConfiguration.nodeId
}

func SetGenerationSpeed(generationSpeed int) {
	appConfiguration.generationSpeed = generationSpeed
}