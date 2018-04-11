package main

import (
	"log"
	"sync"

	"github.com/ic2hrmk/fib/logger/config"
	"github.com/ic2hrmk/fib/logger/server"
)

var Build string

func printConfigScreen() {
	log.Println("Fibonacci generator v.", Build)
	log.Println("Log path:			", config.GetFilePath())
	log.Println("Flow speed:			", config.GetFlowSpeed())
	log.Println("Buffer size:		", config.GetBufferSize())
	log.Println("Logger address:		", config.GetLoggerAddress())
}


func main() {
	printConfigScreen()

	server := server.NewServer(
		config.GetLoggerAddress(),
		config.GetFlowSpeed(),
		config.GetBufferSize(),
	)

	var wg sync.WaitGroup
	wg.Add(1)
	server.ListenAndServe()
	wg.Wait()
}
