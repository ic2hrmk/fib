package config

import (
	"flag"
	"log"

	"github.com/ic2hrmk/fib/common"
)

type configuration struct {
	//	Generator settings
	generationSpeed float64

	//	Logger communication settings
	nodeId        [8]byte
	loggerAddress string
}

func init() {
	//	TODO: add logger IP address configuration
	flag.Float64Var(&appConfiguration.generationSpeed, "generation_speed", -1, "speed of Fibonacci number generation (number/s)")
	flag.Parse()

	//	Validation
	if appConfiguration.generationSpeed == -1 {
		log.Fatal("generation speed wasn't configured")
	}

	copy(appConfiguration.nodeId[:], common.GenerateRandomString(8))
	appConfiguration.loggerAddress = common.DefaultLoggerIP
}

var appConfiguration configuration

func GetGenerationSpeed() float64 {
	return appConfiguration.generationSpeed
}

func GetNodeId() [8]byte {
	return appConfiguration.nodeId
}

func GetLoggerAddress() string {
	return appConfiguration.loggerAddress
}