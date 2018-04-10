package main

import (
	"flag"
	"fmt"
)

var Build string

var generationSpeed int

func init() {
	//	TODO: add logger IP address configuration
	flag.IntVar(&generationSpeed, "generation_speed", 0, "speed of Fibonacci number generation (number/s)")
	flag.Parse()
}

func printSplashScreen() {
	fmt.Printf("Fibonacci generator v.%s", Build)
}

func main() {

}