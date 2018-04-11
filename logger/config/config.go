package config

import (
	"flag"

	"github.com/ic2hrmk/fib/common"
)

type configuration struct {
	//	Logger settings
	filePath   string
	flowSpeed  float64
	bufferSize int64

	//	Logger communication settings
	loggerAddress string
}

func init() {
	//	TODO: add logger IP address configuration
	flag.Float64Var(&appConfiguration.flowSpeed, "flow_speed", -1, "receiving flow speed (number/s)")
	flag.Int64Var(&appConfiguration.bufferSize, "buffer_size", -1, "buffer size (bytes)")
	flag.StringVar(&appConfiguration.filePath, "file_path", "", "path to save logs")
	flag.Parse()

	appConfiguration.loggerAddress = common.DefaultLoggerIP
}

var appConfiguration configuration

func GetFlowSpeed() float64 {
	return appConfiguration.flowSpeed
}

func GetFilePath() string {
	return appConfiguration.filePath
}

func GetBufferSize() int64 {
	return appConfiguration.bufferSize
}

func GetLoggerAddress() string {
	return appConfiguration.loggerAddress
}