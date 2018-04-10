package main

import (
	"log"
	"sync"
	"flag"
	"time"

	"github.com/ic2hrmk/fib/generator/config"
	"github.com/ic2hrmk/fib/generator/model"
	"github.com/ic2hrmk/fib/common/contracts"
	protocol "github.com/ic2hrmk/fib/generator/transmitter"
)

var Build string

func init() {
	//	TODO: add logger IP address configuration

	var generationSpeed float64

	flag.Float64Var(&generationSpeed, "generation_speed", -1, "speed of Fibonacci number generation (number/s)")
	flag.Parse()

	if generationSpeed == -1 {
		log.Fatal("generation speed wasn't configured")
	}

	config.SetGenerationSpeed(generationSpeed)
}

func printConfigScreen() {
	log.Println("Fibonacci generator v.", Build)
	log.Println("Generation speed: ", config.GetGenerationSpeed())
	log.Println("Unique node ID:	  ", config.GetNodeId())
	log.Println("Logger address:   ", config.GetLoggerAddress())
}

func getGeneratorTactTime(generationSpeed float64) time.Duration {
	return time.Duration(float64(1) / float64(generationSpeed)) * time.Millisecond
}

func main() {
	var err error

	//	Print utility configuration
	printConfigScreen()

	//	Setup transmission part
	transmitter := &protocol.Transmitter{}
	transmitter, err = protocol.NewTransmitter()
	if err != nil {
		log.Fatal(err)
	}

	//	Setup generator part
	fibonacciGenerator := model.NewFibonacciGenerator()

	//	Connect transmitter and generator
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		ticker := time.NewTicker(getGeneratorTactTime(config.GetGenerationSpeed()))
		for range ticker.C {
			fibNumber := fibonacciGenerator.Current()
			fibonacciGenerator.Next()

			payload := contracts.Payload{
				NodeId: config.GetNodeId(),
				FibNumber: fibNumber,
			}

			err = transmitter.Send(payload)
			if err != nil {
				log.Println("ERROR: ", err.Error())
			} else {
				log.Println("sent", fibNumber)
			}
		}
	}()
	wg.Wait()
}