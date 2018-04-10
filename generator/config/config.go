package config

import (
	"github.com/ic2hrmk/fib/common"
)

type configuration struct {
	//	Generator settings
	generationSpeed float64

	//	Logger communication settings
	nodeId        string
	loggerAddress string
}

const (
	nodeIdLength    = 8
	defaultLoggerIP = "127.0.0.1:10000"
)

func init() {
	appConfiguration.generationSpeed = 0
	appConfiguration.nodeId = common.GenerateRandomString(nodeIdLength)
	appConfiguration.loggerAddress = defaultLoggerIP
}

var appConfiguration configuration

func GetGenerationSpeed() float64 {
	return appConfiguration.generationSpeed
}

func GetNodeId() string {
	return appConfiguration.nodeId
}

func GetLoggerAddress() string {
	return appConfiguration.loggerAddress
}

func SetGenerationSpeed(generationSpeed float64) {
	appConfiguration.generationSpeed = generationSpeed
}