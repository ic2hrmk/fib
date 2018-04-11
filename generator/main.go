package main

import (
	"log"
	"sync"

	"github.com/ic2hrmk/fib/generator/config"
	"github.com/ic2hrmk/fib/generator/generator"
	"github.com/ic2hrmk/fib/generator/scheduler"
	protocol "github.com/ic2hrmk/fib/generator/transmitter"
	_ "github.com/ic2hrmk/fib/common"
	"github.com/ic2hrmk/fib/common"
)

var Build string

func printConfigScreen() {
	log.Println("Fibonacci generator v.", Build)
	log.Println("Generation speed: ", config.GetGenerationSpeed())
	log.Println("Unique node ID:	  ", config.GetNodeId())
	log.Println("Logger address:   ", config.GetLoggerAddress())
}

func main() {
	//	Print utility configuration
	printConfigScreen()

	//	Scheduler setup
	scheduler := scheduler.Scheduler{
		Transmitter:  protocol.NewTransmitter(config.GetNodeId(), config.GetLoggerAddress()),
		Generator:    generator.NewFibonacciGenerator(),
		TactDuration: common.GetGeneratorTactTime(config.GetGenerationSpeed()),
	}

	//	Connect transmitter and generator
	var wg sync.WaitGroup
	wg.Add(1)
	go scheduler.Run()
	wg.Wait()
}