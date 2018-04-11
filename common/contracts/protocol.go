package contracts

import "fmt"

type ITransmitter interface {
	Send(Payload) error
}

type IReceiver interface {
	ListenAndServe() error
}

type Payload struct {
	NodeId [8]byte
	Number uint64
}

func (p Payload) String() string {
	return fmt.Sprintf("payload [number=%020d]", p.Number)
}