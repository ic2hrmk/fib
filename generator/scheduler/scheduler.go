package scheduler

import (
	"time"
	"log"

	"github.com/ic2hrmk/fib/common/contracts"
)

type Scheduler struct {
	Transmitter  contracts.ITransmitter
	Generator    contracts.INumberGenerator
	TactDuration time.Duration
}

func (s Scheduler) Run() {
	go func() {
		var err error
		ticker := time.NewTicker(s.TactDuration)
		for range ticker.C {
			payload := contracts.Payload{
				Number: s.Generator.Current(),
			}

			s.Generator.Next()

			err = s.Transmitter.Send(payload)
			if err != nil {
				log.Println("ERROR: ", err.Error())
			} else {
				log.Println("TRANSMISSION: ", payload)
			}
		}
	}()
}
